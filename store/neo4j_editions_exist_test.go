package store

import (
	"testing"
	"github.com/ONSdigital/dp-bolt/boltmock"
	. "github.com/smartystreets/goconvey/convey"
	"context"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"fmt"
)

func TestNeoDataStore_EditionExistsSuccess(t *testing.T) {
	Convey("given bolt.QueryForResult return count 1 for edition", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) (int, error) {
					return 1, mapResult(
						&dpbolt.Result{
							Data: []interface{}{int64(1)},
						})
				},
			},
		}

		store := &NeoDataStore{bolt: db}

		Convey("then EditionExists should return exists true and no error", func() {
			exists, err := store.EditionExists(context.Background(), testCodeListID, testEdition)
			So(err, ShouldBeNil)
			So(exists, ShouldBeTrue)
			So(db.QueryForResultCalls, ShouldHaveLength, 1)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		})
	})
}

func TestNeoDataStore_EditionExistsMoreThanOneResult(t *testing.T) {
	Convey("given bolt.QueryForResult return count > 1 for edition", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) (int, error) {
					return 1, mapResult(
						&dpbolt.Result{
							Data: []interface{}{int64(2)},
						})
				},
			},
		}

		store := &NeoDataStore{bolt: db}

		Convey("then EditionExists should return exists true and no error", func() {
			exists, err := store.EditionExists(context.Background(), testCodeListID, testEdition)
			So(err.Error(), ShouldEqual, "editionExists: multiple editions found")
			So(exists, ShouldBeFalse)
			So(db.QueryForResultCalls, ShouldHaveLength, 1)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		})
	})
}

func TestNeoDataStore_EditionExistsQueryForResultError(t *testing.T) {
	Convey("given bolt.QueryForResult returns an error", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult dpbolt.ResultMapper) (int, error) {
					return 0, errTest
				},
			},
		}

		store := &NeoDataStore{bolt: db}

		Convey("then EditionExists should return exists false and the expected error", func() {
			exists, err := store.EditionExists(context.Background(), testCodeListID, testEdition)
			So(err, ShouldEqual, errTest)
			So(exists, ShouldBeFalse)
			So(db.QueryForResultCalls, ShouldHaveLength, 1)
			So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, testCodeListID, testEdition))
			So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		})
	})
}
