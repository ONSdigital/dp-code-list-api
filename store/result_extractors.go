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
	return &count, func(r *dpbolt.Result) error {
		var ok bool
		count, ok = r.Data[0].(int64)
		if !ok {
			return errors.New("extract row result error: failed to cast result to int64")
		}
		return nil
	}
}

func codesResultExtractor(codeListID string, edition string) (*[]models.Code, dpbolt.ResultExtractor) {
	var results []models.Code
	return &results, func(r *dpbolt.Result) error {
		codeVal, codeLabel, id, err := getCodeData(r.Data)
		if err != nil {
			return err
		}

		results = append(results, models.Code{
			ID:    strconv.FormatInt(id, 10),
			Code:  codeVal,
			Label: codeLabel,
			Links: models.CodeLinks{
				Self: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes/%s", codeListID, edition, codeVal),
				},
				CodeList: models.Link{
					Href: fmt.Sprintf("/code-lists/%s", codeListID),
				},
			},
		})
		return nil
	}
}

func codeResultExtractor(codeModel *models.Code, codeListID string, edition string) dpbolt.ResultExtractor {
	extractor := func(r *dpbolt.Result) error {
		codeVal, codeLabel, id, err := getCodeData(r.Data)
		if err != nil {
			return err
		}

		codeModel.ID = strconv.FormatInt(id, 10)
		codeModel.Code = codeVal
		codeModel.Label = codeLabel
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

func getCodeData(data []interface{}) (codeVal string, codeLabel string, id int64, err error) {
	var ok bool

	if len(data) == 0 {
		err = errCodeNotFound
		return
	}

	node, ok := data[0].(graph.Node)
	if !ok {
		t := reflect.TypeOf(data[0]).String()
		err = errors.Errorf("row.Data[0] incorrect type - expected \"graph.Node\", actual %q", t)
		return
	}
	id = node.NodeIdentity

	codeVal, ok = node.Properties["value"].(string)
	if !ok {
		t := reflect.TypeOf(node.Properties["value"]).String()
		err = errors.Errorf("node.Properties[\"value\"] incorrect type - expected \"string\", actual %q", t)
		return
	}

	rel, ok := data[1].(graph.Relationship)
	if !ok {
		t := reflect.TypeOf(data[1]).String()
		err = errors.Errorf("row.Data[1] incorrect type - expected \"graph.Relationship\", actual %q", t)
		return
	}

	codeLabel, ok = rel.Properties["label"].(string)
	if !ok {
		t := reflect.TypeOf(rel.Properties["label"]).String()
		err = errors.Errorf("rel.Properties[\"label\"] incorrect type - expected \"string\", actual %q", t)
		return
	}
	return
}
