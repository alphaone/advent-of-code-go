package day15

import (
	"slices"
	"strings"
)

type (
	Grid  [][]rune
	Coord struct{ l, c int }
)

func (c Coord) add(other Coord) Coord {
	return Coord{c.l + other.l, c.c + other.c}
}

func (g *Grid) findCoord(lookingfor rune) []Coord {
	res := make([]Coord, 0)
	for l, line := range *g {
		for r, rune := range line {
			if rune == lookingfor {
				res = append(res, Coord{l, r})
			}
		}
	}
	return res
}

func (g Grid) String() string {
	res := ""
	for _, line := range g {
		res += string(line) + "\n"
	}
	return res
}

var directions = map[rune]Coord{
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

func (g *Grid) move(pos Coord, dir Coord) Coord {
	if !g.canMove(pos, dir) {
		return pos
	}

	g.pushBlock(pos, dir)
	return pos.add(dir)
}

func (g *Grid) canMove(pos Coord, dir Coord) bool {
	blocks := g.blocksInDir(pos, dir)
	return slices.Contains(blocks, '.')
}

func (g *Grid) blocksInDir(pos Coord, dir Coord) []rune {
	next := pos.add(dir)
	if next.l < 0 || next.l >= len(*g) || next.c < 0 || next.c >= len((*g)[0]) {
		return []rune{}
	}
	if (*g)[next.l][next.c] == '#' {
		return []rune{}
	}

	return append([]rune{(*g)[next.l][next.c]}, g.blocksInDir(next, dir)...)
}

func (g *Grid) pushBlock(pos Coord, dir Coord) {
	cur := (*g)[pos.l][pos.c]
	nextPos := pos.add(dir)
	if (*g)[nextPos.l][nextPos.c] != '.' {
		g.pushBlock(nextPos, dir)
	}
	(*g)[nextPos.l][nextPos.c] = cur
}

func parse(input string) (Grid, []rune) {
	parts := strings.Split(input, "\n\n")

	grid := make(Grid, 0)
	for _, line := range strings.Split(parts[0], "\n") {
		grid = append(grid, []rune(line))
	}

	return grid, []rune(strings.Join(strings.Split(parts[1], "\n"), ""))
}

func solvePartA(g Grid, movements []rune) int {
	pos := g.findCoord('@')[0]
	g[pos.l][pos.c] = '.'

	for i, m := range movements {
		pos = g.move(pos, directions[m])
	}

	sum := 0
	for _, box := range g.findCoord('O') {
		sum += box.l*100 + box.c
	}

	return sum
}
