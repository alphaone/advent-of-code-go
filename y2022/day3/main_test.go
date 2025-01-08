package day3

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"vJrwpWtwJgWrhcsFMMfFFhFp",
	"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
	"PmmdzqPrVvPwwTWBwg",
	"wMqvLMZHhHMvwLHjbvcjnnSBnvTQFn",
	"ttgJtRGJQctTZtZT",
	"CrZsJsPPZsGzwwsLwLmpwMDw",
}

type findWrongItemTestCases struct {
	name  string
	input string
	exp   rune
}

func TestFindWrongItem(t *testing.T) {
	for _, tc := range []findWrongItemTestCases{
		{"Example line 0", example[0], 'p'},
		{"Example line 1", example[1], 'L'},
		{"Example line 2", example[2], 'P'},
		{"Example line 3", example[3], 'v'},
		{"Example line 4", example[4], 't'},
		{"Example line 5", example[5], 's'},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp, findWrongItem(tc.input))
		})
	}
}

func TestRuneScore(t *testing.T) {
	assert.Equal(t, 1, score('a'))
	assert.Equal(t, 2, score('b'))
	assert.Equal(t, 26, score('z'))
	assert.Equal(t, 27, score('A'))
	assert.Equal(t, 52, score('Z'))
}

func TestSolveExample(t *testing.T) {
	assert.Equal(t, 157, solve(example))
}

func TestSolve(t *testing.T) {
	assert.Equal(t, 7568, solve(utils.LoadStrings("input.txt")))
}

func TestFindBadge(t *testing.T) {
	assert.Equal(t, 'r', findBadge([]string{
		"vJrwpWtwJgWrhcsFMMfFFhFp",
		"jqHRNqRjqzjGDLGLrsFMfFZSrLrFZsSL",
		"PmmdzqPrVvPwwTWBwg",
	}))
}

func TestExampleSolveB(t *testing.T) {
	assert.Equal(t, 70, solveB(example))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 2780, solveB(utils.LoadStrings("input.txt")))
}
