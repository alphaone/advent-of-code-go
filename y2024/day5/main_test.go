package day5

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleStr = `47|53
97|13
97|61
97|47
75|29
61|13
75|53
29|13
97|29
53|29
61|53
97|53
61|29
47|13
75|47
97|75
47|61
75|61
47|29
75|13
53|13

75,47,61,53,29
97,61,53,29,13
75,29,13
75,97,47,61,53
61,13,29
97,13,75,29,47
`

func TestParsePageOrder(t *testing.T) {
	expected := PageOrder{
		29: []int{13},
		47: []int{53, 13, 61, 29},
		53: []int{29, 13},
		61: []int{13, 53, 29},
		75: []int{29, 53, 47, 61, 13},
		97: []int{13, 61, 47, 29, 53, 75},
	}
	actual, _ := parseInput(exampleStr)
	assert.Equal(t, expected, actual)
}

func TestParsePages(t *testing.T) {
	expected := [][]int{
		{75, 47, 61, 53, 29},
		{97, 61, 53, 29, 13},
		{75, 29, 13},
		{75, 97, 47, 61, 53},
		{61, 13, 29},
		{97, 13, 75, 29, 47},
	}
	_, actual := parseInput(exampleStr)
	assert.Equal(t, expected, actual)
}

func TestInRightOrder(t *testing.T) {
	exampleRules, _ := parseInput(exampleStr)

	assert.True(t, inRightOrder([]int{75, 47, 61, 53, 29}, exampleRules))
	assert.True(t, inRightOrder([]int{97, 61, 53, 29, 13}, exampleRules))
	assert.True(t, inRightOrder([]int{75, 29, 13}, exampleRules))

	assert.False(t, inRightOrder([]int{75, 97, 47, 61, 53}, exampleRules))
	assert.False(t, inRightOrder([]int{61, 13, 29}, exampleRules))
	assert.False(t, inRightOrder([]int{97, 13, 75, 29, 47}, exampleRules))
}

func TestSolveExampleA(t *testing.T) {
	rules, pageLists := parseInput(exampleStr)
	assert.Equal(t, 143, solvePartA(pageLists, rules))
}

func TestSolveA(t *testing.T) {
	rules, pageLists := parseInput(utils.LoadString("input.txt"))
	assert.Equal(t, 6951, solvePartA(pageLists, rules))
}

func TestSort(t *testing.T) {
	exampleRules, _ := parseInput(exampleStr)
	assert.Equal(t, []int{97, 75, 47, 61, 53}, sortByRules([]int{75, 97, 47, 61, 53}, exampleRules))
	assert.Equal(t, []int{61, 29, 13}, sortByRules([]int{61, 13, 29}, exampleRules))
	assert.Equal(t, []int{97, 75, 47, 29, 13}, sortByRules([]int{97, 13, 75, 29, 47}, exampleRules))

	// correct
	assert.Equal(t, []int{75, 47, 61, 53, 29}, sortByRules([]int{75, 47, 61, 53, 29}, exampleRules))
	assert.Equal(t, []int{97, 61, 53, 29, 13}, sortByRules([]int{97, 61, 53, 29, 13}, exampleRules))
	assert.Equal(t, []int{75, 29, 13}, sortByRules([]int{75, 29, 13}, exampleRules))
}

func TestSolveExampleB(t *testing.T) {
	rules, pageLists := parseInput(exampleStr)
	assert.Equal(t, 123, solvePartB(pageLists, rules))
}

func TestSolveB(t *testing.T) {
	rules, pageLists := parseInput(utils.LoadString("input.txt"))
	assert.Equal(t, 4121, solvePartB(pageLists, rules))
}
