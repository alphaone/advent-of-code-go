package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRotateClockwise(t *testing.T) {
	assert.Equal(t, []string{
		"741",
		"852",
		"963",
	}, AsStrings(RotateClockwise(AsRunes([]string{
		"123",
		"456",
		"789",
	}))))
}

func TestRotateClockwiseMulti(t *testing.T) {
	start := LoadStrings("../y2023/day14/input.txt")

	res := AsRunes(start)
	for i := 0; i < 4*10_000; i++ {
		res = RotateClockwise(res)
	}

	assert.Equal(t, start, AsStrings(res))
}

func TestRotateCounterClockwiseMulti(t *testing.T) {
	start := LoadStrings("../y2023/day14/input.txt")

	res := AsRunes(start)
	for i := 0; i < 4*10_000; i++ {
		res = RotateCounterClockwise(res)
	}

	assert.Equal(t, start, AsStrings(res))
}

func TestMakeSequence(t *testing.T) {
	assert.Equal(t, []int{}, MakeSequence(0, 0, false))
	assert.Equal(t, []int{1}, MakeSequence(1, 1, false))
	assert.Equal(t, []int{2, 3}, MakeSequence(2, 2, false))
	assert.Equal(t, []int{0, 1, 2, 3, 4}, MakeSequence(0, 5, false))
	assert.Equal(t, []int{6, 5, 4}, MakeSequence(6, 3, true))
}

func TestFrequencies(t *testing.T) {
	assert.Equal(t, map[string]int{"a": 3}, Frequencies([]string{"a", "a", "a"}))
	assert.Equal(t, map[string]int{"a": 2, "b": 2, "c": 1}, Frequencies([]string{"a", "b", "a", "c", "b"}))
}
