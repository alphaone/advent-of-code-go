package day18

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"R 6 (#70c710)",
	"D 5 (#0dc571)",
	"L 2 (#5713f0)",
	"D 2 (#d2c081)",
	"R 2 (#59c680)",
	"D 2 (#411b91)",
	"L 5 (#8ceee2)",
	"U 2 (#caa173)",
	"L 1 (#1b58a2)",
	"U 2 (#caa171)",
	"R 2 (#7807d2)",
	"U 3 (#a77fa3)",
	"L 2 (#015232)",
	"U 2 (#7a21e3)",
}

func TestParseInput(t *testing.T) {
	coords, perimeter := parseInputA(example)
	assert.Equal(t, []Pos{
		{0, 0},
		{6, 0},
		{6, 5},
		{4, 5},
		{4, 7},
		{6, 7},
		{6, 9},
		{1, 9},
		{1, 7},
		{0, 7},
		{0, 5},
		{2, 5},
		{2, 2},
		{0, 2},
		{0, 0},
	}, coords)
	assert.Equal(t, 38, perimeter)
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 62, SolvePartA(example))
}

func TestSolvePartA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 62500, SolvePartA(input))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 952408144115, SolvePartB(example))
}

func TestSolvePartB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 122109860712709, SolvePartB(input))
}
