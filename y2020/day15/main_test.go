package day15

import (
	"strings"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestNext(t *testing.T) {
	g := game{0, 3, 6}
	g = g.next()
	assert.Equal(t, game{0, 3, 6, 0}, g)
	g = g.next()
	assert.Equal(t, game{0, 3, 6, 0, 3}, g)
	g = g.next()
	assert.Equal(t, game{0, 3, 6, 0, 3, 3}, g)
	g = g.next()
	assert.Equal(t, game{0, 3, 6, 0, 3, 3, 1}, g)
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
