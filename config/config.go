package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	BindAddr                string        `envconfig:"BIND_ADDR"`
	GracefulShutdownTimeout time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	MongoConfig             MongoConfig
}

// MongoConfig contains the config required to connect to MongoDB.
type MongoConfig struct {
	BindAddr            string `envconfig:"MONGODB_ADDR"`
	Database            string `envconfig:"MONGODB_DATABASE"`
	CodelistsCollection string `envconfig:"MONGODB_CODELISTS_COLLECTION"`
	CodesCollection     string `envconfig:"MONGODB_CODES_COLLECTION"`
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
		MongoConfig: MongoConfig{
			BindAddr:            "localhost:27017",
			Database:            "codelists",
			CodelistsCollection: "codelists",
			CodesCollection:     "codes",
		},
	}

	return cfg, envconfig.Process("", cfg)
}
