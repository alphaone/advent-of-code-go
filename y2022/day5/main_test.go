package day5

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleStack string = `    [D]    
[N] [C]    
[Z] [M] [P]
 1   2   3
`

func TestParseStack(t *testing.T) {
	expected := stack{
		{'N', 'Z'},
		{'D', 'C', 'M'},
		{'P'},
	}

	assert.Equal(t, expected, parseStack(exampleStack))
}

var exampleInstructions string = `move 1 from 2 to 1
move 3 from 1 to 3
move 2 from 2 to 1
move 1 from 1 to 2
`

func TestParseInstructions(t *testing.T) {
	assert.Equal(t, []instruction{
		{amount: 1, from: 1, to: 0},
		{amount: 3, from: 0, to: 2},
		{amount: 2, from: 1, to: 0},
		{amount: 1, from: 0, to: 1},
	}, parseInstructions(exampleInstructions))
}

func TestSolveExampleA(t *testing.T) {
	stack, ins := parse(exampleStack + "\n" + exampleInstructions)
	assert.Equal(t, "CMZ", solveA(stack, ins, moveA))
}

func TestSolveA(t *testing.T) {
	stack, ins := parse(utils.LoadString("input.txt"))
	assert.Equal(t, "MQTPGLLDN", solveA(stack, ins, moveA))
}

func TestSolveExampleB(t *testing.T) {
	stack, ins := parse(exampleStack + "\n" + exampleInstructions)
	assert.Equal(t, "MCD", solveA(stack, ins, moveB))
}

func TestSolveB(t *testing.T) {
	stack, ins := parse(utils.LoadString("input.txt"))
	assert.Equal(t, "LVZPSTTCZ", solveA(stack, ins, moveB))
}
