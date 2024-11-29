package day11

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestCount(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	res := 0
	for _, line := range input {
		for _, char := range line {
			if char == '#' {
				res++
			}
		}
	}
	assert.Equal(t, 442, res)
}

var example = []string{
	"...#......",
	".......#..",
	"#.........",
	"..........",
	"......#...",
	".#........",
	".........#",
	"..........",
	".......#..",
	"#...#.....",
}

func TestDistance(t *testing.T) {
	galaxies := Galaxies(example)
	lineRifts := []int{3, 7}
	colRifts := []int{2, 5, 8}
	assert.Equal(t, 9, Distance(2, lineRifts, colRifts, galaxies[4], galaxies[8]))
	assert.Equal(t, 15, Distance(2, lineRifts, colRifts, galaxies[0], galaxies[6]))
	assert.Equal(t, 17, Distance(2, lineRifts, colRifts, galaxies[2], galaxies[5]))

	assert.Equal(t, 6, Distance(2, lineRifts, colRifts, galaxies[0], galaxies[1]))
	assert.Equal(t, 15, Distance(2, lineRifts, colRifts, galaxies[0], galaxies[5]))
	assert.Equal(t, 19, Distance(2, lineRifts, colRifts, galaxies[1], galaxies[7]))

	assert.Equal(t, 11, Distance(3, lineRifts, colRifts, galaxies[4], galaxies[8]))

	assert.Equal(t, 4, Distance(1, lineRifts, colRifts, galaxies[7], galaxies[8]))
	assert.Equal(t, 5, Distance(2, lineRifts, colRifts, galaxies[7], galaxies[8]))
	assert.Equal(t, 6, Distance(3, lineRifts, colRifts, galaxies[7], galaxies[8]))
	assert.Equal(t, 13, Distance(10, lineRifts, colRifts, galaxies[7], galaxies[8]))
}

func TestPairs(t *testing.T) {
	galaxies := Pairs(Galaxies(example))
	assert.Equal(t, 36, len(galaxies))
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 374, solvePart(example, 2))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 9974721, solvePart(input, 2))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 1030, solvePart(example, 10))
	assert.Equal(t, 8410, solvePart(example, 100))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 702770569197, solvePart(input, 1_000_000))
}
