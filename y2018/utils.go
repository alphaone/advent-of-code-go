package y2018

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func ReadFileAsIntSlice(filename string) []int {
	dat, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(dat), "\n")

	var res []int
	for _, v := range lines {
		i, _ := strconv.Atoi(v)
		res = append(res, i)
	}
	return res
}
