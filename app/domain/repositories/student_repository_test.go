package repositories

import (
	"college/app/domain/entities"
	"testing"

	"github.com/stretchr/testify/mock"
)

type StudentRepositoryMock struct {
	mock.Mock
}

func (m *StudentRepositoryMock) AddStudent(student *entities.Student) {
	m.Called(student)
}

func (m *StudentRepositoryMock) GetStudents() []*entities.Student {
	args := m.Called()

	return args.Get(0).([]*entities.Student)
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

func getStudentMock() *entities.Student {
	return &entities.Student{
		Name:        "Guilherme",
		Email:       "guilherme@email.com",
		CPF:         12345678910,
		PhoneNumber: "15123457891",
	}
}
