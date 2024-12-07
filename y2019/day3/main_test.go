package day3

import (
	"log"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

type Coord [2]int

func (c Coord) add(other Coord) Coord {
	return Coord{c[0] + other[0], c[1] + other[1]}
}

func (c Coord) length() int {
	return utils.AbsInt(c[0]) + utils.AbsInt(c[1])
}

func parsePath(line string) []Coord {
	segments := strings.Split(line, ",")

	cur := Coord{0, 0}
	res := []Coord{}

	for _, x := range segments {
		var path []Coord
		path, cur = walk(cur, x)
		res = append(res, path...)
	}
	return res
}

var directions = map[string]Coord{
	"L": {-1, 0},
	"R": {1, 0},
	"U": {0, 1},
	"D": {0, -1},
}

func walk(cur Coord, segment string) ([]Coord, Coord) {
	r := regexp.MustCompile(`([RULD])(\d+)`)
	x := r.FindStringSubmatch(segment)
	dir := x[1]
	length, _ := strconv.Atoi(x[2])

	path := make([]Coord, 0)
	for i := 0; i < length; i++ {
		cur = cur.add(directions[dir])
		path = append(path, cur)
	}

	return path, cur
}

func intersections(path1 []Coord, path2 []Coord) []Coord {
	res := make([]Coord, 0)
	for _, c := range path1 {
		if slices.Contains(path2, c) {
			res = append(res, c)
		}
	}
	return res
}

func solvePartA(input []string) int {
	path1 := parsePath(input[0])
	path2 := parsePath(input[1])

	log.Print(len(path1), len(path2))

	minimum := Coord{99999, 99999}
	for _, c := range intersections(path1, path2) {
		if c.length() < minimum.length() {
			minimum = c
		}
	}
	return minimum.length()
}

func TestParsePath(t *testing.T) {
	assert.Equal(t, []Coord{
		{1, 0}, {2, 0}, {3, 0}, {3, 1}, {3, 2}, {2, 2}, {2, 1}, {2, 0}, {2, -1},
	}, parsePath("R3,U2,L1,D3"))
}

func TestIntersections(t *testing.T) {
	assert.Equal(t, []Coord{
		{2, 0}, {4, 0},
	}, intersections([]Coord{
		{1, 0}, {2, 0}, {3, 0}, {4, 0},
	}, []Coord{
		{2, 0}, {4, 0}, {5, 0},
	}))
}

func TestSolvePartA(t *testing.T) {
	input := utils.LoadStrings("input.txt")
	assert.Equal(t, 316, solvePartA(input))
}
