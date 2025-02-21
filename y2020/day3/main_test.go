package day3

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"..##.......",
	"#...#...#..",
	".#....#..#.",
	"..#.#...#.#",
	".#...##..#.",
	"..#.##.....",
	".#.#.#....#",
	".#........#",
	"#.##...#...",
	"#...##....#",
	".#..#...#.#",
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 7, solveA(parse(exampleInput)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 228, solveA(parse(utils.LoadStrings("input.txt"))))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 336, solveB(parse(exampleInput)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 6818112000, solveB(parse(utils.LoadStrings("input.txt"))))
}
