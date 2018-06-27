package store

import (
	"context"
	"fmt"

	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/go-ns/log"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/pkg/errors"
)

const (
	getCodeListsQuery = "MATCH (i) where i:_code_list RETURN distinct labels(i) as labels, i"
	getCodeListQuery  = `MATCH (i:_code_list {code:"%s"}) RETURN i`
)

// NeoDataStore represents the necessary information to access
// neo4j
type NeoDataStore struct {
	pool DBPool
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
func CreateNeoDataStore(addr string, conns int) (n NeoDataStore, err error) {
	store, err := bolt.NewClosableDriverPool(addr, conns)
	if err != nil {
		return
	}

	n = NeoDataStore{
		pool: store,
	}

	return
}

// GetCodeLists returns a list of code lists
func (n NeoDataStore) GetCodeLists(ctx context.Context) (*models.CodeListResults, error) {
	conn, err := n.pool.OpenPool()
	if err != nil {
		return nil, errors.WithMessage(err, "could not retrieve connection from pool")
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.ErrorCtx(ctx, errors.WithMessage(err, "could not close neo4j connection"), nil)
		}
	}()

	rows, err := conn.QueryNeo(getCodeListsQuery, nil)
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
	conn, err := n.pool.OpenPool()
	if err != nil {
		return nil, errors.WithMessage(err, "could not retrieve connection from pool")
	}

	defer func() {
		if err := conn.Close(); err != nil {
			log.ErrorCtx(ctx, errors.WithMessage(err, "could not close neo4j connection"), nil)
		}
	}()

	query := fmt.Sprintf(getCodeListQuery, code)

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
		codeList.Name = props["label"].(string)

		if edition > latest {
			latest = edition
		}
	}

	codeList.Links.Latest = models.Link{
		Href: fmt.Sprintf("/code-lists/%s/editions/%d", code, latest),
		ID:   fmt.Sprintf("%d", latest),
	}

	return codeList, nil
}
