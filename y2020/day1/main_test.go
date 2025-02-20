package day1

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleInput = []int{
	1721,
	979,
	366,
	299,
	675,
	1456,
}

func TestFind(t *testing.T) {
	a, b := findTwo(exampleInput)
	assert.Equal(t, 1721, a)
	assert.Equal(t, 299, b)
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 514579, solveA(exampleInput))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 436404, solveA(utils.LoadInts("input.txt")))
}

func TestFind3(t *testing.T) {
	a, b, c := findThree(exampleInput)
	assert.Equal(t, 979, a)
	assert.Equal(t, 366, b)
	assert.Equal(t, 675, c)
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 241861950, solveB(exampleInput))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 274879808, solveB(utils.LoadInts("input.txt")))
}
