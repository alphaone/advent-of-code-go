package day15

import (
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils/sliceutils"
)

type game []int

func parse(input string) game {
	g := game{}
	for _, x := range strings.Split(input, ",") {
		i, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		g = append(g, i)
	}
	return g
}

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
