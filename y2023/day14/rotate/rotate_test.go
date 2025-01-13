package rotate

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

func TestRotateClockwise(t *testing.T) {
	assert.Equal(t, []string{
		"741",
		"852",
		"963",
	}, utils.AsStrings(RotateClockwise(utils.AsRunes([]string{
		"123",
		"456",
		"789",
	}))))
}

func TestRotateClockwiseMulti(t *testing.T) {
	start := utils.LoadStrings("../input.txt")

	res := utils.AsRunes(start)
	for i := 0; i < 4*10_000; i++ {
		res = RotateClockwise(res)
	}

	assert.Equal(t, start, utils.AsStrings(res))
}

func TestRotateCounterClockwiseMulti(t *testing.T) {
	start := utils.LoadStrings("../input.txt")

	res := utils.AsRunes(start)
	for i := 0; i < 4*10_000; i++ {
		res = RotateCounterClockwise(res)
	}

	assert.Equal(t, start, utils.AsStrings(res))
}
