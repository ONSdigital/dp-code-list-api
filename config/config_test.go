package config

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetReturnsDefaultValues(t *testing.T) {
	t.Parallel()
	Convey("When a loading a configuration, default values are return", t, func() {
		cfg, err := Get()
		So(err, ShouldBeNil)
		So(cfg, ShouldResemble, &Configuration{
			BindAddr:                   ":22400",
			CodeListAPIURL:             "http://localhost:22400",
			DatasetAPIURL:              "http://localhost:22000",
			GracefulShutdownTimeout:    time.Second * 5,
			HealthCheckInterval:        time.Second * 30,
			HealthCheckCriticalTimeout: time.Second * 90,
			DefaultOffset:              0,
			DefaultLimit:               20,
			DefaultMaxLimit:            1000,
		})
	})
}
