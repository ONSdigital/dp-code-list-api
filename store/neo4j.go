package store

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/go-ns/log"
	bolt "github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/pkg/errors"
)

const (
	getCodeListsQuery = "MATCH (i) WHERE i:_%s%s RETURN distinct labels(i) as labels, i"
	getCodeListQuery  = "MATCH (i:_%s:`_name_%s`) RETURN i"
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
func (n *NeoDataStore) Close() error {
	return n.pool.Close()
}

// CreateNeoDataStore allows the creation of a NeoDataStore
func CreateNeoDataStore(addr, codelistLabel string, conns int) (n *NeoDataStore, err error) {
	store, err := bolt.NewClosableDriverPool(addr, conns)
	if err != nil {
		return
	}

	n = &NeoDataStore{
		pool:          store,
		codeListLabel: codelistLabel,
	}

	return
}

// GetCodeLists returns a list of code lists
func (n *NeoDataStore) GetCodeLists(ctx context.Context, filterBy string) (*models.CodeListResults, error) {
	logData := log.Data{}
	if len(filterBy) > 0 {
		logData = log.Data{"filter_by": filterBy}
	}
	log.InfoCtx(ctx, "about to query neo4j for code lists", logData)

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
		labels := row[1].(graph.Node).Labels

		name := props["label"].(string)
		neoLabel := strings.Replace(labels[1], `_name_`, "", -1)

		codeList := &models.CodeList{
			Links: models.CodeListLink{
				Self: &models.Link{
					Href: fmt.Sprintf("/code-lists/%s", neoLabel),
					ID:   neoLabel,
				},
				Codes: &models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions", neoLabel),
				},
			},
			Name: name,
		}

		codeLists.Items = append(codeLists.Items, *codeList)
	}

	return codeLists, nil
}

// GetCodeList returns an individual code list for a given code
func (n *NeoDataStore) GetCodeList(ctx context.Context, code string) (*models.CodeList, error) {
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

	var hasMultipleEditions bool

	var row []interface{}

	codeList := &models.CodeList{
		Links: models.CodeListLink{
			Self: &models.Link{
				Href: fmt.Sprintf("/code-lists/%s", code),
				ID:   code,
			},
			Codes: &models.Link{
				Href: fmt.Sprintf("/code-lists/%s/editions", code),
			},
		},
	}

	latest := 0
	count := 0
	var latestEdition string
	for row, _, err = rows.NextNeo(); err == nil; row, _, err = rows.NextNeo() {
		count++
		if count > 1 {
			hasMultipleEditions = true
		}

		props := row[0].(graph.Node).Properties
		edition := props["edition"].(string)

		if !hasMultipleEditions {
			latestEdition = edition
		} else {
			editionInt, err := strconv.Atoi(edition)
			if err != nil || editionInt > latest {
				latestEdition = edition
			}
		}

		name := props["label"].(string)

		codeList.Name = name
	}

	if count == 0 {
		return nil, datastore.NOT_FOUND
	}

	codeList.Links.Latest = &models.Link{
		Href: fmt.Sprintf("/code-lists/%s/editions/%s", code, latestEdition),
		ID:   latestEdition,
	}

	return codeList, nil
}
