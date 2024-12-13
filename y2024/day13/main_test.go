package day13

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = `Button A: X+94, Y+34
Button B: X+22, Y+67
Prize: X=8400, Y=5400

Button A: X+26, Y+66
Button B: X+67, Y+21
Prize: X=12748, Y=12176

Button A: X+17, Y+86
Button B: X+84, Y+37
Prize: X=7870, Y=6450

Button A: X+69, Y+23
Button B: X+27, Y+71
Prize: X=18641, Y=10279
`

func TestParse(t *testing.T) {
	assert.Equal(t, []Machine{{
		A: Coord{x: 94, y: 34}, B: Coord{x: 22, y: 67}, Prize: Coord{x: 8400, y: 5400},
	}, {
		A: Coord{x: 26, y: 66}, B: Coord{x: 67, y: 21}, Prize: Coord{x: 12748, y: 12176},
	}, {
		A: Coord{x: 17, y: 86}, B: Coord{x: 84, y: 37}, Prize: Coord{x: 7870, y: 6450},
	}, {
		A: Coord{x: 69, y: 23}, B: Coord{x: 27, y: 71}, Prize: Coord{x: 18641, y: 10279},
	}}, parseInput(example))
}

func TestSolveEquation(t *testing.T) {
	a, b := parseInput(example)[0].solve(0)
	assert.Equal(t, 80, a)
	assert.Equal(t, 40, b)

	a, b = parseInput(example)[1].solve(0)
	assert.Equal(t, -1, a)
	assert.Equal(t, -1, b)

	a, b = parseInput(example)[2].solve(0)
	assert.Equal(t, 38, a)
	assert.Equal(t, 86, b)
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 480, solve(parseInput(example), 0))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 29201, solve(parseInput(utils.LoadString("input.txt")), 0))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 875318608908, solve(parseInput(example), 10000000000000))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 104140871044942, solve(parseInput(utils.LoadString("input.txt")), 10000000000000))
}
