package day2

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var input = []string{
	"forward 5",
	"down 5",
	"forward 8",
	"up 3",
	"down 8",
	"forward 2 ",
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, Coord{15, 10, 0}, solvePartA(input))
	assert.Equal(t, 150, solvePartA(input).Score())
}

func TestParseLine(t *testing.T) {
	assert.Equal(t, Command{"forward", 5}, parseLine("forward 5"))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 1635930, solvePartA(input).Score())
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, Coord{15, 60, 10}, solvePartB(input))
	assert.Equal(t, 900, solvePartB(input).Score())
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 1781819478, solvePartB(input).Score())
}
