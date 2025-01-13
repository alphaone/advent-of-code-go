package day4

import (
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils/sliceutils"
)

type Assignment struct {
	first  []int
	second []int
}

func parse(input []string) []Assignment {
	res := make([]Assignment, 0)
	for _, line := range input {
		res = append(res, parseLine(line))
	}
	return res
}

func parseLine(line string) Assignment {
	parts := strings.Split(line, ",")
	if len(parts) != 2 {
		panic("line does not have two elves")
	}
	return Assignment{first: expandRange(parts[0]), second: expandRange(parts[1])}
}

func solveA(input []string) int {
	parsed := parse(input)

	contained := 0
	for _, assignment := range parsed {
		if fullyContained(assignment) {
			contained++
		}
	}
	return contained
}

func solveB(input []string) int {
	parsed := parse(input)

	contained := 0
	for _, assignment := range parsed {
		if overlaps(assignment) {
			contained++
		}
	}
	return contained
}

func expandRange(rangeStr string) []int {
	// start, end := rangeStr.split('-')
	parts := strings.Split(rangeStr, "-")
	if len(parts) != 2 {
		panic("malformed range")
	}
	start, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	end, err := strconv.Atoi(parts[1])
	if err != nil {
		panic(err)
	}

	res := make([]int, 0)
	for i := start; i <= end; i++ {
		res = append(res, i)
	}
	return res
}

func fullyContained(a Assignment) bool {
	i := sliceutils.Intersection(a.first, a.second)
	return len(i) == len(a.first) || len(i) == len(a.second)
}

func overlaps(a Assignment) bool {
	i := sliceutils.Intersection(a.first, a.second)
	return len(i) > 0
}
