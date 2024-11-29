package day23

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"#.#####################",
	"#.......#########...###",
	"#######.#########.#.###",
	"###.....#.>.>.###.#.###",
	"###v#####.#v#.###.#.###",
	"###.>...#.#.#.....#...#",
	"###v###.#.#.#########.#",
	"###...#.#.#.......#...#",
	"#####.#.#.#######.#.###",
	"#.....#.#.#.......#...#",
	"#.#####.#.#.#########v#",
	"#.#...#...#...###...>.#",
	"#.#.#v#######v###.###v#",
	"#...#.>.#...>.>.#.###.#",
	"#####v#.#.###v#.#.###.#",
	"#.....#...#...#.#.#...#",
	"#.#########.###.#.#.###",
	"#...###...#...#...#.###",
	"###.###.#.###v#####v###",
	"#...#...#.#.>.>.#.>.###",
	"#.###.###.#.###.#.#v###",
	"#.....###...###...#...#",
	"#####################.#",
}

func TestFindStartAndEnd(t *testing.T) {
	assert.Equal(t, Pos{Line: 0, Col: 1}, FindStart(example))
	assert.Equal(t, Pos{Line: 22, Col: 21}, FindEnd(example))
}

func TestFindPointsOfInterest(t *testing.T) {
	assert.Equal(t, []Pos{
		{Line: 0, Col: 1},
		{Line: 22, Col: 21},
		{Line: 3, Col: 11},
		{Line: 5, Col: 3},
		{Line: 11, Col: 21},
		{Line: 13, Col: 5},
		{Line: 13, Col: 13},
		{Line: 19, Col: 13},
		{Line: 19, Col: 19},
	}, FindPointsOfInterest(example))
}

func TestBuildGraph(t *testing.T) {
	graph := BuildGraph(FindPointsOfInterest(example), example, possibleDirectionsA)
	assert.Equal(t, map[Pos]int{{5, 3}: 15}, graph[Pos{0, 1}])
	assert.Equal(t, map[Pos]int{{3, 11}: 22, {13, 5}: 22}, graph[Pos{5, 3}])
	assert.Equal(t, 9, len(graph))
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 94, solvePartA(example))
}

func TestSoltionPartA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 2106, solvePartA(input))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 154, solvePartB(example))
}

func TestSoltionPartB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 6350, solvePartB(input))
}
