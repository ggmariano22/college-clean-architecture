package useCases

import "college/app/domain/entities"

var Students []*entities.Student

type StudentRepository interface {
	AddStudent(student *entities.Student)
	GetStudents() []*entities.Student
}

type StudentUseCase struct {
	studentRepository StudentRepository
}

func NewStudentUseCase(repository StudentRepository) *StudentUseCase {
	return &StudentUseCase{
		studentRepository: repository,
	}
}

func (uc *StudentUseCase) AddStudent(student *entities.Student) {
	uc.studentRepository.AddStudent(student)
}

func (uc *StudentUseCase) GetStudents() []*entities.Student {
	return uc.studentRepository.GetStudents()
}
