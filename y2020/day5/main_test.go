package day5

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseSeat(t *testing.T) {
	assert.Equal(t, 357, parseSeat("FBFBBFFRLR"))
	assert.Equal(t, 567, parseSeat("BFFFBBFRRR"))
	assert.Equal(t, 119, parseSeat("FFFBBBFRRR"))
	assert.Equal(t, 820, parseSeat("BBFFBBFRLL"))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 858, solveA(utils.LoadStrings("input.txt")))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 557, solveB(utils.LoadStrings("input.txt")))
}
