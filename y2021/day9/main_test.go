package day9

import (
	"slices"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/alphaone/advent/utils/sliceutils"
	"github.com/stretchr/testify/assert"
)

type (
	heightmap [][]int
	coord     struct{ row, col int }
)

func (h heightmap) neighbors(c coord) []int {
	result := []int{}
	for _, n := range h.neighborCoords(c) {
		result = append(result, h[n.row][n.col])
	}
	return result
}

func (h heightmap) neighborCoords(c coord) []coord {
	res := make([]coord, 0)
	for _, d := range []coord{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		n := coord{c.row + d.row, c.col + d.col}
		if n.row >= 0 && n.row < len(h) && n.col >= 0 && n.col < len(h[0]) {
			res = append(res, n)
		}
	}
	return res
}

func parse(input []string) heightmap {
	result := make([][]int, len(input))
	for i, line := range input {
		result[i] = make([]int, len(line))
		for j, c := range line {
			result[i][j] = int(c - '0')
		}
	}
	return result
}

func (h heightmap) lowpoints() []coord {
	result := []coord{}

	for i, row := range h {
		for j, cell := range row {
			nb := h.neighbors(coord{i, j})
			if sliceutils.All(nb, func(a int) bool { return a > cell }) {
				result = append(result, coord{i, j})
			}
		}
	}

	return result
}

func (h heightmap) flood(pos coord, acc *[]coord) {
	if h[pos.row][pos.col] == 9 {
		return
	}

	*acc = append(*acc, pos)

	neighbors := h.neighborCoords(pos)
	for _, n := range neighbors {
		if !slices.Contains(*acc, n) {
			h.flood(n, acc)
		}
	}
}

func solveA(input heightmap) int {
	riskLevel := 0
	for _, lp := range input.lowpoints() {
		riskLevel += input[lp.row][lp.col] + 1
	}
	return riskLevel
}

func solveB(input heightmap) int {
	basins := []int{}
	for _, lp := range input.lowpoints() {
		basin := []coord{}
		input.flood(lp, &basin)
		basins = append(basins, len(basin))
	}

	slices.SortFunc(basins, func(a, b int) int { return b - a })
	res := 1
	for i := range 3 {
		res *= basins[i]
	}
	return res
}

var example = []string{
	"2199943210",
	"3987894921",
	"9856789892",
	"8767896789",
	"9899965678",
}

func TestParse(t *testing.T) {
	assert.Equal(t, []int{2, 1, 9, 9, 9, 4, 3, 2, 1, 0}, parse(example)[0])
}

func TestNeighbors(t *testing.T) {
	assert.Equal(t, []int{1, 3}, parse(example).neighbors(coord{0, 0}))
	assert.Equal(t, []int{9, 5, 7, 9}, parse(example).neighbors(coord{2, 1}))
}

func TestLowPoints(t *testing.T) {
	assert.Equal(t, []coord{{0, 1}, {0, 9}, {2, 2}, {4, 6}}, parse(example).lowpoints())
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 15, solveA(parse(example)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 468, solveA(parse(utils.LoadStrings("input.txt"))))
}

func TestFlood(t *testing.T) {
	basin := []coord{}
	parse(example).flood(coord{0, 1}, &basin)
	assert.Equal(t, []coord{{0, 1}, {0, 0}, {1, 0}}, basin)
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 1134, solveB(parse(example)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 1280496, solveB(parse(utils.LoadStrings("input.txt"))))
}
