package day3

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils"
)

type (
	Coord   struct{ x, y int }
	Segment struct {
		start  Coord
		end    Coord
		length int
	}
)

func (c Coord) add(other Coord) Coord {
	return Coord{c.x + other.x, c.y + other.y}
}

func (c Coord) subtract(other Coord) Coord {
	return Coord{c.x - other.x, c.y - other.y}
}

func (c Coord) scale(length int) Coord {
	return Coord{c.x * length, c.y * length}
}

func (c Coord) length() int {
	return utils.AbsInt(c.x) + utils.AbsInt(c.y)
}

func (c Coord) lengthAlongPath(path []Segment) int {
	res := 0
	for _, s := range path {
		if !isOn(s, c) {
			res += s.length
		} else {
			res += c.subtract(s.start).length()
			return res
		}
	}

	return res
}

var r = regexp.MustCompile(`([RULD])(\d+)`)

func parsePath(line string) []Segment {
	strSegments := strings.Split(line, ",")

	cur := Coord{}
	res := []Segment{}

	for _, s := range strSegments {
		x := r.FindStringSubmatch(s)
		dir := x[1]
		length, _ := strconv.Atoi(x[2])

		next := cur.add(directions[dir].scale(length))
		res = append(res, Segment{
			start:  cur,
			end:    next,
			length: length,
		})
		cur = next
	}
	return res
}

var directions = map[string]Coord{
	"L": {-1, 0},
	"R": {1, 0},
	"U": {0, 1},
	"D": {0, -1},
}

func intersections(path1 []Segment, path2 []Segment) []Coord {
	res := make([]Coord, 0)
	for _, s1 := range path1 {
		for _, s2 := range path2 {
			res = append(res, intersection(s1, s2)...)
		}
	}
	return res
}

func intersection(s1 Segment, s2 Segment) []Coord {
	// one needs to be verical the other horizontal
	if (s1.start.x == s1.end.x) == (s2.start.x == s2.end.x) {
		return []Coord{}
	}

	var vertical, horizontal Segment
	if s1.start.x == s1.end.x {
		vertical = s1
		horizontal = s2
	} else {
		vertical = s2
		horizontal = s1
	}

	return intersectionHV(horizontal, vertical)
}

func isOn(s Segment, p Coord) bool {
	if s.start.x == s.end.x {
		minY := min(s.start.y, s.end.y)
		maxY := max(s.start.y, s.end.y)

		if p.x == s.start.x && minY <= p.y && p.y <= maxY {
			return true
		}
	} else {
		minX := min(s.start.x, s.end.x)
		maxX := max(s.start.x, s.end.x)
		if p.y == s.start.y && minX <= p.x && p.x <= maxX {
			return true
		}
	}
	return false
}

func intersectionHV(horizontal Segment, vertical Segment) []Coord {
	minX := min(horizontal.start.x, horizontal.end.x)
	maxX := max(horizontal.start.x, horizontal.end.x)
	minY := min(vertical.start.y, vertical.end.y)
	maxY := max(vertical.start.y, vertical.end.y)

	if minX <= vertical.start.x && vertical.start.x <= maxX &&
		minY <= horizontal.start.y && horizontal.start.y <= maxY {
		return []Coord{{vertical.start.x, horizontal.start.y}}
	}
	return []Coord{}
}

func solvePartA(input []string) int {
	path1 := parsePath(input[0])
	path2 := parsePath(input[1])

	minLength := 999999
	for _, c := range intersections(path1, path2) {
		if c.length() == 0 {
			continue
		}
		length := c.length()
		if length < minLength {
			minLength = length
		}
	}
	return minLength
}

func solvePartB(input []string) int {
	path1 := parsePath(input[0])
	path2 := parsePath(input[1])

	minLength := 999999
	for _, c := range intersections(path1, path2) {
		if c.length() == 0 {
			continue
		}

		length := c.lengthAlongPath(path1) + c.lengthAlongPath(path2)
		if length < minLength {
			minLength = length
		}
	}
	return minLength
}
