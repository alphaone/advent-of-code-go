package day12

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	assert.True(t, isValid(parseLine("# 1")))

	assert.True(t, isValid(parseLine("#.#.### 1,1,3")))
	assert.False(t, isValid(parseLine("#.#.### 1,2,3")))

	assert.True(t, isValid(parseLine("#....######..#####. 1,6,5")))
	assert.False(t, isValid(parseLine("#....###.###..#####. 1,6,5")))

	assert.False(t, isValid(parseLine("#....##?###..#####. 1,6,5")))
}

func TestCountOptions(t *testing.T) {
	// assert.Equal(t, 1, countOptions(Report{Diagnose: "#", Checksum: []int{1}}))
	// assert.Equal(t, 1, countOptions(Report{Diagnose: ".#.", Checksum: []int{1}}))
	// assert.Equal(t, 1, countOptions(Report{Diagnose: ".?.", Checksum: []int{1}}))
	// assert.Equal(t, 1, countOptions(Report{Diagnose: "???.###", Checksum: []int{1, 1, 3}}))
	// assert.Equal(t, 4, countOptions(Report{Diagnose: ".??..??...?##.", Checksum: []int{1, 1, 3}}))
	// assert.Equal(t, 10, countOptions(Report{Diagnose: "?###????????", Checksum: []int{3, 2, 1}}))
	assert.Equal(t, 2, countOptions(0, Report{Diagnose: ".??.##.", Checksum: []int{1, 2}}))
}

var example = []string{
	"???.### 1,1,3",
	".??..??...?##. 1,1,3",
	"?#?#?#?#?#?#?#? 1,3,1,6",
	"????.#...#... 4,1,1",
	"????.######..#####. 1,6,5",
	"?###???????? 3,2,1",
}

// func TestExampleA(t *testing.T) {
// 	assert.Equal(t, 21, solvePartA(example))
// }
//
// func TestSolutionA(t *testing.T) {
// 	input := utils.LoadStrings("input.txt")
// 	assert.Equal(t, 7599, solvePartA(input))
// }
