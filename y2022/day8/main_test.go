package day8

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"30373",
	"25512",
	"65332",
	"33549",
	"35390",
}

func TestParse(t *testing.T) {
	exp := treepatch{
		{3, 0, 3, 7, 3},
		{2, 5, 5, 1, 2},
		{6, 5, 3, 3, 2},
		{3, 3, 5, 4, 9},
		{3, 5, 3, 9, 0},
	}
	assert.Equal(t, exp, parse(example))
}

type treesInDirTestcase struct {
	name string
	pos  coord
	dir  coord
	exp  []int
}

func TestTreeInDir(t *testing.T) {
	tp := parse(example)
	for _, tc := range []treesInDirTestcase{
		{name: "top-left1", pos: coord{0, 0}, dir: coord{1, 0}, exp: []int{2, 6, 3, 3}},
		{name: "top-left2", pos: coord{0, 0}, dir: coord{-1, 0}, exp: []int{}},
		{name: "top-left3", pos: coord{0, 0}, dir: coord{0, 1}, exp: []int{0, 3, 7, 3}},

		{name: "center down", pos: coord{2, 2}, dir: coord{1, 0}, exp: []int{5, 3}},
		{name: "center up", pos: coord{2, 2}, dir: coord{-1, 0}, exp: []int{5, 3}},
		{name: "center left", pos: coord{2, 2}, dir: coord{0, 1}, exp: []int{3, 2}},
		{name: "center right", pos: coord{2, 2}, dir: coord{0, -1}, exp: []int{5, 6}},
	} {
		assert.Equal(t, tc.exp, tp.treesInDir(tc.pos, tc.dir))
	}
}

type visibleTestcase struct {
	name string
	pos  coord
	exp  bool
}

func TestVisible(t *testing.T) {
	for _, tc := range []visibleTestcase{
		{"top-left", coord{0, 0}, true},
		{"top-right", coord{0, 4}, true},
		{"bottom-left", coord{4, 0}, true},
		{"bottom-right", coord{4, 4}, true},

		{"interior1", coord{1, 1}, true},
		{"interior3", coord{1, 3}, false},
		{"center", coord{2, 2}, false},
	} {
		tp := parse(example)
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp, tp.visible(tc.pos))
		})
	}
}

func TestSolveExample(t *testing.T) {
	assert.Equal(t, 21, solve(example))
}

func TestSolve(t *testing.T) {
	assert.Equal(t, 1845, solve(utils.LoadStrings("input.txt")))
}

type scenicScoreTestcase struct {
	name string
	pos  coord
	exp  int
}

func TestScenicScore(t *testing.T) {
	for _, tc := range []scenicScoreTestcase{
		{"top-left", coord{0, 0}, 0},
		{"example1", coord{1, 2}, 4},
		{"example2", coord{3, 2}, 8},
	} {
		tp := parse(example)
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp, tp.scenicScore(tc.pos))
		})
	}
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 8, solveB(example))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 230112, solveB(utils.LoadStrings("input.txt")))
}
