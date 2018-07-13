package store

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ONSdigital/dp-bolt/boltmock"
	"errors"
	"context"
	"fmt"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/ONSdigital/dp-code-list-api/models"
	"strconv"
)

var (
	testCodeListID = "666" // the number of the best \m/
	testEdition    = "2018"
	testCode       = "99"
	errTest        = errors.New("error happened yo")
)

func TestNeoDataStore_GetCodeEditionExistsErrors(t *testing.T) {
	Convey("given edition exists return an error", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					return errTest
				},
			},
		}

		neoStore := NeoDataStore{db: db}
		code, err := neoStore.GetCode(context.Background(), testCodeListID, testEdition, testCode)

		Convey("then the expected error is returned", func() {
			So(code, ShouldBeNil)
			So(err, ShouldEqual, errTest)

		})

		Convey("and db.QueryForResult is called the expected number of times", func() {
			So(db.QueryForResultCalls, ShouldHaveLength, 1)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		})

	})

	Convey("given edition exists returns an unexpected expected type", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					// call the mapResult passing in the a Result that will cause a int64 cast error.
					return mapResult(
						&dpbolt.Result{
							Data:  []interface{}{"hello"}, // expected type is int64 return string to force cast error.
							Meta:  nil,
							Index: 0,
						})
				},
			},
		}

		neoStore := NeoDataStore{db: db}

		code, err := neoStore.GetCode(context.Background(), testCodeListID, testEdition, testCode)

		Convey("then the expected error is returned", func() {
			So(code, ShouldBeNil)
			So(err.Error(), ShouldEqual, "extract row result error: failed to cast result to int64")
		})

		Convey("and db.QueryForResult is called the expected number of times", func() {
			So(db.QueryForResultCalls, ShouldHaveLength, 1)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		})
	})

	Convey("given edition exists returns false", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					// call the mapResult passing in the a Result that will cause a int64 cast error.
					return mapResult(
						&dpbolt.Result{
							Data:  []interface{}{int64(0)},
							Meta:  nil,
							Index: 0,
						})
				},
			},
		}

		neoStore := NeoDataStore{db: db}

		code, err := neoStore.GetCode(context.Background(), testCodeListID, testEdition, testCode)

		Convey("then the expected error is returned", func() {
			So(code, ShouldBeNil)
			So(err, ShouldEqual, datastore.ErrEditionNotFound)
		})

		Convey("and db.QueryForResult is called the expected number of times", func() {
			So(db.QueryForResultCalls, ShouldHaveLength, 1)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		})
	})

	Convey("given more than one edition exists", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					// call the mapResult passing in the a Result that will cause a int64 cast error.
					return mapResult(
						&dpbolt.Result{
							Data:  []interface{}{int64(2)},
							Meta:  nil,
							Index: 0,
						})
				},
			},
		}

		neoStore := NeoDataStore{db: db}

		code, err := neoStore.GetCode(context.Background(), testCodeListID, testEdition, testCode)

		Convey("then the expected error is returned", func() {
			So(code, ShouldBeNil)
			So(err.Error(), ShouldResemble, "editionExists: multiple editions found")
		})

		Convey("and db.QueryForResult is called the expected number of times", func() {
			So(db.QueryForResultCalls, ShouldHaveLength, 1)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		})
	})
}

func TestNeoDataStore_GetCodeErrors(t *testing.T) {
	Convey("given query for code return an error", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					return mapResult(
						&dpbolt.Result{
							Data:  []interface{}{int64(1)},
							Meta:  nil,
							Index: 0,
						})
				},
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					return errTest
				},
			},
		}

		neoStore := NeoDataStore{db: db}
		code, err := neoStore.GetCode(context.Background(), testCodeListID, testEdition, testCode)

		Convey("then the expected error is returned", func() {
			So(code, ShouldBeNil)
			So(err.Error(), ShouldEqual, "extract result returned an error: error happened yo")

		})

		Convey("and db.QueryForResult is called the expected number of times", func() {
			So(db.QueryForResultCalls, ShouldHaveLength, 2)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)
			So(db.QueryForResultCalls[1].Query, ShouldEqual, fmt.Sprintf(getCodeQuery, testCodeListID, testEdition, testCode))
			So(db.QueryForResultCalls[1].Params, ShouldBeNil)
		})
	})
}

func TestNeoDataStore_GetCodeNotFound(t *testing.T) {
	Convey("given no code is found", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					return mapResult(
						&dpbolt.Result{
							Data:  []interface{}{int64(1)},
							Meta:  nil,
							Index: 0,
						})
				},
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					return mapResult(
						&dpbolt.Result{
							Data: []interface{}{},
						})
				},
			},
		}

		neoStore := NeoDataStore{db: db}
		code, err := neoStore.GetCode(context.Background(), testCodeListID, testEdition, testCode)

		Convey("then the expected error is returned", func() {
			So(code, ShouldBeNil)
			So(err.Error(), ShouldEqual, datastore.ErrCodeNotFound.Error())

		})

		Convey("and db.QueryForResult is called the expected number of times", func() {
			So(db.QueryForResultCalls, ShouldHaveLength, 2)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)
			So(db.QueryForResultCalls[1].Query, ShouldEqual, fmt.Sprintf(getCodeQuery, testCodeListID, testEdition, testCode))
			So(db.QueryForResultCalls[1].Params, ShouldBeNil)
		})
	})
}

func TestNeoDataStore_GetCodeSuccess(t *testing.T) {
	node := graph.Node{
		NodeIdentity: testNodeIdentity,
		Properties:   map[string]interface{}{"value": testNodeValue},
	}

	rel := graph.Relationship{
		Properties: map[string]interface{}{"label": testRelationshipLabel},
	}
	Convey("given getCode is successful", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					return mapResult(
						&dpbolt.Result{
							Data:  []interface{}{int64(1)},
							Meta:  nil,
							Index: 0,
						})
				},
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) error {
					return mapResult(&dpbolt.Result{
						Data: []interface{}{node, rel},
					})
				},
			},
		}

		neoStore := NeoDataStore{db: db}
		code, err := neoStore.GetCode(context.Background(), testCodeListID, testEdition, testCode)

		Convey("then the expected code is return and error is nil", func() {
			So(err, ShouldBeNil)
			expected := &models.Code{
				ID:    strconv.FormatInt(testNodeIdentity, 10),
				Code:  testNodeValue,
				Label: testRelationshipLabel,
				Links: models.CodeLinks{
					Self: models.Link{
						Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes/%s", testCodeListID, testEdition, testNodeValue),
					},
					CodeList: models.Link{
						Href: fmt.Sprintf("/code-lists/%s", testCodeListID),
					},
				},
			}
			So(code, ShouldResemble, expected)

		})

		Convey("and db.QueryForResult is called the expected number of times", func() {
			So(db.QueryForResultCalls, ShouldHaveLength, 2)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)
			So(db.QueryForResultCalls[1].Query, ShouldEqual, fmt.Sprintf(getCodeQuery, testCodeListID, testEdition, testCode))
			So(db.QueryForResultCalls[1].Params, ShouldBeNil)
		})
	})
}
