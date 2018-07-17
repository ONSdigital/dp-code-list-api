package models

// CodeListResults contains an array of code lists which can be paginated
type CodeListResults struct {
	Count      int        `json:"count"`
	Offset     int        `json:"offset"`
	Limit      int        `json:"limit"`
	Items      []CodeList `json:"items"`
	TotalCount int        `json:"number_of_results"`
}

// CodeList containing links to all possible codes
type CodeList struct {
	ID    string       `json:"-"`
	Name  string       `json:"name"`
	Links CodeListLink `json:"links"`
}

// CodeListLink contains links for a code list resource
type CodeListLink struct {
	Self     *Link `json:"self"`
	Editions *Link `json:"editions"`
	Latest   *Link `json:"latest,omitempty"`
}
