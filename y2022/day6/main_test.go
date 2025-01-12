package day6

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

type startOfSeqTestCases struct {
	name  string
	input string
	exp   int
}

func TestStartOfPacket(t *testing.T) {
	for _, tc := range []startOfSeqTestCases{
		{name: "Example 1", input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", exp: 7},
		{name: "Example 2", input: "bvwbjplbgvbhsrlpgdmjqwftvncz", exp: 5},
		{name: "Example 3", input: "nppdvjthqldpwncqszvftbrmjlhg", exp: 6},
		{name: "Example 4", input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", exp: 10},
		{name: "Example 5", input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", exp: 11},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp, startOfSeq(tc.input, 4))
		})
	}
}

func TestStartOfMessage(t *testing.T) {
	for _, tc := range []startOfSeqTestCases{
		{name: "Example 1", input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb", exp: 19},
		{name: "Example 2", input: "bvwbjplbgvbhsrlpgdmjqwftvncz", exp: 23},
		{name: "Example 3", input: "nppdvjthqldpwncqszvftbrmjlhg", exp: 23},
		{name: "Example 4", input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", exp: 29},
		{name: "Example 5", input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", exp: 26},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp, startOfSeq(tc.input, 14))
		})
	}
}

func TestSolve(t *testing.T) {
	assert.Equal(t, 1042, startOfSeq(utils.LoadString("input.txt"), 4))
	assert.Equal(t, 2980, startOfSeq(utils.LoadString("input.txt"), 14))
}
