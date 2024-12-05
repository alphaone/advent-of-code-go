package day5

import (
	"slices"
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils"
)

func solvePartA(pageLists [][]int, rules PageOrder) int {
	res := 0
	for _, pl := range pageLists {
		if inRightOrder(pl, rules) {
			res += middle(pl)
		}
	}

	return res
}

func solvePartB(pageLists [][]int, rules PageOrder) int {
	res := 0
	for _, pl := range pageLists {
		if inRightOrder(pl, rules) {
			continue
		}

		sorted := sortByRules(pl, rules)

		res += middle(sorted)
	}

	return res
}

type PageOrder map[int][]int

func parseInput(input string) (PageOrder, [][]int) {
	parts := strings.Split(input, "\n\n")

	return parsePageOrder(parts[0]), parsePages(parts[1])
}

func parsePages(input string) [][]int {
	lines := strings.Split(input, "\n")

	res := make([][]int, 0)
	for _, l := range lines {
		if strings.TrimSpace(l) == "" {
			continue
		}
		res = append(res, utils.ParseNumbers(l))
	}
	return res
}

func parsePageOrder(input string) PageOrder {
	lines := strings.Split(input, "\n")

	res := make(PageOrder)
	for _, l := range lines {
		fields := strings.Split(l, "|")
		a, _ := strconv.Atoi(fields[0])
		b, _ := strconv.Atoi(fields[1])

		_, ok := res[a]
		if !ok {
			res[a] = make([]int, 0)
		}
		res[a] = append(res[a], b)
	}

	return res
}

func inRightOrder(pageList []int, rules PageOrder) bool {
	return slices.Equal(pageList, sortByRules(pageList, rules))
}

func sortByRules(pageList []int, rules PageOrder) []int {
	copiedList := make([]int, len(pageList))
	for i, x := range pageList {
		copiedList[i] = x
	}

	slices.SortFunc(copiedList, func(a, b int) int {
		if slices.Contains(rules[a], b) {
			return -1
		}
		return 0
	})
	return copiedList
}

func middle(pageList []int) int {
	return pageList[len(pageList)/2]
}
