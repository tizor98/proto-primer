package repository

import (
	"context"
	"github.com/tizor98/proto-primer/models"
)

type IQuestionRepository interface {
	SetQuestion(ctx context.Context, question *models.Question) error
	GetQuestionsPerTest(ctx context.Context, testId string) ([]*models.Question, error)
}

func QuestionRepository(ctx context.Context) IQuestionRepository {
	db := getDBConn(ctx)
	return &questionRepoImp{db: db}
}
