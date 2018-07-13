package store

import (
	"context"
	"fmt"
	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/go-ns/log"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/pkg/errors"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
)

//go:generate moq -out mock/neo4j.go -pkg mock . DBPool Conn Rows

const (
	getCodeListsQuery       = "MATCH (i) WHERE i:_%s%s RETURN distinct labels(i) as labels, i"
	getCodeListQuery        = "MATCH (i:_%s:`_name_%s`) RETURN i"
	getCodeListEditionQuery = "MATCH (i:_%s:`_name_%s` {edition:" + `"%s"` + "}) RETURN i"
	countEditions           = "MATCH (cl:_code_list:`_name_%s`) WHERE cl.edition = %q RETURN count(*)"
	getCodesQuery           = "MATCH (c:_code) -[r:usedBy]->(cl:_code_list: `_name_%s`) WHERE cl.edition = %q RETURN c, r"
	getCodeQuery            = "MATCH (c:_code) -[r:usedBy]->(cl:_code_list: `_name_%s`) WHERE cl.edition = %q AND c.value = %q RETURN c, r"
)

var (
	errEditionNotFound = errors.New("edition does not exist")
)

type Conn bolt.Conn
type Rows bolt.Rows

type BoltDB interface {
	QueryForResult(query string, params map[string]interface{}, extractResult dpbolt.ResultMapper) (int, error)
	QueryForResults(query string, params map[string]interface{}, extractResult dpbolt.ResultMapper) (int, error)
	Close() error
}

// NeoDataStore represents the necessary information to access
// neo4j
type NeoDataStore struct {
	pool          DBPool
	codeListLabel string
	db            BoltDB
}

// DBPool contains the methods to control access to the Neo4J
// database pool
type DBPool interface {
	OpenPool() (bolt.Conn, error)
	Close() error
}

// Close is a wrapper for the neo pool close
func (n *NeoDataStore) Close() error {
	n.db.Close()
	return n.pool.Close()
}

// CreateNeoDataStore allows the creation of a NeoDataStore
func CreateNeoDataStore(addr, codelistLabel string, conns int) (n *NeoDataStore, err error) {
	store, err := bolt.NewClosableDriverPool(addr, conns)
	if err != nil {
		return
	}

	boltDB := dpbolt.New(store)

	n = &NeoDataStore{
		pool:          store,
		codeListLabel: codelistLabel,
		db:            boltDB,
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
	codeListEditionsMap, mapper := codeListsMapper()

	_, err := n.db.QueryForResults(query, nil, mapper)
	if err != nil {
		return nil, err
	}

	codeLists := &models.CodeListResults{}
	for _, codelist := range *codeListEditionsMap {
		codeLists.Items = append(codeLists.Items, *codelist)
	}

	return codeLists, nil
}

func (n *NeoDataStore) GetCodeList(ctx context.Context, code string) (*models.CodeList, error) {
	log.InfoCtx(ctx, "about to query neo4j for code list", log.Data{"code_list_id": code})

	query := fmt.Sprintf(getCodeListQuery, n.codeListLabel, code)
	codeList := &models.CodeList{}
	mapper := codeListMapper(codeList, code)

	count, err := n.db.QueryForResult(query, nil, mapper)
	if err != nil {
		return nil, err
	}
	if count == 0 {
		return nil, datastore.NOT_FOUND
	}

	return codeList, nil
}

func (n *NeoDataStore) GetEditions(ctx context.Context, codeListID string) (*models.Editions, error) {
	log.InfoCtx(ctx, "about to query neo4j for code list editions", log.Data{"code_list_id": codeListID})

	conn, err := n.pool.OpenPool()
	if err != nil {
		return nil, errors.WithMessage(err, "could not retrieve connection from pool")
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.ErrorCtx(ctx, errors.WithMessage(err, "could not close neo4j connection"), nil)
		}
	}()

	query := fmt.Sprintf(getCodeListQuery, n.codeListLabel, codeListID)

	rows, err := conn.QueryNeo(query, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "could not run neo4j query")
	}

	var row []interface{}

	editions := &models.Editions{}
	for row, _, err = rows.NextNeo(); err == nil; row, _, err = rows.NextNeo() {
		props := row[0].(graph.Node).Properties

		edition := props["edition"].(string)

		editionModel := models.Edition{
			Edition: edition,
			Label:   props["label"].(string),
			Links: models.EditionLinks{
				Self: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s", codeListID, edition),
					ID:   edition,
				},
				Editions: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions", codeListID),
				},
				Codes: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes", codeListID, edition),
				},
			},
		}

		editions.Items = append(editions.Items, editionModel)
	}

	if len(editions.Items) == 0 {
		return nil, datastore.NOT_FOUND
	}

	editions.NumberOfResults = len(editions.Items)

	return editions, nil
}

func (n *NeoDataStore) GetEdition(ctx context.Context, codeListID, edition string) (*models.Edition, error) {
	log.InfoCtx(ctx, "about to query neo4j for code list edition", log.Data{"code_list_id": codeListID, "edition": edition})

	conn, err := n.pool.OpenPool()
	if err != nil {
		return nil, errors.WithMessage(err, "could not retrieve connection from pool")
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.ErrorCtx(ctx, errors.WithMessage(err, "could not close neo4j connection"), nil)
		}
	}()

	query := fmt.Sprintf(getCodeListEditionQuery, n.codeListLabel, codeListID, edition)

	rows, err := conn.QueryNeo(query, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "could not run neo4j query")
	}

	var row []interface{}

	count := 0
	editionModel := &models.Edition{}
	for row, _, err = rows.NextNeo(); err == nil; row, _, err = rows.NextNeo() {
		count++
		if count > 1 {
			return nil, errors.Errorf("more than one unique edition found for codelist: %s, edition: %s", codeListID, edition)
		}

		props := row[0].(graph.Node).Properties

		editionModel = &models.Edition{
			Edition: edition,
			Label:   props["label"].(string),
			Links: models.EditionLinks{
				Self: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s", codeListID, edition),
					ID:   edition,
				},
				Editions: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions", codeListID),
				},
				Codes: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes", codeListID, edition),
				},
			},
		}
	}

	if count == 0 {
		return nil, datastore.NOT_FOUND
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

	var count int
	results, rowMapper := codesResultMapper(codeListID, edition)

	count, err = n.db.QueryForResults(fmt.Sprintf(getCodesQuery, codeListID, edition), nil, rowMapper)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, datastore.NOT_FOUND
	}

	return &models.CodeResults{
		Items: *results,
		Count: len(*results),
	}, nil
}

func (n *NeoDataStore) GetCode(ctx context.Context, codeListID, edition string, code string) (*models.Code, error) {
	editionExists, err := n.EditionExists(ctx, codeListID, edition)
	if err != nil {
		return nil, err
	}
	if !editionExists {
		return nil, datastore.ErrEditionNotFound
	}

	var count int
	codeModel := &models.Code{}
	extractor := codeResultMapper(codeModel, codeListID, edition)

	count, err = n.db.QueryForResult(fmt.Sprintf(getCodeQuery, codeListID, edition, code), nil, extractor)
	if err != nil {
		return nil, err
	}

	if count == 0 {
		return nil, datastore.ErrCodeNotFound
	}

	return codeModel, nil
}

func (n *NeoDataStore) EditionExists(ctx context.Context, codeListID string, edition string) (bool, error) {
	data := log.Data{"codelist_id": codeListID, "edition": edition}
	log.InfoCtx(ctx, "checking edition exists", data)

	query := fmt.Sprintf(countEditions, codeListID, edition)

	count, extractor := countEditionMapper()
	_, err := n.db.QueryForResult(query, nil, extractor)
	if err != nil {
		return false, err
	}

	// 0 or 1 is a valid response - more than 1... we have bigger issues.
	if *count > 1 {
		return false, errors.New("editionExists: multiple editions found")
	}
	return *count == 1, nil
}
