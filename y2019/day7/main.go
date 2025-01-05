package day7

import (
	"fmt"
	"slices"
	"strconv"
	"strings"
	"sync"

	"github.com/alphaone/advent/utils"
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
		output: make(chan int),
		input:  make(chan int),
	}
}

func solve(p Program, phaseOptions []int) int {
	maxSignal := 0

	for _, phases := range combinations(phaseOptions) {
		chans := make([]chan int, 5)
		for i := range chans {
			chans[i] = make(chan int, 2)
		}

		var wg sync.WaitGroup
		for i, phase := range phases {
			wg.Add(1)
			pc := p.copy()
			pc.name = fmt.Sprintf("amp-%d", phase)
			pc.input = chans[i]
			pc.output = chans[(i+1)%5]
			pc.input <- phase

			go func() {
				defer wg.Done()
				pc.run()
			}()
		}
		chans[0] <- 0

		wg.Wait()
		output := <-chans[0]
		maxSignal = max(output, maxSignal)
	}

	return maxSignal
}

// more or less same as day 5, but with channels
type Program struct {
	name   string
	memory []int
	output chan int
	input  chan int
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
		// log.Println("waiting for input", p.name)
		*params[0] = <-p.input
		// log.Println("input", p.name, *params[0])
	case 4:
		// output
		p.output <- *params[0]
		// log.Println("output", p.name, *params[0])
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

// all combinations of ints
func combinations(options []int) [][]int {
	if len(options) == 1 {
		return [][]int{options}
	}

	res := make([][]int, 0)
	for i, x := range options {
		restOptions := utils.RemoveIndex(options, i)
		restCombinations := combinations(restOptions)
		for _, restComb := range restCombinations {
			c := append([]int{x}, restComb...)
			res = append(res, c)
		}
	}
	return res
}
