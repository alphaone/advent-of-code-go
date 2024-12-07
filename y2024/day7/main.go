package day7

import (
	"math"
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils"
)

type Equation struct {
	numbers []int
	result  int
}

func (e Equation) IsSolvable(outcomeFn OutcomeFn) bool {
	for _, o := range allOutcomes(e.numbers, e.result, outcomeFn) {
		if e.result == o {
			return true
		}
	}
	return false
}

func allOutcomes(numbers []int, limit int, outcomeFn OutcomeFn) []int {
	if len(numbers) == 1 {
		return numbers
	}
	a := numbers[0]
	b := numbers[1]

	res := []int{}
	for _, o := range outcomeFn(a, b) {
		if o > limit {
			continue
		}

		tail := append([]int{o}, numbers[2:]...)
		res = append(res, allOutcomes(tail, limit, outcomeFn)...)
	}

	return res
}

type OutcomeFn func(a, b int) []int

func outcomesA(a, b int) []int {
	return []int{a + b, a * b}
}

func outcomesB(a, b int) []int {
	return []int{a + b, a * b, concatInts(a, b)}
}

func concatInts(lhs, rhs int) int {
	return int(float64(lhs)*math.Pow(10, math.Ceil(math.Log10(float64(rhs+1))))) + rhs
}

func parseInput(input []string) []Equation {
	res := make([]Equation, 0)
	for _, line := range input {
		parts := strings.Split(line, ":")
		testValue, _ := strconv.Atoi(parts[0])
		numbers := utils.ParseNumbers(parts[1])
		res = append(res, Equation{result: testValue, numbers: numbers})
	}
	return res
}

func solvePart(eqs []Equation, outcomeFn OutcomeFn) int {
	res := 0
	for _, e := range eqs {
		if e.IsSolvable(outcomeFn) {
			res += e.result
		}
	}
	return res
}
