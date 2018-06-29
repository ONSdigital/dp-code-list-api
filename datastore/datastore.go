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
	GetCodeLists(ctx context.Context, filterBy string) (*models.CodeListResults, error)
	GetCodeList(ctx context.Context, codeListID string) (*models.CodeList, error)
	GetEditions(ctx context.Context, codeListID string) (*models.Editions, error)
	GetEdition(ctx context.Context, codeListID, edition string) (*models.Edition, error)
}
