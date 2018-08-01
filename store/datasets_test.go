package store

import (
	"context"
	"io"
	"testing"

	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/dp-code-list-api/store/mock"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/pkg/errors"
	. "github.com/smartystreets/goconvey/convey"
)

func TestCodeDatasets(t *testing.T) {
	Convey("given a code with related datasets exists in neo4j", t, func() {
		index := 0
		rows := &mock.RowsMock{
			NextNeoFunc: func() ([]interface{}, map[string]interface{}, error) {
				index++
				if index == 1 {
					return []interface{}{int64(1)}, nil, nil
				}
				if index == 2 {
					return nil, nil, io.EOF
				}
				if index == 3 {
					return []interface{}{
						graph.Node{
							Properties: map[string]interface{}{
								"dataset_id":   "cpih01",
								"edition":      "time-series",
								"version":      int64(1),
								"is_published": true,
							},
						},
						graph.Relationship{
							Properties: map[string]interface{}{
								"label": "Overall index",
							},
						},
					}, nil, nil
				}
				if index == 4 {
					return []interface{}{
						graph.Node{
							Properties: map[string]interface{}{
								"dataset_id":   "cpih01",
								"edition":      "time-series",
								"version":      int64(3),
								"is_published": true,
							},
						},
						graph.Relationship{
							Properties: map[string]interface{}{
								"label": "Overall index",
							},
						},
					}, nil, nil
				}
				if index == 5 {
					return []interface{}{
						graph.Node{
							Properties: map[string]interface{}{
								"dataset_id":   "mid-year-pop-est",
								"edition":      "time-series",
								"version":      int64(2),
								"is_published": true,
							},
						},
						graph.Relationship{
							Properties: map[string]interface{}{
								"label": "Overall index 2",
							},
						},
					}, nil, nil
				}
				return nil, nil, io.EOF
			},
			CloseFunc: closeNoErr,
		}
		conn := &mock.ConnMock{
			CloseFunc: closeNoErr,
			QueryNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
				return rows, nil
			},
		}
		pool := &mock.DBPoolMock{
			OpenPoolFunc: func() (golangNeo4jBoltDriver.Conn, error) {
				return conn, nil
			},
		}
		neo := NeoDataStore{
			pool:          pool,
			codeListLabel: "",
		}

		Convey("when I request the datasets from the store", func() {
			datasets, err := neo.GetCodeDatasets(context.Background(), "my-code-list-id", "2017", "my-code")
			Convey("then the datasets are returned correctly", func() {
				So(err, ShouldBeNil)
				So(datasets.Count, ShouldEqual, 2)

				dataset1 := datasets.Items[0]
				So(dataset1.DimensionLabel, ShouldEqual, "Overall index")
				So(dataset1.Links.Self.ID, ShouldEqual, "cpih01")
				So(dataset1.Links.Self.Href, ShouldEqual, "/datasets/cpih01")
				So(dataset1.Editions[0].Links.Self.ID, ShouldEqual, "time-series")
				So(dataset1.Editions[0].Links.Self.Href, ShouldEqual, "/datasets/cpih01/editions/time-series")
				So(dataset1.Editions[0].Links.LatestVersion.ID, ShouldEqual, "3")
				So(dataset1.Editions[0].Links.LatestVersion.Href, ShouldEqual, "/datasets/cpih01/editions/time-series/versions/3")
				So(dataset1.Editions[0].Links.DatasetDimension.ID, ShouldEqual, "my-code-list-id")
				So(dataset1.Editions[0].Links.DatasetDimension.Href, ShouldEqual, "/datasets/cpih01/editions/time-series/versions/3/dimensions/my-code-list-id")

				dataset2 := datasets.Items[1]
				So(dataset2.DimensionLabel, ShouldEqual, "Overall index 2")
				So(dataset2.Links.Self.ID, ShouldEqual, "mid-year-pop-est")
				So(dataset2.Links.Self.Href, ShouldEqual, "/datasets/mid-year-pop-est")
				So(dataset2.Editions[0].Links.Self.ID, ShouldEqual, "time-series")
				So(dataset2.Editions[0].Links.Self.Href, ShouldEqual, "/datasets/mid-year-pop-est/editions/time-series")
				So(dataset2.Editions[0].Links.LatestVersion.ID, ShouldEqual, "2")
				So(dataset2.Editions[0].Links.LatestVersion.Href, ShouldEqual, "/datasets/mid-year-pop-est/editions/time-series/versions/2")
				So(dataset2.Editions[0].Links.DatasetDimension.ID, ShouldEqual, "my-code-list-id")
				So(dataset2.Editions[0].Links.DatasetDimension.Href, ShouldEqual, "/datasets/mid-year-pop-est/editions/time-series/versions/2/dimensions/my-code-list-id")
			})
		})
	})

	Convey("given a request for code datasets is made for an edition that does not exist", t, func() {
		rows := &mock.RowsMock{
			NextNeoFunc: func() ([]interface{}, map[string]interface{}, error) {
				return nil, nil, io.EOF
			},
			CloseFunc: closeNoErr,
		}
		conn := &mock.ConnMock{
			CloseFunc: closeNoErr,
			QueryNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
				return rows, nil
			},
		}
		pool := &mock.DBPoolMock{
			OpenPoolFunc: func() (golangNeo4jBoltDriver.Conn, error) {
				return conn, nil
			},
		}
		neo := NeoDataStore{
			pool:          pool,
			codeListLabel: "",
		}
		Convey("when the request is made", func() {
			datasets, err := neo.GetCodeDatasets(context.Background(), "", "", "")
			Convey("then an error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err, ShouldEqual, datastore.ErrEditionNotFound)
				So(datasets, ShouldBeNil)
			})
		})
	})

	Convey("given a request for code datasets is made for an edition that has a bad store", t, func() {
		rows := &mock.RowsMock{
			NextNeoFunc: func() ([]interface{}, map[string]interface{}, error) {
				return nil, nil, testErr
			},
			CloseFunc: closeNoErr,
		}
		conn := &mock.ConnMock{
			CloseFunc: closeNoErr,
			QueryNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
				return rows, nil
			},
		}
		pool := &mock.DBPoolMock{
			OpenPoolFunc: func() (golangNeo4jBoltDriver.Conn, error) {
				return conn, nil
			},
		}
		neo := NeoDataStore{
			pool:          pool,
			codeListLabel: "",
		}
		Convey("when the request is made", func() {
			datasets, err := neo.GetCodeDatasets(context.Background(), "", "", "")
			Convey("then an error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "extractRowResults: rows.NextNeo return an error: error")
				So(datasets, ShouldBeNil)
			})
		})
	})

	Convey("given a request for code datasets is made, but a connection cannot be aquired", t, func() {
		index := 0
		rows := &mock.RowsMock{
			NextNeoFunc: func() ([]interface{}, map[string]interface{}, error) {
				index++
				if index == 1 {
					return []interface{}{int64(1)}, nil, nil
				}
				if index == 2 {
					return nil, nil, io.EOF
				}
				return nil, nil, io.EOF
			},
			CloseFunc: closeNoErr,
		}
		conn := &mock.ConnMock{
			CloseFunc: closeNoErr,
			QueryNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
				return rows, nil
			},
		}
		poolIndex := 0
		pool := &mock.DBPoolMock{
			OpenPoolFunc: func() (golangNeo4jBoltDriver.Conn, error) {
				if poolIndex == 0 {
					poolIndex++
					return conn, nil
				}
				return conn, errors.New("cannot open pool")
			},
		}
		neo := NeoDataStore{
			pool:          pool,
			codeListLabel: "",
		}
		Convey("when the request is made", func() {
			datasets, err := neo.GetCodeDatasets(context.Background(), "", "", "")
			Convey("then an error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "GetCodeDatasets: error opening neo4j connection: cannot open pool")
				So(datasets, ShouldBeNil)
			})
		})
	})

	Convey("given a request for code datasets is made, but a neo query cannot be executed", t, func() {
		index := 0
		rows := &mock.RowsMock{
			NextNeoFunc: func() ([]interface{}, map[string]interface{}, error) {
				index++
				if index == 1 {
					return []interface{}{int64(1)}, nil, nil
				}
				if index == 2 {
					return nil, nil, io.EOF
				}
				return nil, nil, io.EOF
			},
			CloseFunc: closeNoErr,
		}
		queryIndex := 0
		conn := &mock.ConnMock{
			CloseFunc: closeNoErr,
			QueryNeoFunc: func(query string, params map[string]interface{}) (golangNeo4jBoltDriver.Rows, error) {
				if queryIndex == 0 {
					queryIndex++
					return rows, nil
				}
				return rows, errors.New("cannot execute query")
			},
		}
		pool := &mock.DBPoolMock{
			OpenPoolFunc: func() (golangNeo4jBoltDriver.Conn, error) {
				return conn, nil
			},
		}
		neo := NeoDataStore{
			pool:          pool,
			codeListLabel: "",
		}
		Convey("when the request is made", func() {
			datasets, err := neo.GetCodeDatasets(context.Background(), "", "", "")
			Convey("then an error is returned", func() {
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldEqual, "getCode: conn.QueryNeo returned an error: cannot execute query")
				So(datasets, ShouldBeNil)
			})
		})
	})

}
