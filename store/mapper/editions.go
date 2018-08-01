package mapper

import (
	"fmt"
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-code-list-api/models"
		"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

const (
	editionsURI = "/code-lists/%s/editions"
	editionURI  = "/code-lists/%s/editions/%s"
)

func Editions(editions *models.Editions, codeListID string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		var err error
		var node graph.Node

		node, err = getNode(r.Data[0])
		if err != nil {
			return err
		}

		var edition string
		edition, err = getStringProperty("edition", node.Properties)
		if err != nil {
			return err
		}

		var label string
		label, err = getStringProperty("label", node.Properties)
		if err != nil {
			return err
		}

		editionModel := models.Edition{
			Edition: edition,
			Label:   label,
			Links: models.EditionLinks{
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
			},
		}

		editions.Items = append(editions.Items, editionModel)
		return nil
	}
}

func Edition(editionModel *models.Edition, codeListID string, edition string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		node, err := getNode(r.Data[0])
		if err != nil {
			return err
		}

		var label string
		label, err = getStringProperty("label", node.Properties)
		if err != nil {
			return err
		}

		editionModel.Edition = edition
		editionModel.Label = label
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
