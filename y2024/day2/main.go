package day2

import "github.com/alphaone/advent/utils"

func solvePartA(reports []Report) int {
	res := 0
	for _, r := range reports {
		if r.Safe() {
			res += 1
		}
	}
	return res
}

func solvePartB(reports []Report) int {
	res := 0
	for _, r := range reports {
		if r.SafeDampened() {
			res += 1
		}
	}
	return res
}

func parseInput(input [][]int) []Report {
	var res []Report
	for _, x := range input {
		res = append(res, Report(x))
	}
	return res
}

type Report []int

func (r Report) Safe() bool {
	return r.Increasing() || r.Decreasing()
}

func (r Report) Increasing() bool {
	for i, y := range r {
		if i == 0 {
			continue
		}
		x := r[i-1]
		if y-x < 1 || y-x > 3 {
			return false
		}
	}
	return true
}

func (r Report) Decreasing() bool {
	for i, y := range r {
		if i == 0 {
			continue
		}
		x := r[i-1]
		if x-y < 1 || x-y > 3 {
			return false
		}
	}
	return true
}

func (r Report) Damp() []Report {
	var res []Report
	for i := range r {
		res = append(res, utils.RemoveIndex(r, i))
	}
	return res
}

func (r Report) SafeDampened() bool {
	for _, dampened := range r.Damp() {
		if dampened.Safe() {
			return true
		}
	}
	return false
}
