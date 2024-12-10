package day5

import (
	"log"
	"slices"
	"strconv"
	"strings"
)

func parseInput(input string) Program {
	mem := make([]int, 0)
	for _, n := range strings.Split(input, ",") {
		i, _ := strconv.Atoi(n)
		mem = append(mem, i)
	}
	return Program{
		memory: mem,
		output: []int{},
	}
}

func solvePartA(p Program) int {
	p.run()
	return p.output[len(p.output)-1]
}

type Program struct {
	memory []int
	output []int
}

func (p *Program) get(pos int, immediate bool) int {
	if immediate {
		return (*p).memory[pos]
	}
	addr := (*p).memory[pos]
	return (*p).memory[addr]
}

func (p *Program) set(pos int, val int) {
	(*p).memory[pos] = val
}

func (p *Program) step(pos int) int {
	op, parameterModes := parseOp(p.get(pos, true))
	if op == 99 {
		return -1
	}

	instructionLength := 0
	switch op {
	case 1:
		// addition
		addr := p.get(pos+3, true)
		sum := p.get(pos+1, parameterModes[0] == 1) + p.get(pos+2, parameterModes[1] == 1)
		p.set(addr, sum)
		instructionLength = 4
	case 2:
		// multiplication
		addr := p.get(pos+3, true)
		product := p.get(pos+1, parameterModes[0] == 1) * p.get(pos+2, parameterModes[1] == 1)
		p.set(addr, product)
		instructionLength = 4
	case 3:
		// input
		addr := p.get(pos+1, true)
		p.set(addr, 1)
		instructionLength = 2
	case 4:
		// output
		o := p.get(pos+1, parameterModes[0] == 1)
		p.appendOutput(o)
		instructionLength = 2
	}
	return pos + instructionLength
}

func (p *Program) appendOutput(val int) {
	p.output = append(p.output, val)
}

func (p *Program) run() {
	pos := 0
	for {
		pos = p.step(pos)
		if pos == -1 {
			break
		}
		log.Print(p.output)
	}
}

func parseOp(i int) (int, map[int]int) {
	op := i % 100
	parameterModes := make(map[int]int)

	parameters := []rune(strconv.Itoa(i / 100))
	slices.Reverse(parameters)
	for i, v := range parameters {
		parameterModes[i] = int(v) - 48
	}
	return op, parameterModes
}
