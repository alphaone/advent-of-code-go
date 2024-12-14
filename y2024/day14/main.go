package day14

import (
	"log"
	"regexp"
	"strconv"
	"strings"
)

type (
	Coord struct{ X, Y int }
	Robot struct{ pos, v Coord }
)

func (c Coord) move(v Coord, limit Coord, times int) Coord {
	newX := (c.X + times*v.X + times*limit.X) % limit.X
	newY := (c.Y + times*v.Y + times*limit.Y) % limit.Y
	return Coord{newX, newY}
}

func (c Coord) q(limit Coord) int {
	midX := (limit.X) / 2
	midY := (limit.Y) / 2

	if c.X < midX && c.Y < midY {
		return 1
	}
	if c.X > midX && c.Y < midY {
		return 2
	}
	if c.X < midX && c.Y > midY {
		return 3
	}
	if c.X > midX && c.Y > midY {
		return 4
	}
	return 0
}

func visual(robots []Robot, limit Coord) string {
	grid := make([][]int, limit.X)
	for x := range limit.X {
		grid[x] = make([]int, limit.Y)
	}

	for _, r := range robots {
		grid[r.pos.X][r.pos.Y]++
	}

	res := ""
	for y := range limit.Y {
		line := ""
		for x := range limit.X {
			if grid[x][y] == 0 {
				line += " "
			} else {
				line += strconv.Itoa(grid[x][y])
			}
		}
		res += line + "\n"
	}
	return res
}

var r = regexp.MustCompile(`p=(\d+),(\d+) v=(-?\d+),(-?\d+)`)

func Parse(input []string) []Robot {
	res := make([]Robot, 0)
	for _, line := range input {
		m := r.FindStringSubmatch(line)
		posX, _ := strconv.Atoi(m[1])
		posY, _ := strconv.Atoi(m[2])
		vX, _ := strconv.Atoi(m[3])
		vY, _ := strconv.Atoi(m[4])
		res = append(res, Robot{
			pos: Coord{posX, posY},
			v:   Coord{vX, vY},
		})
	}
	return res
}

func solvePartA(robots []Robot, limit Coord) int {
	for i, r := range robots {
		r.pos = r.pos.move(r.v, limit, 100)
		robots[i] = r
	}

	quadrants := make(map[int]int)
	for _, r := range robots {
		q := r.pos.q(limit)
		if q == 0 {
			continue
		}
		quadrants[q]++
	}
	return quadrants[1] * quadrants[2] * quadrants[3] * quadrants[4]
}

func SolvePartB(robots []Robot, limit Coord) int {
	x := 1
	for x < 100000 {
		for i, r := range robots {
			r.pos = r.pos.move(r.v, limit, 1)
			robots[i] = r
		}
		if !strings.Contains(visual(robots, limit), "11111111") {
			x++
			continue
		}
		log.Print("\033[H\033[2J")
		log.Print(visual(robots, limit))
		log.Print(x)
		return x

	}

	return -1
}
