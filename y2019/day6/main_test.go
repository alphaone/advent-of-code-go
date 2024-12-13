package day6

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"COM)B",
	"B)C",
	"C)D",
	"D)E",
	"E)F",
	"B)G",
	"G)H",
	"D)I",
	"E)J",
	"J)K",
	"K)L",
}

func TestCount(t *testing.T) {
	assert.Equal(t, 3, parse(example).count("D"))
	assert.Equal(t, 7, parse(example).count("L"))
	assert.Equal(t, 0, parse(example).count("COM"))
}

func TestPath(t *testing.T) {
	assert.Equal(t, []string{"C", "B", "COM"}, parse(example).path("D"))
	assert.Equal(t, []string{"K", "J", "E", "D", "C", "B", "COM"}, parse(example).path("L"))
	assert.Equal(t, []string{}, parse(example).path("COM"))
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 42, solvePartA(parse(example)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 308790, solvePartA(parse(utils.LoadStrings("input.txt"))))
}

var exampleB = []string{
	"COM)B",
	"B)C",
	"C)D",
	"D)E",
	"E)F",
	"B)G",
	"G)H",
	"D)I",
	"E)J",
	"J)K",
	"K)L",
	"K)YOU",
	"I)SAN",
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 4, solvePartB(parse(exampleB)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 472, solvePartB(parse(utils.LoadStrings("input.txt"))))
}
