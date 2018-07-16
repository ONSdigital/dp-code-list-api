package store

import (
	"context"
	"fmt"
	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/go-ns/log"
	"github.com/pkg/errors"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-code-list-api/store/mapper"
)

const (
	getCodeListsQuery       = "MATCH (i) WHERE i:_%s%s RETURN distinct labels(i) as labels, i"
	getCodeListQuery        = "MATCH (i:_%s:`_name_%s`) RETURN i"
	getCodeListEditionQuery = "MATCH (i:_%s:`_name_%s` {edition:" + `"%s"` + "}) RETURN i"
	countEditions           = "MATCH (cl:_code_list:`_name_%s`) WHERE cl.edition = %q RETURN count(*)"
	getCodesQuery           = "MATCH (c:_code) -[r:usedBy]->(cl:_code_list: `_name_%s`) WHERE cl.edition = %q RETURN c, r"
	getCodeQuery            = "MATCH (c:_code) -[r:usedBy]->(cl:_code_list: `_name_%s`) WHERE cl.edition = %q AND c.value = %q RETURN c, r"
)

type BoltDB interface {
	QueryForResult(query string, params map[string]interface{}, extractResult dpbolt.ResultMapper) error
	QueryForResults(query string, params map[string]interface{}, extractResult dpbolt.ResultMapper) error
	Close() error
}

// NeoDataStore represents the necessary information to access
// neo4j
type NeoDataStore struct {
	codeListLabel string
	bolt          BoltDB
}

// Close is a wrapper for the neo pool close
func (n *NeoDataStore) Close() error {
	log.Info("attempting to close store.", nil)
	return n.bolt.Close()
}

// CreateNeoDataStore allows the creation of a NeoDataStore
func CreateNeoDataStore(boltDB BoltDB, codelistLabel string) (n *NeoDataStore, err error) {
	n = &NeoDataStore{
		codeListLabel: codelistLabel,
		bolt:          boltDB,
	}

	return
}

// GetCodeLists returns a list of code lists
func (n *NeoDataStore) GetCodeLists(ctx context.Context, filterBy string) (*models.CodeListResults, error) {
	logData := log.Data{}
	if len(filterBy) > 0 {
		logData["filter_by"] = filterBy
		filterBy = ":_" + filterBy
	}
	log.InfoCtx(ctx, "about to query neo4j for code lists", logData)

	query := fmt.Sprintf(getCodeListsQuery, n.codeListLabel, filterBy)
	codeListEditionsMap := make(map[string]*models.CodeList)

	err := n.bolt.QueryForResults(query, nil, mapper.CodeLists(codeListEditionsMap))
	if err != nil {
		return nil, err
	}

	codeLists := &models.CodeListResults{}
	for _, codelist := range codeListEditionsMap {
		codeLists.Items = append(codeLists.Items, *codelist)
	}

	return codeLists, nil
}

func (n *NeoDataStore) GetCodeList(ctx context.Context, code string) (*models.CodeList, error) {
	log.InfoCtx(ctx, "about to query neo4j for code list", log.Data{"code_list_id": code})

	query := fmt.Sprintf(getCodeListQuery, n.codeListLabel, code)
	codeList := &models.CodeList{}

	err := n.bolt.QueryForResult(query, nil, mapper.CodeList(codeList, code))
	if err != nil && err == dpbolt.ErrNoResults {
		return nil, datastore.NOT_FOUND
	}

	if err != nil {
		return nil, err
	}

	return codeList, nil
}

func (n *NeoDataStore) GetEditions(ctx context.Context, codeListID string) (*models.Editions, error) {
	log.InfoCtx(ctx, "about to query neo4j for code list editions", log.Data{"code_list_id": codeListID})

	query := fmt.Sprintf(getCodeListQuery, n.codeListLabel, codeListID)

	editions := &models.Editions{}
	err := n.bolt.QueryForResults(query, nil, mapper.Editions(editions, codeListID))
	if err != nil && err == dpbolt.ErrNoResults {
		return nil, datastore.NOT_FOUND
	}
	if err != nil {
		return nil, err
	}

	editions.NumberOfResults = len(editions.Items)
	return editions, nil
}

func (n *NeoDataStore) GetEdition(ctx context.Context, codeListID, edition string) (*models.Edition, error) {
	log.InfoCtx(ctx, "about to query neo4j for code list edition", log.Data{"code_list_id": codeListID, "edition": edition})

	query := fmt.Sprintf(getCodeListEditionQuery, n.codeListLabel, codeListID, edition)

	editionModel := &models.Edition{}
	err := n.bolt.QueryForResult(query, nil, mapper.Edition(editionModel, codeListID, edition))
	if err != nil && err == dpbolt.ErrNoResults {
		return nil, datastore.NOT_FOUND
	}

	if err != nil {
		return nil, err
	}

	return editionModel, nil
}

func (n *NeoDataStore) GetCodes(ctx context.Context, codeListID, edition string) (*models.CodeResults, error) {
	editionExists, err := n.EditionExists(ctx, codeListID, edition)
	if err != nil {
		return nil, err
	}
	if !editionExists {
		return nil, datastore.ErrEditionNotFound
	}

	codeResults := &models.CodeResults{}
	query := fmt.Sprintf(getCodesQuery, codeListID, edition)

	err = n.bolt.QueryForResults(query, nil, mapper.Codes(codeResults, codeListID, edition))
	if err != nil && err == dpbolt.ErrNoResults {
		return nil, datastore.NOT_FOUND
	}

	if err != nil {
		return nil, err
	}

	codeResults.Count = len(codeResults.Items)
	return codeResults, nil
}

func (n *NeoDataStore) GetCode(ctx context.Context, codeListID, edition string, code string) (*models.Code, error) {
	editionExists, err := n.EditionExists(ctx, codeListID, edition)
	if err != nil {
		return nil, err
	}
	if !editionExists {
		return nil, datastore.ErrEditionNotFound
	}

	codeModel := &models.Code{}
	query := fmt.Sprintf(getCodeQuery, codeListID, edition, code)

	err = n.bolt.QueryForResult(query, nil,  mapper.Code(codeModel, codeListID, edition))
	if err != nil && err == dpbolt.ErrNoResults {
		return nil, datastore.ErrCodeNotFound
	}

	if err != nil {
		return nil, err
	}

	return codeModel, nil
}

func (n *NeoDataStore) EditionExists(ctx context.Context, codeListID string, edition string) (bool, error) {
	data := log.Data{"codelist_id": codeListID, "edition": edition}
	log.InfoCtx(ctx, "checking edition exists", data)

	query := fmt.Sprintf(countEditions, codeListID, edition)

	count, extractor := mapper.EditionCount()
	err := n.bolt.QueryForResult(query, nil, extractor)
	if err != nil {
		return false, err
	}

	// 0 or 1 is a valid response - more than 1... we have bigger issues.
	if *count > 1 {
		return false, errors.New("editionExists: multiple editions found")
	}
	return *count == 1, nil
}
