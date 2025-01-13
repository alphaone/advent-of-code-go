package day14

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/alphaone/advent/y2023/day14/rotate"
	"github.com/stretchr/testify/assert"
)

var example1 = []string{
	"O....#....",
	"O.OO#....#",
	".....##...",
	"OO.#O....O",
	".O.....O#.",
	"O.#..O.#.#",
	"..O..#O..O",
	".......O..",
	"#....###..",
	"#OO..#....",
}

func TestMoveLineWest(t *testing.T) {
	assert.Equal(t, []rune("OOOO....##"), MoveLineWest([]rune("OO.O.O..##")))
	assert.Equal(t, []rune("OOO.#OO..."), MoveLineWest([]rune("OO.O#O..O.")))
}

func TestMoveNorth(t *testing.T) {
	assert.Equal(t, []string{
		"OOOO.#.O..",
		"OO..#....#",
		"OO..O##..O",
		"O..#.OO...",
		"........#.",
		"..#....#.#",
		"..O..#.O.O",
		"..O.......",
		"#....###..",
		"#....#....",
	}, utils.AsStrings(rotate.RotateClockwise(MoveWest(rotate.RotateCounterClockwise(utils.AsRunes(example1))))))
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 136, solvePartA(utils.AsRunes(example1)))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 110128, solvePartA(utils.AsRunes(input)))
}

func TestCycle(t *testing.T) {
	expected := []string{
		".....#....",
		"....#...O#",
		"...OO##...",
		".OO#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#....",
		"......OOOO",
		"#...O###..",
		"#..OO#....",
	}
	assert.Equal(t, expected, utils.AsStrings(rotate.RotateClockwise(cycleThrough(rotate.RotateCounterClockwise(utils.AsRunes(example1)), 1))))

	expected = []string{
		".....#....",
		"....#...O#",
		".....##...",
		"..O#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#...O",
		".......OOO",
		"#..OO###..",
		"#.OOO#...O",
	}
	assert.Equal(t, expected, utils.AsStrings(rotate.RotateClockwise(cycleThrough(rotate.RotateCounterClockwise(utils.AsRunes(example1)), 2))))

	expected = []string{
		".....#....",
		"....#...O#",
		".....##...",
		"..O#......",
		".....OOO#.",
		".O#...O#.#",
		"....O#...O",
		".......OOO",
		"#...O###.O",
		"#.OOO#...O",
	}
	assert.Equal(t, expected, utils.AsStrings(rotate.RotateClockwise(cycleThrough(rotate.RotateCounterClockwise(utils.AsRunes(example1)), 3))))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 64, solvePartB(utils.AsRunes(example1), 1_000_000_000))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 103861, solvePartB(utils.AsRunes(input), 1_000_000_000))
}
