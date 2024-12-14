package day14

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"p=0,4 v=3,-3",
	"p=6,3 v=-1,-3",
	"p=10,3 v=-1,2",
	"p=2,0 v=2,-1",
	"p=0,0 v=1,3",
	"p=3,0 v=-2,-2",
	"p=7,6 v=-1,-3",
	"p=3,0 v=-1,-2",
	"p=9,3 v=2,3",
	"p=7,3 v=-1,2",
	"p=2,4 v=2,-3",
	"p=9,5 v=-3,-3",
}

func TestParse(t *testing.T) {
	assert.Equal(t, []Robot{{
		pos: Coord{X: 0, Y: 4}, v: Coord{X: 3, Y: -3},
	}, {
		pos: Coord{X: 6, Y: 3}, v: Coord{X: -1, Y: -3},
	}, {
		pos: Coord{X: 10, Y: 3}, v: Coord{X: -1, Y: 2},
	}, {
		pos: Coord{X: 2, Y: 0}, v: Coord{X: 2, Y: -1},
	}, {
		pos: Coord{X: 0, Y: 0}, v: Coord{X: 1, Y: 3},
	}, {
		pos: Coord{X: 3, Y: 0}, v: Coord{X: -2, Y: -2},
	}, {
		pos: Coord{X: 7, Y: 6}, v: Coord{X: -1, Y: -3},
	}, {
		pos: Coord{X: 3, Y: 0}, v: Coord{X: -1, Y: -2},
	}, {
		pos: Coord{X: 9, Y: 3}, v: Coord{X: 2, Y: 3},
	}, {
		pos: Coord{X: 7, Y: 3}, v: Coord{X: -1, Y: 2},
	}, {
		pos: Coord{X: 2, Y: 4}, v: Coord{X: 2, Y: -3},
	}, {
		pos: Coord{X: 9, Y: 5}, v: Coord{X: -3, Y: -3},
	}}, Parse(example))
}

func TestMove(t *testing.T) {
	assert.Equal(t, Coord{5, 5}, Coord{0, 0}.move(Coord{5, 5}, Coord{10, 10}, 1))
	assert.Equal(t, Coord{9, 9}, Coord{0, 0}.move(Coord{-1, -1}, Coord{10, 10}, 1))
	assert.Equal(t, Coord{0, 9}, Coord{8, 7}.move(Coord{2, 2}, Coord{10, 10}, 1))

	assert.Equal(t, Coord{4, 1}, Coord{2, 4}.move(Coord{2, -3}, Coord{11, 7}, 1))
	assert.Equal(t, Coord{2, 4}, Coord{2, 4}.move(Coord{2, -3}, Coord{11, 7}, 0))
	assert.Equal(t, Coord{1, 3}, Coord{2, 4}.move(Coord{2, -3}, Coord{11, 7}, 5))
}

func TestSolveExamplePartA(t *testing.T) {
	assert.Equal(t, 12, solvePartA(Parse(example), Coord{11, 7}))
}

func TestSolvePartA(t *testing.T) {
	assert.Equal(t, 8179, SolvePartB(Parse(utils.LoadStrings("input.txt")), Coord{101, 103}))
}
