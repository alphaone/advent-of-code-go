package day9

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []int{35, 20, 15, 25, 47, 40, 62, 55, 65, 95, 102, 117, 150, 182, 127, 219, 299, 277, 309, 576}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 127, firstInvalid(example, 5))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 2089807806, firstInvalid(utils.LoadInts("input.txt"), 25))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 62, encryptionWeakness(example, 127))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 245848639, encryptionWeakness(utils.LoadInts("input.txt"), 2089807806))
}
