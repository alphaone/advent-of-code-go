package day2

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert.Equal(t, LineA{1, 3, 'a', "abcde"}, parseA("1-3 a: abcde"))
	assert.Equal(t, LineB{1, 3, 'a', "abcde"}, parseB("1-3 a: abcde"))
}

func TestIsValidA(t *testing.T) {
	assert.True(t, LineA{1, 3, 'a', "abcde"}.IsValid())
	assert.False(t, LineA{1, 3, 'b', "cdefg"}.IsValid())
	assert.True(t, LineA{2, 9, 'c', "ccccccccc"}.IsValid())
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 586, solve(parseA, utils.LoadStrings("input.txt")))
}

func TestIsValidB(t *testing.T) {
	assert.True(t, LineB{1, 3, 'a', "abcde"}.IsValid())
	assert.False(t, LineB{1, 3, 'b', "cdefg"}.IsValid())
	assert.False(t, LineB{2, 9, 'c', "ccccccccc"}.IsValid())
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 352, solve(parseB, utils.LoadStrings("input.txt")))
}
