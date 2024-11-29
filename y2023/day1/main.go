package day1

import (
	"strings"
)

var (
	numberDigits = map[string]int{
		`0`: 0, `1`: 1, `2`: 2, `3`: 3, `4`: 4, `5`: 5, `6`: 6, `7`: 7, `8`: 8, `9`: 9,
	}
	wordDigits = map[string]int{
		`0`: 0, `1`: 1, `2`: 2, `3`: 3, `4`: 4, `5`: 5, `6`: 6, `7`: 7, `8`: 8, `9`: 9,
		`one`: 1, `two`: 2, `three`: 3, `four`: 4, `five`: 5, `six`: 6, `seven`: 7, `eight`: 8, `nine`: 9,
	}
)

func solvePartA(input []string) int {
	result := 0
	for _, line := range input {
		first := leftMostDigit(line, numberDigits)
		last := rightMostDigit(line, numberDigits)
		result += first*10 + last
	}
	return result
}

func solvePartB(input []string) int {
	result := 0
	for _, line := range input {
		first := leftMostDigit(line, wordDigits)
		last := rightMostDigit(line, wordDigits)
		result += first*10 + last
	}
	return result
}

func rightMostDigit(line string, digits map[string]int) int {
	max := -1
	res := 0
	for str, value := range digits {
		if idx := strings.LastIndex(line, str); idx > max {
			max = idx
			res = value
		}
	}
	return res
}

func leftMostDigit(line string, digits map[string]int) int {
	min := len(line)
	res := 0
	for str, value := range digits {
		if idx := strings.Index(line, str); idx != -1 && idx < min {
			min = idx
			res = value
		}
	}
	return res
}
