package dec1

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func SumUp(freqs []int) int {
	res := 0

	for _, v := range freqs {
		res += v
	}

	return res
}

func SumUpB(freqs []int) int {
	res := 0
	seen := map[int]bool{0: true}

	for {
		for _, v := range freqs {
			res += v
			if _, ok := seen[res]; !ok {
				seen[res] = true
			} else {
				return res
			}
		}
	}
}

func readInput() []int {
	dat, _ := ioutil.ReadFile("input.txt")
	lines := strings.Split(string(dat), "\n")

	var res []int
	for _, v := range lines {
		i, _ := strconv.Atoi(v)
		res = append(res, i)
	}
	return res
}
