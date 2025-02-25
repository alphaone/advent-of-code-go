package day7

import (
	"maps"
	"regexp"
	"slices"
	"strconv"
)

type tree map[string]map[string]int

func parse(input []string) tree {
	res := make(tree)

	reOuter := regexp.MustCompile(`(\w+ \w+) bags contain (.*)`)
	reContent := regexp.MustCompile(`(\d+) (\w+ \w+) bags?`)
	for _, line := range input {
		m := reOuter.FindAllStringSubmatch(line, -1)
		outer := m[0][1]
		content := m[0][2]
		m = reContent.FindAllStringSubmatch(content, -1)
		res[outer] = make(map[string]int)
		for _, x := range m {
			n, _ := strconv.Atoi(x[1])
			res[outer][x[2]] = n
		}
	}

	return res
}

func solveA(input tree, bag string) []string {
	candidates := make(map[string]bool)
	for k, v := range input {
		if _, ok := v[bag]; ok {
			candidates[k] = true
		}
	}
	for k := range candidates {
		for _, c := range solveA(input, k) {
			candidates[c] = true
		}
	}

	return slices.Collect(maps.Keys(candidates))
}

func solveB(input tree, bag string) int {
	res := 0
	for k, v := range input[bag] {
		res += v * (1 + solveB(input, k))
	}

	return res
}
