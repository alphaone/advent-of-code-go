package sliceutils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

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
