package repository

import (
	"context"
	"database/sql"
	"github.com/tizor98/proto-primer/models"
	"github.com/tizor98/proto-primer/utils"
)

type questionRepoImp struct {
	db *sql.DB
}

func (q *questionRepoImp) SetQuestion(ctx context.Context, question *models.Question) error {
	_, err := q.db.ExecContext(ctx,
		"INSERT INTO QUESTIONS (id, test_id, question, answer) VALUES ($1, $2, $3, $4)",
		question.Id,
		question.TestId,
		question.Question,
		question.Answer,
	)
	return err
}

func (q *questionRepoImp) GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error) {
	rows, err := q.db.QueryContext(ctx,
		"SELECT ID, TEST_ID, QUESTION, ANSWER FROM QUESTIONS WHERE test_id = $1",
		testId,
	)
	if err != nil {
		return nil, err
	}
	defer utils.CloseElement(rows)

	questions := make([]*models.Question, 0)
	for rows.Next() {
		question := new(models.Question)
		if err := rows.Scan(&question.Id, &question.TestId, &question.Question, &question.Answer); err != nil {
			return nil, err
		}

		questions = append(questions, question)
	}
	return questions, nil
}
