package models

import (
	"errors"
	"fmt"
	"strconv"

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
	ID            string
	LatestVersion int
	Links         *DatasetEditionLinks `json:"links"`
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

// UpdateLinks updates the Dataset Links for all datasets in items.
func (ds *Datasets) UpdateLinks(datasetAPIhost, codeListID string) error {

	hasErrors := false
	for i, dataset := range ds.Items {
		if err := dataset.UpdateLinks(datasetAPIhost, codeListID); err != nil {
			hasErrors = true
			continue
		}
		ds.Items[i] = dataset
	}

	if hasErrors {
		return errors.New("Error(s) happened updating dataset links")
	}
	return nil
}

// UpdateLinks updates the Dataset Self Link, and all the nested DatasetEdition links.
func (d *Dataset) UpdateLinks(datasetAPIhost, codeListID string) error {

	if d.Links == nil {
		d.Links = &DatasetLinks{}
	}

	if d.ID == "" {
		return errors.New("invalid dataset provided")
	}

	// Update Dataset root-level link
	d.Links.Self = CreateLink(d.ID, fmt.Sprintf(datasetAPIuri, d.ID), datasetAPIhost)

	// Update links in nested Editions
	hasErrors := false
	for i, edition := range d.Editions {
		if err := edition.UpdateLinks(datasetAPIhost, codeListID, d.ID); err != nil {
			hasErrors = true
			continue
		}
		d.Editions[i] = edition
	}

	if hasErrors {
		return errors.New("Error(s) happened updating dataset edition links")
	}
	return nil
}

// UpdateLinks updates the links for the DatasetEdition
func (e *DatasetEdition) UpdateLinks(datasetAPIhost, codeListID, datasetID string) error {

	if e.Links == nil {
		e.Links = &DatasetEditionLinks{}
	}

	// Validate ID (needed by Self, LatestVersion and DatasetDimension links)
	if e.ID == "" {
		return errors.New("datasetEdition with empty ID")
	}

	// Validate datasetID (needed by Self and DatasetDimension links)
	if datasetID == "" {
		return errors.New("empty datasetID provided")
	}

	// Create Self link
	e.Links.Self = CreateLink(e.ID, fmt.Sprintf("/datasets/%s/editions/%s", datasetID, e.ID), datasetAPIhost)

	// Validate LatestVersion (needed by LatestVersion and DatasetDimension links)
	if e.LatestVersion < 0 {
		return errors.New("datasetEdition with invalid LatestVersion")
	}

	// Create LatestVersion link
	versionID := strconv.Itoa(e.LatestVersion)
	e.Links.LatestVersion = CreateLink(versionID, fmt.Sprintf("/datasets/%s/editions/%s/versions/%s", datasetID, e.ID, versionID), datasetAPIhost)

	// Validate codeListID and datasetID (needed by DatasetDimension link)
	if codeListID == "" {
		return errors.New("empty codeListID provided")
	}

	// Create DatasetDimension link
	e.Links.DatasetDimension = CreateLink(codeListID, fmt.Sprintf("/datasets/%s/editions/%s/versions/%s/dimensions/%s", datasetID, e.ID, versionID, codeListID), datasetAPIhost)
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
	if dbDatasetEdition == nil {
		return &DatasetEdition{}
	}
	return &DatasetEdition{
		ID:            dbDatasetEdition.ID,
		LatestVersion: dbDatasetEdition.LatestVersion,
	}
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
