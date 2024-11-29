package day17

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"2413432311323",
	"3215453535623",
	"3255245654254",
	"3446585845452",
	"4546657867536",
	"1438598798454",
	"4457876987766",
	"3637877979653",
	"4654967986887",
	"4564679986453",
	"1224686865563",
	"2546548887735",
	"4322674655533",
}

func TestExampleA(t *testing.T) {
	field := utils.AsIntField(example)
	assert.Equal(t, 102, solvePartA(field))
}

func TestSolutionA(t *testing.T) {
	input := utils.AsIntField(utils.LoadStrings("input.txt"))
	assert.Equal(t, 694, solvePartA(input))
}

func TestExampleB1(t *testing.T) {
	field := utils.AsIntField(example)
	assert.Equal(t, 94, solvePartB(field))
}

var example2 = []string{
	"111111111111",
	"999999999991",
	"999999999991",
	"999999999991",
	"999999999991",
}

func TestExampleB2(t *testing.T) {
	field := utils.AsIntField(example2)
	assert.Equal(t, 71, solvePartB(field))
}

func TestSolutionB(t *testing.T) {
	input := utils.AsIntField(utils.LoadStrings("input.txt"))
	assert.Equal(t, 829, solvePartB(input))
}
