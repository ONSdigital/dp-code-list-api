package models

import (
	"errors"
	"fmt"

	dbmodels "github.com/ONSdigital/dp-graph/v2/models"
)

// Editions represents the editions response model
type Editions struct {
	Items      []Edition `json:"items"`
	Count      int       `json:"count"`
	Offset     int       `json:"offset"`
	Limit      int       `json:"limit"`
	TotalCount int       `json:"total_count"`
}

// Edition represents a single edition response model
type Edition struct {
	ID    string        `json:"edition"`
	Label string        `json:"label"`
	Links *EditionLinks `json:"links"`
}

// EditionLinks represents the links returned for a specific edition
type EditionLinks struct {
	Self     *Link `json:"self"`
	Editions *Link `json:"editions"`
	Codes    *Link `json:"codes"`
}

// UpdateLinks updates the EditionLinks in the Edition struct with the provided codeListID
func (e *Edition) UpdateLinks(codeListID, url string) error {

	if e.ID == "" {
		return errors.New("unable to create links - edition id not provided")
	}

	if e.Links == nil {
		e.Links = &EditionLinks{}
	}

	e.Links.Self = CreateLink(e.ID, fmt.Sprintf(editionURI, codeListID, e.ID), url)
	e.Links.Editions = CreateLink("", fmt.Sprintf(editionsURI, codeListID), url)
	e.Links.Codes = CreateLink("", fmt.Sprintf(codesURI, codeListID, e.ID), url)

	return nil
}

// NewEdition creates an Edition struct from a DB Edition model
func NewEdition(dbEdition *dbmodels.Edition) *Edition {
	if dbEdition == nil {
		return &Edition{}
	}
	return &Edition{
		ID:    dbEdition.ID,
		Label: dbEdition.Label,
	}
}

// NewEditions creates an Editions struct from a DB Editions model
func NewEditions(dbEditions *dbmodels.Editions) *Editions {
	if dbEditions == nil {
		return &Editions{}
	}
	items := []Edition{}
	for _, dbItem := range dbEditions.Items {
		items = append(items, *NewEdition(&dbItem))
	}
	return &Editions{
		Items: items,
	}
}
