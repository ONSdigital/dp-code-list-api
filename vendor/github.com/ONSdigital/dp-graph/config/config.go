package config

import (
	"time"

	"github.com/ONSdigital/dp-graph/graph/driver"
	"github.com/ONSdigital/dp-graph/neo4j"
	"github.com/kelseyhightower/envconfig"
	// "dp-graph/gremgo"
	// "dp-graph/mock"
)

type Configuration struct {
	DriverChoice            string        `envconfig:"GRAPH_DRIVER"` // - driver has to be chosen by the import in the using pacakge
	GracefulShutdownTimeout time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	DatabaseAddress         string        `envconfig:"GRAPH_ADDR"`
	PoolSize                int           `envconfig:"GRAPH_POOL_SIZE"`

	//	CodeListLabel      string        `envconfig:"CODE_LIST_LABEL"`
	Driver driver.Driver
}

var cfg *Configuration

// Get configures the application and returns the configuration
func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		GracefulShutdownTimeout: time.Second * 5,
		DatabaseAddress:         "bolt://localhost:7687",
		PoolSize:                30,

		//	CodeListLabel:      "code_list",
	}

	err := envconfig.Process("", cfg)

	var d driver.Driver

	switch cfg.DriverChoice {
	case "neo4j":
		d, err = neo4j.New(cfg.DatabaseAddress, cfg.PoolSize)
		if err != nil {
			panic(err)
		}
		//
		// case "gremgo":
		// 	d = gremgo.Driver{}
		// default:
		// 	d = mock.Driver{}
	}

	cfg.Driver = d

	return cfg, nil
}
