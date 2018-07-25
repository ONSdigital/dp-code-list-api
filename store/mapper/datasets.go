package mapper

import (
	dpbolt "github.com/ONSdigital/dp-bolt/bolt"
	"github.com/johnnadratowski/golang-neo4j-bolt-driver/structures/graph"
)

type Datasets map[string]datasetData
type DatasetEditions map[string]Versions
type Versions []int

type datasetData struct {
	DimensionLabel string
	Editions       DatasetEditions
}

const (
	datasetsURI = "/code-lists/%s/editions/%s/codes/%s/datasets"
)

func CodesDatasets(datasets Datasets) dpbolt.ResultMapper {

	return func(r *dpbolt.Result) error {

		node := r.Data[0].(graph.Node)
		relationship := r.Data[1].(graph.Relationship)

		vars := node.Properties
		relVars := relationship.Properties

		datasetID := vars["dataset_id"].(string)
		datasetEdition := vars["edition"].(string)
		version := vars["version"].(int)
		dimensionLabel := relVars["label"].(string)

		dataset, ok := datasets[datasetID]
		if !ok {
			dataset = datasetData{
				DimensionLabel: dimensionLabel,
				Editions:       make(DatasetEditions, 0),
			}
		}

		if dataset.Editions[datasetEdition] == nil {
			dataset.Editions[datasetEdition] = make(Versions, 0)
		}

		dataset.Editions[datasetEdition] = append(dataset.Editions[datasetEdition], version)

		datasets[datasetID] = dataset

		return nil
	}
}
