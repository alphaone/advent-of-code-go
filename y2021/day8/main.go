package day8

import (
	"bytes"
	"slices"
	"strings"

	"github.com/alphaone/advent/utils/mathutils"
)

type Line struct {
	numbers []string
	output  []string
}

func (l Line) findWiringOfLength(length int) []byte {
	for _, n := range l.numbers {
		if len(n) == length {
			return []byte(n)
		}
	}
	panic("not found")
}

func (l Line) diff(a []byte, b []byte) int {
	count := 0
	for _, wa := range a {
		if bytes.IndexByte(b, wa) == -1 {
			count++
		}
	}
	for _, wb := range b {
		if bytes.IndexByte(a, wb) == -1 {
			count++
		}
	}

	return count
}

func (l Line) toDigit(wiring []byte) int {
	switch len(wiring) {
	case 2:
		return 1
	case 3:
		return 7
	case 4:
		return 4
	case 7:
		return 8
	case 5:
		// 5,2,3
		if l.diff(l.findWiringOfLength(3), wiring) == 2 {
			return 3
		}
		if l.diff(l.findWiringOfLength(4), wiring) == 3 {
			return 5
		}
		return 2
	case 6:
		// 6,9,0
		if l.diff(l.findWiringOfLength(4), wiring) == 2 {
			return 9
		}
		if l.diff(l.findWiringOfLength(3), wiring) == 3 {
			return 0
		}
		return 6
	}

	panic("not found")
}

func (l Line) Output() int {
	res := 0
	slices.Reverse(l.output)
	for i, n := range l.output {
		d := l.toDigit([]byte(n))
		res += d * mathutils.PowInts(10, i)
	}
	return res
}

func parse(line []string) []Line {
	res := []Line{}
	for _, l := range line {
		parts := strings.Split(l, " | ")

		numbers := strings.Split(parts[0], " ")
		output := strings.Split(parts[1], " ")

		numbers = append(numbers, output...)

		res = append(res, Line{numbers, output})
	}
	return res
}

func solveA(input []Line) int {
	count := 0

	for _, l := range input {
		for _, n := range l.output {
			// 1,4,7 or 8
			if len(n) == 2 || len(n) == 3 || len(n) == 4 || len(n) == 7 {
				count++
			}
		}
	}
	return count
}

func solveB(input []Line) int {
	res := 0
	for _, l := range input {
		res += l.Output()
	}
	return res
}
