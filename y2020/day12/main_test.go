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

func TestRotate(t *testing.T) {
	assert.Equal(t, coord{0, 2}, rotateWaypoint(coord{-2, 0}, 90))
	assert.Equal(t, coord{0, -2}, rotateWaypoint(coord{-2, 0}, -90))

	assert.Equal(t, coord{10, 4}, rotateWaypoint(coord{-4, 10}, 90))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 286, solveB(parseInput(exampleInput)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 25235, solveB(parseInput(utils.LoadStrings("input.txt"))))
}
