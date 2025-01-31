package day10

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"noop",
	"addx 3",
	"addx -5",
}

func TestParse(t *testing.T) {
	exp := []instruction{
		{op: "noop", arg: 0},
		{op: "addx", arg: 3},
		{op: "addx", arg: -5},
	}
	assert.Equal(t, exp, parse(example))
}

func TestTick(t *testing.T) {
	c := cpu{registerX: 1}
	instrutions := parse(example)
	c.tick(&instrutions)
	assert.Equal(t, 1, c.registerX)
	c.tick(&instrutions)
	assert.Equal(t, 1, c.registerX)
	c.tick(&instrutions)
	assert.Equal(t, 4, c.registerX)
	c.tick(&instrutions)
	assert.Equal(t, 4, c.registerX)
	c.tick(&instrutions)
	assert.Equal(t, -1, c.registerX)
}

var largerexample = []string{
	"addx 15", "addx -11", "addx 6", "addx -3", "addx 5", "addx -1", "addx -8",
	"addx 13", "addx 4", "noop", "addx -1", "addx 5", "addx -1", "addx 5", "addx -1",
	"addx 5", "addx -1", "addx 5", "addx -1", "addx -35", "addx 1", "addx 24", "addx -19",
	"addx 1", "addx 16", "addx -11", "noop", "noop", "addx 21", "addx -15", "noop",
	"noop", "addx -3", "addx 9", "addx 1", "addx -3", "addx 8", "addx 1", "addx 5",
	"noop", "noop", "noop", "noop", "noop", "addx -36", "noop", "addx 1", "addx 7",
	"noop", "noop", "noop", "addx 2", "addx 6", "noop", "noop", "noop", "noop", "noop",
	"addx 1", "noop", "noop", "addx 7", "addx 1", "noop", "addx -13", "addx 13", "addx 7",
	"noop", "addx 1", "addx -33", "noop", "noop", "noop", "addx 2", "noop", "noop",
	"noop", "addx 8", "noop", "addx -1", "addx 2", "addx 1", "noop", "addx 17", "addx -9",
	"addx 1", "addx 1", "addx -3", "addx 11", "noop", "noop", "addx 1", "noop",
	"addx 1", "noop", "noop", "addx -13", "addx -19", "addx 1", "addx 3", "addx 26",
	"addx -30", "addx 12", "addx -1", "addx 3", "addx 1", "noop", "noop", "noop",
	"addx -9", "addx 18", "addx 1", "addx 2", "noop", "noop", "addx 9", "noop", "noop",
	"noop", "addx -1", "addx 2", "addx -37", "addx 1", "addx 3", "noop", "addx 15",
	"addx -21", "addx 22", "addx -6", "addx 1", "noop", "addx 2", "addx 1", "noop",
	"addx -10", "noop", "noop", "addx 20", "addx 1", "addx 2", "addx 2", "addx -6",
	"addx -11", "noop", "noop", "noop",
}

func TestSolveLargerExample(t *testing.T) {
	assert.Equal(t, 13140, solve(largerexample))
}

func TestSolve(t *testing.T) {
	assert.Equal(t, 11720, solve(utils.LoadStrings("input.txt")))
}

func TestSolveLargerExampleB(t *testing.T) {
	exp := `
##..##..##..##..##..##..##..##..##..##..40
###...###...###...###...###...###...###.80
####....####....####....####....####....120
#####.....#####.....#####.....#####.....160
######......######......######......####200
#######.......#######.......#######.....240
`
	assert.Equal(t, exp, "\n"+solveB(largerexample))
}

func TestSolveB(t *testing.T) {
	exp := `
####.###...##..###..####.###...##....##.40
#....#..#.#..#.#..#.#....#..#.#..#....#.80
###..#..#.#....#..#.###..#..#.#.......#.120
#....###..#....###..#....###..#.......#.160
#....#.#..#..#.#.#..#....#....#..#.#..#.200
####.#..#..##..#..#.####.#.....##...##..240
`
	assert.Equal(t, exp, "\n"+solveB(utils.LoadStrings("input.txt")))
}
