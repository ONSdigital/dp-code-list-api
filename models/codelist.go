package models

import (
	"errors"
	"fmt"

	dbmodels "github.com/ONSdigital/dp-graph/v2/models"
)

// CodeListResults contains an array of code lists which can be paginated
type CodeListResults struct {
	Items      []CodeList `json:"items"`
	Count      int        `json:"count"`
	Offset     int        `json:"offset"`
	Limit      int        `json:"limit"`
	TotalCount int        `json:"total_count"`
}

// CodeList containing links to all possible codes
type CodeList struct {
	ID    string        `json:"-"`
	Links *CodeListLink `json:"links,omitempty"`
}

// CodeListLink contains links for a code list resource
type CodeListLink struct {
	Self     *Link `json:"self,omitempty"`
	Editions *Link `json:"editions,omitempty"`
}

// UpdateLinks updates the links from
func (c *CodeList) UpdateLinks(url string) error {

	if c.ID == "" {
		return errors.New("unable to create links - codelist id not provided")
	}

	if c.Links == nil {
		c.Links = &CodeListLink{}
	}

	c.Links.Self = CreateLink(c.ID, fmt.Sprintf(codeListURI, c.ID), url)
	c.Links.Editions = CreateLink("", fmt.Sprintf(editionsURI, c.ID), url)

	return nil
}

// NewCodeList creates a new CodeList struct from a CodeList DB model
func NewCodeList(dbCodeList *dbmodels.CodeList) *CodeList {
	if dbCodeList == nil {
		return &CodeList{}
	}
	return &CodeList{
		ID: dbCodeList.ID,
	}
}

// NewCodeListResults creates a new CodeListResult struct from a CodeListResult DB model.
func NewCodeListResults(dbCodeListResults *dbmodels.CodeListResults) *CodeListResults {
	results := &CodeListResults{}
	if dbCodeListResults != nil {
		for _, dbItem := range dbCodeListResults.Items {
			results.Items = append(results.Items, *NewCodeList(&dbItem))
		}
	}
	return results
}
