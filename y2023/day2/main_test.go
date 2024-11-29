package day2

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestParseLine(t *testing.T) {
	input := "Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green"
	expected := Game{
		Id: 1,
		Draws: []Draw{
			{"blue": 3, "red": 4},
			{"red": 1, "green": 2, "blue": 6},
			{"green": 2},
		},
	}
	assert.Equal(t, expected, parseLine(input))
}

func TestExampleA(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	assert.Equal(t, 8, solvePartA(input))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 1734, solvePartA(input))
}

func TestExampleB(t *testing.T) {
	input := []string{
		"Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green",
		"Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue",
		"Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red",
		"Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red",
		"Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green",
	}
	assert.Equal(t, 2286, solvePartB(input))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 70387, solvePartB(input))
}
