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

type NeoDataStore struct {
	pool DBPool
}

type DBPool interface {
	OpenPool() (bolt.Conn, error)
	Close() error
}

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
