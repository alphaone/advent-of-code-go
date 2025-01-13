package day11

import (
	"github.com/alphaone/advent/utils"
	"github.com/alphaone/advent/utils/mathutils"
	"github.com/alphaone/advent/utils/sliceutils"
)

func FindRifts(input []string) ([]int, []int) {
	lineRifts := findEmptyLines(input)
	transposed := sliceutils.Transpose(utils.AsRunes(input))
	colRifts := findEmptyLines(utils.AsStrings(transposed))
	return lineRifts, colRifts
}

func findEmptyLines(input []string) []int {
	res := []int{}
	for i, line := range input {
		if isEmptyLine(line) {
			res = append(res, i)
		}
	}
	return res
}

func isEmptyLine(input string) bool {
	for _, char := range input {
		if char != '.' {
			return false
		}
	}
	return true
}

type Galaxy struct {
	Id   int
	Line int
	Col  int
}

func Galaxies(input []string) []Galaxy {
	var res []Galaxy
	for l, line := range input {
		for c, char := range line {
			if char == '#' {
				res = append(res, Galaxy{Id: len(res), Line: l, Col: c})
			}
		}
	}
	return res
}

func Distance(rate int, lineRifts, colRifts []int, g1, g2 Galaxy) int {
	noOfRifts := len(filterInRange(lineRifts, g1.Line, g2.Line)) + len(filterInRange(colRifts, g1.Col, g2.Col))
	return mathutils.AbsDiffInt(g1.Line, g2.Line) + mathutils.AbsDiffInt(g1.Col, g2.Col) + (noOfRifts * (rate - 1))
}

func filterInRange(elements []int, min, max int) []int {
	if min > max {
		min, max = max, min
	}
	var res []int
	for _, e := range elements {
		if e >= min && e <= max {
			res = append(res, e)
		}
	}
	return res
}

func Pairs(galaxies []Galaxy) [][]Galaxy {
	var res [][]Galaxy
	for i, g1 := range galaxies {
		for j, g2 := range galaxies {
			if i < j {
				res = append(res, []Galaxy{g1, g2})
			}
		}
	}
	return res
}

func solvePart(input []string, rate int) int {
	res := 0
	pairs := Pairs(Galaxies(input))
	lineRifts, colRifts := FindRifts(input)
	for _, pair := range pairs {
		res += Distance(rate, lineRifts, colRifts, pair[0], pair[1])
	}
	return res
}
