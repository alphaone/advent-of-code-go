package day1

import (
	"slices"
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils/mathutils"
	"github.com/alphaone/advent/utils/sliceutils"
)

type Tuple struct {
	A, B int
}

func solvePartA(ts []Tuple) int {
	as := make([]int, len(ts))
	bs := make([]int, len(ts))

	for i, t := range ts {
		as[i] = t.A
		bs[i] = t.B
	}

	slices.SortFunc(as, func(a, b int) int { return a - b })
	slices.SortFunc(bs, func(a, b int) int { return a - b })

	sum := 0
	for i, a := range as {
		b := bs[i]
		diff := mathutils.AbsInt(a - b)

		sum += diff
	}

	return sum
}

func solvePartB(ts []Tuple) int {
	as := make([]int, len(ts))
	bs := make([]int, len(ts))

	for i, t := range ts {
		as[i] = t.A
		bs[i] = t.B
	}

	freqs := sliceutils.Frequencies(bs)

	sum := 0
	for _, a := range as {
		f := freqs[a]
		similarity := a * f
		sum += similarity
	}

	return sum
}

func parseInput(lines []string) []Tuple {
	res := make([]Tuple, 0)
	for _, l := range lines {
		x := strings.SplitN(l, "   ", 2)
		if len(x) != 2 {
			panic("wrong line format")
		}
		a, err := strconv.Atoi(x[0])
		if err != nil {
			panic(err)
		}
		b, err := strconv.Atoi(x[1])
		if err != nil {
			panic(err)
		}
		res = append(res, Tuple{a, b})
	}

	return res
}
