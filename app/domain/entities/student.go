package entities

import (
	"college/app/domain/valueObjects"
	"college/app/handlers/dtos"
)

type Student struct {
	person      *valueObjects.Person
	phoneNumber string
}

func NewStudent(input *dtos.StudentInput) *Student {
	return &Student{
		person:      valueObjects.NewPerson(input.Name, input.Email, input.CPF),
		phoneNumber: input.PhoneNumber,
	}
}

func (s *Student) GetPerson() *valueObjects.Person {
	return s.person
}

func (s *Student) GetPhoneNumber() string {
	return s.phoneNumber
}
