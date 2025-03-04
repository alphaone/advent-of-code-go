package day10

import (
	"maps"
	"slices"
	"sort"

	"github.com/alphaone/advent/utils/mathutils"
)

func solveA(input []int) int {
	input = append(input, 0)
	sort.Ints(input)

	jolts1Count := 0
	joult3Count := 0

	prev := input[0]
	for i := 1; i < len(input); i++ {
		diff := input[i] - prev
		prev = input[i]

		switch diff {
		case 1:
			jolts1Count++
		case 3:
			joult3Count++
		case 2:
		// noop
		default:
			panic("unexpected diff")
		}
	}
	// add built-in adapter
	joult3Count++
	return jolts1Count * joult3Count
}

// build a graph of how a adapter can be connected to which set of other adapters
// 0: {1}, 1: {4}, 4: {5, 6, 7}, 5: {6, 7}, 6: {7}, 7: {10}, 10: {11, 12}, 11: {12}, 12: {15}, 15: {16}, 16: {19}, 19: {},
func buildGraph(input []int) map[int][]int {
	input = append(input, 0)
	sort.Ints(input)

	graph := make(map[int][]int)
	for _, v := range input {
		graph[v] = make([]int, 0)
	}

	for i, v := range input {
		for j := i + 1; j < len(input) && input[j]-v <= 3; j++ {
			graph[v] = append(graph[v], input[j])
		}
	}

	return graph
}

// count the number of possible arrangements of the adapters
//
// give 19 a value 1 initial arrangement
// start from the end,
// look up 16 in the graph, it points to 19, which has 1 arrangement, 16 has now also 1 arrangement
// look up 15 in the graph, it points to 16, which has 1 arrangement, 15 has now also 1 arrangement
// ...
// look up 10 in the graph, it points to 11 and 12, which both have 1 arrangement, 10 has now 2 arrangements
// look up 7 in the graph, it points to 10, which has 2 arrangements, 7 has now also 2 arrangements
// ...
// look up 5 in the graph, it points to 6 and 7, which both have 2 arrangement, 5 has now 4 arrangements
// look up 4 in the graph, it points to 5, 6 and 7, which have 4, 2 and 2 arrangement, 4 has now 8 arrangements
// look up 1 in the graph, it points to 4, which has 8 arrangements, 1 has now also 8 arrangements
// look up 0 in the graph, it points to 1, which has 8 arrangements, 0 has now also 8 arrangements
func countArrangements(graph map[int][]int) int {
	keys := slices.Collect(maps.Keys(graph))
	min, max := mathutils.MinMax(keys)

	arrangements := make(map[int]int)
	arrangements[max] = 1

	sort.Ints(keys)
	slices.Reverse(keys)
	for _, k := range keys[1:] {
		for _, option := range graph[k] {
			arrangements[k] += arrangements[option]
		}
	}
	return arrangements[min]
}

func solveB(input []int) int {
	graph := buildGraph(input)
	return countArrangements(graph)
}
