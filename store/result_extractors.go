package store

import (
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-code-list-api/models"
	"reflect"
	"github.com/pkg/errors"
	"strconv"
	"fmt"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

var errCodeNotFound = errors.New("code not found")


func countEditionExtractor() (*int64, dpbolt.ResultExtractor) {
	var count int64
	extractor := func(r *dpbolt.Result) error {
		var ok bool
		count, ok = r.Data[0].(int64)
		if !ok {
			return errors.New("extract row result error: failed to cast result to int64")
		}
		return nil
	}
	return &count, extractor
}

func codeResultExtractor(codeModel *models.Code, codeListID string, edition string) dpbolt.ResultExtractor {
	extractor := func(r *dpbolt.Result) error {
		if len(r.Data) == 0 {
			return errCodeNotFound
		}

		node, ok := r.Data[0].(graph.Node)
		if !ok {
			t := reflect.TypeOf(r.Data[0]).String()
			return errors.Errorf("row.Data[0] incorrect type - expected \"graph.Node\", actual %q", t)
		}

		codeVal, ok := node.Properties["value"].(string)
		if !ok {
			t := reflect.TypeOf(node.Properties["value"]).String()
			return errors.Errorf("node.Properties[\"value\"] incorrect type - expected \"string\", actual %q", t)
		}

		rel, ok := r.Data[1].(graph.Relationship)
		if !ok {
			t := reflect.TypeOf(r.Data[1]).String()
			return errors.Errorf("row.Data[1] incorrect type - expected \"graph.Relationship\", actual %q", t)
		}

		label, ok := rel.Properties["label"].(string)
		if !ok {
			t := reflect.TypeOf(rel.Properties["label"]).String()
			return errors.Errorf("rel.Properties[\"label\"] incorrect type - expected \"string\", actual %q", t)
		}

		codeModel.ID = strconv.FormatInt(node.NodeIdentity, 10)
		codeModel.Code = codeVal
		codeModel.Label = label
		codeModel.Links = models.CodeLinks{
			Self: models.Link{
				Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes/%s", codeListID, edition, codeVal),
			},
			CodeList: models.Link{
				Href: fmt.Sprintf("/code-lists/%s", codeListID),
			},
		}
		return nil
	}
	return extractor
}
