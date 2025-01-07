package day10

import (
	"slices"

	"github.com/alphaone/advent/y2019/day10/vectors"
)

type Coord struct {
	X, Y int
}
type Map []Coord

func parseMap(input []string) Map {
	coords := make(Map, 0)
	for i, row := range input {
		for j, cell := range row {
			if cell == '.' {
				continue
			}
			coords = append(coords, Coord{X: j, Y: i})
		}
	}

	return coords
}

func (m Map) Visible(pov Coord) map[vectors.Vector]bool {
	vs := make(map[vectors.Vector]bool)
	for _, c := range m {
		if c.X == pov.X && c.Y == pov.Y {
			continue
		}

		vs[vectors.New(c.X-pov.X, c.Y-pov.Y).Reduce()] = true
	}

	return vs
}

func sortVectors(m map[vectors.Vector]bool) []vectors.Vector {
	vs := make([]vectors.Vector, 0)
	for v := range m {
		vs = append(vs, v)
	}
	slices.SortFunc(vs, SortCounterClockwiseFunc)
	return vs
}

func solve(m Map) int {
	max := 0

	for _, pov := range m {
		count := len(m.Visible(pov))
		if count > max {
			max = count
		}
	}
	return max
}

func SortCounterClockwiseFunc(a, b vectors.Vector) int {
	north := vectors.New(0, -1)
	if north.Angle(a) < north.Angle(b) {
		return -1
	} else if north.Angle(a) > north.Angle(b) {
		return 1
	} else {
		return 0
	}
}

func solveB(m Map, n int, pov Coord) int {
	vs := m.Visible(pov)
	sorted := sortVectors(vs)

	c := sorted[n-1]
	c = vectors.New(c.X+pov.X, c.Y+pov.Y)

	return c.X*100 + c.Y
}
