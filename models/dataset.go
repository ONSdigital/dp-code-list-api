package models

import (
	"errors"
	"fmt"

	dbmodels "github.com/ONSdigital/dp-graph/v2/models"
)

// Datasets represents the model returned from the api datasets
// endpoint
type Datasets struct {
	Items      []Dataset `json:"items"`
	Count      int       `json:"count"`
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
	TotalCount int       `json:"total_count"`
}

// Dataset represents an individual model dataset
type Dataset struct {
	ID             string
	Links          *DatasetLinks    `json:"links"`
	DimensionLabel string           `json:"dimension_label"`
	Editions       []DatasetEdition `json:"editions"`
}

type DatasetEdition struct {
	Links *DatasetEditionLinks `json:"links"`
}

type DatasetEditionLinks struct {
	Self             *Link `json:"self"`
	DatasetDimension *Link `json:"dataset_dimension"`
	LatestVersion    *Link `json:"latest_version"`
}

// DatasetLinks represents the links returned specifically for a dataset
type DatasetLinks struct {
	Self *Link `json:"self"`
}

func (ds *Datasets) UpdateLinks(host, datasetAPIhost, codeListID, editionID, codeID string) error {
	for i, dataset := range ds.Items {

		if dataset.ID == "" {
			return errors.New("invalid dataset provided")
		}

		if dataset.Links == nil {
			dataset.Links = &DatasetLinks{}
		}

		l := CreateLink(dataset.ID, fmt.Sprintf(datasetAPIuri, dataset.ID), datasetAPIhost)
		dataset.Links.Self = &Link{
			Href: l.Href,
			ID:   dataset.ID,
		}

		var editions []DatasetEdition
		for _, edition := range dataset.Editions {
			if edition.Links == nil || edition.Links.Self == nil || edition.Links.Self.ID == "" {
				continue
			}

			editionID := edition.Links.Self.ID
			edition.Links.Self = CreateLink(editionID, fmt.Sprintf("/datasets/%s/editions/%s", dataset.ID, editionID), datasetAPIhost)
			//	latestVersion := strconv.Itoa(max(versions))

			if edition.Links == nil || edition.Links.LatestVersion == nil || edition.Links.LatestVersion.ID == "" {
				continue
			}

			versionID := edition.Links.LatestVersion.ID
			edition.Links.LatestVersion = CreateLink(versionID, fmt.Sprintf("/datasets/%s/editions/%s/versions/%s", dataset.ID, editionID, versionID), datasetAPIhost)

			if edition.Links == nil || edition.Links.DatasetDimension == nil || edition.Links.DatasetDimension.ID == "" {
				continue
			}

			dimensionID := edition.Links.DatasetDimension.ID
			edition.Links.DatasetDimension = CreateLink(dimensionID, fmt.Sprintf("/datasets/%s/editions/%s/versions/%s/dimensions/%s", dataset.ID, editionID, versionID, dimensionID), datasetAPIhost)

			editions = append(editions, edition)
		}

		dataset.Editions = editions
		ds.Items[i] = dataset
	}
	return nil
}

// NewDataset creates a new Dataset struct from a database Dataset struct
func NewDataset(dbDataset *dbmodels.Dataset) *Dataset {
	if dbDataset == nil {
		return &Dataset{}
	}
	editions := []DatasetEdition{}
	for _, dbEdition := range dbDataset.Editions {
		editions = append(editions, *NewDatasetEdition(&dbEdition))
	}
	return &Dataset{
		ID:             dbDataset.ID,
		DimensionLabel: dbDataset.DimensionLabel,
		Editions:       editions,
	}
}

// NewDatasetEdition creates a new DatasetEdition struct from a database DatasetEdition
func NewDatasetEdition(dbDatasetEdition *dbmodels.DatasetEdition) *DatasetEdition {
	return &DatasetEdition{}
}

// NewDatasets creates a new Datasets struct from a database Datasets
func NewDatasets(dbDatasets *dbmodels.Datasets) *Datasets {
	if dbDatasets == nil {
		return &Datasets{}
	}
	items := []Dataset{}
	for _, dbItem := range dbDatasets.Items {
		items = append(items, *NewDataset(&dbItem))
	}
	return &Datasets{
		Items: items,
	}
}
