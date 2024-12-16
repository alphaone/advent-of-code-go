package day4

import (
	"errors"
)

func solvePartA(g Grid) int {
	count := 0
	for _, p := range g.allPos('X') {
		count += g.countMas(p)
	}
	return count
}

func solvePartB(g Grid) int {
	count := 0
	for _, p := range g.allPos('A') {
		if g.isCross(p) {
			count += 1
		}
	}
	return count
}

type Coord struct{ l, c int }

type Grid []string

func (c Coord) add(other Coord) Coord {
	return Coord{c.l + other.l, c.c + other.c}
}

func (g Grid) allPos(lookingfor rune) []Coord {
	res := make([]Coord, 0)
	for l, line := range g {
		for r, rune := range line {
			if rune == lookingfor {
				res = append(res, Coord{l, r})
			}
		}
	}
	return res
}

func (g Grid) charAtPos(pos Coord) (rune, error) {
	if pos.l < 0 || pos.c < 0 || pos.l >= len(g) || pos.c >= len(g[0]) {
		return ' ', errors.New("out of bounds")
	}

	return []rune(g[pos.l])[pos.c], nil
}

func (g Grid) wordFromPos(pos Coord, dir Coord, length int) string {
	res := ""
	for range make([]int, length) {
		pos = pos.add(dir)
		v, _ := g.charAtPos(pos)
		res += string(v)
	}
	return res
}

func (g Grid) countMas(pos Coord) int {
	allDirs := []Coord{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	res := 0
	for _, dir := range allDirs {
		mas := g.wordFromPos(pos, dir, 3)
		if mas == "MAS" {
			res += 1
		}
	}

	return res
}

func (g Grid) isCross(pos Coord) bool {
	ul, _ := g.charAtPos(pos.add(Coord{-1, -1}))
	ur, _ := g.charAtPos(pos.add(Coord{-1, 1}))
	ll, _ := g.charAtPos(pos.add(Coord{1, -1}))
	lr, _ := g.charAtPos(pos.add(Coord{1, 1}))

	return (ul == 'M' && lr == 'S' || ul == 'S' && lr == 'M') &&
		(ur == 'M' && ll == 'S' || ur == 'S' && ll == 'M')
}
