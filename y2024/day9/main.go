package day9

import (
	"slices"

	"github.com/alphaone/advent/utils"
)

func layoutFromDense(input string) []int {
	res := make([]int, 0)

	for i, c := range input {
		x := int(c) - 48
		if i%2 == 0 {
			// file
			id := i / 2

			for range x {
				res = append(res, id)
			}

		} else {
			// freespace
			for range x {
				res = append(res, -1)
			}
		}
	}

	return res
}

type block struct {
	id     int
	length int
}

func blockLayoutFromDense(input string) []block {
	res := make([]block, 0)

	for i, c := range input {
		x := int(c) - 48
		if i%2 == 0 {
			// file
			id := i / 2

			res = append(res, block{id, x})

		} else {
			// freespace
			res = append(res, block{-1, x})
		}
	}

	return res
}

func cleanupStepA(layout []int) bool {
	idx1 := slices.Index(layout, -1)
	idx2 := utils.LastIndexFunc(layout, func(e int) bool { return e != -1 })
	if idx1 == -1 || idx2 == -1 || idx2 < idx1 {
		return false
	}

	layout[idx1] = layout[idx2]
	layout[idx2] = -1

	return true
}

func cleanupA(layout []int) {
	for {
		ok := cleanupStepA(layout)
		if !ok {
			break
		}
	}
}

func cleanupStepB(layout []block, maxIdx int) []block {
	idx2 := utils.LastIndexFunc(layout[:maxIdx], func(e block) bool { return e.id != -1 })
	if idx2 == -1 {
		return layout
	}
	b := layout[idx2]
	idx1 := slices.IndexFunc(layout, func(e block) bool { return e.id == -1 && e.length >= b.length })
	if idx1 == -1 || idx2 < idx1 {
		return layout
	}

	diff := layout[idx1].length - b.length

	layout[idx1] = layout[idx2]
	layout[idx2] = block{-1, b.length}
	if diff > 0 {
		layout = utils.Insert(layout, idx1+1, block{-1, diff})
	}

	return layout
}

func cleanupB(layout []block) []block {
	maxIdx := len(layout)
	for {
		layout = cleanupStepB(layout, maxIdx)
		maxIdx -= 1
		if maxIdx <= 0 {
			break
		}
	}
	return layout
}

func checksumA(layout []int) int {
	sum := 0
	for i, x := range layout {
		if x == -1 {
			continue
		}
		sum += i * x
	}
	return sum
}

func checksumB(layout []block) int {
	sum := 0
	i := 0
	for _, x := range layout {
		if x.id == -1 {
			i += x.length
			continue
		}

		for range x.length {
			sum += i * x.id
			i += 1
		}
	}
	return sum
}

func solvePartA(input string) int {
	l := layoutFromDense(input)
	cleanupA(l)
	return checksumA(l)
}

func solvePartB(input string) int {
	l := blockLayoutFromDense(input)
	l = cleanupB(l)
	return checksumB(l)
}
