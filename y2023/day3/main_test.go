package day3

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var testInput = []string{
	"467..114..",
	"...*......",
	"..35..633.",
	"......#...",
	"617*......",
	".....+.58.",
	"..592.....",
	"......755.",
	"...$.*....",
	".664.598..",
}

func TestExamplePartA(t *testing.T) {
	assert.Equal(t, 4361, solvePartA(testInput))
}

func TestSolution(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 521515, solvePartA(input))
}

func TestParsePartNumbers(t *testing.T) {
	expected := []PartNumber{
		{No: 467, Start: Coord{0, 0}},
		{No: 114, Start: Coord{0, 5}},
		{No: 35, Start: Coord{2, 2}},
		{No: 633, Start: Coord{2, 6}},
		{No: 617, Start: Coord{4, 0}},
		{No: 58, Start: Coord{5, 7}},
		{No: 592, Start: Coord{6, 2}},
		{No: 755, Start: Coord{7, 6}},
		{No: 664, Start: Coord{9, 1}},
		{No: 598, Start: Coord{9, 5}},
	}
	assert.Equal(t, expected, parsePartNumbers(testInput))
}

func TestValidPartNumber(t *testing.T) {
	assert.True(t, PartNumber{No: 633, Start: Coord{2, 6}}.Valid(testInput))
	assert.True(t, PartNumber{No: 35, Start: Coord{2, 2}}.Valid(testInput))
	assert.False(t, PartNumber{No: 58, Start: Coord{5, 7}}.Valid(testInput))
}

func TestSuroundingBox(t *testing.T) {
	assert.Equal(t,
		Box{TopLeft: Coord{1, 5}, BottomRight: Coord{3, 9}},
		PartNumber{No: 633, Start: Coord{2, 6}}.SuroundingBox(testInput))
	assert.Equal(t,
		Box{TopLeft: Coord{1, 1}, BottomRight: Coord{3, 4}},
		PartNumber{No: 35, Start: Coord{2, 2}}.SuroundingBox(testInput))
	assert.Equal(t,
		Box{TopLeft: Coord{4, 6}, BottomRight: Coord{6, 9}},
		PartNumber{No: 58, Start: Coord{5, 7}}.SuroundingBox(testInput))
	assert.Equal(t,
		Box{TopLeft: Coord{5, 1}, BottomRight: Coord{7, 5}},
		PartNumber{No: 592, Start: Coord{6, 2}}.SuroundingBox(testInput))
}

func TestSurounding(t *testing.T) {
	assert.Equal(t, "......633..#...", PartNumber{No: 633, Start: Coord{2, 6}}.Surounding(testInput))
	assert.Equal(t, "..*..35.....", PartNumber{No: 35, Start: Coord{2, 2}}.Surounding(testInput))
	assert.Equal(t, ".....58.....", PartNumber{No: 58, Start: Coord{5, 7}}.Surounding(testInput))
	assert.Equal(t, "....+.592......", PartNumber{No: 592, Start: Coord{6, 2}}.Surounding(testInput))
}

func TestValidGear(t *testing.T) {
	parts := parsePartNumbers(testInput)
	assert.Equal(t, 16345, Coord{1, 3}.GearRatio(parts, testInput))
	assert.Equal(t, 451490, Coord{8, 5}.GearRatio(parts, testInput))
	assert.Equal(t, 0, Coord{4, 3}.GearRatio(parts, testInput))
}

func TestExamplePartB(t *testing.T) {
	assert.Equal(t, 467835, solvePartB(testInput))
}

func TestSolutionPartB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 69527306, solvePartB(input))
}
