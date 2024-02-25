package repository

import (
	"context"
	"github.com/tizor98/proto-primer/models"
)

type IEnrollmentRepository interface {
	EnrollStudent(ctx context.Context, enrollment *models.Enrollment) error
	GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error)
}

func EnrollmentRepository(ctx context.Context) IEnrollmentRepository {
	db := getDBConn(ctx)
	return &enrollmentRepoImp{db: db}
}
