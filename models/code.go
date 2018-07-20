package models

// CodeResults contains an array of codes which can be paginated
type CodeResults struct {
	Items      []Code `json:"items"`
	TotalCount int    `json:"number_of_results"`
}

// Code for a single dimensions type
type Code struct {
	ID    string    `json:"-"`
	Code  string    `json:"id"`
	Label string    `json:"label"`
	Links CodeLinks `json:"links"`
}

// CodeLinks contains links for a code resource
type CodeLinks struct {
	CodeList Link `json:"code_list"`
	Self     Link `json:"self"`
}
