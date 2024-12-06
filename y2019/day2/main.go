package day2

import (
	"strconv"
	"strings"
)

func parseInput(input string) Program {
	res := make(Program, 0)
	for _, n := range strings.Split(input, ",") {
		i, _ := strconv.Atoi(n)
		res = append(res, i)
	}
	return res
}

func solvePartA(p Program) int {
	p.setNounAndVerb(12, 2)
	p.run()
	return p.output()
}

func solvePartB(original Program) int {
	for noun := 0; noun < 100; noun++ {
		for verb := 0; verb < 100; verb++ {
			p := original.copy()
			p.setNounAndVerb(noun, verb)
			p.run()
			if p.output() == 19690720 {
				return noun*100 + verb
			}
		}
	}
	panic("nothing found")
}

type Program []int

func (p Program) copy() Program {
	res := make(Program, len(p))
	for i, x := range p {
		res[i] = x
	}
	return res
}

func (p *Program) step(pos int) int {
	op := (*p)[pos]
	if op == 99 {
		return -1
	}

	idxA := (*p)[pos+1]
	idxB := (*p)[pos+2]
	idxC := (*p)[pos+3]

	if op == 1 {
		(*p)[idxC] = (*p)[idxA] + (*p)[idxB]
	} else if op == 2 {
		(*p)[idxC] = (*p)[idxA] * (*p)[idxB]
	}
	return pos + 4
}

func (p *Program) setNounAndVerb(noun int, verb int) {
	(*p)[1] = noun
	(*p)[2] = verb
}

func (p *Program) output() int {
	return (*p)[0]
}

func (p *Program) run() {
	pos := 0
	for {
		pos = p.step(pos)
		if pos == -1 {
			break
		}
	}
}
