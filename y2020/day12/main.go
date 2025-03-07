package day12

import (
	"strconv"

	"github.com/alphaone/advent/utils/mathutils"
)

type instruction struct {
	move   rune
	amount int
}

func parseInput(input []string) []instruction {
	result := make([]instruction, len(input))
	for i, line := range input {
		move := rune(line[0])
		amount, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		result[i] = instruction{move, amount}
	}
	return result
}

type ship struct {
	pos coord
	dir rune
}

type coord struct {
	l, c int
}

func move(s ship, ins instruction) ship {
	switch ins.move {
	case 'N':
		return ship{coord{s.pos.l - ins.amount, s.pos.c}, s.dir}
	case 'S':
		return ship{coord{s.pos.l + ins.amount, s.pos.c}, s.dir}
	case 'E':
		return ship{coord{s.pos.l, s.pos.c + ins.amount}, s.dir}
	case 'W':
		return ship{coord{s.pos.l, s.pos.c - ins.amount}, s.dir}
	case 'L':
		return s.turn(-ins.amount)
	case 'R':
		return s.turn(ins.amount)
	case 'F':
		return move(s, instruction{s.dir, ins.amount})
	}
	panic("unexpected move")
}

func (s ship) turn(angle int) ship {
	angles := []rune{'N', 'E', 'S', 'W'}
	idx := 0
	for i, a := range angles {
		if a == s.dir {
			idx = i
			break
		}
	}
	idx = (idx + angle/90) % 4
	if idx < 0 {
		idx += 4
	}
	return ship{s.pos, angles[idx]}
}

func solveA(input []instruction) int {
	s := ship{coord{0, 0}, 'E'}
	for _, ins := range input {
		s = move(s, ins)
	}
	return mathutils.AbsInt(s.pos.l) + mathutils.AbsInt(s.pos.c)
}

func moveWaypoint(wp coord, ins instruction) coord {
	switch ins.move {
	case 'N':
		return coord{wp.l - ins.amount, wp.c}
	case 'S':
		return coord{wp.l + ins.amount, wp.c}
	case 'E':
		return coord{wp.l, wp.c + ins.amount}
	case 'W':
		return coord{wp.l, wp.c - ins.amount}
	case 'L':
		return rotateWaypoint(wp, -ins.amount)
	case 'R':
		return rotateWaypoint(wp, ins.amount)
	}
	panic("unexpected move")
}

func rotateWaypoint(wp coord, angle int) coord {
	for range 4 + (angle / 90) {
		wp = coord{wp.c, -wp.l}
	}
	return wp
}

func solveB(input []instruction) int {
	s := coord{0, 0}
	wp := coord{-1, 10}
	for _, ins := range input {
		if ins.move == 'F' {
			s = coord{s.l + wp.l*ins.amount, s.c + wp.c*ins.amount}
		} else {
			wp = moveWaypoint(wp, ins)
		}
	}
	return mathutils.AbsInt(s.l) + mathutils.AbsInt(s.c)
}
