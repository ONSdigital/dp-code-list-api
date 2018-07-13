package store

import (
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-code-list-api/models"
	"reflect"
	"github.com/pkg/errors"
	"strconv"
	"fmt"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"strings"
	"github.com/ONSdigital/dp-code-list-api/datastore"
)

var errCodeNotFound = errors.New("code not found")

func codeListsMapper() (*map[string]*models.CodeList, dpbolt.ResultMapper) {
	codeListEditionsMap := map[string]*models.CodeList{}

	return &codeListEditionsMap, func(r *dpbolt.Result) error {
		if len(r.Data) < 2 {
			return errors.Errorf("expected at least two rows, got %d", len(r.Data))
		}
		props := r.Data[1].(graph.Node).Properties
		labels := r.Data[1].(graph.Node).Labels

		name := props["label"].(string)
		edition := props["edition"].(string)

		var label string
		for _, l := range labels {
			if strings.Contains(l, "_name_") {
				label = strings.Replace(l, `_name_`, "", -1)
				break
			}
		}

		if previousEdition, ok := codeListEditionsMap[label]; !ok {
			// If no edition for this label exists yet in the map, then create one
			codeList := &models.CodeList{
				Links: models.CodeListLink{
					Self: &models.Link{
						Href: fmt.Sprintf("/code-lists/%s", label),
						ID:   label,
					},
					Editions: &models.Link{
						Href: fmt.Sprintf("/code-lists/%s/editions", label),
					},
					Latest: &models.Link{
						Href: fmt.Sprintf("/code-lists/%s/editions/%s", label, edition),
						ID:   edition,
					},
				},
				Name: name,
			}

			codeListEditionsMap[label] = codeList

		} else {
			// If an edition already exists for this label, then check to see if this version is more recent
			previousEditionValue, err := strconv.Atoi(previousEdition.Links.Latest.ID)
			if err != nil {
				return nil
			}

			currentEditionValue, err := strconv.Atoi(edition)
			if err != nil {
				return nil
			}

			if currentEditionValue > previousEditionValue {
				previousEdition.Links.Latest = &models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s", label, edition),
					ID:   edition,
				}
			}
		}
		return nil
	}
	return nil, nil
}

func codeListMapper(codeList *models.CodeList, id string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		if len(r.Data) == 0 {
			return datastore.NOT_FOUND
		}

		props := r.Data[0].(graph.Node).Properties
		name := props["label"].(string)

		codeList.Name = name
		codeList.Links = models.CodeListLink{
			Self: &models.Link{
				Href: fmt.Sprintf("/code-lists/%s", id),
				ID:   id,
			},
			Editions: &models.Link{
				Href: fmt.Sprintf("/code-lists/%s/editions", id),
			},
		}
		return nil
	}
}

// countEditionMapper is a result mapper for gett
func countEditionMapper() (*int64, dpbolt.ResultMapper) {
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

func codesResultMapper(codeListID string, edition string) (*[]models.Code, dpbolt.ResultMapper) {
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

func codeResultMapper(codeModel *models.Code, codeListID string, edition string) dpbolt.ResultMapper {
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
				Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes/%s", codeListID, edition, codeVal),
			},
			CodeList: models.Link{
				Href: fmt.Sprintf("/code-lists/%s", codeListID),
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
