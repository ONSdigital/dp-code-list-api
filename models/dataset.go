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
	ID             string      `json:"dataset_id"`
	Edition        string      `json:"edition"`
	Version        int         `json:"version"`
	DimensionLabel string      `json:"dimension_label"`
	Links          DatasetLink `json:"links"`
}

// DatasetLink represents the links returned specifically for a dataset
type DatasetLink struct {
	CodeEdition      Link `json:"code_edition"`
	DatasetVersion   Link `json:"dataset_version"`
	DatasetDimension Link `json:"dataset_dimension"`
}
