package generic

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Length[T any](param T) T {
	return param
}

func TestSample(t *testing.T) {
	var result string = Length("Azmi")
	assert.Equal(t, "Azmi", result)

	var resultNumber int = Length(100)
	assert.Equal(t, 100, resultNumber)
}
