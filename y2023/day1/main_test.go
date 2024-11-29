package day1

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleA = []string{
	"1abc2",
	"pqr3stu8vwx",
	"a1b2c3d4e5f",
	"treb7uchet",
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 142, solvePartA(exampleA))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 56465, solvePartA(input))
}

var exampleB = []string{
	"two1nine",
	"eightwothree",
	"abcone2threexyz",
	"xtwone3four",
	"4nineeightseven2",
	"zoneight234",
	"7pqrstsixteen",
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 281, solvePartB(exampleB))
}

func TestExampleB2(t *testing.T) {
	input := []string{"eighthree"}
	assert.Equal(t, 83, solvePartB(input))

	input = []string{"sevenine"}
	assert.Equal(t, 79, solvePartB(input))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 55902, solvePartB(input))
}
