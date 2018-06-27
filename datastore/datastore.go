package datastore

import (
	"context"
	"errors"

	"github.com/ONSdigital/dp-code-list-api/models"
)

var NOT_FOUND = errors.New("resource not found")

//go:generate moq -out datastoretest/datastore.go -pkg storetest . DataStore

// DataStore used to get resources by the API
type DataStore interface {
	GetCodeLists(ctx context.Context) (*models.CodeListResults, error)
	GetCodeList(ctx context.Context, codeListID string) (*models.CodeList, error)
}
