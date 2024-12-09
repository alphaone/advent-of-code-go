package day8

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleStr = []string{
	"............",
	"........0...",
	".....0......",
	".......0....",
	"....0.......",
	"......A.....",
	"............",
	"............",
	"........A...",
	".........A..",
	"............",
	"............",
}

func TestGroupFrequencies(t *testing.T) {
	expected := map[rune][]Coord{
		'0': {{1, 8}, {2, 5}, {3, 7}, {4, 4}},
		'A': {{5, 6}, {8, 8}, {9, 9}},
	}
	assert.Equal(t, expected, parseInput(exampleStr).groupFrequencies())
}

func TestAntinodes(t *testing.T) {
	assert.Equal(t, []Coord{{1, 3}, {7, 6}}, antinodes([]Coord{{3, 4}, {5, 5}}, Coord{10, 10}))

	freqs := parseInput(exampleStr).groupFrequencies()
	assert.Equal(t, []Coord{
		{2, 4}, {1, 3}, {7, 7}, {10, 10},
	}, antinodes(freqs['A'], Coord{11, 11}))
	assert.Equal(t, []Coord{
		{3, 2}, {1, 3}, {0, 6}, {5, 6}, {4, 9}, {2, 10}, {7, 0}, {6, 3}, {5, 1},
	}, antinodes(freqs['0'], Coord{11, 11}))
}

func TestExampleSolvePartA(t *testing.T) {
	assert.Equal(t, 14, solve(parseInput(exampleStr), antinodes))
}

func TestSolvePartA(t *testing.T) {
	assert.Equal(t, 295, solve(parseInput(utils.LoadStrings("input.txt")), antinodes))
}

func TestAntinodes2(t *testing.T) {
	assert.Equal(t, []Coord{
		{0, 0}, {0, 0}, {1, 3}, {2, 6}, {3, 9}, {1, 3}, {0, 5}, {2, 1}, {4, 2}, {6, 3}, {8, 4}, {10, 5}, {2, 1},
	}, antinodes2([]Coord{{0, 0}, {1, 3}, {2, 1}}, Coord{11, 11}))
}

func TestExampleSolvePartB(t *testing.T) {
	assert.Equal(t, 34, solve(parseInput(exampleStr), antinodes2))
}

func TestSolvePartB(t *testing.T) {
	assert.Equal(t, 1034, solve(parseInput(utils.LoadStrings("input.txt")), antinodes2))
}
