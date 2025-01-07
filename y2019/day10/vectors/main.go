package vectors

import (
	"math"

	"github.com/alphaone/advent/utils"
)

type Vector struct {
	X, Y int
}

func New(x, y int) Vector {
	return Vector{
		X: x,
		Y: y,
	}
}

func (v Vector) ProportionalTo(b Vector) bool {
	return v.X*b.Y == v.Y*b.X
}

func (v Vector) Reduce() Vector {
	gcd := utils.GCD(utils.AbsInt(v.X), utils.AbsInt(v.Y))
	return Vector{
		X: v.X / gcd,
		Y: v.Y / gcd,
	}
}

func (v Vector) Angle(b Vector) float64 {
	return math.Mod(math.Atan2(float64(b.Y), float64(b.X))-math.Atan2(float64(v.Y), float64(v.X))+2*math.Pi, 2*math.Pi)
}

func SortClockwiseFunc(a, b Vector) int {
	if a.Angle(New(0, 1)) < b.Angle(New(0, 1)) {
		return -1
	} else if a.Angle(New(1, 0)) > b.Angle(New(1, 0)) {
		return 1
	} else {
		return 0
	}
}
