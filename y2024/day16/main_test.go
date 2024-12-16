package day16

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"###############",
	"#.......#....E#",
	"#.#.###.#.###.#",
	"#.....#.#...#.#",
	"#.###.#####.#.#",
	"#.#.#.......#.#",
	"#.#.#####.###.#",
	"#...........#.#",
	"###.#.#####.#.#",
	"#...#.....#.#.#",
	"#.#.#.###.#.#.#",
	"#.....#...#.#.#",
	"#.###.#.#.#.#.#",
	"#S..#.....#...#",
	"###############",
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 7036, solvePartA(parse(example)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 72428, solvePartA(parse(utils.LoadStrings("input.txt"))))
}
