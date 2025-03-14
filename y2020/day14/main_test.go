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

func TestMaskAddressApply(t *testing.T) {
	assert.Equal(t, []int{0b1}, mask("0").applyAddress(0b1))
	assert.Equal(t, []int{0b1}, mask("1").applyAddress(0b0))
	assert.Equal(t, []int{0b0}, mask("0").applyAddress(0b0))
	assert.Equal(t, []int{0b0, 0b1}, mask("X").applyAddress(0b1))

	assert.Equal(t, []int{0b11}, mask("00").applyAddress(0b11))
	assert.Equal(t, []int{0b11}, mask("11").applyAddress(0b00))
	assert.Equal(t, []int{0b11}, mask("01").applyAddress(0b10))
	assert.Equal(t, []int{0b11}, mask("10").applyAddress(0b01))
	assert.Equal(t, []int{0b01}, mask("00").applyAddress(0b01))
	assert.Equal(t, []int{0b10}, mask("10").applyAddress(0b00))
	assert.Equal(t, []int{0b10, 0b11}, mask("1X").applyAddress(0b00))
	assert.Equal(t, []int{0b01, 0b11}, mask("X1").applyAddress(0b00))
	assert.Equal(t, []int{0b00, 0b01, 0b10, 0b11}, mask("XX").applyAddress(0b00))

	assert.Equal(t, []int{0b1111}, mask("1111").applyAddress(0b0101))
	assert.Equal(t, []int{0b1101}, mask("1100").applyAddress(0b0101))
	assert.Equal(t, []int{0b0101}, mask("0000").applyAddress(0b0101))

	assert.Equal(t, []int{0b0100, 0b0101}, mask("000X").applyAddress(0b0101))
	assert.Equal(t, []int{0b1100, 0b1101, 0b1110, 0b1111}, mask("11XX").applyAddress(0b0101))

	assert.Equal(t, []int{26, 27, 58, 59}, mask("0X1001X").applyAddress(42))
}

var exampleB = []string{
	"mask = 000000000000000000000000000000X1001X",
	"mem[42] = 100",
	"mask = 00000000000000000000000000000000X0XX",
	"mem[26] = 1",
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 208, solveB(parse(exampleB)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 3608464522781, solveB(parse(utils.LoadStrings("input.txt"))))
}
