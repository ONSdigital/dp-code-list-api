package models

// CodeResults contains an array of code lists which can be paginated
type CodeListResults struct {
	Count      int        `json:"count"`
	Offset     int        `json:"offset"`
	Limit      int        `json:"limit"`
	Items      []CodeList `json:"items"`
	TotalCount int        `json:"total_count"`
}

// CodeList containing links to all possible codes
type CodeList struct {
	ID    string       `json:"id"       bson:"_id"`
	Name  string       `json:"name"     bson:"name"`
	Links CodeListLink `json:"links"    bson:"links"`
}

// CodeListLink contains links for a code list resource
type CodeListLink struct {
	Self  Link `json:"self"     bson:"self"`
	Codes Link `json:"codes"    bson:"codes"`
}
