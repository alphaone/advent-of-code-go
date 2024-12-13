package day13

import (
	"regexp"
	"strconv"
	"strings"
)

type (
	Coord   struct{ x, y int }
	Machine struct {
		A     Coord
		B     Coord
		Prize Coord
	}
)

func parseInput(input string) []Machine {
	machineStrings := strings.Split(input, "\n\n")
	res := make([]Machine, 0)
	for _, str := range machineStrings {
		res = append(res, parseMachine(str))
	}
	return res
}

var r = regexp.MustCompile(`X\+(\d+), Y\+(\d+)\n.*X\+(\d+), Y\+(\d+)\n.*X=(\d+), Y=(\d+)`)

func parseMachine(str string) Machine {
	m := r.FindStringSubmatch(str)
	ax, err := strconv.Atoi(m[1])
	if err != nil {
		panic(err)
	}
	ay, err := strconv.Atoi(m[2])
	if err != nil {
		panic(err)
	}
	bx, err := strconv.Atoi(m[3])
	if err != nil {
		panic(err)
	}
	by, err := strconv.Atoi(m[4])
	if err != nil {
		panic(err)
	}
	px, err := strconv.Atoi(m[5])
	if err != nil {
		panic(err)
	}
	py, err := strconv.Atoi(m[6])
	if err != nil {
		panic(err)
	}
	return Machine{A: Coord{ax, ay}, B: Coord{bx, by}, Prize: Coord{px, py}}
}

// https://www.reddit.com/r/adventofcode/comments/1hd7irq/2024_day_13_an_explanation_of_the_mathematics/
func (m Machine) solve(offset int) (int, int) {
	prize := Coord{m.Prize.x + offset, m.Prize.y + offset}
	det := m.A.x*m.B.y - m.A.y*m.B.x
	a := (prize.x*m.B.y - prize.y*m.B.x) / det
	b := (m.A.x*prize.y - m.A.y*prize.x) / det

	p := Coord{m.A.x*a + m.B.x*b, m.A.y*a + m.B.y*b}
	if p == prize {
		return a, b
	}
	return -1, -1
}

func solve(machines []Machine, offset int) int {
	res := 0
	for _, m := range machines {
		a, b := m.solve(offset)
		if a != -1 && b != -1 {
			res += 3*a + b
		}
	}
	return res
}
