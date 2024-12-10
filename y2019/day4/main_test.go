package day4

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestValid(t *testing.T) {
	assert.Equal(t, true, valid(111111, func(fv int) bool { return fv >= 2 }))
	assert.Equal(t, false, valid(223450, func(fv int) bool { return fv >= 2 }))
	assert.Equal(t, false, valid(123789, func(fv int) bool { return fv >= 2 }))
}

func TestValidB(t *testing.T) {
	assert.Equal(t, true, valid(112233, func(fv int) bool { return fv == 2 }))
	assert.Equal(t, false, valid(123444, func(fv int) bool { return fv == 2 }))
	assert.Equal(t, true, valid(111122, func(fv int) bool { return fv == 2 }))
}

func TestSolvePartA(t *testing.T) {
	assert.Equal(t, 1694, solvePartA(parseInput(utils.LoadString("input.txt"))))
}

func TestSolvePartB(t *testing.T) {
	assert.Equal(t, 1148, solvePartB(parseInput(utils.LoadString("input.txt"))))
}
