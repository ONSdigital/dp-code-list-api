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
	Code  string    `json:"id"      bson:"code"`
	Label string    `json:"label"   bson:"label"`
	Links CodeLinks `json:"links"   bson:"links"`
}

// CodeLinks contains links for a code resource
type CodeLinks struct {
	CodeList Link `json:"code_list"     bson:"code_list"`
	Self     Link `json:"self"          bson:"self"`
}
