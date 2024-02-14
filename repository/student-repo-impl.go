package repository

import (
	"context"
	"database/sql"
	"github.com/tizor98/proto-primer/models"
	"log"
)

type studentRepoImp struct {
	db *sql.DB
}

func (s *studentRepoImp) SetStudent(ctx context.Context, student *models.Student) error {
	var newStudent *models.Student
	err := s.db.QueryRowContext(ctx, "INSERT INTO STUDENT (ID, NAME, AGE) VALUES ($1, $2, $3)", student.Id, student.Name, student.Age).Scan(&newStudent.Id, &newStudent.Name, &newStudent.Age)
	if err != nil {
		log.Fatalln(err)
		return err
	}
	return nil
}

func (s *studentRepoImp) GetStudent(ctx context.Context, id string) (*models.Student, error) {
	var student *models.Student
	err := s.db.QueryRowContext(ctx, "SELECT ID, NAME, AGE FROM STUDENT WHERE ID = $1", id).Scan(&student.Id, &student.Name, &student.Age)
	if err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return student, nil
}
