package mapper

import (
	"testing"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"fmt"
)

func TestGetCount(t *testing.T) {
	Convey("given dpbolt.Result.Data contains an int64 value", t, func() {
		r := &dpbolt.Result{Data: []interface{}{int64(666)}}

		Convey("when count is called", func() {
			count, mapper := GetCount()
			err := mapper(r)

			Convey("then the expected count is set and err is nil", func() {
				So(err, ShouldBeNil)
				So(*count, ShouldEqual, int64(666))
			})
		})
	})

	Convey("given dpbolt.Result.Data is empty", t, func() {
		r := &dpbolt.Result{Data: []interface{}{}}

		Convey("when count is called", func() {
			count, mapper := GetCount()
			err := mapper(r)

			Convey("then the expected err is returned", func() {
				So(err.Error(), ShouldResemble, "get count error: expecting single result value but 0 returned")
				So(*count, ShouldEqual, int64(0))
			})
		})
	})

	Convey("given dpbolt.Result.Data[0] is not type int64", t, func() {
		r := &dpbolt.Result{Data: []interface{}{"I AM NOT int64"}}

		Convey("when count is called", func() {
			count, mapper := GetCount()
			err := mapper(r)

			Convey("then the expected err is returned", func() {
				So(err.Error(), ShouldResemble, "failed to cast value to requested type, expected \"int64\" but was type \"string\"")
				So(*count, ShouldEqual, int64(0))
			})
		})
	})
}

func TestGetNode(t *testing.T) {
	Convey("given an input of type graph.Node", t, func() {
		expected := graph.Node{
			NodeIdentity: 666,
			Properties: map[string]interface{}{
				"key": "value",
			},
		}
		var input interface{} = expected

		Convey("when getNode is called", func() {
			node, err := getNode(input)

			Convey("then the expected node value is returned", func() {
				So(err, ShouldBeNil)
				So(node, ShouldResemble, expected)
			})
		})
	})

	Convey("given the input type is not graph.Node", t, func() {
		var input interface{} = []string{"Hello", "world!"}

		Convey("when getNode is called", func() {
			_, err := getNode(input)

			Convey("then the expected error value is returned", func() {
				So(err.Error(), ShouldEqual, "failed to cast value to requested type, expected \"graph.Node\" but was type \"[]string\"")
			})
		})
	})

	Convey("given input is nil", t, func() {
		Convey("when getNode is called", func() {
			_, err := getNode(nil)

			Convey("then the expected error value is returned", func() {
				So(err, ShouldEqual, ErrInputNil)
			})
		})
	})
}

func TestGetRelationship(t *testing.T) {
	Convey("given input is nil", t, func() {
		Convey("when getRelationship is called", func() {
			_, err := getRelationship(nil)

			Convey("then the expected error value is returned", func() {
				So(err, ShouldEqual, ErrInputNil)
			})
		})
	})

	Convey("given input is not type graph.Relationship", t, func() {
		Convey("when getRelationship is called", func() {
			_, err := getRelationship("I AM THE OWL")

			Convey("then the expected error value is returned", func() {
				So(err.Error(), ShouldEqual, "failed to cast value to requested type, expected \"graph.Relationship\" but was type \"string\"")
			})
		})
	})

	Convey("given input is type graph.Relationship", t, func() {
		Convey("when getRelationship is called", func() {
			r := graph.Relationship{
				Properties: map[string]interface{}{
					"key": "value",
				},
				RelIdentity: 666,
			}

			var in interface{} = r
			out, err := getRelationship(in)

			Convey("then the expected values are returned", func() {
				So(out, ShouldResemble, r)
				So(err, ShouldBeNil)
			})
		})
	})
}

func TestGetStringProperty(t *testing.T) {
	Convey("should return expected error if properties is nil", t, func() {
		val, err := getStringProperty("", nil)
		So(err, ShouldEqual, ErrInputNil)
		So(val, ShouldEqual, "")
	})

	Convey("should return empty string if properties does not contain key value", t, func() {
		val, err := getStringProperty("xxx", map[string]interface{}{"aaa": "aaa"})
		So(err, ShouldBeNil)
		So(val, ShouldBeEmpty)
	})

	Convey("should return casting error value is not type string", t, func() {
		val, err := getStringProperty("a", map[string]interface{}{"a": []string{"not", "a", "string"}})
		So(err.Error(), ShouldEqual, fmt.Sprintf("failed to cast value to requested type, expected %q but was type %q", "string", "[]string"))
		So(val, ShouldBeEmpty)
	})

	Convey("should return expected string value if it exists in properties", t, func() {
		val, err := getStringProperty("a", map[string]interface{}{"a": "aaa"})
		So(err, ShouldBeNil)
		So(val, ShouldEqual, "aaa")
	})
}
