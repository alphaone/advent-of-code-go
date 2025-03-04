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
