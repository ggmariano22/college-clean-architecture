package entities

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateStudent(t *testing.T) {
	want := &Student{
		Name:        "Guilherme",
		Email:       "guilherme@email.com",
		CPF:         12345678910,
		PhoneNumber: "15123457891",
	}

	got := NewStudent("Guilherme", "guilherme@email.com", 12345678910, "15123457891")

	assert.Equal(t, want, got)
}
