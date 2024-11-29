package day18

import (
	"math/big"
	"regexp"
	"strconv"
)

type Pos struct {
	X, Y int
}

func parseInputA(input []string) ([]Pos, int) {
	coords := []Pos{{0, 0}}
	x, y := 0, 0
	perimeter := 0

	for _, line := range input {
		re := regexp.MustCompile(`([RLUD]) (\d+)`)
		res := re.FindStringSubmatch(line)

		dir := res[1][0]
		steps, _ := strconv.Atoi(res[2])

		switch dir {
		case 'R':
			x += steps
		case 'L':
			x -= steps
		case 'U':
			y -= steps
		case 'D':
			y += steps
		}
		perimeter += steps
		coords = append(coords, Pos{x, y})
	}
	return coords, perimeter
}

func convertBtoA(input []string) []string {
	result := []string{}
	for _, line := range input {
		re := regexp.MustCompile(`#(\w{5})(\w)`)
		res := re.FindStringSubmatch(line)

		n := new(big.Int)
		n.SetString(res[1], 16)

		// The last hexadecimal digit encodes the direction to dig: 0 means R, 1 means D, 2 means L, and 3 means U.

		switch res[2] {
		case "0":
			result = append(result, "R "+n.String())
		case "1":
			result = append(result, "D "+n.String())
		case "2":
			result = append(result, "L "+n.String())
		case "3":
			result = append(result, "U "+n.String())
		}
	}
	return result
}

type Loop []Pos

func (l Loop) Len() int { return len(l) }
func (l Loop) Get(i int) Pos {
	return l[saneMod(i, l.Len())]
}

// also works for negative numbers
func saneMod(a, b int) int {
	return (a%b + b) % b
}

// Shoelace formula + perimeter
func calculateArea(perimeter int, positions []Pos) int {
	l := Loop(positions)
	area := 0
	for i := range l {
		area += l.Get(i).X * (l.Get(i+1).Y - l.Get(i-1).Y)
	}

	return area/2 + perimeter/2 + 1
}

func SolvePartA(input []string) int {
	coords, perimeter := parseInputA(input)
	return calculateArea(perimeter, coords)
}

func SolvePartB(input []string) int {
	input = convertBtoA(input)
	coords, perimeter := parseInputA(input)
	return calculateArea(perimeter, coords)
}
