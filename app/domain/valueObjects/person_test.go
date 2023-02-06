package valueObjects

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreatePerson(t *testing.T) {
	person := NewPerson("Guilherme", "guilherme@email.com", 12345678910)

	assert.Equal(t, "Guilherme", person.GetName())
	assert.Equal(t, "guilherme@email.com", person.GetEmail())
	assert.Equal(t, 12345678910, person.GetCpf())
}
