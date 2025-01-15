package day9

import (
	"fmt"

	"github.com/alphaone/advent/utils/mathutils"
	"github.com/alphaone/advent/utils/sliceutils"
)

type coord struct {
	l, c int
}
type instruction struct {
	dir    coord
	length int
}

func (c coord) add(other coord) coord {
	return coord{c.l + other.l, c.c + other.c}
}

func (c coord) dst(other coord) int {
	return mathutils.AbsDiffInt(c.l, other.l) + mathutils.AbsDiffInt(c.c, other.c)
}

func (tail coord) moveTowards(head coord) coord {
	if mathutils.AbsDiffInt(tail.l, head.l) <= 1 && mathutils.AbsDiffInt(tail.c, head.c) <= 1 {
		return tail
	}

	if mathutils.AbsDiffInt(head.l, tail.l) <= 1 {
		return coord{head.l, tail.c + (head.c-tail.c)/2}
	}
	if mathutils.AbsDiffInt(head.c, tail.c) <= 1 {
		return coord{tail.l + (head.l-tail.l)/2, head.c}
	}

	return coord{tail.l + (head.l-tail.l)/2, tail.c + (head.c-tail.c)/2}
}

func parse(input []string) []instruction {
	var res []instruction
	for _, line := range input {
		var dir string
		var length int
		fmt.Sscanf(line, "%s %d", &dir, &length)
		var c coord
		switch dir {
		case "R":
			c = coord{0, 1}
		case "L":
			c = coord{0, -1}
		case "U":
			c = coord{-1, 0}
		case "D":
			c = coord{1, 0}
		}
		res = append(res, instruction{c, length})
	}
	return res
}

func solveA(input []string) int {
	head := coord{0, 0}
	tail := coord{0, 0}

	visited := make(map[coord]bool)

	for _, ins := range parse(input) {
		for range ins.length {
			head = head.add(ins.dir)
			tail = tail.moveTowards(head)
			visited[tail] = true
		}
	}
	return len(visited)
}

func solveB(input []string) int {
	knots := make([]coord, 10)
	for i := range knots {
		knots[i] = coord{0, 0}
	}

	visited := make(map[coord]bool)
	for _, ins := range parse(input) {
		for range ins.length {
			knots[0] = knots[0].add(ins.dir)
			for i := range knots {
				if i == 0 {
					continue
				}

				knots[i] = knots[i].moveTowards(knots[i-1])
			}
			visited[knots[9]] = true
		}
	}
	return len(visited)
}

func visual(knots []coord) string {
	minL, maxL, minC, maxC := 0, 0, 0, 0
	for _, k := range knots {
		if k.l < minL {
			minL = k.l
		}
		if k.l > maxL {
			maxL = k.l
		}
		if k.c < minC {
			minC = k.c
		}
		if k.c > maxC {
			maxC = k.c
		}
	}

	res := make([][]rune, 0)
	for x := minL; x <= maxL; x++ {
		line := make([]rune, 0)
		for y := minC; y <= maxC; y++ {
			line = append(line, '.')
		}
		res = append(res, line)
	}

	adjustL, adjustC := -min(0, minL), -min(0, minC)

	for i, k := range knots {
		res[k.l+adjustL][k.c+adjustC] = rune(i + '0')
	}

	res[adjustL][adjustC] = 'S'

	return string(sliceutils.Join(res, '\n'))
}
