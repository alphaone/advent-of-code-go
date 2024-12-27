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

var example2 = `#######
#...#.#
#.....#
#..OO@#
#..O..#
#.....#
#######

<vv<<^^<<^^`

var largeExample = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
`

func TestParseExample(t *testing.T) {
	grid, movements := Parse(example1)
	assert.Equal(t, "#..O.O.#", string(grid[1]))
	assert.Equal(t, "<^^>>>vv<v>>v<<", string(movements))
}

func TestParseBExample(t *testing.T) {
	grid, movements := ParseB(example1)
	assert.Equal(t, "####@...[]....##", string(grid[2]))
	assert.Equal(t, "<^^>>>vv<v>>v<<", string(movements))
}

func TestSolveExample1(t *testing.T) {
	grid, movements := Parse(example1)
	assert.Equal(t, 2028, solve(grid, movements))
}

func TestSolveLargeExampleA(t *testing.T) {
	grid, movements := Parse(largeExample)
	assert.Equal(t, 10092, solve(grid, movements))
}

func TestSolvePartA(t *testing.T) {
	grid, movements := Parse(utils.LoadString("input.txt"))
	assert.Equal(t, 1538871, solve(grid, movements))
}

func TestSolveSmallExampleB(t *testing.T) {
	grid, movements := ParseB(example2)
	assert.Equal(t, 618, solve(grid, movements))
}

func TestSolveLargeExampleB(t *testing.T) {
	grid, movements := ParseB(largeExample)
	assert.Equal(t, 9021, solve(grid, movements))
}

func TestSolvePartB(t *testing.T) {
	grid, movements := ParseB(utils.LoadString("input.txt"))
	assert.Equal(t, 1543338, solve(grid, movements))
}
