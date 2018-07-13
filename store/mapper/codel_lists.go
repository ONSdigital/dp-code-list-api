package mapper

import (
	"github.com/pkg/errors"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"strings"
	"github.com/ONSdigital/dp-code-list-api/models"
	"fmt"
	"strconv"
	"github.com/ONSdigital/dp-code-list-api/datastore"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

func CodeLists(codeListEditionsMap map[string]*models.CodeList) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
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
}

func CodeList(codeList *models.CodeList, id string) dpbolt.ResultMapper {
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
