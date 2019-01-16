package driver

import (
	"context"
	"errors"

	"github.com/ONSdigital/dp-code-list-api/models"
)

var ErrNotFound = errors.New("not found")

type Driver interface {
	Close(ctx context.Context) error

	GetCodeLists(ctx context.Context, filterBy string) (*models.CodeListResults, error)
	GetCodeList(ctx context.Context, codeListID string) (*models.CodeList, error)
	GetEditions(ctx context.Context, codeListID string) (*models.Editions, error)
	GetEdition(ctx context.Context, codeListID, edition string) (*models.Edition, error)
	GetCodes(ctx context.Context, codeListID, edition string) (*models.CodeResults, error)
	GetCode(ctx context.Context, codeListID, edition string, code string) (*models.Code, error)
	GetCodeDatasets(ctx context.Context, codeListID, edition string, code string) (*models.Datasets, error)
}
