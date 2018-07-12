package models

// CodeResults contains an array of codes which can be paginated
type CodeResults struct {
	Count      int    `json:"count"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	Items      []Code `json:"items"`
	TotalCount int    `json:"total_count"`
}

// Code for a single dimensions type
type Code struct {
	ID    string    `json:"-"`
	Code  string    `json:"id"`
	Label string    `json:"label"`
	Links CodeLinks `json:"links"`
	Type  string    `json:"type,omitempty"`
}

// CodeLinks contains links for a code resource
type CodeLinks struct {
	CodeList Link `json:"code_list"`
	Self     Link `json:"self"`
}
