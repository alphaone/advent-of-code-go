package day9

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestExampleA(t *testing.T) {
	p := parseInput(`109,1,204,-1,1001,100,1,100,1008,100,16,101,1006,101,0,99`)
	p.input = []int{1}
	p.run()
	assert.Equal(t, []int{109, 1, 204, -1, 1001, 100, 1, 100, 1008, 100, 16, 101, 1006, 101, 0, 99}, p.output)

	p = parseInput(`1102,34915192,34915192,7,4,7,99,0`)
	assert.Equal(t, 1219070632396864, solve(p, 1))

	p = parseInput(`104,1125899906842624,99`)
	assert.Equal(t, 1125899906842624, solve(p, 1))
}

func TestSolve(t *testing.T) {
	p := parseInput(utils.LoadString("input.txt"))
	assert.Equal(t, 2453265701, solve(p, 1))
}

func TestSolveB(t *testing.T) {
	p := parseInput(utils.LoadString("input.txt"))
	assert.Equal(t, 80805, solve(p, 2))
}
