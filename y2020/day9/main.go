package day9

import "github.com/alphaone/advent/utils/mathutils"

func firstInvalid(input []int, preambleLength int) int {
	for i := preambleLength; i < len(input); i++ {
		if !isValid(input[i], input[i-preambleLength:i]) {
			return input[i]
		}
	}

	return -1
}

func isValid(target int, preamble []int) bool {
	for i, a := range preamble {
		for _, b := range preamble[i:] {
			if a+b == target {
				return true
			}
		}
	}

	return false
}

func contiguousSet(input []int, target int) []int {
	for i := 0; i < len(input); i++ {
		for j := i + 1; j < len(input); j++ {
			sum := mathutils.Sum(input[i:j])
			if sum == target {
				return input[i:j]
			}
		}
	}

	return []int{-1, -1}
}

func encryptionWeakness(input []int, target int) int {
	min, max := mathutils.MinMax(contiguousSet(input, target))
	return min + max
}
