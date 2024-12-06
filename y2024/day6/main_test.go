package day6

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = Grid(utils.AsRunes([]string{
	"....#.....",
	".........#",
	"..........",
	"..#.......",
	".......#..",
	"..........",
	".#..^.....",
	"........#.",
	"#.........",
	"......#...",
}))

func TestFindStart(t *testing.T) {
	expected := Coord{6, 4}
	assert.Equal(t, expected, example.findCoord('^'))
}

func TestCharAtPos(t *testing.T) {
	val, err := example.charAtPos(Coord{0, 4})
	assert.NoError(t, err)
	assert.Equal(t, '#', val)

	val, err = example.charAtPos(Coord{0, 6})
	assert.NoError(t, err)
	assert.Equal(t, '.', val)

	val, err = example.charAtPos(Coord{-1, 4})
	assert.Error(t, err)
}

func TestWalk(t *testing.T) {
	pos := Coord{6, 4}
	dir := Coord{-1, 0}
	for i := 0; i < 13; i++ {
		pos, dir = example.walk(pos, dir)
	}

	assert.Equal(t, Coord{l: 3, r: 8}, pos)
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 41, solvePartA(example))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 5453, solvePartA(Grid(utils.AsRunes(utils.LoadStrings("input.txt")))))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 6, solvePartB(example))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 2188, solvePartB(Grid(utils.AsRunes(utils.LoadStrings("input.txt")))))
}

func TestIsLoop(t *testing.T) {
	assert.Equal(t, false, example.isLoop(Coord{6, 4}, Coord{-1, 0}))

	grid := example.alterAt(Coord{6, 3}, '#')
	assert.Equal(t, true, grid.isLoop(Coord{6, 4}, Coord{-1, 0}))
}
