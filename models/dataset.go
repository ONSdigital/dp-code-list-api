package models

// Datasets represents the model returned from the api datasets
// endpoint
type Datasets struct {
	Items           []Dataset `json:"datasets"`
	NumberOfResults int       `json:"number_of_results"`
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
