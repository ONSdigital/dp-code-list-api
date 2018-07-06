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
		So(cfg.BindAddr, ShouldEqual, ":22400")
		So(cfg.GracefulShutdownTimeout, ShouldEqual, time.Second*5)
		So(cfg.Neo4jDatabaseAddress, ShouldEqual, "bolt://localhost:7687")
		So(cfg.Neo4jPoolSize, ShouldEqual, 30)
		So(cfg.Neo4jCodeListLabel, ShouldEqual, "code_list")
	})
}
