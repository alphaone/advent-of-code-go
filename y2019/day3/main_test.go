package day3

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestParsePath(t *testing.T) {
	assert.Equal(t, []Segment{
		{start: Coord{x: 0, y: 0}, end: Coord{x: 8, y: 0}, length: 8},
		{start: Coord{x: 8, y: 0}, end: Coord{x: 8, y: 5}, length: 5},
		{start: Coord{x: 8, y: 5}, end: Coord{x: 3, y: 5}, length: 5},
		{start: Coord{x: 3, y: 5}, end: Coord{x: 3, y: 2}, length: 3},
	}, parsePath("R8,U5,L5,D3"))
	assert.Equal(t, []Segment{
		{start: Coord{x: 0, y: 0}, end: Coord{x: 0, y: 7}, length: 7},
		{start: Coord{x: 0, y: 7}, end: Coord{x: 6, y: 7}, length: 6},
		{start: Coord{x: 6, y: 7}, end: Coord{x: 6, y: 3}, length: 4},
		{start: Coord{x: 6, y: 3}, end: Coord{x: 2, y: 3}, length: 4},
	}, parsePath("U7,R6,D4,L4"))
}

func TestIntersections(t *testing.T) {
	assert.Equal(t, []Coord{
		{x: 0, y: 0}, {x: 6, y: 5}, {x: 3, y: 3},
	}, intersections([]Segment{
		{start: Coord{x: 0, y: 0}, end: Coord{x: 8, y: 0}, length: 8},
		{start: Coord{x: 8, y: 0}, end: Coord{x: 8, y: 5}, length: 5},
		{start: Coord{x: 8, y: 5}, end: Coord{x: 3, y: 5}, length: 5},
		{start: Coord{x: 3, y: 5}, end: Coord{x: 3, y: 2}, length: 3},
	}, []Segment{
		{start: Coord{x: 0, y: 0}, end: Coord{x: 0, y: 7}, length: 7},
		{start: Coord{x: 0, y: 7}, end: Coord{x: 6, y: 7}, length: 6},
		{start: Coord{x: 6, y: 7}, end: Coord{x: 6, y: 3}, length: 4},
		{start: Coord{x: 6, y: 3}, end: Coord{x: 2, y: 3}, length: 4},
	}))
}

func TestLength2(t *testing.T) {
	assert.Equal(t, 20, Coord{3, 3}.lengthAlongPath([]Segment{
		{start: Coord{x: 0, y: 0}, end: Coord{x: 8, y: 0}, length: 8},
		{start: Coord{x: 8, y: 0}, end: Coord{x: 8, y: 5}, length: 5},
		{start: Coord{x: 8, y: 5}, end: Coord{x: 3, y: 5}, length: 5},
		{start: Coord{x: 3, y: 5}, end: Coord{x: 3, y: 2}, length: 3},
	}))
	assert.Equal(t, 20, Coord{3, 3}.lengthAlongPath([]Segment{
		{start: Coord{x: 0, y: 0}, end: Coord{x: 0, y: 7}, length: 7},
		{start: Coord{x: 0, y: 7}, end: Coord{x: 6, y: 7}, length: 6},
		{start: Coord{x: 6, y: 7}, end: Coord{x: 6, y: 3}, length: 4},
		{start: Coord{x: 6, y: 3}, end: Coord{x: 2, y: 3}, length: 4},
	}))
	assert.Equal(t, 15, Coord{6, 5}.lengthAlongPath([]Segment{
		{start: Coord{x: 0, y: 0}, end: Coord{x: 8, y: 0}, length: 8},
		{start: Coord{x: 8, y: 0}, end: Coord{x: 8, y: 5}, length: 5},
		{start: Coord{x: 8, y: 5}, end: Coord{x: 3, y: 5}, length: 5},
		{start: Coord{x: 3, y: 5}, end: Coord{x: 3, y: 2}, length: 3},
	}))
	assert.Equal(t, 15, Coord{6, 5}.lengthAlongPath([]Segment{
		{start: Coord{x: 0, y: 0}, end: Coord{x: 0, y: 7}, length: 7},
		{start: Coord{x: 0, y: 7}, end: Coord{x: 6, y: 7}, length: 6},
		{start: Coord{x: 6, y: 7}, end: Coord{x: 6, y: 3}, length: 4},
		{start: Coord{x: 6, y: 3}, end: Coord{x: 2, y: 3}, length: 4},
	}))
}

func TestSolvePartA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 316, solvePartA(input))
}

func TestSolvePartB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 16368, solvePartB(input))
}
