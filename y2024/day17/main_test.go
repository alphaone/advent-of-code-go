package day17

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleStrA = `Register A: 729
Register B: 0
Register C: 0

Program: 0,1,5,4,3,0
`

func TestParse(t *testing.T) {
	assert.Equal(t, computer{
		A: 729, B: 0, C: 0, Program: []int{0, 1, 5, 4, 3, 0}, Output: []int{},
	}, parse(exampleStrA))
}

func TestOperations(t *testing.T) {
	c := computer{C: 9, Program: []int{2, 6}}
	c.run()
	assert.Equal(t, 1, c.B)

	c = computer{A: 10, Program: []int{5, 0, 5, 1, 5, 4}}
	c.run()
	assert.Equal(t, []int{0, 1, 2}, c.Output)

	c = computer{A: 2024, Program: []int{0, 1, 5, 4, 3, 0}}
	c.run()
	assert.Equal(t, 0, c.A)
	assert.Equal(t, []int{4, 2, 5, 6, 7, 7, 7, 7, 3, 1, 0}, c.Output)
}

func TestSolveExamplePartA(t *testing.T) {
	assert.Equal(t, "4,6,3,5,6,3,5,2,1,0", solvePartA(parse(exampleStrA)))
}

func TestSolvePartA(t *testing.T) {
	assert.Equal(t, "4,0,4,7,1,2,7,1,6", solvePartA(parse(utils.LoadString("input.txt"))))
}
