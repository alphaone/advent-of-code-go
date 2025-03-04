package day10

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleInput = []int{16, 10, 15, 5, 1, 11, 7, 19, 6, 12, 4}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 35, solveA(exampleInput))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 2059, solveA(utils.LoadInts("input.txt")))
}

func TestGraph(t *testing.T) {
	expected := map[int][]int{
		0: {1}, 1: {4}, 4: {5, 6, 7}, 5: {6, 7}, 6: {7}, 7: {10}, 10: {11, 12}, 11: {12}, 12: {15}, 15: {16}, 16: {19}, 19: {},
	}
	assert.Equal(t, expected, buildGraph(exampleInput))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 8, solveB(exampleInput))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 86812553324672, solveB(utils.LoadInts("input.txt")))
}
