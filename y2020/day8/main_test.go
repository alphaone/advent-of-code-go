package day8

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}

func TestParse(t *testing.T) {
	assert.Equal(t, []cmd{
		{op: "nop", val: 0},
		{op: "acc", val: 1},
		{op: "jmp", val: 4},
		{op: "acc", val: 3},
		{op: "jmp", val: -3},
		{op: "acc", val: -99},
		{op: "acc", val: 1},
		{op: "jmp", val: -4},
		{op: "acc", val: 6},
	}, parse(example))
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 5, solveA(parse(example)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 1586, solveA(parse(utils.LoadStrings("input.txt"))))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 8, solveB(parse(example)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 703, solveB(parse(utils.LoadStrings("input.txt"))))
}
