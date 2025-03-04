package day10

import "sort"

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
