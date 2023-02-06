package entities

import (
	"college/app/domain/valueObjects"
	"college/app/handlers/dtos"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStudent(t *testing.T) {
	input := &dtos.StudentInput{
		CommonPersonInput: dtos.CommonPersonInput{
			Name:  "Guilherme",
			Email: "guilherme@email.com",
			CPF:   12345678910,
		},
		PhoneNumber: "12345-6789",
	}

	want := &Student{
		person:      valueObjects.NewPerson(input.Name, input.Email, input.CPF),
		phoneNumber: input.PhoneNumber,
	}

	got := NewStudent(input)

	assert.Equal(t, want, got)
}
