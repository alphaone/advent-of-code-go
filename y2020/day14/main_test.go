package day14

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"mask = XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X",
	"mem[8] = 11",
	"mem[7] = 101",
	"mem[8] = 0",
}

func TestParse(t *testing.T) {
	expected := []instruction{
		{op: "mask", mask: "XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X"},
		{op: "mem", addr: 8, value: 11},
		{op: "mem", addr: 7, value: 101},
		{op: "mem", addr: 8, value: 0},
	}
	assert.Equal(t, expected, parse(example))
}

func TestMaskApply(t *testing.T) {
	assert.Equal(t, 0b1111, mask("1111").apply(0b0101))
	assert.Equal(t, 0b0000, mask("0000").apply(0b0101))
	assert.Equal(t, 0b1100, mask("1100").apply(0b0101))

	assert.Equal(t, 0b0101, mask("XXXX").apply(0b0101))
	assert.Equal(t, 0b1001, mask("10XX").apply(0b0101))

	assert.Equal(t, 73, mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X").apply(11))
	assert.Equal(t, 101, mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X").apply(101))
	assert.Equal(t, 64, mask("XXXXXXXXXXXXXXXXXXXXXXXXXXXXX1XXXX0X").apply(0))
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 165, solveA(parse(example)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 12610010960049, solveA(parse(utils.LoadStrings("input.txt"))))
}
