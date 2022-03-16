package main

import "strconv"

func main() {

}

func replaceDivisiblesby5(numbers []string) []string {
	result := []string{}
	for index := range numbers{
		number, err := strconv.Atoi(numbers[index])
		if err == nil {
			if number %5 == 0{
				result = append(result, "Buzz")
			}else{
				result = append(result, numbers[index])
			}
		}
	}
	return result
}
