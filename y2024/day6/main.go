package day6

import (
	"errors"
	"sync"
	"sync/atomic"

	"golang.org/x/exp/maps"
	"golang.org/x/exp/slices"
)

type Coord struct {
	l, c int
}

type (
	Grid [][]rune
	Path map[Coord][]rune
)

func (c Coord) add(other Coord) Coord {
	return Coord{c.l + other.l, c.c + other.c}
}

func (c Coord) rotateClockwise() Coord {
	switch c {
	case Coord{-1, 0}:
		return Coord{0, 1}
	case Coord{0, 1}:
		return Coord{1, 0}
	case Coord{1, 0}:
		return Coord{0, -1}
	case Coord{0, -1}:
		return Coord{-1, 0}
	}
	panic("cannot rotate")
}

func (c Coord) AsRune() rune {
	switch c {
	case Coord{-1, 0}:
		return 'N'
	case Coord{0, 1}:
		return 'E'
	case Coord{1, 0}:
		return 'S'
	case Coord{0, -1}:
		return 'W'
	}
	panic("invalid dir")
}

func (g *Grid) findCoord(lookingfor rune) Coord {
	for l, line := range *g {
		for r, rune := range line {
			if rune == lookingfor {
				return Coord{l, r}
			}
		}
	}
	panic("not found")
}

func (g *Grid) charAtPos(pos Coord) (rune, error) {
	if pos.l < 0 || pos.c < 0 || pos.l >= len(*g) || pos.c >= len((*g)[0]) {
		return ' ', errors.New("out of bounds")
	}

	return []rune((*g)[pos.l])[pos.c], nil
}

func (g *Grid) alterAt(pos Coord, replace rune) Grid {
	res := make([][]rune, len(*g))
	for l, line := range *g {
		newLine := make([]rune, len(line))
		for r, char := range line {
			newLine[r] = char
		}
		res[l] = newLine
	}
	res[pos.l][pos.c] = replace
	return res
}

func (g *Grid) walk(pos Coord, dir Coord) (Coord, Coord) {
	target := pos.add(dir)
	char, err := g.charAtPos(target)
	if err != nil {
		return pos, Coord{0, 0}
	}
	if char == '.' || char == '^' {
		return target, dir
	}
	if char == '#' {
		return pos, dir.rotateClockwise()
	}
	panic("unknown obstacle")
}

func (g *Grid) pathWalked() Path {
	pos := g.findCoord('^')
	dir := Coord{-1, 0}
	path := Path{pos: []rune{'N'}}
	for {
		pos, dir = g.walk(pos, dir)
		if dir.l == 0 && dir.c == 0 {
			break
		}

		_, ok := path[pos]
		if !ok {
			path[pos] = make([]rune, 0)
		}
		path[pos] = append(path[pos], dir.AsRune())
	}
	return path
}

func (g *Grid) isLoop(pos Coord, dir Coord) bool {
	path := Path{pos: []rune{'N'}}

	for {
		pos, dir = g.walk(pos, dir)

		if dir.l == 0 && dir.c == 0 {
			return false
		}

		seen, ok := path[pos]
		if !ok {
			path[pos] = make([]rune, 0)
		}
		if ok && slices.Contains(seen, dir.AsRune()) {
			return true
		}

		path[pos] = append(path[pos], dir.AsRune())
	}
}

func solvePartA(g Grid) int {
	return len(g.pathWalked())
}

func solvePartB(g Grid) int {
	pos := g.findCoord('^')
	path := g.pathWalked()

	var count int32 = 0
	var wg sync.WaitGroup
	for _, c := range maps.Keys(path) {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if pos.l == c.l && pos.c == c.c {
				return
			}

			grid := g.alterAt(c, '#')
			if grid.isLoop(pos, Coord{-1, 0}) {
				atomic.AddInt32(&count, 1)
			}
		}()
	}
	wg.Wait()

	return int(count)
}
