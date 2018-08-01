package mapper

import (
	"fmt"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
	"strconv"
	"testing"
)

var (
	testCodeListID        = "666" // the number of the best \m/
	testEdition           = "2018"
	testCode              = "99"
	errTest               = errors.New("error happened yo")
	testNodeIdentity      = int64(666) // the number of the best \m/
	testNodeValue         = "node value"
	testRelationshipLabel = "relationship label"
)

func TestCodeResultExtractor_Success(t *testing.T) {
	Convey("given a valid result", t, func() {
		node := graph.Node{
			NodeIdentity: testNodeIdentity,
			Properties:   map[string]interface{}{"value": testNodeValue},
		}

		rel := graph.Relationship{
			Properties: map[string]interface{}{"label": testRelationshipLabel},
		}

		actual := &models.Code{}
		extractor := Code(actual, testCodeListID, testEdition)

		Convey("when extractor is called", func() {
			err := extractor(
				&dpbolt.Result{
					Data:  []interface{}{node, rel},
					Meta:  nil,
					Index: 0,
				})

			Convey("then no error is returned and code is populated with the expected values", func() {
				So(err, ShouldBeNil)

				expected := &models.Code{
					ID:    strconv.FormatInt(testNodeIdentity, 10),
					Code:  testNodeValue,
					Label: testRelationshipLabel,
					Links: models.CodeLinks{
						Self: models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes/%s", testCodeListID, testEdition, testNodeValue),
						},
						Datasets: models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes/%s/datasets", testCodeListID, testEdition, testNodeValue),
						},
						CodeList: models.Link{
							Href: fmt.Sprintf("/code-lists/%s", testCodeListID),
						},
					},
				}

				So(expected, ShouldResemble, actual)
			})
		})
	})
}

func TestCodeResultExtractor_BadTypes(t *testing.T) {
	node := graph.Node{
		Properties: map[string]interface{}{
			"value": testCode,
		},
	}

	Convey("given data.0 is not type graph.Node", t, func() {
		actual := &models.Code{}
		extractor := Code(actual, testCodeListID, testEdition)

		Convey("when extractor is called", func() {
			err := extractor(
				&dpbolt.Result{
					Data:  []interface{}{"not graph.Node"},
					Meta:  nil,
					Index: 0,
				})

			Convey("then expected error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "failed to cast value to requested type, expected \"graph.Node\" but was type \"string\"")
				So(actual.ID, ShouldBeBlank)
				So(actual.Code, ShouldBeBlank)
				So(actual.Label, ShouldBeBlank)
				So(actual.Links, ShouldResemble, models.CodeLinks{})
			})
		})
	})

	Convey("given node.Properties.value is not type string", t, func() {
		actual := &models.Code{}
		extractor := Code(actual, testCodeListID, testEdition)

		Convey("when extractor is called", func() {
			err := extractor(
				&dpbolt.Result{
					Data:  []interface{}{graph.Node{Properties: map[string]interface{}{"value": 666}}},
					Meta:  nil,
					Index: 0,
				})

			Convey("then expected error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "failed to cast value to requested type, expected \"string\" but was type \"int\"")
				So(actual.ID, ShouldBeBlank)
				So(actual.Code, ShouldBeBlank)
				So(actual.Label, ShouldBeBlank)
				So(actual.Links, ShouldResemble, models.CodeLinks{})
			})
		})
	})

	Convey("given data.1 is not type graph.Relationship", t, func() {
		actual := &models.Code{}
		extractor := Code(actual, testCodeListID, testEdition)

		Convey("when extractor is called", func() {
			err := extractor(
				&dpbolt.Result{
					Data:  []interface{}{node, "not graph.Relationship"},
					Meta:  nil,
					Index: 0,
				})

			Convey("then expected error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "failed to cast value to requested type, expected \"graph.Relationship\" but was type \"string\"")
				So(actual.ID, ShouldBeBlank)
				So(actual.Code, ShouldBeBlank)
				So(actual.Label, ShouldBeBlank)
				So(actual.Links, ShouldResemble, models.CodeLinks{})
			})
		})
	})

	Convey("given relationship.Properties.label is not type string", t, func() {
		actual := &models.Code{}
		extractor := Code(actual, testCodeListID, testEdition)

		Convey("when extractor is called", func() {
			err := extractor(
				&dpbolt.Result{
					Data:  []interface{}{node, graph.Relationship{Properties: map[string]interface{}{"label": 999}}},
					Meta:  nil,
					Index: 0,
				})

			Convey("then expected error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "failed to cast value to requested type, expected \"string\" but was type \"int\"")
				So(actual.ID, ShouldBeBlank)
				So(actual.Code, ShouldBeBlank)
				So(actual.Label, ShouldBeBlank)
				So(actual.Links, ShouldResemble, models.CodeLinks{})
			})
		})
	})
}
