package day1

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestExampleA(t *testing.T) {
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	assert.Equal(t, 7, solvePartA(input))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadInts("input.txt")
	assert.Equal(t, 1559, solvePartA(input))
}

func TestExampleB(t *testing.T) {
	input := []int{
		199,
		200,
		208,
		210,
		200,
		207,
		240,
		269,
		260,
		263,
	}
	assert.Equal(t, 5, solvePartB(input))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadInts("input.txt")
	assert.Equal(t, 1600, solvePartB(input))
}
