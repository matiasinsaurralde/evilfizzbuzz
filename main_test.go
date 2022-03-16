package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateNumbers(t *testing.T) {
	numbers := make([]string, 101)
	GenerateNumbers(numbers)

	assert.Len(t, numbers, 101)
	assert.Equal(t, "1", numbers[1])
	assert.Equal(t, "50", numbers[50])
	assert.Equal(t, "100", numbers[100])
}
