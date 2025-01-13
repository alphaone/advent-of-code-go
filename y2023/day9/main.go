package day9

import (
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils/mathutils"
)

func derivative(line []int) []int {
	result := make([]int, len(line)-1)
	for i := 0; i < len(result); i++ {
		result[i] = line[i+1] - line[i]
	}
	return result
}

func allZeros(line []int) bool {
	for _, x := range line {
		if x != 0 {
			return false
		}
	}
	return true
}

func extrapolateNext(line []int) int {
	stack := [][]int{}
	for {
		stack = append([][]int{line}, stack...)
		if allZeros(line) {
			break
		}
		line = derivative(line)
	}
	stack[0] = append(stack[0], 0)

	for i, cur := range stack[1:] {
		lastLine := stack[i]
		nextValue := lastLine[len(lastLine)-1] + cur[len(cur)-1]
		stack[i+1] = append(cur, nextValue)
	}

	return stack[len(stack)-1][len(stack[len(stack)-1])-1]
}

func extrapolatePrev(line []int) int {
	stack := [][]int{}
	for {
		stack = append([][]int{line}, stack...)
		if allZeros(line) {
			break
		}
		line = derivative(line)
	}
	stack[0] = append(stack[0], 0)

	for i, cur := range stack[1:] {
		lastLine := stack[i]
		nextValue := cur[0] - lastLine[0]
		stack[i+1] = append([]int{nextValue}, cur...)
	}
	return stack[len(stack)-1][0]
}

func solvePartA(input []string) int {
	var results []int
	for _, stringLine := range input {
		var intLine []int
		for _, x := range strings.Split(stringLine, " ") {
			converted, _ := strconv.Atoi(x)
			intLine = append(intLine, converted)
		}

		results = append(results, extrapolateNext(intLine))
	}

	return mathutils.Sum(results)
}

func solvePartB(input []string) int {
	var results []int
	for _, stringLine := range input {
		var intLine []int
		for _, x := range strings.Split(stringLine, " ") {
			converted, _ := strconv.Atoi(x)
			intLine = append(intLine, converted)
		}

		results = append(results, extrapolatePrev(intLine))
	}

	return mathutils.Sum(results)
}
