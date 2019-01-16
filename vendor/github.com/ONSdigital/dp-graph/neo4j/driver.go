package neo4j

import (
	"context"
	"io"

	"github.com/ONSdigital/dp-graph/graph/driver"
	"github.com/ONSdigital/dp-graph/neo4j/mapper"
	"github.com/ONSdigital/go-ns/log"
	bolt "github.com/ONSdigital/golang-neo4j-bolt-driver"
	"github.com/pkg/errors"
)

type Neo4j struct {
	pool bolt.ClosableDriverPool
	//	apiHost    string
	//	mapper *mapper.Mapper
}

func New(dbAddr string, size int) (d *Neo4j, err error) {
	pool, err := bolt.NewClosableDriverPool(dbAddr, size)
	if err != nil {
		log.Error(err, nil)
		return nil, err
	}

	return &Neo4j{
		pool: pool,
	}, nil
}

func (n *Neo4j) Close(ctx context.Context) error {
	return n.pool.Close()
}

func (n *Neo4j) exec(query string, mapp mapper.ResultMapper, single bool) error {
	c, err := n.pool.OpenPool()
	if err != nil {
		return err
	}
	defer c.Close()

	rows, err := c.QueryNeo(query, nil)
	if err != nil {
		return errors.WithMessage(err, "error executing neo4j query")
	}
	defer rows.Close()

	index := 0
	numOfResults := 0
results:
	for {
		data, meta, nextNeoErr := rows.NextNeo()
		if nextNeoErr != nil {
			if nextNeoErr != io.EOF {
				return errors.WithMessage(nextNeoErr, "extractResults: rows.NextNeo() return unexpected error")
			}
			break results
		}

		numOfResults++
		if single && index > 0 {
			return errors.WithMessage(err, "non unique results")
		}

		if mapp != nil {
			if err := mapp(&mapper.Result{Data: data, Meta: meta, Index: index}); err != nil {
				return errors.WithMessage(err, "mapResult returned an error")
			}
		}
		index++
	}

	if numOfResults == 0 {
		return driver.ErrNotFound
	}

	return nil
}
