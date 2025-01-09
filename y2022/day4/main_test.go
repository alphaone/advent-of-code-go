package day4

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestParse(t *testing.T) {
	assert.Equal(t, Assignment{first: []int{2, 3, 4}, second: []int{6, 7, 8}}, parseLine("2-4,6-8"))
}

func TestFullyContains(t *testing.T) {
	assert.False(t, fullyContained(Assignment{first: []int{2, 3, 4}, second: []int{6, 7, 8}}))
	assert.True(t, fullyContained(Assignment{first: []int{2, 3, 4, 5, 6, 7, 8}, second: []int{3, 4, 5, 6, 7}}))
	assert.True(t, fullyContained(Assignment{first: []int{6}, second: []int{4, 5, 6}}))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 515, solveA(utils.LoadStrings("input.txt")))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 883, solveB(utils.LoadStrings("input.txt")))
}
