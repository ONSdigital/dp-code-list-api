package mapper

import (
	"fmt"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/pkg/errors"
	"reflect"
	"strconv"
)

var errCodeNotFound = errors.New("code not found")

const (
	codesURI = "/code-lists/%s/editions/%s/codes"
	codeURI  = "/code-lists/%s/editions/%s/codes/%s"
)

func Codes(results *models.CodeResults, codeListID string, edition string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		codeVal, codeLabel, id, err := getCodeData(r.Data)
		if err != nil {
			return err
		}

		results.Items = append(results.Items, models.Code{
			ID:    strconv.FormatInt(id, 10),
			Code:  codeVal,
			Label: codeLabel,
			Links: models.CodeLinks{
				Self: models.Link{
					Href: fmt.Sprintf(codeURI, codeListID, edition, codeVal),
				},
				Datasets: models.Link{
					Href: fmt.Sprintf(datasetsURI, codeListID, edition, codeVal),
				},
				CodeList: models.Link{
					Href: fmt.Sprintf(codeListURI, codeListID),
				},
			},
		})
		return nil
	}
}

func Code(codeModel *models.Code, codeListID string, edition string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		codeVal, codeLabel, id, err := getCodeData(r.Data)
		if err != nil {
			return err
		}

		codeModel.ID = strconv.FormatInt(id, 10)
		codeModel.Code = codeVal
		codeModel.Label = codeLabel
		codeModel.Links = models.CodeLinks{
			Self: models.Link{
				Href: fmt.Sprintf(codeURI, codeListID, edition, codeVal),
			},
			Datasets: models.Link{
				Href: fmt.Sprintf(datasetsURI, codeListID, edition, codeVal),
			},
			CodeList: models.Link{
				Href: fmt.Sprintf(codeListURI, codeListID),
			},
		}
		return nil
	}
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
