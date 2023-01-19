package repositories

import "college/app/domain/entities"

type StudentRepository struct {
	students []*entities.Student
}

func NewStudentRepository() *StudentRepository {
	return &StudentRepository{}
}

func (sr *StudentRepository) AddStudent(student *entities.Student) {
	sr.students = append(sr.students, student)
}

func (sr *StudentRepository) GetStudents() []*entities.Student {
	return sr.students
}
