package day6

import (
	"strings"

	"github.com/alphaone/advent/utils"
)

type Orbits map[string]string

func (o Orbits) count(p string) int {
	return len(o.path(p))
}

func (o Orbits) path(p string) []string {
	c, ok := o[p]
	if !ok {
		return []string{}
	}

	return append([]string{c}, o.path(c)...)
}

func parse(input []string) Orbits {
	res := make(Orbits)
	for _, line := range input {
		parts := strings.Split(line, ")")
		c := parts[0]
		o := parts[1]
		res[o] = c
	}
	return res
}

func solvePartA(orbits Orbits) int {
	res := 0
	for o := range orbits {
		res += orbits.count(o)
	}
	return res
}

func solvePartB(o Orbits) int {
	youPath := o.path("YOU")
	sanPath := o.path("SAN")

	for {
		youOrbit := youPath[len(youPath)-1]
		sanOrbit := sanPath[len(sanPath)-1]
		if youOrbit == sanOrbit {
			youPath = utils.RemoveIndex(youPath, len(youPath)-1)
			sanPath = utils.RemoveIndex(sanPath, len(sanPath)-1)
		} else {
			break
		}
	}

	return len(youPath) + len(sanPath)
}
