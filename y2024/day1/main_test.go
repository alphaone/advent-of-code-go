package day1

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleA = []string{
	"3   4",
	"4   3",
	"2   5",
	"1   3",
	"3   9",
	"3   3",
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 11, solvePartA(parseInput(exampleA)))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 1580061, solvePartA(parseInput(input)))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 31, solvePartB(parseInput(exampleA)))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 23046913, solvePartB(parseInput(input)))
}

func TestParse(t *testing.T) {
	res := parseInput(exampleA)
	expected := []Tuple{{3, 4}, {4, 3}, {2, 5}, {1, 3}, {3, 9}, {3, 3}}

	assert.Equal(t, expected, res)
}
