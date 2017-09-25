package datastore

import (
	"errors"
	"github.com/ONSdigital/dp-code-list-api/models"
)

var NOT_FOUND = errors.New("resource not found")

//go:generate moq -out datastoretest/datastore.go -pkg storetest . DataStore

// DataStore used to get resources by the API
type DataStore interface {
	GetCodes(codeListID string) (*models.CodeResults, error)
	GetCode(codeListID, CodeID string) (*models.Code, error)
	GetCodeLists() (*models.CodeListResults, error)
	GetCodeList(codeListID string) (*models.CodeList, error)
}
