package day2

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleA = []Report{
	{7, 6, 4, 2, 1},
	{1, 2, 7, 8, 9},
	{9, 7, 6, 2, 1},
	{1, 3, 2, 4, 5},
	{8, 6, 4, 4, 1},
	{1, 3, 6, 7, 9},
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 2, solvePartA(exampleA))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadIntCols("input.txt")
	assert.Equal(t, 356, solvePartA(parseInput(input)))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 4, solvePartB(exampleA))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadIntCols("input.txt")
	assert.Equal(t, 413, solvePartB(parseInput(input)))
}

var tests = []struct {
	name   string
	report Report
	safe   bool
}{
	{name: "increasing", report: Report{1, 2, 3}, safe: true},
	{name: "not increasing", report: Report{1, 2, 2}, safe: false},
	{name: "increasing too steep", report: Report{1, 2, 6}, safe: false},
	{name: "decreasing", report: Report{3, 2, 1}, safe: true},
	{name: "not decreasing", report: Report{4, 2, 2}, safe: false},
	{name: "decreasing too steep", report: Report{6, 2, 1}, safe: false},
}

func TestSafe(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			assert.Equal(t, tt.safe, tt.report.Safe())
		})
	}
}

func TestDamp(t *testing.T) {
	input := Report{1, 2, 3, 4}
	expected := []Report{
		{2, 3, 4},
		{1, 3, 4},
		{1, 2, 4},
		{1, 2, 3},
	}

	assert.Equal(t, expected, input.Damp())
}
