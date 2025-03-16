package day15

import (
	"strings"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	g := game{0, 3, 6}
	assert.Equal(t, 0, g.play(4))
	assert.Equal(t, 3, g.play(5))
	assert.Equal(t, 3, g.play(6))
	assert.Equal(t, 1, g.play(7))
}

func TestSolveExampleA(t *testing.T) {
	g := game([]int{0, 3, 6})
	assert.Equal(t, 436, g.play(2020))
	g = game([]int{3, 1, 2})
	assert.Equal(t, 1836, g.play(2020))
}

func TestSolveA(t *testing.T) {
	input := strings.TrimSpace(utils.LoadString("input.txt"))
	g := parse(input)
	assert.Equal(t, 870, g.play(2020))
}

func TestSolveExampleB(t *testing.T) {
	g := game([]int{0, 3, 6})
	assert.Equal(t, 175594, g.play(30_000_000))
}

func TestSolveB(t *testing.T) {
	input := strings.TrimSpace(utils.LoadString("input.txt"))
	g := parse(input)
	assert.Equal(t, 9136, g.play(30_000_000))
}
