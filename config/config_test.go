package config

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRetrunsDefaultValues(t *testing.T) {
	t.Parallel()
	Convey("When a loading a configuration, default values are return", t, func() {
		configuration, err := Get()
		So(err, ShouldBeNil)
		So(configuration.BindAddr, ShouldEqual, ":22400")
		So(configuration.GracefulShutdownTimeout, ShouldEqual, time.Second*5)
		So(configuration.Neo4jDatabaseAddress, ShouldEqual, "bolt://localhost:7687")
		So(configuration.Neo4jPoolSize, ShouldEqual, 30)
		So(configuration.Neo4jCodeListLabel, ShouldEqual, "code_list")
	})
}
