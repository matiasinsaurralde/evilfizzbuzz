package main

import (
	"strconv"

	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateNumbers(t *testing.T) {
	numbers := make([]int, 101)
	GenerateNumbers(numbers)

	assert.Len(t, numbers, 101)
	assert.Equal(t, 1, numbers[1])
	assert.Equal(t, 50, numbers[50])
	assert.Equal(t, 100, numbers[100])
}

func FizzBuzz3And5(list []int) []string {
	strList := make([]string, len(list))
	for i, num := range list {
		if num%15 == 0 {
			strList[i] = "FizzBuzz"
		} else {
			strList[i] = strconv.Itoa(num)
		}
	}
	return strList
}

func FizzBuzz3(list []int) []string {
	strList := make([]string, len(list))
	for i, num := range list {
		if num%3 == 0 {
			strList[i] = "Fizz"
		} else {
			strList[i] = strconv.Itoa(num)
		}
	}
	return strList
}

func TestEvilFizzBuzz_FizzBuzz3(t *testing.T) {
	var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	var ExpectedList = []string{"1", "2", "Fizz", "4", "5", "Fizz", "7", "8", "Fizz", "10", "11", "Fizz"}
	assert.Equal(t, ExpectedList, FizzBuzz3(list))
}

func TestEvilFizzBuzz_FizzBuzz3And5(t *testing.T) {
	var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 30, 60, 75, 90, 100}
	var ExpectedList = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
		"FizzBuzz", "16", "FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz", "100"}
	assert.Equal(t, ExpectedList, FizzBuzz3And5(list))
}

func TestGetCommaDelimitedString(t *testing.T) {
	slice := []string{"1", "2", "Fizz"}
	expectation := "1,2,Fizz"
	actual := GetCommaDelimitedString(slice)
	assert.Equal(t, expectation, actual)
}
