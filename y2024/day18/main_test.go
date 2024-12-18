package day18

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"5,4", "4,2", "4,5", "3,0", "2,1", "6,3", "2,4", "1,5", "0,6", "3,3", "2,6", "5,1",
	"1,2", "5,5", "2,5", "6,5", "1,4", "0,4", "6,4", "1,1", "6,1", "1,0", "0,5", "1,6", "2,0",
}

func TestParse(t *testing.T) {
	assert.Equal(t, []Coord{{4, 5}, {2, 4}, {5, 4}}, parse(example[:3]))
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 22, solvePartA(parse(example[:12]), Coord{6, 6}))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 340, solvePartA(parse(utils.LoadStrings("input.txt")[:1024]), Coord{70, 70}))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, Coord{1, 6}, solvePartB(parse(example), Coord{6, 6}))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, Coord{l: 32, c: 34}, solvePartB(parse(utils.LoadStrings("input.txt")), Coord{70, 70}))
}
