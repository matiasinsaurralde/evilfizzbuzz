package main

import "strconv"

func GenerateNumbers(numbers []string) {
	// generate numbers from 1 to 100
	for i := range numbers {
		numbers[i] = strconv.Itoa(i)
	}

}

func main() {}
