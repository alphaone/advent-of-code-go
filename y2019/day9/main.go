package day9

import (
	"slices"
	"strconv"
	"strings"
)

func parseInput(input string) Program {
	mem := make([]int, 1e6)
	for idx, n := range strings.Split(strings.TrimSpace(input), ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		mem[idx] = i
	}
	return Program{
		memory: mem,
		output: []int{},
		input:  []int{},
	}
}

func solve(p Program, input int) int {
	p.input = []int{input}
	p.run()
	return p.output[0]
}

type Program struct {
	name         string
	memory       []int
	relativeBase int
	output       []int
	input        []int
}

func (p Program) copy() Program {
	mem := make([]int, len(p.memory))
	copy(mem, p.memory)
	return Program{name: p.name, memory: mem}
}

func (p *Program) step(pos int) int {
	op := parseOp((*p).memory[pos])
	if op == 99 {
		return -1
	}

	params := p.params(op, pos)
	pos += len(params)

	switch op {
	case 1:
		// addition
		*params[2] = *params[0] + *params[1]
	case 2:
		// multiplication
		*params[2] = *params[0] * *params[1]
	case 3:
		// input
		head, tail := p.input[0], p.input[1:]
		*params[0] = head
		p.input = tail
	case 4:
		// output
		p.output = append(p.output, *params[0])
	case 5:
		// jump if true
		if *params[0] != 0 {
			return *params[1]
		}
	case 6:
		// jump if false
		if *params[0] == 0 {
			return *params[1]
		}
	case 7:
		// less than
		if *params[0] < *params[1] {
			*params[2] = 1
		} else {
			*params[2] = 0
		}
	case 8:
		// equal
		if *params[0] == *params[1] {
			*params[2] = 1
		} else {
			*params[2] = 0
		}
	case 9:
		// relative base offset
		p.relativeBase += *params[0]
	}
	return pos + 1
}

var noOfParams = map[int]int{
	1: 3, 2: 3,
	3: 1, 4: 1,
	5: 2, 6: 2,
	7: 3, 8: 3,
	9:  1,
	99: 0,
}

func (p *Program) params(op int, pos int) []*int {
	parameterModes := parseParamModes((*p).memory[pos])

	params := make([]*int, 0)
	for i := range noOfParams[op] {
		var val *int
		switch parameterModes[i] {
		case 0: // position mode
			addr := (*p).memory[pos+i+1]
			val = &(*p).memory[addr]
		case 1: // direct mode
			val = &(*p).memory[pos+i+1]
		case 2: // relative mode
			addr := (*p).memory[pos+i+1] + p.relativeBase
			val = &(*p).memory[addr]
		}
		params = append(params, val)
	}
	return params
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

func parseOp(fullop int) int {
	return fullop % 100
}

func parseParamModes(fullop int) map[int]int {
	parameterModes := make(map[int]int)
	parameters := []rune(strconv.Itoa(fullop / 100))
	slices.Reverse(parameters)
	for i, v := range parameters {
		parameterModes[i] = int(v) - 48
	}
	return parameterModes
}
