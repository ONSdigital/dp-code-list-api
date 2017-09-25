package models

// CodeResults contains an array of codes which can be paginated
type CodeResults struct {
	Count        int    `json:"count"`
	Start        int    `json:"start_index"`
	ItemsPerPage int    `json:"items_per_page"`
	Items        []Code `json:"items"`
	TotalCount   int    `json:"total_count"`
}

// Code for a single dimensions type
type Code struct {
	ID    string    `json:"-"`
	Code  string    `json:"code"    bson:"code"`
	Label string    `json:"label"   bson:"label"`
	Links CodeLinks `json:"links"   bson:"links"`
}

// CodeLinks contains links for a code resource
type CodeLinks struct {
	CodeList Href `json:"code_list"`
	Self     Href `json:"self"`
}
