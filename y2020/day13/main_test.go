package day13

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = `939
7,13,x,x,59,x,31,19`

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 295, solveA(parse(example)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 2845, solveA(parse(utils.LoadString("input.txt"))))
}

func TestCombineBus(t *testing.T) {
	assert.Equal(t, bus{35, 14}, combineBus(bus{7, 0}, bus{5, 1}))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 1068781, solveB(parseB(example)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 487905974205117, solveB(parseB(utils.LoadString("input.txt"))))
}
