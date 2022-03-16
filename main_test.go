package main

import (
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHelloWorld(t *testing.T) {

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

func TestEvilFizzBuzz_FizzBuzz3And5(t *testing.T) {
	var list = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 30, 60, 75, 90, 100}
	var ExpectedList = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10", "11", "12", "13", "14",
		"FizzBuzz", "16", "FizzBuzz", "FizzBuzz", "FizzBuzz", "FizzBuzz", "100"}
	assert.Equal(t, ExpectedList, FizzBuzz3And5(list))
}
