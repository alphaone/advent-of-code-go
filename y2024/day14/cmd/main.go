package main

import (
	"github.com/alphaone/advent/utils"
	"github.com/alphaone/advent/y2024/day14"
)

func main() {
	day14.SolvePartB(day14.Parse(utils.LoadStrings("input.txt")), day14.Coord{X: 101, Y: 103})
}
