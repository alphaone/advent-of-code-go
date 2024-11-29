package day3

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"00100",
	"11110",
	"10110",
	"10111",
	"10101",
	"01111",
	"00111",
	"11100",
	"10000",
	"11001",
	"00010",
	"01010",
}

func TestTranspose(t *testing.T) {
	expected := []string{
		"011110011100",
		"010001010101",
		"111111110000",
		"011101100011",
		"000111100100",
	}
	actual := Transpose(example)
	assert.Equal(t, expected, actual)
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 198, solvePartA(example))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 2954600, solvePartA(input))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 230, solvePartB(example))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 1662846, solvePartB(input))
}
