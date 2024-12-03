package day3

import (
	"regexp"
	"strconv"
	"strings"
)

type Multiplication struct {
	a, b int
}

func (m Multiplication) Product() int {
	return m.a * m.b
}

func parseMultiplicationsA(line string) []Multiplication {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(line, -1)
	res := make([]Multiplication, 0)
	for _, m := range matches {
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		res = append(res, Multiplication{a, b})
	}
	return res
}

func solve(input string, parseFn func(string) []Multiplication) int {
	mults := parseFn(input)
	results := 0
	for _, r := range mults {
		results += r.Product()
	}
	return results
}

func parseMultiplicationsB(line string) []Multiplication {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)`)
	matches := r.FindAllStringSubmatch(line, -1)
	indexes := r.FindAllStringSubmatchIndex(line, -1)

	res := make([]Multiplication, 0)
	for i, m := range matches {
		start := indexes[i][0]
		if strings.LastIndex(line[:start], "don't()") > strings.LastIndex(line[:start], "do()") {
			continue
		}
		a, _ := strconv.Atoi(m[1])
		b, _ := strconv.Atoi(m[2])
		res = append(res, Multiplication{a, b})
	}
	return res
}

func parseMultiplicationsB2(line string) []Multiplication {
	r := regexp.MustCompile(`mul\((\d{1,3}),(\d{1,3})\)|don't\(\)|do\(\)`)
	matches := r.FindAllStringSubmatch(line, -1)
	res := make([]Multiplication, 0)

	multEnabled := true
	for _, m := range matches {
		if strings.Contains(m[0], "don't()") {
			multEnabled = false
		} else if strings.Contains(m[0], "do()") {
			multEnabled = true
		} else if multEnabled {
			a, _ := strconv.Atoi(m[1])
			b, _ := strconv.Atoi(m[2])
			res = append(res, Multiplication{a, b})
		}
	}
	return res
}
