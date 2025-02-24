package day5

import (
	"sort"
	"strconv"
	"strings"
)

func solveA(input []string) int {
	max := 0
	for _, s := range input {
		s := parseSeat(s)
		if s > max {
			max = s
		}
	}
	return max
}

func solveB(input []string) int {
	allSeats := make([]int, 0)
	for _, s := range input {
		allSeats = append(allSeats, parseSeat(s))
	}
	sort.Ints(allSeats)

	for i, s := range allSeats {
		if allSeats[i+1] != s+1 {
			return s + 1
		}
	}

	panic("no seat found")
}

func parseSeat(s string) int {
	s = strings.ReplaceAll(s, "F", "0")
	s = strings.ReplaceAll(s, "B", "1")
	s = strings.ReplaceAll(s, "L", "0")
	s = strings.ReplaceAll(s, "R", "1")

	id, err := strconv.ParseInt(s, 2, 64)
	if err != nil {
		panic(err)
	}

	return int(id)
}
