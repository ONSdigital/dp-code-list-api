package store

import (
	"context"
	"fmt"

	"strconv"

	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/dp-code-list-api/store/mapper"
	"github.com/ONSdigital/go-ns/log"
	"github.com/pkg/errors"
)

const (
	getCodeListsQuery       = "MATCH (i) WHERE i:_%s%s RETURN distinct labels(i) as labels"
	getCodeListQuery        = "MATCH (i:_code_list:`_%s_%s`) RETURN i"
	codeListExistsQuery     = "MATCH (cl:_code_list:`_%s_%s`) RETURN count(*)"
	getCodeListEditionQuery = "MATCH (i:_code_list:`_%s_%s` {edition:" + `"%s"` + "}) RETURN i"
	countEditions           = "MATCH (cl:_code_list:`_%s_%s`) WHERE cl.edition = %q RETURN count(*)"
	getCodesQuery           = "MATCH (c:_code) -[r:usedBy]->(cl:_code_list: `_%s_%s`) WHERE cl.edition = %q RETURN c, r"
	getCodeQuery            = "MATCH (c:_code) -[r:usedBy]->(cl:_code_list: `_%s_%s`) WHERE cl.edition = %q AND c.value = %q RETURN c, r"
	getCodeDatasets         = "MATCH (d)<-[inDataset]-(c:_code)-[r:usedBy]->(cl:_code_list:`_code_list_%s`) WHERE (cl.edition=" + `"%s"` + ") AND (c.value=" + `"%s"` + ") AND (d.is_published=true) RETURN d,r"
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
	codeListResults := &models.CodeListResults{}

	prefix := "_" + n.codeListLabel + "_"
	err := n.bolt.QueryForResults(query, nil, mapper.CodeLists(codeListResults, prefix))
	if err != nil && err == dpbolt.ErrNoResults {
		return nil, datastore.ErrCodeListsNotFound
	}
	if err != nil {
		return nil, err
	}

	return codeListResults, nil
}

func (n *NeoDataStore) GetCodeList(ctx context.Context, code string) (*models.CodeList, error) {
	log.InfoCtx(ctx, "about to query neo4j for code list", log.Data{"code_list_id": code})

	query := fmt.Sprintf(codeListExistsQuery, n.codeListLabel, code)
	count, mapper := mapper.CodeListCount()

	err := n.bolt.QueryForResult(query, nil, mapper)
	if err != nil && err == dpbolt.ErrNoResults {
		return nil, datastore.ErrCodeListNotFound
	}

	if err != nil {
		return nil, err
	}

	// from a Neo4j POV Codelists are't actually a thing a codeList exists if there is 1 or more edition nodes.
	if *count == 0 {
		return nil, datastore.ErrCodeListNotFound
	}

	return &models.CodeList{
		Links: models.CodeListLink{
			Self: &models.Link{
				ID:   code,
				Href: fmt.Sprintf("/code-lists/%s", code),
			},
			Editions: &models.Link{
				Href: fmt.Sprintf("/code-lists/%s/editions", code),
			},
		},
	}, nil
}

func (n *NeoDataStore) GetEditions(ctx context.Context, codeListID string) (*models.Editions, error) {
	log.InfoCtx(ctx, "about to query neo4j for code list editions", log.Data{"code_list_id": codeListID})

	query := fmt.Sprintf(getCodeListQuery, n.codeListLabel, codeListID)

	editions := &models.Editions{}
	err := n.bolt.QueryForResults(query, nil, mapper.Editions(editions, codeListID))
	if err != nil && err == dpbolt.ErrNoResults {
		return nil, datastore.ErrEditionsNotFound
	}
	if err != nil {
		return nil, err
	}

	editions.Count = len(editions.Items)
	editions.Limit = len(editions.Items)
	editions.TotalCount = len(editions.Items)

	return editions, nil
}

func (n *NeoDataStore) GetEdition(ctx context.Context, codeListID, edition string) (*models.Edition, error) {
	log.InfoCtx(ctx, "about to query neo4j for code list edition", log.Data{"code_list_id": codeListID, "edition": edition})

	query := fmt.Sprintf(getCodeListEditionQuery, n.codeListLabel, codeListID, edition)

	editionModel := &models.Edition{}
	err := n.bolt.QueryForResult(query, nil, mapper.Edition(editionModel, codeListID, edition))
	if err != nil && err == dpbolt.ErrNoResults {
		return nil, datastore.ErrEditionNotFound
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
	query := fmt.Sprintf(getCodesQuery, n.codeListLabel, codeListID, edition)

	err = n.bolt.QueryForResults(query, nil, mapper.Codes(codeResults, codeListID, edition))
	if err != nil && err == dpbolt.ErrNoResults {
		return nil, datastore.ErrCodesNotFound
	}

	if err != nil {
		return nil, err
	}

	count := len(codeResults.Items)
	codeResults.TotalCount = count
	codeResults.Count = count
	codeResults.Limit = count
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
	query := fmt.Sprintf(getCodeQuery, n.codeListLabel, codeListID, edition, code)

	err = n.bolt.QueryForResult(query, nil, mapper.Code(codeModel, codeListID, edition))
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

	query := fmt.Sprintf(countEditions, n.codeListLabel, codeListID, edition)

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

func (n *NeoDataStore) GetCodeDatasets(ctx context.Context, codeListID, edition, code string) (*models.Datasets, error) {
	editionExists, err := n.EditionExists(ctx, codeListID, edition)
	if err != nil {
		return nil, err
	}
	if !editionExists {
		return nil, datastore.ErrEditionNotFound
	}

	query := fmt.Sprintf(getCodeDatasets, codeListID, edition, code)
	datasets := make(mapper.Datasets, 0)

	err = n.bolt.QueryForResults(query, nil, mapper.CodesDatasets(datasets))
	if err != nil && err == dpbolt.ErrNoResults {
		return nil, datastore.ErrDatasetsNotFound
	} else if err != nil {
		return nil, err
	}

	datasetsModel := createDatasetsResponseModel(datasets, codeListID)

	datasetsModel.TotalCount = len(datasetsModel.Items)
	datasetsModel.Count = len(datasetsModel.Items)
	datasetsModel.Limit = len(datasetsModel.Items)

	return datasetsModel, err
}

func createDatasetsResponseModel(datasets mapper.Datasets, codeListID string) *models.Datasets {

	datasetsModel := &models.Datasets{}
	for datasetID, dataset := range datasets {

		editions := make([]models.DatasetEdition, 0)

		for editionID, versions := range dataset.Editions {

			latestVersion := strconv.Itoa(max(versions))

			editions = append(editions, models.DatasetEdition{
				Links: models.DatasetEditionLinks{
					Self: models.Link{
						ID:   editionID,
						Href: fmt.Sprintf("/datasets/%s/editions/%s", datasetID, editionID),
					},
					LatestVersion: models.Link{
						ID:   latestVersion,
						Href: fmt.Sprintf("/datasets/%s/editions/%s/versions/%s", datasetID, editionID, latestVersion),
					},
					DatasetDimension: models.Link{
						ID:   codeListID,
						Href: fmt.Sprintf("/datasets/%s/editions/%s/versions/%s/dimensions/%s", datasetID, editionID, latestVersion, codeListID),
					},
				},
			})
		}

		datasetsModel.Items = append(datasetsModel.Items, models.Dataset{

			DimensionLabel: dataset.DimensionLabel,
			Links: models.DatasetLinks{
				Self: models.Link{
					ID:   datasetID,
					Href: fmt.Sprintf("/datasets/%s", datasetID),
				},
			},
			Editions: editions,
		})
	}

	return datasetsModel
}

func max(input []int) (max int) {

	for _, value := range input {
		if value > max {
			max = value
		}
	}

	return max
}
