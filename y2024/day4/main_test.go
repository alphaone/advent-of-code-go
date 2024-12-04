package day4

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = Grid{
	"MMMSXXMASM",
	"MSAMXMSMSA",
	"AMXSXMAAMM",
	"MSAMASMSMX",
	"XMASAMXAMM",
	"XXAMMXXAMA",
	"SMSMSASXSS",
	"SAXAMASAAA",
	"MAMMMXMMMM",
	"MXMXAXMASX",
}

func TestFindStart(t *testing.T) {
	expected := []Coord{{0, 4}, {0, 5}, {1, 4}}
	assert.Equal(t, expected, example.allPos('X')[:3])
}

func TestCharAtPos(t *testing.T) {
	val, err := example.charAtPos(Coord{0, 4})
	assert.NoError(t, err)
	assert.Equal(t, 'X', val)

	val, err = example.charAtPos(Coord{-1, 4})
	assert.Error(t, err)
}

func TestWordFromPos(t *testing.T) {
	val := example.wordFromPos(Coord{0, 4}, Coord{1, 1}, 3)
	assert.Equal(t, "MAS", val)
}

func TestCountMas(t *testing.T) {
	assert.Equal(t, 0, example.countMas(Coord{0, 3}))
	assert.Equal(t, 1, example.countMas(Coord{0, 4}))
	assert.Equal(t, 1, example.countMas(Coord{0, 5}))
	assert.Equal(t, 2, example.countMas(Coord{4, 6}))
}

func TestSolveExampleA(t *testing.T) {
	res := solvePartA(example)
	assert.Equal(t, 18, res)
}

func TestSolveA(t *testing.T) {
	grid := Grid(utils.LoadStrings("input.txt"))
	res := solvePartA(grid)
	assert.Equal(t, 2504, res)
}

func TestSolveExampleB(t *testing.T) {
	res := solvePartB(example)
	assert.Equal(t, 9, res)
}

func TestSolveB(t *testing.T) {
	grid := Grid(utils.LoadStrings("input.txt"))
	res := solvePartB(grid)
	assert.Equal(t, 1923, res)
}
