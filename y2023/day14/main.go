package day14

import (
	"fmt"
	"sort"
	"strings"

	"github.com/alphaone/advent/utils"
)

func MoveLineWest(s []rune) []rune {
	// TODO: build utils function for slice split
	r := strings.Split(string(s), "#")
	if len(r) == 1 {
		return moveStones(s)
	}

	res := [][]rune{}
	for _, v := range r {
		res = append(res, moveStones([]rune(v)))
	}
	return utils.Join(res, '#')
}

func moveStones(s []rune) []rune {
	sort.Slice(s, func(i, j int) bool { return s[i] > s[j] })
	return s
}

func MoveWest(input [][]rune) [][]rune {
	result := [][]rune{}
	for _, line := range input {
		result = append(result, MoveLineWest(line))
	}
	return result
}

func solvePartA(input [][]rune) int {
	input = utils.RotateCounterClockwise(input)
	moved := MoveWest(input)

	return Weight(utils.RotateClockwise(moved))
}

func Weight(input [][]rune) int {
	bottom := len(input)

	res := 0
	for i, line := range input {
		x := strings.Count(string(line), "O")
		res += x * (bottom - i)
	}
	return res
}

func cycleThrough(input [][]rune, count int) [][]rune {
	cache := make(map[string]int)
	rest := 0

	for i := 0; i < count; i++ {
		input = oneCycle(input)

		if prevIteration, ok := cache[strings.Join(utils.AsStrings(input), "")]; ok {
			period := i - prevIteration
			rest = (count - i - 1) % period
			fmt.Printf("Found cycle at %d (prev: %d); Period: %d, rest cycles: %d\n", i, prevIteration, period, rest)
			break
		}
		cacheKey := strings.Join(utils.AsStrings(input), "")
		cache[cacheKey] = i
	}

	for i := 0; i < rest; i++ {
		input = oneCycle(input)
	}

	return input
}

func oneCycle(input [][]rune) [][]rune {
	input = MoveWest(input)
	input = MoveWest(utils.RotateClockwise(input))
	input = MoveWest(utils.RotateClockwise(input))
	input = MoveWest(utils.RotateClockwise(input))
	return utils.RotateClockwise(input)
}

func solvePartB(input [][]rune, count int) int {
	input = utils.RotateCounterClockwise(input)
	cycled := cycleThrough(input, count)
	cycled = utils.RotateClockwise(cycled)
	return Weight(cycled)
}
