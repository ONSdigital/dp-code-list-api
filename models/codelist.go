package models

// CodeListResults contains an array of code lists which can be paginated
type CodeListResults struct {
	Items      []CodeList `json:"items"`
	TotalCount int        `json:"number_of_results"`
}

// CodeList containing links to all possible codes
type CodeList struct {
	ID    string       `json:"-"`
	Links CodeListLink `json:"links,omitempty"`
}

// CodeListLink contains links for a code list resource
type CodeListLink struct {
	Self     *Link `json:"self,omitempty"`
	Editions *Link `json:"editions,omitempty"`
}
