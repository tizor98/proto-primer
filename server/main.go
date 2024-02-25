package main

import (
	"context"
	student "github.com/tizor98/proto-primer/proto-student"
	test "github.com/tizor98/proto-primer/proto-test"
	"github.com/tizor98/proto-primer/repository"
	"github.com/tizor98/proto-primer/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"io"
	"log"
	"net"
)

func main() {
	listener, err := net.Listen("tcp", ":5060")
	if err != nil {
		log.Fatalln(err)
	}
	defer closeElement(listener)

	studentRepo := repository.StudentRepository(context.Background())
	testRepo := repository.TestRepository(context.Background())
	questionRepo := repository.QuestionRepository(context.Background())
	enrollmentRepo := repository.EnrollmentRepository(context.Background())

	studentService := service.NewStudentService(studentRepo)
	testService := service.NewTestService(testRepo, questionRepo, enrollmentRepo)

	server := grpc.NewServer()
	student.RegisterStudentServiceServer(server, studentService)
	test.RegisterTestServiceServer(server, testService)

	reflection.Register(server)

	if err := server.Serve(listener); err != nil {
		log.Fatalln(err)
	}
}

func closeElement(closable io.Closer) {
	err := closable.Close()
	if err != nil {
		log.Fatalln(err)
	}
}
