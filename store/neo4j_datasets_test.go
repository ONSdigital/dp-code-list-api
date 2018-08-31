package store

import (
	"context"
	"fmt"
	"testing"

	"github.com/ONSdigital/dp-bolt/bolt"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-bolt/boltmock"
	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/dp-code-list-api/store/mapper"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNeoDataStore_GetCodeDatasets(t *testing.T) {
	Convey("given a code with related datasets exists in neo4j", t, func() {

		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					return mapResult(
						&dpbolt.Result{Data: []interface{}{int64(1)}},
					)
				},
			},
			QueryForResultsFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					mapResult(&dpbolt.Result{Data: []interface{}{
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
					}})
					mapResult(&dpbolt.Result{Data: []interface{}{
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
					}})
					mapResult(&dpbolt.Result{Data: []interface{}{
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
					}})
					return nil
				},
			},
		}

		neo := testStore
		neo.bolt = db

		Convey("when I request the datasets from the store", func() {
			datasets, err := neo.GetCodeDatasets(context.Background(), "my-code-list-id", "2017", "my-code")
			Convey("then the datasets are returned correctly", func() {
				So(err, ShouldBeNil)
				So(datasets.Count, ShouldEqual, 2)

				checkCPIH := func(dataset *models.Dataset) {
					So(dataset.DimensionLabel, ShouldEqual, "Overall index")
					So(dataset.Links.Self.ID, ShouldEqual, "cpih01")
					So(dataset.Links.Self.Href, ShouldEqual, "/datasets/cpih01")
					So(dataset.Editions[0].Links.Self.ID, ShouldEqual, "time-series")
					So(dataset.Editions[0].Links.Self.Href, ShouldEqual, "/datasets/cpih01/editions/time-series")
					So(dataset.Editions[0].Links.LatestVersion.ID, ShouldEqual, "3")
					So(dataset.Editions[0].Links.LatestVersion.Href, ShouldEqual, "/datasets/cpih01/editions/time-series/versions/3")
					So(dataset.Editions[0].Links.DatasetDimension.ID, ShouldEqual, "my-code-list-id")
					So(dataset.Editions[0].Links.DatasetDimension.Href, ShouldEqual, "/datasets/cpih01/editions/time-series/versions/3/dimensions/my-code-list-id")
				}
				checkMidYearPop := func(dataset *models.Dataset) {
					So(dataset.DimensionLabel, ShouldEqual, "Overall index 2")
					So(dataset.Links.Self.ID, ShouldEqual, "mid-year-pop-est")
					So(dataset.Links.Self.Href, ShouldEqual, "/datasets/mid-year-pop-est")
					So(dataset.Editions[0].Links.Self.ID, ShouldEqual, "time-series")
					So(dataset.Editions[0].Links.Self.Href, ShouldEqual, "/datasets/mid-year-pop-est/editions/time-series")
					So(dataset.Editions[0].Links.LatestVersion.ID, ShouldEqual, "2")
					So(dataset.Editions[0].Links.LatestVersion.Href, ShouldEqual, "/datasets/mid-year-pop-est/editions/time-series/versions/2")
					So(dataset.Editions[0].Links.DatasetDimension.ID, ShouldEqual, "my-code-list-id")
					So(dataset.Editions[0].Links.DatasetDimension.Href, ShouldEqual, "/datasets/mid-year-pop-est/editions/time-series/versions/2/dimensions/my-code-list-id")
				}

				for _, ds := range datasets.Items {
					switch {
					case ds.Links.Self.ID == "cpih01":
						checkCPIH(&ds)
					case ds.Links.Self.ID == "mid-year-pop-est":
						checkMidYearPop(&ds)
					default:
						//fail if any dataset does not match the two that were provided
						So(ds, ShouldBeNil)
					}
				}
			})
		})
	})
}

func TestNeoDataStore_GetCodeDatasetsEditionNotFound(t *testing.T) {
	Convey("should return expected error if no addition exists", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult bolt.ResultMapper) error {
					return mapResult(&bolt.Result{Data: []interface{}{int64(0)}})
				},
			},
		}

		neo := testStore
		neo.bolt = db

		datasets, err := neo.GetCodeDatasets(context.Background(), testCodeListID, testEdition, testCode)

		So(datasets, ShouldBeNil)
		So(err, ShouldNotBeNil)
		So(db.QueryForResultCalls, ShouldHaveLength, 1)
		So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, codeListLabel, testCodeListID, testEdition))
		So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		So(db.QueryForResultsCalls, ShouldHaveLength, 0)
	})
}

func TestNeoDataStore_GetCodeDatasetsEditionExitsError(t *testing.T) {
	Convey("should return expected error if EditionExists returns an error", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				boltmock.ErrQueryFunc,
			},
		}

		neo := NeoDataStore{codeListLabel: codeListLabel, bolt: db, mapper: &mapper.Mapper{Host: ""}}

		datasets, err := neo.GetCodeDatasets(context.Background(), testCodeListID, testEdition, testCode)

		So(datasets, ShouldBeNil)
		So(err, ShouldNotBeNil)
		So(err, ShouldEqual, boltmock.Err)
		So(db.QueryForResultCalls, ShouldHaveLength, 1)
		So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, codeListLabel, testCodeListID, testEdition))
		So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		So(db.QueryForResultsCalls, ShouldHaveLength, 0)
	})
}

func TestNeoDataStore_GetCodeDatasetsQueryForDatasetsError(t *testing.T) {
	Convey("should return expected error if bolt.QueryForResults errors", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult bolt.ResultMapper) error {
					return mapResult(&bolt.Result{Data: []interface{}{int64(1)}})
				},
			},

			QueryForResultsFuncs: []boltmock.QueryFunc{
				boltmock.ErrQueryFunc,
			},
		}

		neo := NeoDataStore{codeListLabel: codeListLabel, bolt: db, mapper: &mapper.Mapper{Host: ""}}

		datasets, err := neo.GetCodeDatasets(context.Background(), testCodeListID, testEdition, testCode)

		So(datasets, ShouldBeNil)
		So(err, ShouldNotBeNil)
		So(err, ShouldEqual, boltmock.Err)
		So(db.QueryForResultCalls, ShouldHaveLength, 1)
		So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, codeListLabel, testCodeListID, testEdition))
		So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		So(db.QueryForResultsCalls, ShouldHaveLength, 1)
	})
}

func TestNeoDataStore_GetCodeDatasetsNoResults(t *testing.T) {
	Convey("should return expected error if no datasets are found", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult bolt.ResultMapper) error {
					return mapResult(&bolt.Result{Data: []interface{}{int64(1)}})
				},
			},

			QueryForResultsFuncs: []boltmock.QueryFunc{
				boltmock.NoResultFunc,
			},
		}

		neo := NeoDataStore{codeListLabel: codeListLabel, bolt: db, mapper: &mapper.Mapper{Host: ""}}

		datasets, err := neo.GetCodeDatasets(context.Background(), testCodeListID, testEdition, testCode)

		So(datasets, ShouldBeNil)
		So(err, ShouldNotBeNil)
		So(err, ShouldEqual, datastore.ErrDatasetsNotFound)
		So(db.QueryForResultCalls, ShouldHaveLength, 1)
		So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, codeListLabel, testCodeListID, testEdition))
		So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		So(db.QueryForResultsCalls, ShouldHaveLength, 1)
	})
}
