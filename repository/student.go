package repository

import (
	"context"
	"github.com/tizor98/proto-primer/models"
)

type IStudentRepository interface {
	GetStudent(ctx context.Context, id string) (*models.Student, error)
	SetStudent(ctx context.Context, student *models.Student) error
}

func StudentRepository(ctx context.Context) IStudentRepository {
	db := getDBConn(ctx)
	return &studentRepoImp{db: db}
}
