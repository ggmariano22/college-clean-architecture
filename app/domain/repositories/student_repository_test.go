package repositories

import (
	"college/app/domain/entities"
	"college/app/handlers/dtos"
	"errors"
	"testing"

	"github.com/stretchr/testify/mock"
)

type StudentRepositoryMock struct {
	mock.Mock
}

func (m *StudentRepositoryMock) AddStudent(student *dtos.StudentInput) {
	m.Called(student)
}

func (m *StudentRepositoryMock) GetStudents() []*entities.Student {
	args := m.Called()

	return args.Get(0).([]*entities.Student)
}

func (m *StudentRepositoryMock) GetStudentByCPF(cpf int) (*entities.Student, error) {
	args := m.Called(cpf)

	return args.Get(0).(*entities.Student), args.Error(1)
}

func TestAddStudent(t *testing.T) {
	studentMock := getStudentMock()

	mockRepo := new(StudentRepositoryMock)

	mockRepo.On("AddStudent", studentMock)

	mockRepo.AddStudent(studentMock)

	mockRepo.AssertExpectations(t)
}

func TestGetStudents(t *testing.T) {
	mockRepo := new(StudentRepositoryMock)

	mockRepo.On("GetStudents").Return([]*entities.Student{})

	mockRepo.GetStudents()

	mockRepo.AssertExpectations(t)
}

func TestGetStudentByCPF(t *testing.T) {
	mockRepo := new(StudentRepositoryMock)

	mockRepo.On("GetStudentByCPF", 123456).Return(&entities.Student{}, nil)

	mockRepo.GetStudentByCPF(123456)

	mockRepo.AssertExpectations(t)
}

func TestGetStudentByCPFWithError(t *testing.T) {
	mockRepo := new(StudentRepositoryMock)

	mockRepo.On("GetStudentByCPF", 123456).Return(&entities.Student{}, errors.New("Student with CPF 123456 not found."))

	mockRepo.GetStudentByCPF(123456)

	mockRepo.AssertExpectations(t)
}

func getStudentMock() *dtos.StudentInput {
	input := &dtos.StudentInput{
		CommonPersonInput: dtos.CommonPersonInput{
			Name:  "Guilherme",
			Email: "guilherme@email.com",
			CPF:   12345678910,
		},
		PhoneNumber: "12345-6789",
	}

	return input
}
