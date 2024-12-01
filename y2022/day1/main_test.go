package day1

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleA = `1000
2000
3000

4000

5000
6000

7000
8000
9000

10000
`

func TestExampleA(t *testing.T) {
	assert.Equal(t, 24000, solvePartA(exampleA))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadString("input.txt")
	assert.Equal(t, 70296, solvePartA(input))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 45000, solvePartB(exampleA))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadString("input.txt")
	assert.Equal(t, 205381, solvePartB(input))
}

func TestParse(t *testing.T) {
	res := parseInput(exampleA)
	expected := CaloriesByElf{
		[]int{1000, 2000, 3000},
		[]int{4000},
		[]int{5000, 6000},
		[]int{7000, 8000, 9000},
		[]int{10000},
	}

	assert.Equal(t, expected, res)
}

func TestSumByElf(t *testing.T) {
	res := sumByElf(parseInput(exampleA))
	expected := []int{6000, 4000, 11000, 24000, 10000}

	assert.Equal(t, expected, res)
}
