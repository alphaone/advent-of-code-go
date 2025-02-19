package day10

import (
	"fmt"
	"sort"
)

func solveA(input []string) int {
	result := 0
	for _, s := range input {
		c, _, err := FirstIllegalChar(s)
		if err == nil {
			result += score(c)
		}
	}
	return result
}

func solveB(input []string) int {
	scores := make([]int, 0)

	for _, s := range input {
		_, stack, err := FirstIllegalChar(s)
		if err != nil {
			scores = append(scores, autoCompleteScore(string(stack)))
		}
	}

	sort.Ints(scores)

	return scores[len(scores)/2]
}

func FirstIllegalChar(s string) (rune, []rune, error) {
	return firstIllegalChar(s, 0, []rune{})
}

func firstIllegalChar(s string, idx int, stack []rune) (rune, []rune, error) {
	if len(s) == 0 {
		return 0, stack, fmt.Errorf("incomplete")
	}
	current := rune(s[0])

	if isOpening(current) {
		return firstIllegalChar(s[1:], idx+1, append([]rune{complement(current)}, stack...))
	}

	if isClosing(current) {
		if len(stack) == 0 {
			return 0, stack, fmt.Errorf("stack empty")
		}
		nextToClose, stack := stack[0], stack[1:]
		if current == nextToClose {
			return firstIllegalChar(s[1:], idx+1, stack)
		} else {
			return current, stack, nil
		}
	}

	panic("Invalid character")
}

func isOpening(r rune) bool {
	return r == '(' || r == '[' || r == '{' || r == '<'
}

func isClosing(r rune) bool {
	return r == ')' || r == ']' || r == '}' || r == '>'
}

func complement(opening rune) rune {
	switch opening {
	case '(':
		return ')'
	case '[':
		return ']'
	case '{':
		return '}'
	case '<':
		return '>'
	}
	panic("Invalid opening character")
}

func score(r rune) int {
	switch r {
	case ')':
		return 3
	case ']':
		return 57
	case '}':
		return 1197
	case '>':
		return 25137
	}
	panic("Invalid closing character")
}

func autoCompleteScore(s string) int {
	res := 0

	for _, c := range s {
		res *= 5
		switch c {
		case ')':
			res += 1
		case ']':
			res += 2
		case '}':
			res += 3
		case '>':
			res += 4
		}
	}

	return res
}
