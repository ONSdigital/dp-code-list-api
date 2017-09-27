package config

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"time"
)

func TestGetRetrunsDefaultValues(t *testing.T) {
	t.Parallel()
	Convey("When a loading a configuration, default values are return", t, func() {
		configuration, error := Get()
		So(error, ShouldBeNil)
		So(configuration.BindAddr, ShouldEqual, ":22400")
		So(configuration.MongoDBURL, ShouldEqual, "localhost:27017")
		So(configuration.GracefulShutdownTimeout, ShouldEqual, time.Second*5)
	})
}
