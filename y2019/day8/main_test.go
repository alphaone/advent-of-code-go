package day8

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = `123456789012`

func TestParse(t *testing.T) {
	assert.Equal(t, [][]int{{1, 2, 3, 4, 5, 6}, {7, 8, 9, 0, 1, 2}}, parseInput(example, 3, 2))
}

func TestExampleA(t *testing.T) {
	layers := parseInput(example, 3, 2)
	assert.Equal(t, 1, solveA(layers))
}

func TestSolveA(t *testing.T) {
	layers := parseInput(utils.LoadString("input.txt"), 25, 6)
	assert.Equal(t, 2125, solveA(layers))
}

func TestStack(t *testing.T) {
	layers := parseInput("0222112222120000", 2, 2)
	assert.Equal(t, []int{0, 1, 1, 0}, stack(layers))
}

func TestExampleB(t *testing.T) {
	layers := parseInput("0222112222120000", 2, 2)
	assert.Equal(t, " X\nX \n", solveB(layers, 2))
}

func TestSolveB(t *testing.T) {
	layers := parseInput(utils.LoadString("input.txt"), 25, 6)
	expected := `
  XX X   XXXXX X  X XXXX 
   X X   X   X X  X X    
   X  X X   X  XXXX XXX  
   X   X   X   X  X X    
X  X   X  X    X  X X    
 XX    X  XXXX X  X X    
`

	assert.Equal(t, expected, "\n"+solveB(layers, 25))
}
