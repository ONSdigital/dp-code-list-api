package models

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
	Links          DatasetLinks     `json:"links"`
	DimensionLabel string           `json:"dimension_label"`
	Editions       []DatasetEdition `json:"editions"`
}

type DatasetEdition struct {
	Links DatasetEditionLinks `json:"links"`
}

type DatasetEditionLinks struct {
	Self             Link `json:"self"`
	DatasetDimension Link `json:"dataset_dimension"`
	LatestVersion    Link `json:"latest_version"`
}

// DatasetLink represents the links returned specifically for a dataset
type DatasetLinks struct {
	Self Link `json:"self"`
}
