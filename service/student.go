package service

import (
	"context"
	"errors"
	"github.com/tizor98/proto-primer/models"
	student "github.com/tizor98/proto-primer/proto-student"
	"github.com/tizor98/proto-primer/repository"
	"github.com/tizor98/proto-primer/utils"
	"log"
)

var (
	ErrStudentNotFound = errors.New("student was not found")
)

type IStudentService interface {
	student.StudentServiceServer
}

type studentService struct {
	repo repository.IStudentRepository
	student.UnimplementedStudentServiceServer
}

func NewStudentService(repo repository.IStudentRepository) IStudentService {
	return &studentService{repo: repo}
}

func (s *studentService) GetStudent(ctx context.Context, req *student.GetStudentRequest) (*student.Student, error) {
	studentModel, err := s.repo.GetStudent(ctx, req.GetId())
	if err != nil {
		log.Fatalln(err)
	}
	if studentDto, ok := utils.Mapper(studentModel, &student.Student{}); ok {
		return studentDto, nil
	}
	return nil, ErrStudentNotFound
}

func (s *studentService) SetStudent(ctx context.Context, req *student.Student) (*student.SetStudentResponse, error) {
	studentModel := &models.Student{
		Id:   req.GetId(),
		Name: req.GetName(),
		Age:  req.GetAge(),
	}

	studentId, err := s.repo.SetStudent(ctx, studentModel)
	if err != nil {
		return nil, err
	}
	return &student.SetStudentResponse{Id: studentId}, nil
}
