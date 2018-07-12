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
			QueryForResultFuncs: []boltmock.QueryForResultFunc{
				func(query string, params map[string]interface{}, resultExtractor dpbolt.ResultExtractor) error {
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
			QueryForResultFuncs: []boltmock.QueryForResultFunc{
				func(query string, params map[string]interface{}, resultExtractor dpbolt.ResultExtractor) error {
					// call the resultExtractor passing in the a Result that will cause a int64 cast error.
					return resultExtractor(
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
			QueryForResultFuncs: []boltmock.QueryForResultFunc{
				func(query string, params map[string]interface{}, resultExtractor dpbolt.ResultExtractor) error {
					// call the resultExtractor passing in the a Result that will cause a int64 cast error.
					return resultExtractor(
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
			QueryForResultFuncs: []boltmock.QueryForResultFunc{
				func(query string, params map[string]interface{}, resultExtractor dpbolt.ResultExtractor) error {
					// call the resultExtractor passing in the a Result that will cause a int64 cast error.
					return resultExtractor(
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
			QueryForResultFuncs: []boltmock.QueryForResultFunc{
				func(query string, params map[string]interface{}, resultExtractor dpbolt.ResultExtractor) error {
					return resultExtractor(
						&dpbolt.Result{
							Data:  []interface{}{int64(1)},
							Meta:  nil,
							Index: 0,
						})
				},
				func(query string, params map[string]interface{}, resultExtractor dpbolt.ResultExtractor) error {
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
