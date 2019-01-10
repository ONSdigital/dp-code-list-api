package store

import (
	"context"
	"fmt"
	"testing"

	"github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-bolt/boltmock"
	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/ONSdigital/dp-code-list-api/store/mapper"
	"github.com/ONSdigital/golang-neo4j-bolt-driver/structures/graph"
	. "github.com/smartystreets/goconvey/convey"
)

func TestNeoDataStore_GetEditionQueryForResultError(t *testing.T) {
	Convey("given bolt.QueryForResult returns an error", t, func() {
		mockDB := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				boltmock.ErrQueryFunc,
			},
		}

		store := &NeoDataStore{bolt: mockDB, codeListLabel: "xxx", mapper: &mapper.Mapper{Host: ""}}

		Convey("when get edition is called", func() {
			ed, err := store.GetEdition(context.Background(), testCodeListID, testEdition)

			Convey("then the expected error is returned", func() {
				So(err, ShouldEqual, boltmock.Err)
				So(ed, ShouldBeNil)

				So(mockDB.QueryForResultCalls, ShouldHaveLength, 1)
				So(mockDB.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(getCodeListEditionQuery, "xxx", testCodeListID, testEdition))
				So(mockDB.QueryForResultCalls[0].Params, ShouldBeNil)
			})
		})
	})
}

func TestNeoDataStore_GetEditionNoResults(t *testing.T) {
	Convey("given bolt.QueryForResult returns no results", t, func() {
		mockDB := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				boltmock.NoResultFunc,
			},
		}

		store := &NeoDataStore{bolt: mockDB, codeListLabel: "xxx", mapper: &mapper.Mapper{Host: ""}}

		Convey("when get edition is called", func() {
			ed, err := store.GetEdition(context.Background(), testCodeListID, testEdition)

			Convey("then the expected error is returned", func() {
				So(err, ShouldEqual, datastore.ErrEditionNotFound)
				So(ed, ShouldBeNil)

				So(mockDB.QueryForResultCalls, ShouldHaveLength, 1)
				So(mockDB.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(getCodeListEditionQuery, "xxx", testCodeListID, testEdition))
				So(mockDB.QueryForResultCalls[0].Params, ShouldBeNil)
			})
		})
	})
}

func TestNeoDataStore_GetEditionSuccess(t *testing.T) {
	Convey("given bolt.QueryForResult is successful", t, func() {
		mockDB := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult bolt.ResultMapper) error {
					return mapResult(&bolt.Result{
						Data: []interface{}{
							graph.Node{
								Properties: map[string]interface{}{"label": "editionLabel"},
							},
						},
						Index: 0,
						Meta:  nil,
					})
				},
			},
		}

		store := &NeoDataStore{bolt: mockDB, codeListLabel: "xxx", mapper: &mapper.Mapper{Host: ""}}

		Convey("when get edition is called", func() {
			ed, err := store.GetEdition(context.Background(), testCodeListID, testEdition)

			Convey("then the expected error is returned", func() {
				So(err, ShouldBeNil)
				So(ed, ShouldResemble, &models.Edition{
					Edition: testEdition,
					Label:   "editionLabel",
					Links: models.EditionLinks{
						Self: models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions/%s", testCodeListID, testEdition),
							ID:   testEdition,
						},
						Editions: models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions", testCodeListID),
						},
						Codes: models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes", testCodeListID, testEdition),
						},
					},
				})

				So(mockDB.QueryForResultCalls, ShouldHaveLength, 1)
				So(mockDB.QueryForResultCalls[0].Query, ShouldEqual, fmt.Sprintf(getCodeListEditionQuery, "xxx", testCodeListID, testEdition))
				So(mockDB.QueryForResultCalls[0].Params, ShouldBeNil)
			})
		})
	})
}
