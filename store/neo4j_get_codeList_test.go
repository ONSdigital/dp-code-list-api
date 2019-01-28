package store

import (
	"context"
	"fmt"
	"testing"

	"github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-bolt/boltmock"
	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/dp-code-list-api/models"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNeoDataStore_GetCodeListSuccess(t *testing.T) {
	Convey("given bolt.QueryForResult returns no error", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult bolt.ResultMapper) error {
					return mapResult(&bolt.Result{
						//Data: []interface{}{graph.Node{Properties: map[string]interface{}{"label": "testLabel"}}},
						Data: []interface{}{int64(1)},
					})
				},
			},
		}

		store := testStore
		store.bolt = db

		Convey("when get code list is called", func() {
			cl, err := store.GetCodeList(context.Background(), testCodeListID)

			Convey("then the expected code list is returned", func() {
				So(err, ShouldBeNil)
				So(cl, ShouldResemble, &models.CodeList{
					Links: models.CodeListLink{
						Self: &models.Link{
							Href: fmt.Sprintf("/code-lists/%s", testCodeListID),
							ID:   testCodeListID,
						},
						Editions: &models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions", testCodeListID),
						},
					},
				})

				So(db.QueryForResultCalls, ShouldHaveLength, 1)
				So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(codeListExistsQuery, store.codeListLabel, testCodeListID))
				So(db.QueryForResultCalls[0].Params, ShouldBeNil)
			})
		})
	})
}

func TestNeoDataStore_GetCodeListNoResults(t *testing.T) {
	Convey("given bolt.QueryForResult returns no results error", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				boltmock.NoResultFunc,
			},
		}

		store := testStore
		store.bolt = db

		Convey("when get code list is called", func() {
			cl, err := store.GetCodeList(context.Background(), testCodeListID)

			Convey("then the expected error is returned", func() {
				So(err, ShouldEqual, datastore.ErrCodeListNotFound)
				So(cl, ShouldBeNil)

				So(db.QueryForResultCalls, ShouldHaveLength, 1)
				So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(codeListExistsQuery, store.codeListLabel, testCodeListID))
				So(db.QueryForResultCalls[0].Params, ShouldBeNil)
			})
		})
	})
}

func TestNeoDataStore_GetCodeListQueryForResultError(t *testing.T) {
	Convey("given bolt.QueryForResult returns an error", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				boltmock.ErrQueryFunc,
			},
		}

		store := testStore
		store.bolt = db

		Convey("when get code list is called", func() {
			cl, err := store.GetCodeList(context.Background(), testCodeListID)

			Convey("then the expected error is returned", func() {
				So(err, ShouldEqual, boltmock.Err)
				So(cl, ShouldBeNil)

				So(db.QueryForResultCalls, ShouldHaveLength, 1)
				So(db.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(codeListExistsQuery, store.codeListLabel, testCodeListID))
				So(db.QueryForResultCalls[0].Params, ShouldBeNil)
			})
		})
	})
}