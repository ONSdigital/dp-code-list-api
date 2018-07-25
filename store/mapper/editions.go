package mapper

import (
	"fmt"

	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-code-list-api/models"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
	"github.com/pkg/errors"
)

const (
	editionsURI = "/code-lists/%s/editions"
	editionURI  = "/code-lists/%s/editions/%s"
)

func Editions(editions *models.Editions, codeListID string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		props := r.Data[0].(graph.Node).Properties
		edition := props["edition"].(string)
		editionModel := models.Edition{
			Edition: edition,
			Label:   props["label"].(string),
			Links: models.EditionLinks{
				Self: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s", codeListID, edition),
					ID:   edition,
				},
				Editions: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions", codeListID),
				},
				Codes: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s/codes", codeListID, edition),
				},
			},
		}

		editions.Items = append(editions.Items, editionModel)
		return nil
	}
}

func Edition(editionModel *models.Edition, codeListID string, edition string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		props := r.Data[0].(graph.Node).Properties

		editionModel.Edition = edition
		editionModel.Label = props["label"].(string)
		editionModel.Links = models.EditionLinks{
			Self: models.Link{
				Href: fmt.Sprintf(editionURI, codeListID, edition),
				ID:   edition,
			},
			Editions: models.Link{
				Href: fmt.Sprintf(editionsURI, codeListID),
			},
			Codes: models.Link{
				Href: fmt.Sprintf(codesURI, codeListID, edition),
			},
		}
		return nil
	}
}

// EditionCount is a result mapper for gett
func EditionCount() (*int, dpbolt.ResultMapper) {
	var count int
	return &count, func(r *dpbolt.Result) error {
		countInt64, ok := r.Data[0].(int64)
		if !ok {
			return errors.New("extract row result error: failed to cast result to int64")
		}

		count = int(countInt64)
		return nil
	}
}
