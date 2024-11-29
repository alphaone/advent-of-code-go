package utils

import (
	"regexp"
	"strconv"
)

var reNumber = regexp.MustCompile(`\d+`)

func ParseNumbers(s string) []int {
	var numbers []int
	for _, n := range reNumber.FindAllString(s, -1) {
		i, _ := strconv.Atoi(n)
		numbers = append(numbers, i)
	}
	return numbers
}
