package day10

import (
	"strings"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/alphaone/advent/y2019/day10/vectors"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	s := strings.Split(".#..#\n.....\n#####\n....#\n...##", "\n")
	m := parseMap(s)

	assert.Equal(t, Map{
		{X: 1, Y: 0},
		{X: 4, Y: 0},
		{X: 0, Y: 2},
		{X: 1, Y: 2},
		{X: 2, Y: 2},
		{X: 3, Y: 2},
		{X: 4, Y: 2},
		{X: 4, Y: 3},
		{X: 3, Y: 4},
		{X: 4, Y: 4},
	}, m)
}

type countVisibleTestCases struct {
	name  string
	input string
	pov   Coord
	exp   int
}

func TestCountVisible(t *testing.T) {
	for _, tc := range []countVisibleTestCases{
		{name: "example 1", input: ".#..#\n.....\n#####\n....#\n...##", pov: Coord{3, 4}, exp: 8},
		{name: "example 2", input: ".#..#\n.....\n#####\n....#\n...##", pov: Coord{4, 4}, exp: 7},
		{name: "example 3", input: ".#..#\n.....\n#####\n....#\n...##", pov: Coord{0, 2}, exp: 6},

		{
			name: "example 4", pov: Coord{0, 0}, exp: 7,
			input: "#.........\n...A......\n...B..a...\n.EDCG....a\n..F.c.b...\n.....c....\n..efd.c.gb\n.......c..\n....f...c.\n...e..d..c",
		},
		{
			name: "example 5", pov: Coord{5, 8}, exp: 33,
			input: "......#.#.\n#..#.#....\n..#######.\n.#.#.###..\n.#..#.....\n..#....#.#\n#..#....#.\n.##.#..###\n##...#..#.\n.#....####",
		},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := parseMap(strings.Split(tc.input, "\n"))
			assert.Equal(t, tc.exp, len(m.Visible(tc.pov)))
		})
	}
}

func TestVisible(t *testing.T) {
	m := parseMap(strings.Split(".#..#\n.....\n#####\n....#\n...##", "\n"))
	assert.Equal(t, map[vectors.Vector]bool{
		{X: 1, Y: -4}:  true,
		{X: 1, Y: -2}:  true,
		{X: 1, Y: -1}:  true,
		{X: 1, Y: 0}:   true,
		{X: 0, Y: -1}:  true,
		{X: -1, Y: -1}: true,
		{X: -1, Y: -2}: true,
		{X: -3, Y: -2}: true,
	}, m.Visible(Coord{3, 4}))
}

type exampleSolveTestCases struct {
	name  string
	input string
	exp   int
}

func TestExampleSolve(t *testing.T) {
	for _, tc := range []exampleSolveTestCases{
		{name: "example 1", exp: 8, input: ".#..#\n.....\n#####\n....#\n...##"},
		{name: "example 2", exp: 33, input: "......#.#.\n#..#.#....\n..#######.\n.#.#.###..\n.#..#.....\n..#....#.#\n#..#....#.\n.##.#..###\n##...#..#.\n.#....####"},
		{name: "example 3", exp: 41, input: ".#..#..###\n####.###.#\n....###.#.\n..###.##.#\n##.##.#.#.\n....###..#\n..#.#..#.#\n#..#.#.###\n.##...##.#\n.....#.#.."},
	} {
		t.Run(tc.name, func(t *testing.T) {
			m := parseMap(strings.Split(tc.input, "\n"))
			assert.Equal(t, tc.exp, solve(m))
		})
	}
}

func TestSolve(t *testing.T) {
	s := utils.LoadStrings("input.txt")
	assert.Equal(t, 326, solve(parseMap(s)))
}

type vectorOrderingTestCases struct {
	name  string
	input map[vectors.Vector]bool
	exp   []vectors.Vector
}

func TestVectorOrdering(t *testing.T) {
	for _, tc := range []vectorOrderingTestCases{
		{
			name: "example 1",
			input: map[vectors.Vector]bool{
				{X: 0, Y: 1}:   true,
				{X: 1, Y: 0}:   true,
				{X: 0, Y: -1}:  true,
				{X: -1, Y: 0}:  true,
				{X: -1, Y: 1}:  true,
				{X: -1, Y: -1}: true,
				{X: 1, Y: 1}:   true,
				{X: 1, Y: -1}:  true,
			},
			exp: []vectors.Vector{
				{X: 0, Y: -1},
				{X: 1, Y: -1},
				{X: 1, Y: 0},
				{X: 1, Y: 1},
				{X: 0, Y: 1},
				{X: -1, Y: 1},
				{X: -1, Y: 0},
				{X: -1, Y: -1},
			},
		},
		{
			name: "example 2",
			input: map[vectors.Vector]bool{
				{X: 0, Y: -1}: true,
				{X: 1, Y: -3}: true,
				{X: 1, Y: -2}: true,
				{X: 1, Y: -1}: true,
				{X: 2, Y: -3}: true,
			},
			exp: []vectors.Vector{
				{X: 0, Y: -1},
				{X: 1, Y: -3},
				{X: 1, Y: -2},
				{X: 2, Y: -3},
				{X: 1, Y: -1},
			},
		},
	} {
		assert.Equal(t, tc.exp, sortVectors(tc.input))
	}
}

func TestExampleBOrdering(t *testing.T) {
	s := strings.Split(".#....#####...#..\n##...##.#####..##\n##...#...#.#####.\n..#.....X...###..\n..#.#.....#....##", "\n")
	m := parseMap(s)
	assert.Equal(t,
		[]vectors.Vector{
			{X: 0, Y: -1},
			{X: 1, Y: -3},
			{X: 1, Y: -2},
			{X: 2, Y: -3},
			{X: 1, Y: -1},
			{X: 3, Y: -2},
			{X: 2, Y: -1},
			{X: 3, Y: -1},
			{X: 7, Y: -2},
		},
		sortVectors(m.Visible(Coord{8, 3}))[:9])
}

func TestSolveB(t *testing.T) {
	s := utils.LoadStrings("input.txt")
	assert.Equal(t, 1623, solveB(parseMap(s), 200, Coord{22, 28}))
}
