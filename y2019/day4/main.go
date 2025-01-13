package day4

import (
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils/sliceutils"
)

func parseInput(input string) (int, int) {
	parts := strings.Split(strings.TrimSpace(input), "-")
	a, _ := strconv.Atoi(parts[0])
	b, _ := strconv.Atoi(parts[1])
	return a, b
}

func solvePartB(from, to int) int {
	res := 0
	for n := from; n <= to; n++ {
		if valid(n, func(fv int) bool { return fv == 2 }) {
			res += 1
		}
	}
	return res
}

func solvePartA(from, to int) int {
	res := 0
	for n := from; n <= to; n++ {
		if valid(n, func(fv int) bool { return fv >= 2 }) {
			res += 1
		}
	}
	return res
}

func valid(input int, freqPred func(int) bool) bool {
	runeArray := []rune(strconv.Itoa(input))
	if len(runeArray) != 6 {
		return false
	}

	decreasing := false
	for i, d := range runeArray {
		if i > 0 && runeArray[i-1] > d {
			decreasing = true
		}
	}

	freqs := sliceutils.Frequencies(runeArray)
	foundDouble := false
	for _, v := range freqs {
		if freqPred(v) {
			foundDouble = true
		}
	}

	return foundDouble && !decreasing
}
