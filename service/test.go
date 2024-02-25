package service

import (
	"context"
	"errors"
	"github.com/tizor98/proto-primer/models"
	proto_student "github.com/tizor98/proto-primer/proto-student"
	test "github.com/tizor98/proto-primer/proto-test"
	"github.com/tizor98/proto-primer/repository"
	"github.com/tizor98/proto-primer/utils"
	"io"
	"log"
	"time"
)

var (
	ErrTestNotFound = errors.New("test was not found")
)

type ITestService interface {
	test.TestServiceServer
}

type testService struct {
	testRepo       repository.ITestRepository
	questionRepo   repository.IQuestionRepository
	enrollmentRepo repository.IEnrollmentRepository
	test.UnimplementedTestServiceServer
}

func NewTestService(
	testRepo repository.ITestRepository,
	questionRepo repository.IQuestionRepository,
	enrollmentRepo repository.IEnrollmentRepository,
) ITestService {
	return &testService{
		testRepo:       testRepo,
		questionRepo:   questionRepo,
		enrollmentRepo: enrollmentRepo,
	}
}

func (t *testService) GetTest(ctx context.Context, req *test.GetTestRequest) (*test.Test, error) {
	testModel, err := t.testRepo.GetTest(ctx, req.GetId())
	if err != nil {
		log.Fatalln(err)
	}
	if studentDto, ok := utils.Mapper(testModel, &test.Test{}); ok {
		return studentDto, nil
	}
	return nil, ErrTestNotFound
}

func (t *testService) SetTest(ctx context.Context, req *test.Test) (*test.SetTestResponse, error) {
	testModel := &models.Test{
		Id:   req.GetId(),
		Name: req.GetName(),
	}

	testId, err := t.testRepo.SetTest(ctx, testModel)
	if err != nil {
		return nil, err
	}
	return &test.SetTestResponse{Id: testId, Name: testModel.Name}, nil
}

func (t *testService) SetQuestions(stream test.TestService_SetQuestionsServer) error {
	for {
		questionMsg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&test.ClientStreamResponse{Ok: true})
		}
		if err != nil {
			return err
		}
		question := &models.Question{
			Id:       questionMsg.Id,
			Question: questionMsg.Question,
			Answer:   questionMsg.Answer,
			TestId:   questionMsg.TestId,
		}
		err = t.questionRepo.SetQuestion(stream.Context(), question)
		if err != nil {
			return stream.SendAndClose(&test.ClientStreamResponse{Ok: false})
		}
	}
}

func (t *testService) EnrollStudents(stream test.TestService_EnrollStudentsServer) error {
	for {
		enrollMsg, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&test.ClientStreamResponse{Ok: true})
		}
		if err != nil {
			return err
		}

		if enrollModel, ok := utils.Mapper(enrollMsg, &models.Enrollment{}); ok {
			if err = t.enrollmentRepo.EnrollStudent(stream.Context(), enrollModel); err != nil {
				return stream.SendAndClose(&test.ClientStreamResponse{Ok: false})
			}
		}
	}
}

func (t *testService) GetStudentsPerTest(req *test.GetStudentPerTestRequest, stream test.TestService_GetStudentsPerTestServer) error {
	students, err := t.enrollmentRepo.GetStudentsPerTest(stream.Context(), req.GetTestId())
	if err != nil {
		return err
	}

	for _, student := range students {
		time.Sleep(time.Second * 2)
		if studentDto, ok := utils.Mapper(student, &proto_student.Student{}); ok {
			if err := stream.Send(studentDto); err != nil {
				return err
			}
		}
	}
	return nil
}

func (t *testService) TakeTest(stream test.TestService_TakeTestServer) error {
	var questions []*models.Question
	var totalQuestions int32
	var points int32
	var currQuestion int32
	for {
		req, err := stream.Recv()
		if err != nil {
			return err
		}

		var prevResponse = test.ANSWER_STATUS_INCORRECT
		if len(req.GetAnswer()) == 0 && questions == nil {
			questions, err = t.questionRepo.GetQuestionsPerTest(stream.Context(), req.GetTestId())
			if err != nil {
				return err
			}
			totalQuestions = int32(len(questions))
		} else if questions != nil && currQuestion < totalQuestions {
			if req.GetAnswer() == questions[currQuestion].Answer {
				points++
				prevResponse = test.ANSWER_STATUS_CORRECT
			}
			currQuestion++
		}

		if currQuestion >= totalQuestions {
			var testStatus = test.TEST_STATUS_FAIL
			if points == totalQuestions {
				testStatus = test.TEST_STATUS_PASS
			}
			response := &test.TakeTestResponse{TakeTest: &test.TakeTestResponse_Results{Results: &test.TakeTestResults{
				PrevAnswer:     prevResponse,
				TotalPoints:    points,
				TotalQuestions: totalQuestions,
				Status:         testStatus,
			}}}
			err := stream.Send(response)
			return err
		}

		questionToAsk := questions[currQuestion]
		response := &test.TakeTestResponse{TakeTest: &test.TakeTestResponse_Question{Question: &test.TakeTestQuestion{
			QuestionId:      questionToAsk.Id,
			Question:        questionToAsk.Question,
			CurrentQuestion: currQuestion + 1,
			TotalQuestions:  totalQuestions,
			PrevAnswer:      prevResponse,
		}}}

		if err = stream.Send(response); err != nil {
			return err
		}

	}
}
