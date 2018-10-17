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

func (m *Mapper) Editions(editions *models.Editions, codeListID string) dpbolt.ResultMapper {
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
				Self:     *models.CreateLink(edition, fmt.Sprintf(editionURI, codeListID, edition), m.Host),
				Editions: *models.CreateLink("", fmt.Sprintf(editionsURI, codeListID), m.Host),
				Codes:    *models.CreateLink("", fmt.Sprintf(codesURI, codeListID, edition), m.Host),
			},
		}

		editions.Items = append(editions.Items, editionModel)
		return nil
	}
}

func (m *Mapper) Edition(editionModel *models.Edition, codeListID string, edition string) dpbolt.ResultMapper {
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
			Self:     *models.CreateLink(edition, fmt.Sprintf(editionURI, codeListID, edition), m.Host),
			Editions: *models.CreateLink("", fmt.Sprintf(editionsURI, codeListID), m.Host),
			Codes:    *models.CreateLink("", fmt.Sprintf(codesURI, codeListID, edition), m.Host),
		}
		return nil
	}
}
