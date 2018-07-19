package store

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ONSdigital/dp-bolt/boltmock"
	"github.com/ONSdigital/dp-bolt/bolt"
	"context"
	"fmt"
	"github.com/ONSdigital/dp-code-list-api/datastore"
)

func TestNeoDataStore_GetCodeDatasetsEditionNotFound(t *testing.T) {
	Convey("should return expected error if no addition exists", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult bolt.ResultMapper) error {
					return mapResult(&bolt.Result{Data: []interface{}{int64(0)}})
				},
			},
		}

		neo := NeoDataStore{codeListLabel: codeListLabel, bolt: db}

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

		neo := NeoDataStore{codeListLabel: codeListLabel, bolt: db}

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

		neo := NeoDataStore{codeListLabel: codeListLabel, bolt: db}

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

		neo := NeoDataStore{codeListLabel: codeListLabel, bolt: db}

		datasets, err := neo.GetCodeDatasets(context.Background(), testCodeListID, testEdition, testCode)

		So(datasets, ShouldBeNil)
		So(err, ShouldNotBeNil)
		So(err, ShouldEqual, datastore.NOT_FOUND)
		So(db.QueryForResultCalls, ShouldHaveLength, 1)
		So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(countEditions, codeListLabel, testCodeListID, testEdition))
		So(db.QueryForResultCalls[0].Params, ShouldBeNil)
		So(db.QueryForResultsCalls, ShouldHaveLength, 1)
	})
}
