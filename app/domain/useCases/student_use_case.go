package useCases

import "college/app/domain/entities"

var Students []*entities.Student

type IStudentRepository interface {
	AddStudent(student *entities.Student)
	GetStudents() []*entities.Student
}

type StudentUseCase struct {
	studentRepository IStudentRepository
}

func NewStudentUseCase(repository IStudentRepository) *StudentUseCase {
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
