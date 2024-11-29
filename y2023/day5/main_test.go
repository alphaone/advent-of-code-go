package day5

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"seeds: 79 14 55 13",
	"",
	"seed-to-soil map:",
	"50 98 2",
	"52 50 48",
	"",
	"soil-to-fertilizer map:",
	"0 15 37",
	"37 52 2",
	"39 0 15",
	"",
	"fertilizer-to-water map:",
	"49 53 8",
	"0 11 42",
	"42 0 7",
	"57 7 4",
	"",
	"water-to-light map:",
	"88 18 7",
	"18 25 70",
	"",
	"light-to-temperature map:",
	"45 77 23",
	"81 45 19",
	"68 64 13",
	"",
	"temperature-to-humidity map:",
	"0 69 1",
	"1 0 69",
	"",
	"humidity-to-location map:",
	"60 56 37",
	"56 93 4",
}

func TestParseGarden(t *testing.T) {
	expected := Garden{
		Seeds:      []int{79, 14, 55, 13},
		SeedRanges: []SeedRange{{Start: 79, Count: 14}, {Start: 55, Count: 13}},
		ConverionPipeline: []ConversionMap{
			{
				Name: "seed-to-soil map:",
				Ranges: []LocationRange{
					{DescStart: 50, SrcStart: 98, Count: 2},
					{DescStart: 52, SrcStart: 50, Count: 48},
				},
			},
			{
				Name: "soil-to-fertilizer map:",
				Ranges: []LocationRange{
					{DescStart: 0, SrcStart: 15, Count: 37},
					{DescStart: 37, SrcStart: 52, Count: 2},
					{DescStart: 39, SrcStart: 0, Count: 15},
				},
			},
			{
				Name: "fertilizer-to-water map:",
				Ranges: []LocationRange{
					{DescStart: 49, SrcStart: 53, Count: 8},
					{DescStart: 0, SrcStart: 11, Count: 42},
					{DescStart: 42, SrcStart: 0, Count: 7},
					{DescStart: 57, SrcStart: 7, Count: 4},
				},
			},
			{
				Name: "water-to-light map:",
				Ranges: []LocationRange{
					{DescStart: 88, SrcStart: 18, Count: 7},

					{DescStart: 18, SrcStart: 25, Count: 70},
				},
			},
			{
				Name: "light-to-temperature map:",
				Ranges: []LocationRange{
					{DescStart: 45, SrcStart: 77, Count: 23},
					{DescStart: 81, SrcStart: 45, Count: 19},
					{DescStart: 68, SrcStart: 64, Count: 13},
				},
			},
			{
				Name: "temperature-to-humidity map:",
				Ranges: []LocationRange{
					{DescStart: 0, SrcStart: 69, Count: 1},
					{DescStart: 1, SrcStart: 0, Count: 69},
				},
			},
			{
				Name: "humidity-to-location map:",
				Ranges: []LocationRange{
					{DescStart: 60, SrcStart: 56, Count: 37},
					{DescStart: 56, SrcStart: 93, Count: 4},
				},
			},
		},
	}
	assert.Equal(t, expected, parseGarden(example))
}

func TestParseMap(t *testing.T) {
	input := []string{
		"seed-to-soil map:",
		"50 98 2",
		"52 50 48",
	}
	expected := ConversionMap{
		Name: "seed-to-soil map:",
		Ranges: []LocationRange{
			{DescStart: 50, SrcStart: 98, Count: 2},
			{DescStart: 52, SrcStart: 50, Count: 48},
		},
	}
	assert.Equal(t, expected, parseMap(input))
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 35, solvePartA(example))
}

func TestSolutionA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 178159714, solvePartA(input))
}

func TestConvertRange(t *testing.T) {
	rng := LocationRange{50, 98, 2}
	assert.Equal(t, 50, rng.Convert(98))
	assert.Equal(t, 51, rng.Convert(99))

	assert.Equal(t, 52, rng.Convert(52))
}

func TestConvertMap(t *testing.T) {
	m := ConversionMap{
		Name: "seed-to-soil map:",
		Ranges: []LocationRange{
			{DescStart: 50, SrcStart: 98, Count: 2},
			{DescStart: 52, SrcStart: 50, Count: 48},
		},
	}
	assert.Equal(t, 81, m.Convert(79))
	assert.Equal(t, 14, m.Convert(14))
	assert.Equal(t, 57, m.Convert(55))
	assert.Equal(t, 13, m.Convert(13))
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 46, solvePartB(example))
}

func TestSolutionB(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 100165128, solvePartB(input))
}
