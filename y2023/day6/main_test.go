package day6

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOutcomes(t *testing.T) {
	assert.Equal(t, []int{0, 6, 10, 12, 12, 10, 6, 0}, outcomes(7))
}

func TestAccelaration(t *testing.T) {
	assert.Equal(t, 0, accelaration(0, 7))
	assert.Equal(t, 6, accelaration(1, 7))
	assert.Equal(t, 10, accelaration(2, 7))
}

var exampleA = []Race{
	{time: 7, distance: 9},
	{time: 15, distance: 40},
	{time: 30, distance: 200},
}

func TestExample(t *testing.T) {
	assert.Equal(t, 288, solve(exampleA))
}

// Time:        46     85     75     82
// Distance:   208   1412   1257   1410
var inputA = []Race{
	{time: 46, distance: 208},
	{time: 85, distance: 1412},
	{time: 75, distance: 1257},
	{time: 82, distance: 1410},
}

func TestSolutionA(t *testing.T) {
	assert.Equal(t, 1108800, solve(inputA))
}

var exampleB = []Race{
	{time: 71530, distance: 940200},
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 71503, solve(exampleB))
}

// Time:        46857582
// Distance:   208141212571410
var inputB = []Race{
	{time: 46857582, distance: 208141212571410},
}

func TestSolutionB(t *testing.T) {
	assert.Equal(t, 36919753, solve(inputB))
}
