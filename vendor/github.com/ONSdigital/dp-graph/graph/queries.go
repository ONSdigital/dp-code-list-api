package graph

import (
	"context"

	"github.com/ONSdigital/dp-code-list-api/models"
)

func (db *DB) GetCodeLists(ctx context.Context, filterBy string) (*models.CodeListResults, error) {
	return db.driver.GetCodeLists(ctx, filterBy)
}

func (db *DB) GetCodeList(ctx context.Context, codeListID string) (*models.CodeList, error) {
	return db.driver.GetCodeList(ctx, codeListID)
}

func (db *DB) GetEditions(ctx context.Context, codeListID string) (*models.Editions, error) {
	return db.driver.GetEditions(ctx, codeListID)
}

func (db *DB) GetEdition(ctx context.Context, codeListID, edition string) (*models.Edition, error) {
	return db.driver.GetEdition(ctx, codeListID, edition)
}

func (db *DB) GetCodes(ctx context.Context, codeListID, edition string) (*models.CodeResults, error) {
	return db.driver.GetCodes(ctx, codeListID, edition)
}

func (db *DB) GetCode(ctx context.Context, codeListID, edition string, code string) (*models.Code, error) {
	return db.driver.GetCode(ctx, codeListID, edition, code)
}

func (db *DB) GetCodeDatasets(ctx context.Context, codeListID, edition string, code string) (*models.Datasets, error) {
	return db.driver.GetCodeDatasets(ctx, codeListID, edition, code)
}
