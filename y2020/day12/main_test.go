package day12

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"F10",
	"N3",
	"F7",
	"R90",
	"F11",
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 25, solveA(parseInput(exampleInput)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 1533, solveA(parseInput(utils.LoadStrings("input.txt"))))
}
