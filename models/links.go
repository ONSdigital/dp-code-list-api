package models

// Href contains the id and a link to a resource
type Href struct {
	ID   string `json:"id"     bson:"id"`
	Href string `json:"href"   bson:"href" `
}
