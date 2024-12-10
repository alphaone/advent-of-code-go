package day10

import (
	"errors"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleStr = []string{
	"89010123",
	"78121874",
	"87430965",
	"96549874",
	"45678903",
	"32019012",
	"01329801",
	"10456732",
}

type (
	ElevationMap [][]int
	Coord        struct{ l, r int }
)

func (c Coord) add(other Coord) Coord {
	return Coord{c.l + other.l, c.r + other.r}
}

func parseInput(input []string) ElevationMap {
	res := make(ElevationMap, 0)
	for _, line := range input {
		mapLine := make([]int, 0)
		for _, r := range line {
			mapLine = append(mapLine, int(r)-48)
		}
		res = append(res, mapLine)
	}
	return res
}

func (m ElevationMap) findCoords(lookingfor int) []Coord {
	res := make([]Coord, 0)
	for l, line := range m {
		for r, n := range line {
			if n == lookingfor {
				res = append(res, Coord{l, r})
			}
		}
	}
	return res
}

func (m ElevationMap) elevationAt(pos Coord) (int, error) {
	if pos.l < 0 || pos.r < 0 || pos.l >= len(m) || pos.r >= len((m)[0]) {
		return -1, errors.New("out of bounds")
	}

	return (m)[pos.l][pos.r], nil
}

// find ways to walk to summit, counting different paths
// recursive!
func (m ElevationMap) walk(pos Coord) map[Coord]int {
	curElevation, _ := m.elevationAt(pos)

	if curElevation == 9 {
		return map[Coord]int{pos: 1}
	}

	coordMap := make(map[Coord]int)
	dirs := []Coord{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	for _, d := range dirs {
		nextPos := pos.add(d)
		nextElevation, _ := m.elevationAt(nextPos)
		if nextElevation == curElevation+1 {
			for k, v := range m.walk(nextPos) {
				coordMap[k] += v
			}
		}
	}

	return coordMap
}

func solvePartA(m ElevationMap) int {
	res := 0
	for _, startingPos := range m.findCoords(0) {
		summits := m.walk(startingPos)
		res += len(summits)
	}
	return res
}

func solvePartB(m ElevationMap) int {
	res := 0
	for _, startingPos := range m.findCoords(0) {
		summits := m.walk(startingPos)
		for _, v := range summits {
			res += v
		}
	}
	return res
}

func TestParse(t *testing.T) {
	assert.Equal(t, ElevationMap{{1, 2, 3}, {4, 5, 6}}, parseInput([]string{"123", "456"}))
}

func TestFindTrailStarts(t *testing.T) {
	assert.Equal(t, []Coord{
		{l: 0, r: 2}, {l: 0, r: 4}, {l: 2, r: 4}, {l: 4, r: 6}, {l: 5, r: 2}, {l: 5, r: 5}, {l: 6, r: 0}, {l: 6, r: 6}, {l: 7, r: 1},
	}, parseInput(exampleStr).findCoords(0))
}

func TestWalk(t *testing.T) {
	assert.Equal(t, map[Coord]int{
		{l: 0, r: 1}: 4,
		{l: 3, r: 0}: 4,
		{l: 3, r: 4}: 4,
		{l: 4, r: 5}: 4,
		{l: 5, r: 4}: 4,
	}, parseInput(exampleStr).walk(Coord{0, 2}))
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 36, solvePartA(parseInput(exampleStr)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 514, solvePartA(parseInput(utils.LoadStrings("input.txt"))))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 81, solvePartB(parseInput(exampleStr)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 1162, solvePartB(parseInput(utils.LoadStrings("input.txt"))))
}
