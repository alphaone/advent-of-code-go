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
		match := reOuter.FindAllStringSubmatch(line, -1)
		outer := match[0][1]
		content := match[0][2]

		match = reContent.FindAllStringSubmatch(content, -1)
		res[outer] = make(map[string]int)
		for _, m := range match {
			n, _ := strconv.Atoi(m[1])
			res[outer][m[2]] = n
		}
	}

	return res
}

func solveA(input tree, target string) []string {
	candidates := make(map[string]bool)
	for k, v := range input {
		if _, ok := v[target]; ok {
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

func solveB(input tree, target string) int {
	res := 0
	for k, v := range input[target] {
		res += v * (1 + solveB(input, k))
	}

	return res
}
