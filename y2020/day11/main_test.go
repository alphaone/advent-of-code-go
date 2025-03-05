package day11

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleInput = []string{
	"L.LL.LL.LL",
	"LLLLLLL.LL",
	"L.L.L..L..",
	"LLLL.LL.LL",
	"L.LL.LL.LL",
	"L.LLLLL.LL",
	"..L.L.....",
	"LLLLLLLLLL",
	"L.LLLLLL.L",
	"L.LLLLL.LL",
}

func TestCountNeighbors(t *testing.T) {
	input := []string{
		"#.##.##.##",
		"#######.##",
		"#.#.#..#..",
		"####.##.##",
		"#.##.##.##",
		"#.#####.##",
		"..#.#.....",
		"##########",
		"#.######.#",
		"#.#####.##",
	}
	g := parseInput(input)
	assert.Equal(t, 2, g.countOccupiedNeighbors(0, 0))
}

func TestCountVisible(t *testing.T) {
	input := []string{
		".......#.",
		"...#.....",
		".#.......",
		".........",
		"..#L....#",
		"....#....",
		".........",
		"#........",
		"...#.....",
	}
	g := parseInput(input)
	assert.Equal(t, 8, g.countVisibleOccupiedNeighbors(4, 3))

	input = []string{
		".............",
		".L.L.#.#.#.#.",
		".............",
	}
	g = parseInput(input)
	assert.Equal(t, 0, g.countVisibleOccupiedNeighbors(1, 1))
}

func TestStep(t *testing.T) {
	input := parseInput(exampleInput)
	expected := `
#.##.##.##
#######.##
#.#.#..#..
####.##.##
#.##.##.##
#.#####.##
..#.#.....
##########
#.######.#
#.#####.##
`

	step1 := input.step()
	assert.Equal(t, expected, "\n"+step1.String())

	expected = `
#.LL.L#.##
#LLLLLL.L#
L.L.L..L..
#LLL.LL.L#
#.LL.LL.LL
#.LLLL#.##
..L.L.....
#LLLLLLLL#
#.LLLLLL.L
#.#LLLL.##
`

	step2 := step1.step()
	assert.Equal(t, expected, "\n"+step2.String())
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 37, solveA(parseInput(exampleInput)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 2283, solveA(parseInput(utils.LoadStrings("input.txt"))))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 26, solveB(parseInput(exampleInput)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 2054, solveB(parseInput(utils.LoadStrings("input.txt"))))
}
