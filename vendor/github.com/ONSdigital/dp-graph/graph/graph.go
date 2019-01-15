package graph

import (
	"errors"

	"github.com/ONSdigital/dp-graph/config"
	"github.com/ONSdigital/dp-graph/graph/driver"
)

type DB struct {
	driver driver.Driver
}

func (db *DB) Close() error {
	return db.driver.Close()
}

func New() (*DB, error) {
	cfg, err := config.Get()
	if err != nil {
		return nil, err
	}

	if cfg.Driver == nil {
		return nil, errors.New("no graph driver configured - GRAPH_DRIVER")
	}

	return &DB{
		driver: cfg.Driver,
	}, nil
}
