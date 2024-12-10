package day5

import (
	"slices"
	"strconv"
	"strings"
)

func parseInput(input string) Program {
	mem := make([]int, 0)
	for _, n := range strings.Split(strings.TrimSpace(input), ",") {
		i, err := strconv.Atoi(n)
		if err != nil {
			panic(err)
		}
		mem = append(mem, i)
	}
	return Program{
		memory: mem,
		output: []int{},
	}
}

func solve(p Program) int {
	p.run()
	return p.output[len(p.output)-1]
}

type Program struct {
	memory []int
	output []int
	input  int
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
		*params[0] = p.input
	case 4:
		// output
		p.appendOutput(*params[0])
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
	}
	return pos + 1
}

var noOfParams = map[int]int{
	1: 3, 2: 3,
	3: 1, 4: 1,
	5: 2, 6: 2,
	7: 3, 8: 3,
	99: 0,
}

func (p *Program) params(op int, pos int) []*int {
	parameterModes := parseParamModes((*p).memory[pos])

	params := make([]*int, 0)
	for i := range noOfParams[op] {
		var val *int
		if parameterModes[i] == 0 { // position mode
			addr := (*p).memory[pos+i+1]
			val = &(*p).memory[addr]
		} else { // direct mode
			val = &(*p).memory[pos+i+1]
		}
		params = append(params, val)
	}
	return params
}

func (p *Program) appendOutput(val int) {
	p.output = append(p.output, val)
}

func (p *Program) run() {
	p.output = []int{}
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
