package day11

import (
	"log"
	"math"
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils"
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

func blink(stones *[]int) {
	for i := 0; i < len(*stones); i++ {
		s1, s2 := applyRule((*stones)[i])
		(*stones)[i] = s1
		if s2 != nil {
			*stones = utils.Insert(*stones, i+1, *s2)
			i++
		}
	}
}

func applyRule(n int) (int, *int) {
	if n == 0 {
		return 1, nil
	}

	a, b := splitNumber(n)
	if b != nil {
		return a, b
	}

	return n * 2024, nil
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
	for i := range n {
		log.Print(i, len(stones))
		blink(&stones)
	}
	return len(stones)
}
