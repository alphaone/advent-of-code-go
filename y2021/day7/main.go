package day7

import (
	"slices"

	"github.com/alphaone/advent/utils/mathutils"
)

func SumUpDistance(input []int, alignment int) int {
	result := 0
	for _, v := range input {
		result += mathutils.AbsDiffInt(v, alignment)
	}
	return result
}

func FindShortestDistance(input []int) (int, int) {
	minDistance := -1
	minDistanceAlignment := -1
	max := slices.Max(input)

	for alignment := 0; alignment < max; alignment++ {
		distance := SumUpDistance(input, alignment)
		if minDistance == -1 || distance < minDistance {
			minDistance = distance
			minDistanceAlignment = alignment
		}
	}
	return minDistance, minDistanceAlignment
}

func SumUpFuel(input []int, alignment int) int {
	result := 0
	for _, v := range input {
		distance := mathutils.AbsDiffInt(v, alignment)
		result += sumFromAToB(0, distance)
	}
	return result
}

func FindLowestFuel(input []int) (int, int) {
	minFuel := -1
	minFuelAlignment := -1
	max := slices.Max(input)

	for alignment := 0; alignment < max; alignment++ {
		distance := SumUpFuel(input, alignment)
		if minFuel == -1 || distance < minFuel {
			minFuel = distance
			minFuelAlignment = alignment
		}
	}
	return minFuel, minFuelAlignment
}

func sumFromAToB(a, b int) int {
	return (b * (b + 1) / 2) - (a * (a - 1) / 2)
}
