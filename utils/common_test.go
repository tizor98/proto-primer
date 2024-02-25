package utils

import (
	"github.com/tizor98/proto-primer/models"
	test "github.com/tizor98/proto-primer/proto-test"
	"testing"
)

func TestMapper(t *testing.T) {
	id := "2312312"
	question := "Wdadsadads"
	answer := "DDasda d"
	testId := "test1"

	if questionDto, ok := Mapper(&models.Question{
		Id:       id,
		Question: question,
		Answer:   answer,
		TestId:   testId,
	}, &test.Question{}); ok {
		if questionDto.Id != id ||
			questionDto.Question != question ||
			questionDto.Answer != answer ||
			questionDto.TestId != testId {
			t.Error("structs does not hava equals values")
		}
	}
}
