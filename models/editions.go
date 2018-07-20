package models

// Editions represents the editions response model
type Editions struct {
	Items      []Edition `json:"items"`
	Count      int       `json:"count"`
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
	TotalCount int       `json:"total_count"`
}

// Edition represents a single edition response model
type Edition struct {
	Edition string       `json:"edition"`
	Label   string       `json:"label"`
	Links   EditionLinks `json:"links"`
}

// EditionLinks reprsents the links returned for a specific edition
type EditionLinks struct {
	Self     Link `json:"self"`
	Editions Link `json:"editions"`
	Codes    Link `json:"codes"`
}
