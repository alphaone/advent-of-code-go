package day12

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"RRRRIICCFF",
	"RRRRIICCCF",
	"VVRRRCCFFF",
	"VVRCCCJFFF",
	"VVVVCJJCFE",
	"VVIVCCJJEE",
	"VVIIICJJEE",
	"MIIIIIJJEE",
	"MIIISIJEEE",
	"MMMISSJEEE",
}

func TestNeighbors(t *testing.T) {
	g := parseGrid(example)
	assert.Equal(t, []Coord{{0, 3}, {2, 3}, {1, 2}}, g.findNeighors(Coord{1, 3}, 'R'))
}

func TestSolveExamplePartA(t *testing.T) {
	assert.Equal(t, 1930, solve(parseGrid(example)))
}

func TestSolvePartA(t *testing.T) {
	assert.Equal(t, 1359028, solve(parseGrid(utils.LoadStrings("input.txt"))))
}
