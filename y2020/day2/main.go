package day2

import (
	"fmt"

	"github.com/alphaone/advent/utils/sliceutils"
)

type LineA struct {
	from     int
	to       int
	char     rune
	password string
}

func parseA(line string) validatable {
	var from, to int
	var char rune
	var password string
	_, _ = fmt.Sscanf(line, "%d-%d %c: %s", &from, &to, &char, &password)
	return LineA{from, to, char, password}
}

type LineB struct {
	first    int
	second   int
	char     rune
	password string
}

func parseB(line string) validatable {
	var first, second int
	var char rune
	var password string
	_, _ = fmt.Sscanf(line, "%d-%d %c: %s", &first, &second, &char, &password)
	return LineB{first, second, char, password}
}

func (l LineA) IsValid() bool {
	freqs := sliceutils.Frequencies([]rune(l.password))

	return l.from <= freqs[l.char] && freqs[l.char] <= l.to
}

func (l LineB) IsValid() bool {
	return l.password[l.first-1] == byte(l.char) != (l.password[l.second-1] == byte(l.char))
}

type validatable interface {
	IsValid() bool
}

func solve(parse func(string) validatable, input []string) int {
	res := 0

	for _, line := range input {
		parsed := parse(line)
		if parsed.IsValid() {
			res++
		}
	}

	return res
}
