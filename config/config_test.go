package config

import (
	"testing"
	"time"

	. "github.com/smartystreets/goconvey/convey"
)

func TestGetRetrunsDefaultValues(t *testing.T) {
	t.Parallel()
	Convey("When a loading a configuration, default values are return", t, func() {
		configuration, error := Get()
		So(error, ShouldBeNil)
		So(configuration.BindAddr, ShouldEqual, ":22400")
		So(configuration.MongoConfig.BindAddr, ShouldEqual, "localhost:27017")
		So(configuration.MongoConfig.Database, ShouldEqual, "codelists")
		So(configuration.MongoConfig.CodelistsCollection, ShouldEqual, "codelists")
		So(configuration.MongoConfig.CodesCollection, ShouldEqual, "codes")
		So(configuration.GracefulShutdownTimeout, ShouldEqual, time.Second*5)
	})
}
