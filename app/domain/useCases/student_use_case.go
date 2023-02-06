package useCases

import (
	"college/app/domain/entities"
	"college/app/handlers/dtos"
)

type IStudentRepository interface {
	AddStudent(student *dtos.StudentInput)
	GetStudents() []*entities.Student
	GetStudentByCPF(cpf int) (*entities.Student, error)
}

type StudentUseCase struct {
	studentRepository IStudentRepository
}

func NewStudentUseCase(repository IStudentRepository) *StudentUseCase {
	return &StudentUseCase{
		studentRepository: repository,
	}
}

func (uc *StudentUseCase) AddStudent(student *dtos.StudentInput) {
	uc.studentRepository.AddStudent(student)
}

func (uc *StudentUseCase) GetStudents() []*dtos.StudentOutput {
	output := []*dtos.StudentOutput{}

	students := uc.studentRepository.GetStudents()

	for _, student := range students {
		output = append(output, &dtos.StudentOutput{
			Name:  student.GetPerson().GetName(),
			Email: student.GetPerson().GetEmail(),
		})
	}

	return output
}

func (uc *StudentUseCase) GetStudentByCPF(cpf int) (*dtos.StudentOutput, error) {
	student, err := uc.studentRepository.GetStudentByCPF(cpf)

	if err != nil {
		return nil, err
	}

	return &dtos.StudentOutput{
		Name:  student.GetPerson().GetName(),
		Email: student.GetPerson().GetEmail(),
	}, nil
}
