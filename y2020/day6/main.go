package day6

import (
	"strings"
)

type (
	group []form
	form  []rune
)

func parse(input string) []group {
	groups := strings.Split(strings.TrimSpace(input), "\n\n")
	res := make([]group, 0, len(groups))

	for _, group := range groups {
		people := strings.Split(group, "\n")
		forms := make([]form, 0, len(people))

		for _, person := range people {
			f := form([]rune(person))
			forms = append(forms, f)
		}
		res = append(res, forms)
	}

	return res
}

func solveA(groups []group) int {
	sum := 0
	for _, group := range groups {
		m := make(map[rune]bool)
		for _, form := range group {
			for _, q := range form {
				m[q] = true
			}
		}
		sum += len(m)
	}
	return sum
}

func solveB(groups []group) int {
	sum := 0
	for _, group := range groups {
		m := make(map[rune]int)
		for _, form := range group {
			for _, q := range form {
				m[q] += 1
			}
		}
		for _, v := range m {
			if v == len(group) {
				sum += 1
			}
		}
	}
	return sum
}
