package config

import (
	"github.com/kelseyhightower/envconfig"
	"time"
)

type Configuration struct {
	BindAddr                string        `envconfig:"BIND_ADDR"`
	MongoDBURL              string        `envconfig:"MONGODB_ADDR"`
	GracefulShutdownTimeout time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
}

var cfg *Configuration

// Get configures the application and returns the configuration
func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		BindAddr:                ":22400",
		MongoDBURL:              "localhost:27017",
		GracefulShutdownTimeout: time.Second * 5,
	}

	return cfg, envconfig.Process("", cfg)
}
