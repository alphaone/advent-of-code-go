package day5

import (
	"regexp"
	"strconv"

	"github.com/alphaone/advent/utils"
)

type Pos struct {
	X, Y int
}

type Segment struct {
	Start, End Pos
}

// 0,9 -> 5,9
func ParseLine(line string) Segment {
	re := regexp.MustCompile(`(\d+),(\d+) -> (\d+),(\d+)`)
	res := re.FindStringSubmatch(line)
	if len(res) == 0 {
		panic("Invalid line")
	}
	x1, _ := strconv.Atoi(res[1])
	y1, _ := strconv.Atoi(res[2])
	x2, _ := strconv.Atoi(res[3])
	y2, _ := strconv.Atoi(res[4])
	return Segment{
		Start: Pos{x1, y1},
		End:   Pos{x2, y2},
	}
}

func (s Segment) Line() bool {
	return s.Start.X == s.End.X || s.Start.Y == s.End.Y
}

type Grid map[Pos]int

func (g Grid) Fill(s Segment) {
	if s.Start.X == s.End.X {
		// fmt.Printf("Vertical: %v\n", s)
		for _, y := range utils.MakeSequence(s.Start.Y, utils.AbsDiffInt(s.Start.Y, s.End.Y)+1, s.Start.Y > s.End.Y) {
			g[Pos{s.Start.X, y}]++
		}
	} else if s.Start.Y == s.End.Y {
		// fmt.Printf("Horizontal: %v\n", s)
		for _, x := range utils.MakeSequence(s.Start.X, utils.AbsDiffInt(s.Start.X, s.End.X)+1, s.Start.X > s.End.X) {
			g[Pos{x, s.Start.Y}]++
		}
	} else if utils.AbsDiffInt(s.Start.X, s.End.X) == utils.AbsDiffInt(s.Start.Y, s.End.Y) {
		// fmt.Printf("Diagonal: %v\n", s)
		for i, x := range utils.MakeSequence(s.Start.X, utils.AbsDiffInt(s.Start.X, s.End.X)+1, s.Start.X > s.End.X) {
			step := 1
			if s.Start.Y > s.End.Y {
				step = -1
			}
			g[Pos{x, s.Start.Y + step*i}]++
		}
	}
}

func (g Grid) Count() int {
	count := 0
	for _, v := range g {
		if v > 1 {
			count++
		}
	}
	return count
}

func (g Grid) Print() {
	for y := 0; y < 10; y++ {
		for x := 0; x < 10; x++ {
			if v := g[Pos{x, y}]; v > 0 {
				print(v)
			} else {
				print(".")
			}
		}
		println()
	}
}

func solvePartA(input []string) int {
	grid := Grid{}
	segments := []Segment{}
	for _, line := range input {
		s := ParseLine(line)
		if s.Line() {
			segments = append(segments, s)
		}
	}

	for _, s := range segments {
		grid.Fill(s)
	}

	return grid.Count()
}

func solvePartB(input []string) int {
	grid := Grid{}
	segments := []Segment{}
	for _, line := range input {
		s := ParseLine(line)
		segments = append(segments, s)
	}
	for _, s := range segments {
		grid.Fill(s)
	}
	return grid.Count()
}
