package dec11

import (
	"fmt"
	"math"
)

type Coords struct {
	x int
	y int
}

func hundredsDigit(number int) int {
	return int(math.Mod(math.Floor(float64(number)/100), 10))
}

func powerLevel(serialNumber int, cell Coords) int {
	rackId := cell.x + 10
	return hundredsDigit((rackId*cell.y+serialNumber)*rackId) - 5
}

func squarePowerLevel(serialNumber int, size int, cell Coords) int {
	res := 0

	for x := cell.x; x < cell.x+size; x++ {
		for y := cell.y; y < cell.y+size; y++ {
			res += powerLevel(serialNumber, Coords{x, y})
		}
	}

	return res
}

func bestSquareOfSize(serialNumber, size int) (Coords, int) {
	maxPower := -1
	maxCoords := Coords{-1, -1}
	for x := 0; x < 302-size; x++ {
		for y := 0; y < 302-size; y++ {
			currentCoords := Coords{x, y}
			currentPower := squarePowerLevel(serialNumber, size, currentCoords)
			if currentPower > maxPower {
				maxCoords = currentCoords
				maxPower = currentPower
			}
		}
	}

	return maxCoords, maxPower
}

func bestSquare(serialNumber int) (Coords, int) {
	maxPower := -1
	maxCoords := Coords{-1, -1}
	maxSize := -1
	for size := 1; size < 20; size++ {
		currentCoords, currentPower := bestSquareOfSize(serialNumber, size)
		if currentPower > maxPower {
			maxPower = currentPower
			maxCoords = currentCoords
			maxSize = size
		}
		fmt.Printf("Size: %v -> %v (%v)\n", size, currentPower, currentCoords)
	}
	return maxCoords, maxSize
}
