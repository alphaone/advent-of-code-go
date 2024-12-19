package day19

import "strings"

type onsen struct {
	towels   []string
	patterns []string
}

func parse(input string) onsen {
	parts := strings.Split(strings.TrimSpace(input), "\n\n")

	towels := strings.Split(parts[0], ", ")
	patterns := strings.Split(parts[1], "\n")
	return onsen{towels, patterns}
}

func countPaths(pattern string, towels []string, cache map[string]int) int {
	if val, found := cache[pattern]; found {
		return val
	}

	res := 0
	for _, t := range towels {
		if t == pattern {
			res += 1
		}
		if strings.HasPrefix(pattern, t) {
			res += countPaths(pattern[len(t):], towels, cache)
		}
	}
	cache[pattern] = res
	return res
}

func solvePartA(onsen onsen) int {
	cache := make(map[string]int)
	count := 0
	for _, p := range onsen.patterns {
		if countPaths(p, onsen.towels, cache) > 0 {
			count++
		}
	}

	return count
}

func solvePartB(onsen onsen) int {
	cache := make(map[string]int)
	count := 0
	for _, p := range onsen.patterns {
		count += countPaths(p, onsen.towels, cache)
	}

	return count
}
