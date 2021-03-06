package datastore

import (
	"context"

	"github.com/ONSdigital/dp-graph/v2/models"
)

//go:generate moq -out datastoretest/datastore.go -pkg storetest . DataStore

// DataStore used to get resources by the API
type DataStore interface {
	GetCodeLists(ctx context.Context, filterBy string) (*models.CodeListResults, error)
	GetCodeList(ctx context.Context, code string) (*models.CodeList, error)
	GetEditions(ctx context.Context, codeListID string) (*models.Editions, error)
	GetEdition(ctx context.Context, codeListID, editionID string) (*models.Edition, error)
	CountCodes(ctx context.Context, codeListID, edition string) (int64, error)
	GetCodes(ctx context.Context, codeListID, editionID string) (*models.CodeResults, error)
	GetCode(ctx context.Context, codeListID, editionID string, codeID string) (*models.Code, error)
	GetCodeDatasets(ctx context.Context, codeListID, edition string, code string) (*models.Datasets, error)
}
