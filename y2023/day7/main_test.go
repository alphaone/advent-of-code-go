package day7

import (
	"sort"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	line := "32T3K 765"
	expected := Game{
		Bid:  765,
		Hand: []rune("32T3K"),
	}
	assert.Equal(t, expected, parseLine(line))
}

func TestHandType(t *testing.T) {
	assert.Equal(t, 1, Hand([]rune("32T3K")).Type(false))
	assert.Equal(t, 3, Hand([]rune("T55J5")).Type(false))
	assert.Equal(t, 2, Hand([]rune("KK677")).Type(false))
	assert.Equal(t, 2, Hand([]rune("KTJJT")).Type(false))
	assert.Equal(t, 3, Hand([]rune("QQQJA")).Type(false))

	assert.Equal(t, 1, Hand([]rune("32T3K")).Type(true))
	assert.Equal(t, 5, Hand([]rune("T55J5")).Type(true))
	assert.Equal(t, 2, Hand([]rune("KK677")).Type(true))
	assert.Equal(t, 5, Hand([]rune("KTJJT")).Type(true))
	assert.Equal(t, 5, Hand([]rune("QQQJA")).Type(true))
}

func TestLessThan(t *testing.T) {
	assert.False(t, Hand([]rune("KK677")).LessThan(Hand([]rune("KTJJT")), false))
	assert.True(t, Hand([]rune("KTJJT")).LessThan(Hand([]rune("KK677")), false))
	// assert.True(t, Hand([]rune("KK677")).LessThan(Hand([]rune("T55J5"))))
	// assert.True(t, Hand([]rune("T55J5")).LessThan(Hand([]rune("QQQJA"))))
}

func TestGameSort(t *testing.T) {
	games := []Game{
		{Hand: []rune("32T3K")},
		{Hand: []rune("T55J5")},
		{Hand: []rune("KK677")},
		{Hand: []rune("KTJJT")},
		{Hand: []rune("QQQJA")},
	}
	sort.Sort(GamesByValue(games))
	assert.Equal(t, []Game{
		{Hand: []rune("32T3K")},
		{Hand: []rune("KTJJT")},
		{Hand: []rune("KK677")},
		{Hand: []rune("T55J5")},
		{Hand: []rune("QQQJA")},
	}, games)
	sort.Sort(GamesByValueWithJoker(games))
	assert.Equal(t, []Game{
		{Hand: []rune("32T3K")},
		{Hand: []rune("KK677")},
		{Hand: []rune("T55J5")},
		{Hand: []rune("QQQJA")},
		{Hand: []rune("KTJJT")},
	}, games)
}

var exampleA = []string{
	"32T3K 765",
	"T55J5 684",
	"KK677 28",
	"KTJJT 220",
	"QQQJA 483",
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 6440, solvePartA(exampleA))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 251806792, solvePartA(input))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 5905, solvePartB(exampleA))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 252113488, solvePartB(input))
}
