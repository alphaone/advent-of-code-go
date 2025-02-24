package day6

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleInput = `abc

a
b
c

ab
ac

a
a
a
a

b`

func TestParse(t *testing.T) {
	expected := []group{
		{[]rune{'a', 'b', 'c'}},
		{[]rune{'a'}, []rune{'b'}, []rune{'c'}},
		{[]rune{'a', 'b'}, []rune{'a', 'c'}},
		{[]rune{'a'}, []rune{'a'}, []rune{'a'}, []rune{'a'}},
		{[]rune{'b'}},
	}
	assert.Equal(t, expected, parse(exampleInput))
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 11, solveA(parse(exampleInput)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 6443, solveA(parse(utils.LoadString("input.txt"))))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 6, solveB(parse(exampleInput)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 3232, solveB(parse(utils.LoadString("input.txt"))))
}
