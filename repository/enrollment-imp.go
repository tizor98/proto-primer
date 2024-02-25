package repository

import (
	"context"
	"database/sql"
	"github.com/tizor98/proto-primer/models"
	"github.com/tizor98/proto-primer/utils"
)

type enrollmentRepoImp struct {
	db *sql.DB
}

func (e *enrollmentRepoImp) EnrollStudent(ctx context.Context, enrollment *models.Enrollment) error {
	_, err := e.db.ExecContext(ctx,
		"INSERT INTO ENROLLMENTS (student_id, test_id) VALUES ($1, $2)",
		enrollment.StudentId,
		enrollment.TestId,
	)
	return err
}

func (e *enrollmentRepoImp) GetStudentsPerTest(ctx context.Context, testId string) ([]*models.Student, error) {
	rows, err := e.db.QueryContext(ctx,
		"SELECT ID, NAME, AGE FROM students WHERE ID IN (SELECT student_id FROM enrollments WHERE test_id = $1)",
		testId,
	)
	if err != nil {
		return nil, err
	}
	defer utils.CloseElement(rows)

	output := make([]*models.Student, 0)
	for rows.Next() {
		student := new(models.Student)
		if err := rows.Scan(&student.Id, &student.Name, &student.Age); err != nil {
			return nil, err
		}
		output = append(output, student)
	}
	return output, nil
}
