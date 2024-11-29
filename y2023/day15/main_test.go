package day15

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestHashing(t *testing.T) {
	assert.Equal(t, 52, Hashing("HASH"))
	assert.Equal(t, 30, Hashing("rn=1"))
	assert.Equal(t, 253, Hashing("cm-"))
	assert.Equal(t, 97, Hashing("qp=3"))

	assert.Equal(t, 0, Hashing("rn"))
	assert.Equal(t, 1, Hashing("qp"))
	assert.Equal(t, 3, Hashing("ot"))
}

var example = "rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7"

func TestExample(t *testing.T) {
	assert.Equal(t, 1320, solvePartA(example))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 511257, solvePartA(input[0]))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 145, solvePartB(example))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 239484, solvePartB(input[0]))
}
