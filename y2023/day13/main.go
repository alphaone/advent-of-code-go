package day13

import (
	"strings"

	"github.com/alphaone/advent/utils"
)

func findSymmetry(input []string) int {
	options := findOptions(input)
	for _, option := range options {
		if checkSymmetry(input, option) {
			return option * 100
		}
	}

	transposed := utils.AsStrings(utils.Transpose(utils.AsRunes(input)))
	optionsT := findOptions(transposed)
	for _, option := range optionsT {
		if checkSymmetry(transposed, option) {
			return option
		}
	}
	panic("no symmetry found")
}

func findOptions(input []string) []int {
	result := []int{}

	for i, line := range input {
		if i == 0 {
			continue
		}
		lastLine := input[i-1]
		if line == lastLine {
			result = append(result, i)
		}
	}
	return result
}

func checkSymmetry(input []string, option int) bool {
	if option == 0 || option == len(input) {
		return true
	}

	line1 := input[option]
	line2 := input[option-1]

	if line1 != line2 {
		return false
	}

	head := input[:option-1]
	tail := input[option+1:]

	var smallerInput []string
	smallerInput = append(smallerInput, head...)
	smallerInput = append(smallerInput, tail...)
	return checkSymmetry(smallerInput, option-1)
}

func solvePartA(input []string) int {
	inputs := strings.Split(strings.Join(input, "\n"), "\n\n")
	result := 0
	for _, input := range inputs {
		result += findSymmetry(strings.Split(input, "\n"))
	}

	return result
}
