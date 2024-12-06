package day2

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestStep(t *testing.T) {
	p := Program{1, 9, 10, 3, 2, 3, 11, 0, 99, 30, 40, 50}
	assert.Equal(t, 4, p.step(0))
	assert.Equal(t, Program{1, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, p)

	assert.Equal(t, 8, p.step(4))
	assert.Equal(t, Program{3500, 9, 10, 70, 2, 3, 11, 0, 99, 30, 40, 50}, p)

	assert.Equal(t, -1, p.step(8))
}

func TestRun(t *testing.T) {
	p := Program{1, 1, 1, 4, 99, 5, 6, 0, 99}
	p.run()
	assert.Equal(t, Program{30, 1, 1, 4, 2, 5, 6, 0, 99}, p)
}

func TestSolvePartA(t *testing.T) {
	input := utils.LoadString("input.txt")
	assert.Equal(t, 5434663, solvePartA(parseInput(input)))
}

func TestSolvePartB(t *testing.T) {
	input := utils.LoadString("input.txt")
	assert.Equal(t, 4559, solvePartB(parseInput(input)))
}
