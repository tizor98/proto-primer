package repository

import (
	"context"
	"github.com/tizor98/proto-primer/models"
)

type ITestRepository interface {
	GetTest(ctx context.Context, id string) (*models.Test, error)
	SetTest(ctx context.Context, test *models.Test) (string, error)
}

func TestRepository(ctx context.Context) ITestRepository {
	db := getDBConn(ctx)
	return &testRepoImp{db: db}
}
