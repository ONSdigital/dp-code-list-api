package graph

import (
	"context"
	"errors"

	"github.com/ONSdigital/dp-graph/config"
	"github.com/ONSdigital/dp-graph/graph/driver"
	"github.com/ONSdigital/dp-graph/mock"
)

type DB struct {
	driver.Driver
}

func New(ctx context.Context) (*DB, error) {
	cfg, err := config.Get()
	if err != nil {
		return nil, err
	}

	if cfg.Driver == nil {
		return nil, errors.New("no graph driver configured - GRAPH_DRIVER")
	}

	return &DB{
		cfg.Driver,
	}, nil
}

//Test sets flags for managing responses from the Mock driver
func Test(backend, query, content bool) *mock.Mock {
	return &mock.Mock{
		IsBackendUnreachable: backend,
		IsQueryInvalid:       query,
		IsContentNotFound:    content,
	}
}
