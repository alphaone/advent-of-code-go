package day15

import "github.com/alphaone/advent/utils/sliceutils"

type game []int

func (g game) next() game {
	last := g[len(g)-1]

	x := sliceutils.LastIndexFunc(g[:len(g)-1], func(x int) bool { return x == last })
	if x == -1 {
		return append(g, 0)
	}

	return append(g, len(g)-1-x)
}

func (g game) play(turns int) int {
	for len(g) < turns {
		g = g.next()
	}
	return g[len(g)-1]
}
