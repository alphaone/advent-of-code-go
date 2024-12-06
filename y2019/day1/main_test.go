package day1

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestFuel(t *testing.T) {
	assert.Equal(t, 2, fuel(12))
	assert.Equal(t, 2, fuel(14))
	assert.Equal(t, 654, fuel(1969))
}

func TestFuel2(t *testing.T) {
	assert.Equal(t, 2, fuel2(12))
	assert.Equal(t, 2, fuel2(14))
	assert.Equal(t, 966, fuel2(1969))
}

func TestSolvePartA(t *testing.T) {
	input := utils.LoadInts("input.txt")
	assert.Equal(t, 3173518, solvePartA(input))
}

func TestSolvePartB(t *testing.T) {
	input := utils.LoadInts("input.txt")
	assert.Equal(t, 4757427, solvePartB(input))
}
