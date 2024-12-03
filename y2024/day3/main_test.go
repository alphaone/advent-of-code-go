package day3

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleA = "xmul(2,4)%&mul[3,7]!@^do_not_mul(5,5)+mul(32,64]then(mul(11,8)mul(8,5))"

func TestParseMultiplicationA(t *testing.T) {
	res := parseMultiplicationsA(exampleA)
	expected := []Multiplication{
		{2, 4}, {5, 5}, {11, 8}, {8, 5},
	}
	assert.Equal(t, expected, res)
}

func TestExampleSolvePartA(t *testing.T) {
	assert.Equal(t, 161, solve(exampleA, parseMultiplicationsA))
}

func TestSolvePartA(t *testing.T) {
	input := utils.LoadString("input.txt")
	assert.Equal(t, 175015740, solve(input, parseMultiplicationsA))
}

var exampleB = "xmul(2,4)&mul[3,7]!^don't()_mul(5,5)+mul(32,64](mul(11,8)undo()?mul(8,5))"

func TestParseMultiplicationB(t *testing.T) {
	res := parseMultiplicationsB(exampleB)
	assert.Equal(t, []Multiplication{{2, 4}, {8, 5}}, res)

	res = parseMultiplicationsB2(exampleB)
	assert.Equal(t, []Multiplication{{2, 4}, {8, 5}}, res)
}

func TestExampleSolvePartB(t *testing.T) {
	assert.Equal(t, 48, solve(exampleB, parseMultiplicationsB))
	assert.Equal(t, 48, solve(exampleB, parseMultiplicationsB2))
}

func TestSolvePartB(t *testing.T) {
	input := utils.LoadString("input.txt")
	assert.Equal(t, 112272912, solve(input, parseMultiplicationsB))
	assert.Equal(t, 112272912, solve(input, parseMultiplicationsB2))
}
