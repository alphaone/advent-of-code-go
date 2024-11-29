package day4

import (
	"sort"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	input := "Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53"
	expected := Card{
		Id:            1,
		Winning:       []int{41, 48, 83, 86, 17},
		Played:        []int{83, 86, 6, 31, 17, 9, 48, 53},
		instanceCount: 1,
	}
	assert.Equal(t, expected, parseLine(input))
}

func TestMatching(t *testing.T) {
	card := Card{
		Id:      1,
		Winning: []int{41, 48, 83, 86, 17},
		Played:  []int{83, 86, 6, 31, 17, 9, 48, 53},
	}

	res := card.Matching()
	sort.Ints(res)
	assert.Equal(t, []int{17, 48, 83, 86}, res)
}

func TestScore(t *testing.T) {
	card := Card{
		Id:      1,
		Winning: []int{41, 48, 83, 86, 17},
		Played:  []int{83, 86, 6, 31, 17, 9, 48, 53},
	}
	assert.Equal(t, 8, card.Score())
}

var input = []string{
	"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53",
	"Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19",
	"Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1",
	"Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83",
	"Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36",
	"Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11",
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 13, solvePartA(input))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 18653, solvePartA(input))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 30, solvePartB(input))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 5921508, solvePartB(input))
}
