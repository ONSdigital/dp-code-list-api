package store

import (
	"testing"
	. "github.com/smartystreets/goconvey/convey"
	"github.com/ONSdigital/dp-bolt/boltmock"
	"github.com/ONSdigital/dp-bolt/bolt"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"context"
	"github.com/ONSdigital/dp-code-list-api/models"
	"fmt"
	"github.com/ONSdigital/dp-code-list-api/datastore"
)

func TestNeoDataStore_GetCodeListsSuccess(t *testing.T) {
	Convey("given bolt.QueryForResults returns no error", t, func() {
		expected := &models.CodeListResults{
			Items: []models.CodeList{
				{
					Name: "c1",
					Links: models.CodeListLink{
						Self: &models.Link{
							Href: fmt.Sprintf("/code-lists/%s", "c1"),
							ID:   "c1",
						},
						Editions: &models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions", "c1"),
						},
						Latest: &models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions/%s", "c1", "e1"),
							ID:   "e1",
						},
					},
				},
				{
					Name: "c2",
					Links: models.CodeListLink{
						Self: &models.Link{
							Href: fmt.Sprintf("/code-lists/%s", "c2"),
							ID:   "c2",
						},
						Editions: &models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions", "c2"),
						},
						Latest: &models.Link{
							Href: fmt.Sprintf("/code-lists/%s/editions/%s", "c2", "e2"),
							ID:   "e2",
						},
					},
				},
			},
		}

		db := &boltmock.DB{
			QueryForResultsFuncs: [] boltmock.QueryFunc{
				func(query string, params map[string]interface{}, mapResult bolt.ResultMapper) error {
					mapResult(&bolt.Result{
						Data: []interface{}{
							[]interface{}{"_name_c1"},
							graph.Node{Properties: map[string]interface{}{"label": "c1", "edition": "e1"}},
						},
					})

					mapResult(&bolt.Result{
						Data: []interface{}{
							[]interface{}{"_name_c2"},
							graph.Node{Properties: map[string]interface{}{"label": "c2", "edition": "e2"}},
						},
					})
					return nil
				},
			},
		}

		store := NeoDataStore{codeListLabel: "666", bolt: db}

		Convey("when getCodeLists is called", func() {
			actual, err := store.GetCodeLists(context.Background(), "")

			Convey("then the expected code lists is returned and error is nil", func() {
				So(err, ShouldBeNil)
				So(actual, ShouldResemble, expected)
				So(db.QueryForResultsCalls, ShouldHaveLength, 1)
				So(db.QueryForResultsCalls[0].Query, ShouldEqual, fmt.Sprintf(getCodeListsQuery, store.codeListLabel, ""))
				So(db.QueryForResultsCalls[0].Params, ShouldBeNil)
			})
		})
	})
}

func TestNeoDataStore_GetCodeListsNoResults(t *testing.T) {
	Convey("given bolt.QueryForResults returns no results error", t, func() {
		db := &boltmock.DB{
			QueryForResultsFuncs: []boltmock.QueryFunc{
				boltmock.NoResultFunc,
			},
		}

		store := NeoDataStore{codeListLabel: "666", bolt: db}

		Convey("when get codeList is caleld", func() {
			res, err := store.GetCodeLists(context.Background(), "")
			So(err, ShouldEqual, datastore.NOT_FOUND)
			So(res, ShouldBeNil)
			So(db.QueryForResultsCalls, ShouldHaveLength, 1)
			So(db.QueryForResultsCalls[0].Query, ShouldEqual, fmt.Sprintf(getCodeListsQuery, store.codeListLabel, ""))
			So(db.QueryForResultsCalls[0].Params, ShouldBeNil)
		})
	})
}

func TestNeoDataStore_GetCodeListsError(t *testing.T) {
	Convey("given bolt.QueryForResults returns an error", t, func() {
		db := &boltmock.DB{
			QueryForResultsFuncs: []boltmock.QueryFunc{
				boltmock.ErrQueryFunc,
			},
		}

		store := NeoDataStore{codeListLabel: "666", bolt: db}

		Convey("when get codeList is caleld", func() {
			res, err := store.GetCodeLists(context.Background(), "")
			So(err, ShouldEqual, boltmock.Err)
			So(res, ShouldBeNil)
			So(db.QueryForResultsCalls, ShouldHaveLength, 1)
			So(db.QueryForResultsCalls[0].Query, ShouldEqual, fmt.Sprintf(getCodeListsQuery, store.codeListLabel, ""))
			So(db.QueryForResultsCalls[0].Params, ShouldBeNil)
		})
	})
}
