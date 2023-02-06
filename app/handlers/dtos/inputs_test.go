package dtos

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPersonInput(t *testing.T) {
	var mock *StudentInput
	want := mock

	got := &StudentInput{
		CommonPersonInput: CommonPersonInput{
			Name:  "Guilherme",
			Email: "guilherme@email.com",
			CPF:   12345678910,
		},
		PhoneNumber: "12345-6789",
	}

	assert.Equal(t, reflect.TypeOf(want), reflect.TypeOf(got))
}
