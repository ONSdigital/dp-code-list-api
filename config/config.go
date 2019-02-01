package config

import (
	"time"

	"github.com/kelseyhightower/envconfig"
)

type Configuration struct {
	BindAddr                string        `envconfig:"BIND_ADDR"`
	CodeListAPIURL          string        `envconfig:"CODE_LIST_API_URL"`
	DatasetAPIURL           string        `envconfig:"DATASET_API_URL"`
	GracefulShutdownTimeout time.Duration `envconfig:"GRACEFUL_SHUTDOWN_TIMEOUT"`
	HealthCheckInterval     time.Duration `envconfig:"HEALTHCHECK_INTERVAL"`
}

var cfg *Configuration

// Get configures the application and returns the configuration
func Get() (*Configuration, error) {
	if cfg != nil {
		return cfg, nil
	}

	cfg = &Configuration{
		BindAddr:                ":22400",
		CodeListAPIURL:          "http://localhost:22400",
		DatasetAPIURL:           "http://localhost:22000",
		GracefulShutdownTimeout: time.Second * 5,
		HealthCheckInterval:     30 * time.Second,
	}

	return cfg, envconfig.Process("", cfg)
}
