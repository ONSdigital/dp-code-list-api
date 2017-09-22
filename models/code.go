package models

// CodeResults
type CodeResults struct {
	Items []Code `json:"items"`
}

// Code for a single dimensions type
type Code struct {
	ID    string    `json:"-"`
	Code  string    `json:"code"`
	Label string    `json:"label"`
	Links CodeLinks `json:"links"`
}

// CodeLinks contains links from a code resource
type CodeLinks struct {
	CodeList Href `json:"code_list"`
	Self     Href `json:"self"`
}
