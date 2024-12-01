package day1

import (
	"slices"
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils"
)

type CaloriesByElf = [][]int

func parseInput(s string) CaloriesByElf {
	s = strings.TrimSpace(s)
	elfs := strings.Split(s, "\n\n")

	res := make(CaloriesByElf, 0)
	for _, x := range elfs {
		caloriesStrs := strings.Split(x, "\n")
		calories := make([]int, 0)
		for _, y := range caloriesStrs {
			c, err := strconv.Atoi(y)
			if err != nil {
				panic(err)
			}
			calories = append(calories, c)
		}
		res = append(res, calories)
	}

	return res
}

func sumByElf(elfs CaloriesByElf) []int {
	res := make([]int, 0)
	for _, elf := range elfs {
		res = append(res, utils.Sum(elf))
	}
	return res
}

func solvePartA(input string) int {
	return slices.Max(sumByElf(parseInput(input)))
}

func solvePartB(input string) int {
	elfs := sumByElf(parseInput(input))
	slices.SortFunc(elfs, func(a, b int) int { return b - a })

	return utils.Sum(elfs[0:3])
}
