package models

// Href contains the id and a link to a resource
type Href struct {
	ID   string `bson:"id"  json:"id"`
	Href string `bson:"href"   json:"href"`
}
