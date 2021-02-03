package models

import (
	"errors"
	"fmt"

	dbmodels "github.com/ONSdigital/dp-graph/v2/models"
)

// CodeResults contains an array of codes which can be paginated
type CodeResults struct {
	Items      []Code `json:"items"`
	Count      int    `json:"count"`
	Offset     int    `json:"offset"`
	Limit      int    `json:"limit"`
	TotalCount int    `json:"total_count"`
}

// Code for a single dimensions type
type Code struct {
	ID    string     `json:"code"`
	Label string     `json:"label"`
	Links *CodeLinks `json:"links"`
}

// CodeLinks contains links for a code resource
type CodeLinks struct {
	CodeList *Link `json:"code_list"`
	Datasets *Link `json:"datasets"`
	Self     *Link `json:"self"`
}

// UpdateLinks updates the links for a code list
func (c *Code) UpdateLinks(host, codeListID, edition string) error {

	if c.ID == "" {
		return errors.New("unable to create links - code ID not provided")
	}

	if c.Links == nil {
		c.Links = &CodeLinks{}
	}

	c.Links.Self = CreateLink(c.ID, fmt.Sprintf(codeURI, codeListID, edition, c.ID), host)
	c.Links.Datasets = CreateLink("", fmt.Sprintf(datasetsURI, codeListID, edition, c.ID), host)
	c.Links.CodeList = CreateLink("", fmt.Sprintf(codeListURI, codeListID), host)

	return nil
}

// NewCode creates a new Code struct from a Code DB model
func NewCode(dbCode *dbmodels.Code) *Code {
	if dbCode == nil {
		return &Code{}
	}
	return &Code{
		ID:    dbCode.Code,
		Label: dbCode.Label,
	}
}

// NewCodeResults creates a new CodeResult struct from a CodeResult DB model.
func NewCodeResults(dbCodes []dbmodels.Code) *CodeResults {
	results := &CodeResults{}
	if dbCodes != nil {
		for _, dbItem := range dbCodes {
			results.Items = append(results.Items, *NewCode(&dbItem))
		}
	}
	return results
}
