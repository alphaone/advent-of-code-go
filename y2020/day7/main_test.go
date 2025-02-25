package day7

import (
	"sort"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"light red bags contain 1 bright white bag, 2 muted yellow bags.",
	"dark orange bags contain 3 bright white bags, 4 muted yellow bags.",
	"bright white bags contain 1 shiny gold bag.",
	"muted yellow bags contain 2 shiny gold bags, 9 faded blue bags.",
	"shiny gold bags contain 1 dark olive bag, 2 vibrant plum bags.",
	"dark olive bags contain 3 faded blue bags, 4 dotted black bags.",
	"vibrant plum bags contain 5 faded blue bags, 6 dotted black bags.",
	"faded blue bags contain no other bags.",
	"dotted black bags contain no other bags.",
}

func TestParse(t *testing.T) {
	expected := tree{
		"bright white": {"shiny gold": 1},
		"dark olive":   {"dotted black": 4, "faded blue": 3},
		"dark orange":  {"bright white": 3, "muted yellow": 4},
		"dotted black": {},
		"faded blue":   {},
		"light red":    {"bright white": 1, "muted yellow": 2},
		"muted yellow": {"faded blue": 9, "shiny gold": 2},
		"shiny gold":   {"dark olive": 1, "vibrant plum": 2},
		"vibrant plum": {"dotted black": 6, "faded blue": 5},
	}
	assert.Equal(t, expected, parse(example))
}

func TestSolveExampleA(t *testing.T) {
	expected := []string{"bright white", "dark orange", "light red", "muted yellow"}
	actual := solveA(parse(example), "shiny gold")
	sort.Strings(actual)
	assert.Equal(t, expected, actual)
}

func TestSolveA(t *testing.T) {
	actual := solveA(parse(utils.LoadStrings("input.txt")), "shiny gold")
	assert.Equal(t, 242, len(actual))
}

func TestSolveExampleB(t *testing.T) {
	actual := solveB(parse(example), "shiny gold")
	assert.Equal(t, 32, actual)
}

func TestSolveB(t *testing.T) {
	actual := solveB(parse(utils.LoadStrings("input.txt")), "shiny gold")
	assert.Equal(t, 176035, actual)
}
