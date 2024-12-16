package day15

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example1 = `########
#..O.O.#
##@.O..#
#...O..#
#.#.O..#
#...O..#
#......#
########

<^^>>>vv<v
>>v<<`

func TestParseExample(t *testing.T) {
	grid, movements := parse(example1)
	assert.Equal(t, "#..O.O.#", string(grid[1]))
	assert.Equal(t, "<^^>>>vv<v>>v<<", string(movements))
}

func TestSolveExample1(t *testing.T) {
	grid, movements := parse(example1)
	assert.Equal(t, 2028, solvePartA(grid, movements))
}

func TestSolvePartA(t *testing.T) {
	grid, movements := parse(utils.LoadString("input.txt"))
	assert.Equal(t, 1538871, solvePartA(grid, movements))
}
