package day9

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleStr = "2333133121414131402"

func TestLayoutFromDense(t *testing.T) {
	assert.Equal(t, []int{
		0, 0, -1, -1, -1, 1, 1, 1, -1, -1, -1, 2, -1, -1, -1, 3, 3, 3, -1, 4, 4, -1, 5, 5, 5, 5, -1, 6, 6, 6, 6, -1, 7, 7, 7, -1, 8, 8, 8, 8, 9, 9,
	}, layoutFromDense(exampleStr))
}

func TestBlockLayoutFromDense(t *testing.T) {
	assert.Equal(t, []block{
		{id: 0, length: 2},
		{id: -1, length: 3},
		{id: 1, length: 3},
		{id: -1, length: 3},
		{id: 2, length: 1},
		{id: -1, length: 3},
		{id: 3, length: 3},
		{id: -1, length: 1},
		{id: 4, length: 2},
		{id: -1, length: 1},
		{id: 5, length: 4},
		{id: -1, length: 1},
		{id: 6, length: 4},
		{id: -1, length: 1},
		{id: 7, length: 3},
		{id: -1, length: 1},
		{id: 8, length: 4},
		{id: -1, length: 0},
		{id: 9, length: 2},
	}, blockLayoutFromDense(exampleStr))
}

func TestCleanup(t *testing.T) {
	layout := []int{0, -1, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, 2}
	cleanupStepA(layout)
	assert.Equal(t, []int{0, 2, -1, 1, 1, 1, -1, -1, -1, -1, 2, 2, 2, 2, -1}, layout)

	cleanupA(layout)
	assert.Equal(t, []int{0, 2, 2, 1, 1, 1, 2, 2, 2, -1, -1, -1, -1, -1, -1}, layout)
}

func TestCleanupB(t *testing.T) {
	layout := blockLayoutFromDense(exampleStr)
	layout = cleanupStepB(layout, len(layout))
	assert.Equal(t, []block{
		{id: 0, length: 2},
		{id: 9, length: 2},
		{id: -1, length: 1},
		{id: 1, length: 3},
		{id: -1, length: 3},
		{id: 2, length: 1},
		{id: -1, length: 3},
		{id: 3, length: 3},
		{id: -1, length: 1},
		{id: 4, length: 2},
		{id: -1, length: 1},
		{id: 5, length: 4},
		{id: -1, length: 1},
		{id: 6, length: 4},
		{id: -1, length: 1},
		{id: 7, length: 3},
		{id: -1, length: 1},
		{id: 8, length: 4},
		{id: -1, length: 0},
		{id: -1, length: 2},
	}, layout)

	layout = cleanupB(layout)
	assert.Equal(t, []block{
		{id: 0, length: 2},
		{id: 9, length: 2},
		{id: 2, length: 1},
		{id: 1, length: 3},
		{id: 7, length: 3},
		{id: -1, length: 1},
		{id: 4, length: 2},
		{id: -1, length: 1},
		{id: 3, length: 3},
		{id: -1, length: 1},
		{id: -1, length: 2},
		{id: -1, length: 1},
		{id: 5, length: 4},
		{id: -1, length: 1},
		{id: 6, length: 4},
		{id: -1, length: 1},
		{id: -1, length: 3},
		{id: -1, length: 1},
		{id: 8, length: 4},
		{id: -1, length: 0},
		{id: -1, length: 2},
	}, layout)
}

func TestChecksum(t *testing.T) {
	assert.Equal(t, 1928, checksumA([]int{0, 0, 9, 9, 8, 1, 1, 1, 8, 8, 8, 2, 7, 7, 7, 3, 3, 3, 6, 4, 4, 6, 5, 5, 5, 5, 6, 6, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1, -1}))
	assert.Equal(t, 2858, checksumB([]block{
		{0, 2},
		{9, 2},
		{2, 1},
		{1, 3},
		{7, 3},
		{-1, 1},
		{4, 2},
		{-1, 1},
		{3, 3},
		{-1, 1},
		{-1, 2},
		{-1, 1},
		{5, 4},
		{-1, 1},
		{6, 4},
		{-1, 5},
		{8, 4},
	}))
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 1928, solvePartA(exampleStr))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 6366665108136, solvePartA(utils.LoadString("input.txt")))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 2858, solvePartB(exampleStr))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 6398065450842, solvePartB(utils.LoadString("input.txt")))
}
