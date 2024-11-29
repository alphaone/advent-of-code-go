package day3

import (
	"strconv"
	"strings"
)

func Transpose(input []string) []string {
	result := make([]string, len(input[0]))
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(input); j++ {
			result[i] += string(input[j][i])
		}
	}
	return result
}

func countOccurrences(input string, char rune) int {
	result := 0
	for _, c := range input {
		if c == char {
			result++
		}
	}
	return result
}

func solvePartA(input []string) int {
	transposed := Transpose(input)
	var gammaRune []rune
	var epsilonRune []rune
	for _, line := range transposed {
		if countOccurrences(line, '1') > countOccurrences(line, '0') {
			gammaRune = append(gammaRune, '1')
			epsilonRune = append(epsilonRune, '0')
		} else {
			gammaRune = append(gammaRune, '0')
			epsilonRune = append(epsilonRune, '1')
		}
	}

	gamma, _ := strconv.ParseInt(string(gammaRune), 2, 64)
	epsilon, _ := strconv.ParseInt(string(epsilonRune), 2, 64)
	return int(gamma * epsilon)
}

func solvePartB(input []string) int {
	oxygen := findWinner(input, func(line string) bool {
		return countOccurrences(line, '1') >= countOccurrences(line, '0')
	})
	co2 := findWinner(input, func(line string) bool {
		return countOccurrences(line, '1') < countOccurrences(line, '0')
	})

	return int(oxygen * co2)
}

func findWinner(input []string, cmp func(string) bool) int {
	var winner string
	for {
		if len(input) == 1 {
			winner += input[0]
			break
		}

		transposed := Transpose(input)
		if cmp(transposed[0]) {
			input = keepLineStartingWith(input, "1")
			winner += "1"
		} else {
			input = keepLineStartingWith(input, "0")
			winner += "0"
		}

		for i, line := range input {
			input[i] = line[1:]
		}
	}
	res, _ := strconv.ParseInt(winner, 2, 64)
	return int(res)
}

func keepLineStartingWith(input []string, prefix string) []string {
	result := make([]string, 0)
	for _, line := range input {
		if strings.HasPrefix(line, prefix) {
			result = append(result, line)
		}
	}
	return result
}
