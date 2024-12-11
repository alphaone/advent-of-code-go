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

func TestBlink(t *testing.T) {
	stones := []int{125, 17}
	blink(&stones)
	assert.Equal(t, []int{253000, 1, 7}, stones)
	blink(&stones)
	assert.Equal(t, []int{253, 0, 2024, 14168}, stones)
	blink(&stones)
	assert.Equal(t, []int{512072, 1, 20, 24, 28676032}, stones)
	blink(&stones)
	assert.Equal(t, []int{512, 72, 2024, 2, 0, 2, 4, 2867, 6032}, stones)
}

func TestSolvePartA(t *testing.T) {
	assert.Equal(t, 185894, solve(25, parseInput(utils.LoadString("input.txt"))))
}

func TestSolvePartB(t *testing.T) {
	assert.Equal(t, 185894, solve(75, parseInput(utils.LoadString("input.txt"))))
}
