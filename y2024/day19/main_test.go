package day19

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleStr = `
r, wr, b, g, bwu, rb, gb, br

brwrr
bggr
gbbr
rrbgbr
ubwu
bwurrg
brgr
bbrgwb
`

func TestCountPaths(t *testing.T) {
	assert.Equal(t, 1, countPaths("bwurrg", parse(exampleStr).towels, make(map[string]int)))
	assert.Equal(t, 2, countPaths("brgr", parse(exampleStr).towels, make(map[string]int)))
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 6, solvePartA(parse(exampleStr)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 311, solvePartA(parse(utils.LoadString("input.txt"))))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 16, solvePartB(parse(exampleStr)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 616234236468263, solvePartB(parse(utils.LoadString("input.txt"))))
}
