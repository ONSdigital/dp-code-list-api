package store

import (
	"context"
	"fmt"

	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/go-ns/log"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/pkg/errors"
)

const (
	getCodeListsQuery       = "MATCH (i) where i:_%s%s RETURN distinct labels(i) as labels, i"
	getCodeListQuery        = `MATCH (i:_%s {code:"%s"}) RETURN i`
	getCodeListEditionQuery = `MATCH (i:_%s {code:"%s",year:%s}) RETURN i`
)

// NeoDataStore represents the necessary information to access
// neo4j
type NeoDataStore struct {
	pool          DBPool
	codeListLabel string
}

// DBPool contains the methods to control access to the Neo4J
// database pool
type DBPool interface {
	OpenPool() (bolt.Conn, error)
	Close() error
}

// Close is a wrapper for the neo pool close
func (n NeoDataStore) Close() error {
	return n.pool.Close()
}

// CreateNeoDataStore allows the creation of a NeoDataStore
func CreateNeoDataStore(addr, codelistLabel string, conns int) (n NeoDataStore, err error) {
	store, err := bolt.NewClosableDriverPool(addr, conns)
	if err != nil {
		return
	}

	n = NeoDataStore{
		pool:          store,
		codeListLabel: codelistLabel,
	}

	return
}

// GetCodeLists returns a list of code lists
func (n NeoDataStore) GetCodeLists(ctx context.Context, filterBy string) (*models.CodeListResults, error) {
	log.InfoCtx(ctx, "about to query neo4j for all code lists", nil)

	conn, err := n.pool.OpenPool()
	if err != nil {
		return nil, errors.WithMessage(err, "could not retrieve connection from pool")
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.ErrorCtx(ctx, errors.WithMessage(err, "could not close neo4j connection"), nil)
		}
	}()

	if len(filterBy) > 0 {
		filterBy = ":_" + filterBy
	}

	query := fmt.Sprintf(getCodeListsQuery, n.codeListLabel, filterBy)

	rows, err := conn.QueryNeo(query, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "could not run neo4j query")
	}

	var row []interface{}

	codeLists := &models.CodeListResults{}
	for row, _, err = rows.NextNeo(); err == nil; row, _, err = rows.NextNeo() {
		if len(row) < 2 {
			return nil, errors.Errorf("expected at least two rows, got %d", len(row))
		}
		props := row[1].(graph.Node).Properties

		code := props["code"].(string)

		codeList, err := n.GetCodeList(ctx, code)
		if err != nil {
			return nil, err
		}

		codeLists.Items = append(codeLists.Items, *codeList)
	}

	return codeLists, nil
}

// GetCodeList returns an individual code list for a given code
func (n NeoDataStore) GetCodeList(ctx context.Context, code string) (*models.CodeList, error) {
	log.InfoCtx(ctx, "about to query neo4j for code list", log.Data{"code_list_id": code})

	conn, err := n.pool.OpenPool()
	if err != nil {
		return nil, errors.WithMessage(err, "could not retrieve connection from pool")
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.ErrorCtx(ctx, errors.WithMessage(err, "could not close neo4j connection"), nil)
		}
	}()

	query := fmt.Sprintf(getCodeListQuery, n.codeListLabel, code)

	rows, err := conn.QueryNeo(query, nil)
	if err != nil {
		return nil, errors.WithMessage(err, "could not run neo4j query")
	}

	var row []interface{}

	codeList := &models.CodeList{
		Links: models.CodeListLink{
			Self: models.Link{
				Href: fmt.Sprintf("/code-lists/%s", code),
				ID:   code,
			},
			Codes: models.Link{
				Href: fmt.Sprintf("/code-lists/%s/editions", code),
			},
		},
	}

	latest := int64(0)
	for row, _, err = rows.NextNeo(); err == nil; row, _, err = rows.NextNeo() {
		props := row[0].(graph.Node).Properties

		edition := props["year"].(int64)
		name := props["label"].(string)

		codeList.Name = name

		if edition > latest {
			latest = edition
		}
	}

	// If the latest value is still 0, then there is no value for this code
	if latest == 0 {
		return nil, datastore.NOT_FOUND
	}

	codeList.Links.Latest = models.Link{
		Href: fmt.Sprintf("/code-lists/%s/editions/%d", code, latest),
		ID:   fmt.Sprintf("%d", latest),
	}

	return codeList, nil
}

func (n NeoDataStore) GetEditions(ctx context.Context, codeListID string) (*models.Editions, error) {
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

		edition := fmt.Sprintf("%d", props["year"].(int64))
		code := props["code"].(string)

		editionModel := models.Edition{
			ID:          code,
			Edition:     edition,
			Label:       props["label"].(string),
			ReleaseDate: props["last_updated"].(string),
			Links: models.EditionLinks{
				Self: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s", code, edition),
					ID:   edition,
				},
				Editions: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions", code),
				},
				Codes: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes", code, edition),
				},
			},
		}

		editions.Items = append(editions.Items, editionModel)
	}

	editions.NumberOfResults = len(editions.Items)

	return editions, nil
}

func (n NeoDataStore) GetEdition(ctx context.Context, codeListID, edition string) (*models.Edition, error) {
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

		edition := fmt.Sprintf("%d", props["year"].(int64))
		code := props["code"].(string)

		editionModel = &models.Edition{
			ID:          code,
			Edition:     edition,
			Label:       props["label"].(string),
			ReleaseDate: props["last_updated"].(string),
			Links: models.EditionLinks{
				Self: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s", code, edition),
					ID:   edition,
				},
				Editions: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions", code),
				},
				Codes: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes", code, edition),
				},
			},
		}
	}

	return editionModel, nil
}
