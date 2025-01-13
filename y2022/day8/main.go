package day8

import (
	"github.com/alphaone/advent/utils"
)

type (
	treepatch [][]int
	coord     struct {
		l, c int
	}
)

func (c coord) add(other coord) coord {
	return coord{c.l + other.l, c.c + other.c}
}

func (tp treepatch) treesInDir(pos, dir coord) []int {
	res := []int{}
	for pos = pos.add(dir); pos.l >= 0 && pos.c >= 0 && pos.l < len(tp) && pos.c < len(tp[0]); pos = pos.add(dir) {
		res = append(res, tp[pos.l][pos.c])
	}
	return res
}

func (tp treepatch) visible(pos coord) bool {
	dirs := []coord{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	cur := tp[pos.l][pos.c]

	for _, dir := range dirs {
		if utils.All(tp.treesInDir(pos, dir), func(t int) bool { return t < cur }) {
			return true
		}
	}
	return false
}

func (tp treepatch) scenicScore(pos coord) int {
	dirs := []coord{{-1, 0}, {0, 1}, {1, 0}, {0, -1}}
	cur := tp[pos.l][pos.c]

	score := 1
	for _, dir := range dirs {
		dst := 0
		for _, t := range tp.treesInDir(pos, dir) {
			dst++
			if t >= cur {
				break
			}
		}
		score *= dst
	}
	return score
}

func parse(input []string) treepatch {
	res := make(treepatch, len(input))
	for i, line := range input {
		trees := make([]int, len(line))
		for j, c := range line {
			trees[j] = int(c) - '0'
		}
		res[i] = trees
	}
	return res
}

func solve(input []string) int {
	tp := parse(input)
	res := 0
	for i := 0; i < len(tp); i++ {
		for j := 0; j < len(tp[0]); j++ {
			if tp.visible(coord{i, j}) {
				res++
			}
		}
	}
	return res
}

func solveB(input []string) int {
	tp := parse(input)
	best := -1
	for i := 0; i < len(tp); i++ {
		for j := 0; j < len(tp[0]); j++ {
			score := tp.scenicScore(coord{i, j})
			if score > best {
				best = score
			}
		}
	}
	return best
}
