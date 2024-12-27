package day15

import (
	"strings"
)

type (
	Grid  [][]rune
	Coord struct{ L, C int }
)

func (c Coord) add(other Coord) Coord {
	return Coord{c.L + other.L, c.C + other.C}
}

func (g *Grid) FindCoord(lookingfor rune) []Coord {
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

func (g Grid) Visual(pos Coord) string {
	res := ""
	for l, line := range g {
		for c, rune := range line {
			if pos.L == l && pos.C == c {
				res += "@"
			} else {
				res += string(rune)
			}
		}
		res += "\n"
	}
	return res
}

var Directions = map[rune]Coord{
	'^': {-1, 0},
	'v': {1, 0},
	'<': {0, -1},
	'>': {0, 1},
}

func (g *Grid) Move(pos Coord, dir Coord) Coord {
	if !g.CanMove(pos, dir) {
		return pos
	}

	g.PushBlock(pos, dir)
	return pos.add(dir)
}

func (g *Grid) CanMove(pos Coord, dir Coord) bool {
	next := pos.add(dir)
	switch (*g)[next.L][next.C] {
	case '#':
		return false
	case '.':
		return true
	case 'O':
		return g.CanMove(next, dir)
	case '[':
		if dir.L == 0 {
			// horizontal
			if dir.C == 1 {
				return g.CanMove(next.add(dir), dir)
			} else {
				panic("cannot move inner block [ to left")
			}
		} else {
			// vertical
			return g.CanMove(next, dir) && g.CanMove(next.add(Coord{0, 1}), dir)
		}
	case ']':
		if dir.L == 0 {
			// horizontal
			if dir.C == -1 {
				return g.CanMove(next.add(dir), dir)
			} else {
				panic("cannot move inner block ] to right")
			}
		} else {
			// vertical
			return g.CanMove(next, dir) && g.CanMove(next.add(Coord{0, -1}), dir)
		}
	}
	panic("unreachable")
}

func (g *Grid) PushBlock(pos Coord, dir Coord) {
	nextPos := pos.add(dir)
	switch (*g)[nextPos.L][nextPos.C] {
	case '.':
	case 'O':
		g.PushBlock(nextPos, dir)
	case '[':
		if dir.L == 0 {
			// horizontal
			g.PushBlock(nextPos, dir)
		} else {
			// vertical
			g.PushBlock(nextPos, dir)
			g.PushBlock(nextPos.add(Coord{0, 1}), dir)
			(*g)[nextPos.L][nextPos.C+1] = '.'
		}
	case ']':
		if dir.L == 0 {
			// horizontal
			g.PushBlock(nextPos, dir)
		} else {
			// vertical
			g.PushBlock(nextPos, dir)
			g.PushBlock(nextPos.add(Coord{0, -1}), dir)
			(*g)[nextPos.L][nextPos.C-1] = '.'
		}
	}
	(*g)[nextPos.L][nextPos.C] = (*g)[pos.L][pos.C]
}

func Parse(input string) (Grid, []rune) {
	parts := strings.Split(input, "\n\n")

	grid := make(Grid, 0)
	for _, line := range strings.Split(parts[0], "\n") {
		grid = append(grid, []rune(line))
	}

	return grid, []rune(strings.Join(strings.Split(parts[1], "\n"), ""))
}

func ParseB(input string) (Grid, []rune) {
	parts := strings.Split(input, "\n\n")

	grid := make(Grid, 0)
	for _, line := range strings.Split(parts[0], "\n") {
		runeLine := make([]rune, 0)
		for _, x := range line {
			switch x {
			case '#':
				runeLine = append(runeLine, '#', '#')
			case 'O':
				runeLine = append(runeLine, '[', ']')
			case '.':
				runeLine = append(runeLine, '.', '.')
			case '@':
				runeLine = append(runeLine, '@', '.')
			}
		}
		grid = append(grid, runeLine)

	}

	return grid, []rune(strings.Join(strings.Split(parts[1], "\n"), ""))
}

func solve(g Grid, movements []rune) int {
	pos := g.FindCoord('@')[0]
	g[pos.L][pos.C] = '.'

	for _, m := range movements {
		pos = g.Move(pos, Directions[m])
	}

	sum := 0
	for _, box := range g.FindCoord('O') {
		sum += box.L*100 + box.C
	}
	for _, box := range g.FindCoord('[') {
		sum += box.L*100 + box.C
	}

	return sum
}
