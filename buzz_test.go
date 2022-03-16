package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Divideby(t *testing.T) {
	numbersToDivide := []string{"5", "10", "3", "1"}

	result := replaceDivisiblesby5(numbersToDivide)
	assert.Equal(t, []string{"Buzz", "Buzz", "3", "1"}, result)
}
