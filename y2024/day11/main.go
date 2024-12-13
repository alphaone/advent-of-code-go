package day11

import (
	"math"
	"strconv"
	"strings"
)

func parseInput(input string) []int {
	parts := strings.Split(strings.TrimSpace(input), " ")

	res := make([]int, len(parts))
	for i, x := range parts {
		d, err := strconv.Atoi(x)
		if err != nil {
			panic(err)
		}
		res[i] = d
	}
	return res
}

func splitNumber(n int) (int, *int) {
	l := int(math.Ceil(math.Log10(float64(n + 1))))
	if l%2 == 0 {
		f := int(math.Pow10(l / 2))
		p1 := n / f
		p2 := n % f
		return p1, &p2
	}
	return n, nil
}

func solve(n int, stones []int) int {
	sum := 0
	for _, cur := range stones {
		sum += blink(cur, n)
	}
	return sum
}

type cachekey = struct{ n, blinksLeft int }

var mem = map[cachekey]int{}

func blink(cur int, blinksLeft int) int {
	if blinksLeft == 0 {
		return 1
	}

	cached, ok := mem[cachekey{cur, blinksLeft}]
	if ok {
		return cached
	}

	res := blinkRecursive(cur, blinksLeft)
	mem[cachekey{cur, blinksLeft}] = res

	return res
}

func blinkRecursive(cur int, blinksLeft int) int {
	blinksLeft--
	if cur == 0 {
		return blink(1, blinksLeft)
	}
	a, b := splitNumber(cur)
	if b != nil {
		return blink(a, blinksLeft) + blink(*b, blinksLeft)
	}
	return blink(cur*2024, blinksLeft)
}
