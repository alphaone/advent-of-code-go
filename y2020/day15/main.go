package day15

import (
	"strconv"
	"strings"
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

func (g game) play(targetTurns int) int {
	turnNumbers := map[int]int{}
	for i, x := range g[:len(g)-1] {
		turnNumbers[x] = i + 1
	}

	prev := g[len(g)-1]
	for turn := len(g); turn < targetTurns; turn++ {
		if lastSeen, ok := turnNumbers[prev]; ok {
			turnNumbers[prev] = turn
			prev = turn - lastSeen
		} else {
			turnNumbers[prev] = turn
			prev = 0
		}
	}

	return prev
}
