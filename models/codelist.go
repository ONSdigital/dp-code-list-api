package models

import (
	"errors"
	"fmt"

	dbmodels "github.com/ONSdigital/dp-graph/v2/models"
)

// CodeListResults contains an array of code lists which can be paginated
type CodeListResults struct {
	Items []CodeList `json:"items"`
	Count int        `json:"count"`

	// Offset is a TODO for future use. Its use will be to do with pagination.
	// It is intended to be used with the other structure members where:
	//   Limit - max number of items we're returning in this response
	//           (e.g. 20, or 50, rather than all we use atm)
	//   Count - how many items are actually present in the response
	//   TotalCount - how many total items there may be
	//                (so the full list size, maybe thousands)
	//   Offset - the number of documents into the full list that this
	//            particular response is starting at.
	//
	// So in a list that has a totalCount of 511, we might set a limit of 100,
	// an offset of 500, and get a response whose count is 11, because its the
	// last 11 documents in the list.
	Offset int `json:"offset"`

	Limit      int `json:"limit"`
	TotalCount int `json:"total_count"`
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
