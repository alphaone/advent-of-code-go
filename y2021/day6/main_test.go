package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	example = []int{3, 4, 3, 1, 2}
	input   = []int{2, 5, 3, 4, 4, 5, 3, 2, 3, 3, 2, 2, 4, 2, 5, 4, 1, 1, 4, 4, 5, 1, 2, 1, 5, 2, 1, 5, 1, 1, 1, 2, 4, 3, 3, 1, 4, 2, 3, 4, 5, 1, 2, 5, 1, 2, 2, 5, 2, 4, 4, 1, 4, 5, 4, 2, 1, 5, 5, 3, 2, 1, 3, 2, 1, 4, 2, 5, 5, 5, 2, 3, 3, 5, 1, 1, 5, 3, 4, 2, 1, 4, 4, 5, 4, 5, 3, 1, 4, 5, 1, 5, 3, 5, 4, 4, 4, 1, 4, 2, 2, 2, 5, 4, 3, 1, 4, 4, 3, 4, 2, 1, 1, 5, 3, 3, 2, 5, 3, 1, 2, 2, 4, 1, 4, 1, 5, 1, 1, 2, 5, 2, 2, 5, 2, 4, 4, 3, 4, 1, 3, 3, 5, 4, 5, 4, 5, 5, 5, 5, 5, 4, 4, 5, 3, 4, 3, 3, 1, 1, 5, 2, 4, 5, 5, 1, 5, 2, 4, 5, 4, 2, 4, 4, 4, 2, 2, 2, 2, 2, 3, 5, 3, 1, 1, 2, 1, 1, 5, 1, 4, 3, 4, 2, 5, 3, 4, 4, 3, 5, 5, 5, 4, 1, 3, 4, 4, 2, 2, 1, 4, 1, 2, 1, 2, 1, 5, 5, 3, 4, 1, 3, 2, 1, 4, 5, 1, 5, 5, 1, 2, 3, 4, 2, 1, 4, 1, 4, 2, 3, 3, 2, 4, 1, 4, 1, 4, 4, 1, 5, 3, 1, 5, 2, 1, 1, 2, 3, 3, 2, 4, 1, 2, 1, 5, 1, 1, 2, 1, 2, 1, 2, 4, 5, 3, 5, 5, 1, 3, 4, 1, 1, 3, 3, 2, 2, 4, 3, 1, 1, 2, 4, 1, 1, 1, 5, 4, 2, 4, 3}
)

func TestAdvance(t *testing.T) {
	assert.Equal(t, []int{0, 1, 1, 2, 1, 0, 0, 0, 0}, ageBucketFromInput([]int{3, 4, 3, 1, 2}))
	// 2,3,2,0,1
	assert.Equal(t, []int{1, 1, 2, 1, 0, 0, 0, 0, 0}, Advance([]int{0, 1, 1, 2, 1, 0, 0, 0, 0}))
	// 1,2,1,6,0,8
	assert.Equal(t, []int{1, 2, 1, 0, 0, 0, 1, 0, 1}, Advance([]int{1, 1, 2, 1, 0, 0, 0, 0, 0}))
	// 0,1,0,5,6,7,8
	assert.Equal(t, []int{2, 1, 0, 0, 0, 1, 1, 1, 1}, Advance([]int{1, 2, 1, 0, 0, 0, 1, 0, 1}))
	// 6,0,6,4,5,6,7,8,8
	assert.Equal(t, []int{1, 0, 0, 0, 1, 1, 3, 1, 2}, Advance([]int{2, 1, 0, 0, 0, 1, 1, 1, 1}))
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 26, solvePartA(example, 18))
	assert.Equal(t, 5934, solvePartA(example, 80))
}

func TestSolvePartA(t *testing.T) {
	assert.Equal(t, 345387, solvePartA(input, 80))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 26984457539, solvePartA(example, 256))
}

func TestSolvePartB(t *testing.T) {
	assert.Equal(t, 1574445493136, solvePartA(input, 256))
}
