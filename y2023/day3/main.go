package day3

import (
	"fmt"
	"regexp"
	"strconv"
)

type PartNumber struct {
	No    int
	Start Coord
}
type (
	Coord struct{ Line, Col int }
	Box   struct{ TopLeft, BottomRight Coord }
)

func solvePartA(input []string) int {
	pns := parsePartNumbers(input)
	var result int
	for _, pn := range pns {
		if pn.Valid(input) {
			result += pn.No
		}
	}
	return result
}

func parsePartNumbers(input []string) []PartNumber {
	var result []PartNumber
	re := regexp.MustCompile(`\d+`)
	for lineNo, line := range input {
		parseIdx := re.FindAllStringIndex(line, -1)
		parseResult := re.FindAllString(line, -1)
		for idx, noString := range parseResult {
			no, _ := strconv.Atoi(noString)
			result = append(result, PartNumber{
				No:    no,
				Start: Coord{lineNo, parseIdx[idx][0]},
			})
		}
	}
	return result
}

func (pn PartNumber) Valid(context []string) bool {
	re := regexp.MustCompile(`[^\d\.]`)
	return re.MatchString(pn.Surounding(context))
}

func (pn PartNumber) SuroundingBox(context []string) Box {
	minLine := max(pn.Start.Line-1, 0)
	maxLine := min(pn.Start.Line+1, len(context)-1)
	minCol := max(pn.Start.Col-1, 0)
	maxCol := min(pn.Start.Col+len(fmt.Sprint(pn.No)), len(context[pn.Start.Line])-1)

	return Box{
		TopLeft:     Coord{minLine, minCol},
		BottomRight: Coord{maxLine, maxCol},
	}
}

func (pn PartNumber) Surounding(context []string) string {
	box := pn.SuroundingBox(context)

	var res string
	for line := box.TopLeft.Line; line <= box.BottomRight.Line; line++ {
		res += context[line][box.TopLeft.Col : box.BottomRight.Col+1]
	}
	return res
}

func findGears(input []string) []Coord {
	var result []Coord
	for lineNo, line := range input {
		for colNo, char := range line {
			if char == '*' {
				result = append(result, Coord{lineNo, colNo})
			}
		}
	}
	return result
}

func (c Coord) GearRatio(parts []PartNumber, context []string) int {
	var foundParts []PartNumber
	for _, part := range parts {
		box := part.SuroundingBox(context)
		if box.Contains(c) {
			foundParts = append(foundParts, part)
		}
	}
	if len(foundParts) == 2 {
		return foundParts[0].No * foundParts[1].No
	} else {
		return 0
	}
}

func (b Box) Contains(c Coord) bool {
	return b.TopLeft.Line <= c.Line && c.Line <= b.BottomRight.Line &&
		b.TopLeft.Col <= c.Col && c.Col <= b.BottomRight.Col
}

func solvePartB(input []string) int {
	parts := parsePartNumbers(input)
	gears := findGears(input)
	var result int
	for _, gear := range gears {
		result += gear.GearRatio(parts, input)
	}
	return result
}
