package day5

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseOp(t *testing.T) {
	op, params := parseOp(1002)
	assert.Equal(t, 2, op)
	assert.Equal(t, map[int]int{0: 0, 1: 1}, params)
}

func TestSolvePartA(t *testing.T) {
	input := utils.LoadString("input.txt")
	assert.Equal(t, 7259358, solvePartA(parseInput(input)))
}
