package day6

import (
	"github.com/alphaone/advent/utils/sliceutils"
)

func startOfSeq(s string, length int) int {
	for i := 0; i < len(s)-length; i++ {
		if sliceutils.IsUnique([]rune(s[i : i+length])) {
			return i + length
		}
	}
	return -1
}
