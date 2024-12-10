package day5

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseOp(t *testing.T) {
	op := parseOp(1002)
	assert.Equal(t, 2, op)
	paramModes := parseParamModes(1002)
	assert.Equal(t, map[int]int{0: 0, 1: 1}, paramModes)

	paramModes = parseParamModes(10002)
	assert.Equal(t, map[int]int{0: 0, 1: 0, 2: 1}, paramModes)

	assert.Equal(t, true, x(10, 0))
	assert.Equal(t, false, x(10, 1))
	assert.Equal(t, true, x(010, 2))
	assert.Equal(t, false, x(100, 2))
}

func x(pMode int, i int) bool {
	return pMode/intPow(10, i)%10 == 0
}

func intPow(x int, y int) int {
	ans := 1
	for i := 0; i < y; i++ {
		ans = ans * x
	}
	return ans
}

func TestSolvePartA(t *testing.T) {
	p := parseInput(utils.LoadString("input.txt"))
	p.input = 1
	assert.Equal(t, 7259358, solve(p))
}

func TestSolveExamplePartB(t *testing.T) {
	p := parseInput("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")
	p.input = 7
	assert.Equal(t, 999, solve(p))
	p = parseInput("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")
	p.input = 8
	assert.Equal(t, 1000, solve(p))
	p = parseInput("3,21,1008,21,8,20,1005,20,22,107,8,21,20,1006,20,31,1106,0,36,98,0,0,1002,21,125,20,4,20,1105,1,46,104,999,1105,1,46,1101,1000,1,20,4,20,1105,1,46,98,99")
	p.input = 9
	assert.Equal(t, 1001, solve(p))

	p = parseInput("3,9,8,9,10,9,4,9,99,-1,8")
	p.input = 8
	p.run()
	assert.Equal(t, []int{1}, p.output)

	p = parseInput("3,9,8,9,10,9,4,9,99,-1,8")
	p.input = 8
	p.input = 7
	p.run()
	assert.Equal(t, []int{0}, p.output)

	p = parseInput("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")
	p.input = 1
	p.run()
	assert.Equal(t, []int{1}, p.output)

	p = parseInput("3,12,6,12,15,1,13,14,13,4,13,99,-1,0,1,9")
	p.input = 0
	p.run()
	assert.Equal(t, []int{0}, p.output)
}

func TestMoreExamplesPartB(t *testing.T) {
	p := parseInput("3,3,1105,-1,9,1101,0,0,12,4,12,99,1")
	p.input = 1
	p.run()
	assert.Equal(t, []int{1}, p.output)

	p = parseInput("3,3,1105,-1,9,1101,0,0,12,4,12,99,1")
	p.input = 0
	p.run()
	assert.Equal(t, []int{0}, p.output)
}

func TestExamplesFromReddit(t *testing.T) {
	p := parseInput("101,-1,7,7,4,7,1105,11,0,99")
	p.input = 1
	p.run()
	assert.Equal(t, []int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, p.output)

	p = parseInput("1,0,3,3,1005,2,10,5,1,0,4,1,99")
	p.input = 1
	p.run()
	assert.Equal(t, []int{0}, p.output)
}

func TestSolvePartB(t *testing.T) {
	p := parseInput(utils.LoadString("input.txt"))
	p.input = 5
	assert.Equal(t, 11826654, solve(p))
}
