package repositories

import (
	"college/app/domain/entities"
	"college/app/handlers/dtos"
	"errors"
	"fmt"
)

type StudentRepository struct {
	students []*entities.Student
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (sr *StudentRepository) AddStudent(student *dtos.StudentInput) {
	sr.students = append(sr.students, entities.NewStudent(student))
}

func (sr *StudentRepository) GetStudents() []*entities.Student {
	return sr.students
}

func (sr *StudentRepository) GetStudentByCPF(cpf int) (*entities.Student, error) {
	for _, student := range sr.GetStudents() {
		if cpf == student.GetPerson().GetCpf() {
			fmt.Println(student)
			return student, nil
		}
	}

	return nil, errors.New(fmt.Sprintf("Student with CPF %d not found", cpf))
}
