package day12

import (
	"log"
	"slices"
)

type (
	Grid  [][]rune
	Coord struct{ l, c int }
)

func (c Coord) add(other Coord) Coord {
	return Coord{c.l + other.l, c.c + other.c}
}

func parseGrid(input []string) Grid {
	res := make(Grid, len(input))
	for l, line := range input {
		runes := make([]rune, len(line))
		for r, rune := range line {
			runes[r] = rune
		}
		res[l] = runes
	}
	return res
}

type Plot struct {
	rune          rune
	area          int
	circumference int
}

func (p Plot) price() int {
	return p.area * p.circumference
}

func solve(grid Grid) int {
	plots := make([]Plot, 0)
	seen := make(map[Coord]bool)
	for l, line := range grid {
		for r, rune := range line {
			_, ok := seen[Coord{l, r}]
			if ok {
				continue
			}

			plotCoords := []Coord{}
			circ := 0
			grid.flood(Coord{l, r}, &plotCoords, &circ)

			plots = append(plots, Plot{rune: rune, area: len(plotCoords), circumference: circ})

			for _, c := range plotCoords {
				seen[c] = true
			}
		}
	}

	sum := 0
	for _, p := range plots {
		log.Printf("%s: A=%d C=%d", string(p.rune), p.area, p.circumference)
		sum += p.price()
	}
	return sum
}

func (g Grid) flood(pos Coord, acc *[]Coord, circ *int) {
	rune := g.runeAt(pos)
	*acc = append(*acc, pos)

	neighbors := g.findNeighors(pos, rune)
	(*circ) += (4 - len(neighbors))
	for _, n := range neighbors {
		if !slices.Contains(*acc, n) {
			g.flood(n, acc, circ)
		}
	}
}

func (g Grid) findNeighors(pos Coord, rune rune) []Coord {
	res := make([]Coord, 0)
	for _, d := range []Coord{{-1, 0}, {0, 1}, {1, 0}, {0, -1}} {
		next := pos.add(d)
		if g.runeAt(next) == rune {
			res = append(res, next)
		}
	}
	return res
}

func (g *Grid) runeAt(pos Coord) rune {
	if pos.l < 0 || pos.c < 0 || pos.l >= len(*g) || pos.c >= len((*g)[0]) {
		return ' '
	}

	return []rune((*g)[pos.l])[pos.c]
}
