package store

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ONSdigital/dp-bolt/boltmock"
	"github.com/ONSdigital/dp-bolt/bolt"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/ONSdigital/dp-code-list-api/models"
	"fmt"
	"context"
	"github.com/ONSdigital/dp-code-list-api/datastore"
)

func TestNeoDataStore_GetEditionsSuccess(t *testing.T) {
	Convey("given bolt.QueryForResults does not return an error", t, func() {
		e1 := &bolt.Result{
			Data:  []interface{}{graph.Node{Properties: map[string]interface{}{"label": "e1", "edition": "e1"}}},
			Index: 0,
			Meta:  nil,
		}

		e2 := &bolt.Result{
			Data:  []interface{}{graph.Node{Properties: map[string]interface{}{"label": "e2", "edition": "e2"}}},
			Index: 0,
			Meta:  nil,
		}

		db := &boltmock.DB{
			QueryForResultsFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult bolt.ResultMapper) error {
					mapResult(e1)
					mapResult(e2)
					return nil
				},
			},
		}

		store := NeoDataStore{bolt: db, codeListLabel: "xxx"}

		expected := &models.Editions{
			Items: []models.Edition{
				{
					Edition: "e1",
					Label:   "e1",
					Links: models.EditionLinks{
						Self: models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions/%s", testCodeListID, "e1"),
							ID:   "e1",
						},
						Editions: models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions", testCodeListID),
						},
						Codes: models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes", testCodeListID, "e1"),
						},
					},
				},
				{
					Edition: "e2",
					Label:   "e2",
					Links: models.EditionLinks{
						Self: models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions/%s", testCodeListID, "e2"),
							ID:   "e2",
						},
						Editions: models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions", testCodeListID),
						},
						Codes: models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes", testCodeListID, "e2"),
						},
					},
				},
			},
			NumberOfResults: 2,
		}

		Convey("when GetEditions is called", func() {
			eds, err := store.GetEditions(context.Background(), testCodeListID)

			Convey("then the expected editions is returned and error is nil", func() {
				So(err, ShouldBeNil)
				So(eds, ShouldResemble, expected)

				So(db.QueryForResultsCalls, ShouldHaveLength, 1)
				So(db.QueryForResultsCalls[0].Query, ShouldEqual, fmt.Sprintf(getCodeListQuery, store.codeListLabel, testCodeListID))
				So(db.QueryForResultsCalls[0].Params, ShouldBeNil)
			})

		})
	})
}

func TestNeoDataStore_GetEditionsQueryForResultsError(t *testing.T) {
	Convey("given bolt.QueryForResults returns an error", t, func() {
		db := &boltmock.DB{
			QueryForResultsFuncs: []boltmock.QueryFunc{
				boltmock.ErrQueryFunc,
			},
		}

		store := NeoDataStore{bolt: db, codeListLabel: "666"}

		Convey("when get editions is called", func() {
			eds, err := store.GetEditions(context.Background(), testCodeListID)

			Convey("then the expected error is returned", func() {
				So(err, ShouldEqual, boltmock.Err)
				So(eds, ShouldBeNil)

				So(db.QueryForResultsCalls, ShouldHaveLength, 1)
				So(db.QueryForResultsCalls[0].Query, ShouldEqual, fmt.Sprintf(getCodeListQuery, store.codeListLabel, testCodeListID))
				So(db.QueryForResultsCalls[0].Params, ShouldBeNil)
			})
		})
	})
}

func TestNeoDataStore_GetEditionsNoResults(t *testing.T) {
	Convey("given bolt.QueryForResults returns no results error", t, func() {
		db := &boltmock.DB{
			QueryForResultsFuncs: []boltmock.QueryFunc{
				boltmock.NoResultFunc,
			},
		}

		store := NeoDataStore{bolt: db, codeListLabel: "666"}

		Convey("when get editions is called", func() {
			eds, err := store.GetEditions(context.Background(), testCodeListID)

			Convey("then the expected error is returned", func() {
				So(err, ShouldEqual, datastore.NOT_FOUND)
				So(eds, ShouldBeNil)

				So(db.QueryForResultsCalls, ShouldHaveLength, 1)
				So(db.QueryForResultsCalls[0].Query, ShouldEqual, fmt.Sprintf(getCodeListQuery, store.codeListLabel, testCodeListID))
				So(db.QueryForResultsCalls[0].Params, ShouldBeNil)
			})
		})
	})
}
