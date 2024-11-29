package day8

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"RL",
	"",
	"AAA = (BBB, CCC)",
	"BBB = (DDD, EEE)",
	"CCC = (ZZZ, GGG)",
	"DDD = (DDD, DDD)",
	"EEE = (EEE, EEE)",
	"GGG = (GGG, GGG)",
	"ZZZ = (ZZZ, ZZZ)",
}

func TestParseNodes(t *testing.T) {
	expected := map[string]Node{
		"AAA": {"AAA", "BBB", "CCC"},
		"BBB": {"BBB", "DDD", "EEE"},
		"CCC": {"CCC", "ZZZ", "GGG"},
		"DDD": {"DDD", "DDD", "DDD"},
		"EEE": {"EEE", "EEE", "EEE"},
		"GGG": {"GGG", "GGG", "GGG"},
		"ZZZ": {"ZZZ", "ZZZ", "ZZZ"},
	}
	actual := parseNodes(example[2:])
	assert.Equal(t, expected, actual)
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 2, solvePartA(example))
}

func TestExampleA2(t *testing.T) {
	example2 := []string{
		"LLR",
		"",
		"AAA = (BBB, BBB)",
		"BBB = (AAA, ZZZ)",
		"ZZZ = (ZZZ, ZZZ)",
	}
	assert.Equal(t, 6, solvePartA(example2))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 18727, solvePartA(input))
}

var exampleB = []string{
	"LR",
	"",
	"11A = (11B, XXX)",
	"11B = (XXX, 11Z)",
	"11Z = (11B, XXX)",
	"22A = (22B, XXX)",
	"22B = (22C, 22C)",
	"22C = (22Z, 22Z)",
	"22Z = (22B, 22B)",
	"XXX = (XXX, XXX)",
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 6, solvePartB(exampleB))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 18024643846273, solvePartB(input))
}
