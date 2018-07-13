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
	Convey("given...", t, func() {
		db := &boltmock.DB{
			QueryForResultFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, resultExtractor dpbolt.ResultExtractor) error {
					return resultExtractor(
						&dpbolt.Result{Data: []interface{}{int64(1)}},
					)
				},
			},
			QueryForResultsFuncs: []boltmock.QueryFunc{
				func(query string, params map[string]interface{}, resultExtractor dpbolt.ResultExtractor) error {
					resultExtractor(&dpbolt.Result{Data: row1})
					resultExtractor(&dpbolt.Result{Data: row2})
					return nil
				},
			},
		}
		store := NeoDataStore{db: db}

		result, err := store.GetCodes(context.Background(), testCodeListID, testEdition)
		So(err, ShouldBeNil)
		So(result, ShouldResemble, &models.CodeResults{Items: codes, Count: 2})
	})

}
