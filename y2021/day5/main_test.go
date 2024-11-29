package day5

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"0,9 -> 5,9",
	"8,0 -> 0,8",
	"9,4 -> 3,4",
	"2,2 -> 2,1",
	"7,0 -> 7,4",
	"6,4 -> 2,0",
	"0,9 -> 2,9",
	"3,4 -> 1,4",
	"0,0 -> 8,8",
	"5,5 -> 8,2",
}

func TestParseLine(t *testing.T) {
	assert.Equal(t, Segment{
		Start: Pos{0, 9},
		End:   Pos{5, 9},
	}, ParseLine("0,9 -> 5,9"))
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 5, solvePartA(example))
}

func TestSolvePartA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 5124, solvePartA(input))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 12, solvePartB(example))
}

func TestSolvePartB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 19771, solvePartB(input))
}
