package day2

import (
	"fmt"

	"github.com/alphaone/advent/utils/sliceutils"
)

type Line struct {
	from     int
	to       int
	char     rune
	password string
}

func (l Line) IsValidA() bool {
	freqs := sliceutils.Frequencies([]rune(l.password))

	return l.from <= freqs[l.char] && freqs[l.char] <= l.to
}

func (l Line) IsValidB() bool {
	return l.password[l.from-1] == byte(l.char) != (l.password[l.to-1] == byte(l.char))
}

func Parse(line string) Line {
	var from, to int
	var char rune
	var password string
	_, _ = fmt.Sscanf(line, "%d-%d %c: %s", &from, &to, &char, &password)
	return Line{from, to, char, password}
}

func solveA(input []string) int {
	res := 0

	for _, line := range input {
		parsed := Parse(line)
		if parsed.IsValidA() {
			res++
		}
	}

	return res
}

func solveB(input []string) int {
	res := 0

	for _, line := range input {
		parsed := Parse(line)
		if parsed.IsValidB() {
			res++
		}
	}

	return res
}
