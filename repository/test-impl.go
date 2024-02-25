package repository

import (
	"context"
	"database/sql"
	"github.com/tizor98/proto-primer/models"
)

type testRepoImp struct {
	db *sql.DB
}

func (t *testRepoImp) GetTest(ctx context.Context, id string) (*models.Test, error) {
	var test = &models.Test{}
	err := t.db.QueryRowContext(ctx, "SELECT ID, NAME FROM TESTS WHERE ID = $1", id).Scan(&test.Id, &test.Name)
	if err != nil {
		return nil, err
	}
	return test, nil
}

func (t *testRepoImp) SetTest(ctx context.Context, test *models.Test) (string, error) {
	_, err := t.db.ExecContext(ctx, "INSERT INTO TESTS (ID, NAME) VALUES ($1, $2)", test.Id, test.Name)
	if err != nil {
		return "", err
	}
	return test.Id, nil
}
