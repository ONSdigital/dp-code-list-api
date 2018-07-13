package store

import (
	"github.com/ONSdigital/dp-code-list-api/models"
	"fmt"
	"testing"
	"github.com/ONSdigital/dp-bolt/boltmock"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"context"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ONSdigital/dp-code-list-api/datastore"
)

var (
	code1 = models.Code{
		ID:    "666",
		Code:  "c1",
		Label: "label1",
		Links: models.CodeLinks{
			Self: models.Link{
				Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes/%s", testCodeListID, testEdition, "c1"),
			},
			CodeList: models.Link{
				Href: fmt.Sprintf("/code-lists/%s", testCodeListID),
			},
		},
	}

	code2 = models.Code{
		ID:    "667",
		Code:  "c2",
		Label: "label2",
		Links: models.CodeLinks{
			Self: models.Link{
				Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes/%s", testCodeListID, testEdition, "c2"),
			},
			CodeList: models.Link{
				Href: fmt.Sprintf("/code-lists/%s", testCodeListID),
			},
		},
	}

	codes = []models.Code{code1, code2}

	row1 = []interface{}{
		graph.Node{
			NodeIdentity: int64(666),
			Properties:   map[string]interface{}{"value": code1.Code},
		},
		graph.Relationship{
			Properties: map[string]interface{}{"label": code1.Label},
		},
	}

	row2 = []interface{}{
		graph.Node{
			NodeIdentity: int64(667),
			Properties:   map[string]interface{}{"value": code2.Code},
		},
		graph.Relationship{
			Properties: map[string]interface{}{"label": code2.Label},
		},
	}
)

func TestNeoDataStore_GetCodesSuccess(t *testing.T) {
	Convey("given get codes is successful", t, func() {
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
					mapResult(&dpbolt.Result{Data: row1})
					mapResult(&dpbolt.Result{Data: row2})
					return nil
				},
			},
		}
		store := NeoDataStore{db: db}

		result, err := store.GetCodes(context.Background(), testCodeListID, testEdition)

		Convey("then the expected codeResult is returned and error is nil", func() {
			So(err, ShouldBeNil)
			So(result, ShouldResemble, &models.CodeResults{Items: codes, Count: 2})
		})

		Convey("and the expected calls are made", func() {
			So(db.QueryForResultCalls, ShouldHaveLength, 1)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)

			So(db.QueryForResultsCalls, ShouldHaveLength, 1)
			So(db.QueryForResultsCalls[0].Query, ShouldEqual, fmt.Sprintf(getCodesQuery, testCodeListID, testEdition))
			So(db.QueryForResultsCalls[0].Params, ShouldBeNil)
		})
	})
}

func TestNeoDataStore_GetCodesEditionExitsError(t *testing.T) {
	Convey("given edition exits returns an error", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					return errTest
				},
			},
		}
		store := NeoDataStore{db: db}

		result, err := store.GetCodes(context.Background(), testCodeListID, testEdition)

		Convey("then codeResult is nil and the expected error is returned", func() {
			So(err, ShouldResemble, errTest)
			So(result, ShouldBeNil)
		})

		Convey("and the expected calls are made", func() {
			So(db.QueryForResultCalls, ShouldHaveLength, 1)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)

			So(db.QueryForResultsCalls, ShouldHaveLength, 0)
		})
	})
}

func TestNeoDataStore_GetCodesEditionExitsFalse(t *testing.T) {
	Convey("given edition exits returns false", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					return mapResult(
						&dpbolt.Result{Data: []interface{}{int64(0)}},
					)
				},
			},
		}
		store := NeoDataStore{db: db}

		result, err := store.GetCodes(context.Background(), testCodeListID, testEdition)

		Convey("then codeResult is nil and the expected error is returned", func() {
			So(err, ShouldResemble, datastore.ErrEditionNotFound)
			So(result, ShouldBeNil)
		})

		Convey("and the expected calls are made", func() {
			So(db.QueryForResultCalls, ShouldHaveLength, 1)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)

			So(db.QueryForResultsCalls, ShouldHaveLength, 0)
		})
	})
}

func TestNeoDataStore_GetCodesResultMapperError(t *testing.T) {
	Convey("given getCodes result mapper returns an error", t, func() {
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
					return errTest
				},
			},
		}
		store := NeoDataStore{db: db}

		result, err := store.GetCodes(context.Background(), testCodeListID, testEdition)

		Convey("then codeResult is nil and the expected error is returned", func() {
			So(err, ShouldResemble, errTest)
			So(result, ShouldBeNil)
		})

		Convey("and the expected calls are made", func() {
			So(db.QueryForResultCalls, ShouldHaveLength, 1)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)

			So(db.QueryForResultsCalls, ShouldHaveLength, 1)
			So(db.QueryForResultsCalls[0].Query, ShouldEqual, fmt.Sprintf(getCodesQuery, testCodeListID, testEdition))
			So(db.QueryForResultsCalls[0].Params, ShouldBeNil)
		})
	})
}
