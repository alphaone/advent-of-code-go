package day10

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example1 = []string{
	".....",
	".S-7.",
	".|.|.",
	".L-J.",
	".....",
}

var example2 = []string{
	"..F7.",
	".FJ|.",
	"SJ.L7",
	"|F--J",
	"LJ...",
}

func TestFindStart(t *testing.T) {
	assert.Equal(t, Point{1, 1, 'S'}, FindStart(example1))
}

func TestOptionsFromStart(t *testing.T) {
	d, c := OptionsFromStart(example1, Point{1, 1, '.'})
	assert.Equal(t, d, 'E')
	assert.Equal(t, c, 'F')
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 4, solvePartA(example1))
}

func TestExampleA2(t *testing.T) {
	assert.Equal(t, 8, solvePartA(example2))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 6778, solvePartA(input))
}

var example3 = []string{
	"...........",
	".S-------7.",
	".|F-----7|.",
	".||.....||.",
	".||.....||.",
	".|L-7.F-J|.",
	".|..|.|..|.",
	".L--J.L--J.",
	"...........",
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 4, solvePartB(example3))
}

var example4 = []string{
	".F----7F7F7F7F-7....",
	".|F--7||||||||FJ....",
	".||.FJ||||||||L7....",
	"FJL7L7LJLJ||LJ.L-7..",
	"L--J.L7...LJS7F-7L7.",
	"....F-J..F7FJ|L7L7L7",
	"....L7.F7||L7|.L7L7|",
	".....|FJLJ|FJ|F7|.LJ",
	"....FJL-7.||.||||...",
	"....L---J.LJ.LJLJ...",
}

func TestExampleB2(t *testing.T) {
	assert.Equal(t, 8, solvePartB(example4))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 433, solvePartB(input))
}
