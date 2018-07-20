package datastore

import (
	"context"
	"errors"

	"github.com/ONSdigital/dp-code-list-api/models"
)

var (
	ErrEditionNotFound   = errors.New("edition not found")
	ErrEditionsNotFound  = errors.New("editions not found")
	ErrCodeNotFound      = errors.New("code not found")
	ErrCodeListNotFound  = errors.New("code list not found")
	ErrCodeListsNotFound = errors.New("code lists not found")
	ErrCodesNotFound     = errors.New("codes not found")
	ErrDatasetsNotFound  = errors.New("datasets not found")
)

//go:generate moq -out datastoretest/datastore.go -pkg storetest . DataStore

// DataStore used to get resources by the API
type DataStore interface {
	GetCodeLists(ctx context.Context, filterBy string) (*models.CodeListResults, error)
	GetCodeList(ctx context.Context, codeListID string) (*models.CodeList, error)
	GetEditions(ctx context.Context, codeListID string) (*models.Editions, error)
	GetEdition(ctx context.Context, codeListID, edition string) (*models.Edition, error)
	GetCodes(ctx context.Context, codeListID, edition string) (*models.CodeResults, error)
	GetCode(ctx context.Context, codeListID, edition string, code string) (*models.Code, error)
	GetCodeDatasets(ctx context.Context, codeListID, edition string, code string) (*models.Datasets, error)
}
