package day3

import (
	"slices"

	"github.com/alphaone/advent/utils/sliceutils"
)

func findWrongItem(rucksack string) rune {
	firstHalf, secondHalf := rucksack[:len(rucksack)/2], rucksack[len(rucksack)/2:]
	res := sliceutils.Intersection([]rune(firstHalf), []rune(secondHalf))
	if len(res) == 0 {
		panic("No wrong item found")
	}

	return res[0]
}

func score(r rune) int {
	if r >= 'A' && r <= 'Z' {
		return int(r-'A') + 27
	}
	return int(r-'a') + 1
}

func solve(input []string) int {
	res := 0
	for _, rucksack := range input {
		res += score(findWrongItem(rucksack))
	}
	return res
}

func findBadge(rucksacks []string) rune {
	res := []rune(rucksacks[0])
	for _, rucksack := range rucksacks[1:] {
		res = sliceutils.Intersection(res, []rune(rucksack))
	}
	if len(res) == 0 {
		panic("No badge found")
	}

	return res[0]
}

func solveB(input []string) int {
	res := 0
	for group := range slices.Chunk(input, 3) {
		res += score(findBadge(group))
	}
	return res
}
