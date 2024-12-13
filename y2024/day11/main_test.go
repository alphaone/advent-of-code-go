package day11

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestSplitN(t *testing.T) {
	a, b := splitNumber(1234)
	assert.Equal(t, 12, a)
	assert.Equal(t, 34, *b)

	a, b = splitNumber(100000)
	assert.Equal(t, 100, a)
	assert.Equal(t, 0, *b)

	a, b = splitNumber(123)
	assert.Equal(t, 123, a)
	assert.Nil(t, b)
}

func TestBlinkRecursive(t *testing.T) {
	stones := []int{125, 17}
	assert.Equal(t, 3, solve(1, stones))
	assert.Equal(t, 4, solve(2, stones))
	assert.Equal(t, 5, solve(3, stones))
	assert.Equal(t, 9, solve(4, stones))
}

func TestSolvePartA(t *testing.T) {
	assert.Equal(t, 185894, solve(25, parseInput(utils.LoadString("input.txt"))))
}

func TestSolvePartB(t *testing.T) {
	assert.Equal(t, 221632504974231, solve(75, parseInput(utils.LoadString("input.txt"))))
}
