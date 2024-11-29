package day16

import (
	"strings"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	`.|...\....`,
	`|.-.\.....`,
	`.....|-...`,
	`........|.`,
	`..........`,
	`.........\`,
	`..../.\\..`,
	`.-.-/..|..`,
	`.|....-|.\`,
	`..//.|....`,
}

func TestAdvance(t *testing.T) {
	field := parseInput(example)
	field = advance(field, 0, 0, 'E')

	expected := []string{
		`>|<<<\....`,
		`|v-.\^....`,
		`.v...|->>>`,
		`.v...v^.|.`,
		`.v...v^...`,
		`.v...v^..\`,
		`.v../2\\..`,
		`<->-/vv|..`,
		`.|<<<2-|.\`,
		`.v//.|.v..`,
	}
	assert.Equal(t, strings.Join(expected, "\n"), strings.Trim(field.String(), "\n"))
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 46, solvePartA(example))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 7860, solvePartA(input))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 51, solvePartB(example))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 8331, solvePartB(input))
}
