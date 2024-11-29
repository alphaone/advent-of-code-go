package day9

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"0 3 6 9 12 15",
	"1 3 6 10 15 21",
	"10 13 16 21 30 45",
}

func TestDerivative(t *testing.T) {
	assert.Equal(t, []int{3, 3, 3, 3, 3}, derivative([]int{0, 3, 6, 9, 12, 15}))
}

func TestExtrapolateNext(t *testing.T) {
	assert.Equal(t, 18, extrapolateNext([]int{0, 3, 6, 9, 12, 15}))
	assert.Equal(t, 28, extrapolateNext([]int{1, 3, 6, 10, 15, 21}))
	assert.Equal(t, 68, extrapolateNext([]int{10, 13, 16, 21, 30, 45}))
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 114, solvePartA(example))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 1762065988, solvePartA(input))
}

func TestExtrapolatePrev(t *testing.T) {
	assert.Equal(t, -3, extrapolatePrev([]int{0, 3, 6, 9, 12, 15}))
	assert.Equal(t, 0, extrapolatePrev([]int{1, 3, 6, 10, 15, 21}))
	assert.Equal(t, 5, extrapolatePrev([]int{10, 13, 16, 21, 30, 45}))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 2, solvePartB(example))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 1066, solvePartB(input))
}
