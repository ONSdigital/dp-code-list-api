package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	BindAddr   string `envconfig:"BIND_ADDR"`
	MongoDBURL string `envconfig:"MONGODB_IMPORTS_ADDR"`
}

var cfg *Configuration

// Get - configures the application and returns the configuration
func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		BindAddr:   ":22400",
		MongoDBURL: "localhost:27017",
	}

	return cfg, envconfig.Process("", cfg)
}
