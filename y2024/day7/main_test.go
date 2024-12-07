package day7

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestIsSolvable(t *testing.T) {
	assert.Equal(t, true, Equation{result: 292, numbers: []int{11, 6, 16, 20}}.IsSolvable(outcomesA))
	assert.Equal(t, false, Equation{result: 21037, numbers: []int{9, 7, 18, 13}}.IsSolvable(outcomesA))
	assert.Equal(t, false, Equation{result: 7290, numbers: []int{6, 8, 6, 15}}.IsSolvable(outcomesA))

	assert.Equal(t, true, Equation{result: 7290, numbers: []int{6, 8, 6, 15}}.IsSolvable(outcomesB))
}

func TestAllOutcomes(t *testing.T) {
	assert.Equal(t, []int{3, 2}, allOutcomes([]int{1, 2}, 10, outcomesA))
	assert.Equal(t, []int{6, 9, 5, 6}, allOutcomes([]int{1, 2, 3}, 10, outcomesA))

	assert.Equal(t, []int{3, 2, 12}, allOutcomes([]int{1, 2}, 15, outcomesB))
	assert.Equal(t, []int{3, 2}, allOutcomes([]int{1, 2}, 10, outcomesB))
}

func TestConcat(t *testing.T) {
	assert.Equal(t, 1234, concatInts(12, 34))
	assert.Equal(t, 100034, concatInts(1000, 34))
}

func TestSolvePartA(t *testing.T) {
	assert.Equal(t, 4364915411363, solvePart(parseInput(utils.LoadStrings("input.txt")), outcomesA))
}

func TestSolvePartB(t *testing.T) {
	assert.Equal(t, 38322057216320, solvePart(parseInput(utils.LoadStrings("input.txt")), outcomesB))
}
