package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	BindAddr                string        `envconfig:"BIND_ADDR"`
	GracefulShutdownTimeout time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	Neo4jDatabaseAddress    string        `envconfig:"NEO4J_ADDR"`
	Neo4jPoolSize           int           `envconfig:"NEO4J_POOL_SIZE"`
	Neo4jCodeListLabel      string        `envconfig:"NEO4J_CODE_LIST_LABEL"`
}

var cfg *Configuration

// Get configures the application and returns the configuration
func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		BindAddr:                ":22400",
		GracefulShutdownTimeout: time.Second * 5,
		Neo4jDatabaseAddress:    "bolt://localhost:7687",
		Neo4jPoolSize:           30,
		Neo4jCodeListLabel:      "code_list",
	}

	return cfg, envconfig.Process("", cfg)
}
