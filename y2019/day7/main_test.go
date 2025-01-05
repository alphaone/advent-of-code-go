package day7

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestCombinations(t *testing.T) {
	res := combinations([]int{1, 2, 3})
	assert.Equal(t, [][]int{{1, 2, 3}, {1, 3, 2}, {2, 1, 3}, {2, 3, 1}, {3, 1, 2}, {3, 2, 1}}, res)
}

func TestExample(t *testing.T) {
	p := parseInput(`3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0`)
	assert.Equal(t, 43210, solve(p, []int{0, 1, 2, 3, 4}))

	p = parseInput(`3,23,3,24,1002,24,10,24,1002,23,-1,23,101,5,23,23,1,24,23,23,4,23,99,0,0`)
	assert.Equal(t, 54321, solve(p, []int{0, 1, 2, 3, 4}))

	p = parseInput(`3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33,1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0`)
	assert.Equal(t, 65210, solve(p, []int{0, 1, 2, 3, 4}))
}

func TestSolvePartA(t *testing.T) {
	p := parseInput(utils.LoadString("input.txt"))
	assert.Equal(t, 18812, solve(p, []int{0, 1, 2, 3, 4}))
}

func TestExampleB(t *testing.T) {
	p := parseInput(`3,26,1001,26,-4,26,3,27,1002,27,2,27,1,27,26,27,4,27,1001,28,-1,28,1005,28,6,99,0,0,5`)
	assert.Equal(t, 139629729, solve(p, []int{5, 6, 7, 8, 9}))

	p = parseInput(`3,52,1001,52,-5,52,3,53,1,52,56,54,1007,54,5,55,1005,55,26,1001,54,-5,54,1105,1,12,1,53,54,53,1008,54,0,55,1001,55,1,55,2,53,55,53,4,53,1001,56,-1,56,1005,56,6,99,0,0,0,0,10`)
	assert.Equal(t, 18216, solve(p, []int{5, 6, 7, 8, 9}))
}

func TestSolvePartB(t *testing.T) {
	p := parseInput(utils.LoadString("input.txt"))
	assert.Equal(t, 25534964, solve(p, []int{5, 6, 7, 8, 9}))
}
