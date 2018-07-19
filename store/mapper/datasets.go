package mapper

import (
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/ONSdigital/dp-code-list-api/models"
	"fmt"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

func CodesDatasets(datasetsModel *models.Datasets, codeListID string, edition string) dpbolt.ResultMapper {
	return func(r *dpbolt.Result) error {
		node := r.Data[0].(graph.Node)
		relationship := r.Data[1].(graph.Relationship)

		vars := node.Properties
		relVars := relationship.Properties

		datasetID := vars["dataset_id"].(string)
		datasetEdition := vars["edition"].(string)
		version := vars["version"].(int64)

		datasetsModel.Items = append(datasetsModel.Items, models.Dataset{
			ID:             datasetID,
			Edition:        datasetEdition,
			Version:        int(version),
			DimensionLabel: relVars["label"].(string),
			Links: models.DatasetLink{
				CodeEdition: models.Link{
					Href: fmt.Sprintf("/code-lists/%s/editions/%s", codeListID, edition),
				},
				DatasetVersion: models.Link{
					Href: fmt.Sprintf("/datasets/%s/editions/%s/versions/%d", datasetID, datasetEdition, version),
				},
				DatasetDimension: models.Link{
					Href: fmt.Sprintf("/datasets/%s/editions/%s/versions/%d/dimensions/%s", datasetID, datasetEdition, version, codeListID),
				},
			},
		})
		return nil
	}
}
