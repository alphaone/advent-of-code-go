package day4

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"7,4,9,5,11,17,23,2,0,14,21,24,10,16,13,6,15,25,12,22,18,20,8,19,3,26,1",
	"",
	"22 13 17 11  0",
	"8  2 23  4 24",
	"21  9 14 16  7",
	"6 10  3 18  5",
	"1 12 20 15 19",
	"",
	"3 15  0  2 22",
	"9 18 13 17  5",
	"19  8  7 25 23",
	"20 11 10 24  4",
	"14 21 16 12  6",
	"",
	"14 21 17 24  4",
	"10 16 15  9 19",
	"18  8 23 26 20",
	"22 11 13  6  5",
	"2  0 12  3  7",
}

func TestParseInput(t *testing.T) {
	expected := Game{
		Numbers: []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24, 10, 16, 13, 6, 15, 25, 12, 22, 18, 20, 8, 19, 3, 26, 1},
		Boards: []Board{
			{
				Numbers: [][]int{
					{22, 13, 17, 11, 0},
					{8, 2, 23, 4, 24},
					{21, 9, 14, 16, 7},
					{6, 10, 3, 18, 5},
					{1, 12, 20, 15, 19},
				},
				Crosses: []Coord{},
			},
			{
				Numbers: [][]int{
					{3, 15, 0, 2, 22},
					{9, 18, 13, 17, 5},
					{19, 8, 7, 25, 23},
					{20, 11, 10, 24, 4},
					{14, 21, 16, 12, 6},
				},
				Crosses: []Coord{},
			},
			{
				Numbers: [][]int{
					{14, 21, 17, 24, 4},
					{10, 16, 15, 9, 19},
					{18, 8, 23, 26, 20},
					{22, 11, 13, 6, 5},
					{2, 0, 12, 3, 7},
				},
				Crosses: []Coord{},
			},
		},
	}
	assert.Equal(t, expected, parseInput(example))
}

func TestBoardCompleted(t *testing.T) {
	board := Board{
		Numbers: [][]int{
			{22, 13, 17, 11, 0},
			{8, 2, 23, 4, 24},
			{21, 9, 14, 16, 7},
			{6, 10, 3, 18, 5},
			{1, 12, 20, 15, 19},
		},
		Crosses: []Coord{},
	}
	assert.False(t, board.Completed())

	board.Crosses = []Coord{{0, 2}, {1, 2}, {2, 2}, {3, 2}, {4, 2}}
	assert.True(t, board.Completed())
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 4512, solvePartA(example))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 23177, solvePartA(input))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 1924, solvePartB(example))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 6804, solvePartB(input))
}
