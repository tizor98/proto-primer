package repository

import (
	"context"
	"database/sql"
	"github.com/tizor98/proto-primer/models"
	"math/rand"
	"strconv"
)

type studentRepoImp struct {
	db *sql.DB
}

func (s *studentRepoImp) SetStudent(ctx context.Context, student *models.Student) (string, error) {
	_, err := s.db.ExecContext(ctx, "INSERT INTO STUDENTS (ID, NAME, AGE) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age)
	if err != nil {
		return "", err
	}
	return student.Id, nil
}

func (s *studentRepoImp) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	var student = &models.Student{}
	err := s.db.QueryRowContext(ctx, "SELECT ID, NAME, AGE FROM STUDENTS WHERE ID = $1", id).Scan(&student.Id, &student.Name, &student.Age)
	if err != nil {
		return nil, err
	}
	return student, nil
}

func generateStudentId() string {
	return strconv.FormatUint(rand.New(rand.NewSource(rand.Int63())).Uint64(), 10)
}
