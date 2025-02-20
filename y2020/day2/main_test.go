package day2

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert.Equal(t, Line{1, 3, 'a', "abcde"}, Parse("1-3 a: abcde"))
}

func TestIsValidA(t *testing.T) {
	assert.True(t, Line{1, 3, 'a', "abcde"}.IsValidA())
	assert.False(t, Line{1, 3, 'b', "cdefg"}.IsValidA())
	assert.True(t, Line{2, 9, 'c', "ccccccccc"}.IsValidA())
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 586, solveA(utils.LoadStrings("input.txt")))
}

func TestIsValidB(t *testing.T) {
	assert.True(t, Line{1, 3, 'a', "abcde"}.IsValidB())
	assert.False(t, Line{1, 3, 'b', "cdefg"}.IsValidB())
	assert.False(t, Line{2, 9, 'c', "ccccccccc"}.IsValidB())
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 352, solveB(utils.LoadStrings("input.txt")))
}
