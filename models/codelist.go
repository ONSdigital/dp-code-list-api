package models

// CodeResults
type CodeListResults struct {
	Items []CodeList `json:"items"`
}

// CodeList containing links to all possible codes
type CodeList struct {
	ID    string       `bson:"_id"   json:"id"`
	Name  string       `bson:"name"   json:"name"`
	Links CodeListLink `bson:"links"   json:"links"`
}

// CodeListLink containing al links to resources
type CodeListLink struct {
	Self  Href `json:"self"`
	Codes Href `json:"codes"`
}
