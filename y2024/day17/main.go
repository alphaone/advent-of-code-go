package day17

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type computer struct {
	A, B, C int
	Program []int
	Output  []int
}

func (c computer) copy() computer {
	cc := computer{
		A:       c.A,
		B:       c.B,
		C:       c.C,
		Program: make([]int, len(c.Program)),
		Output:  []int{},
	}
	copy(cc.Program, c.Program)
	return cc
}

func (c computer) outputAsString() string {
	outputs := make([]string, len(c.Output))
	for i, x := range c.Output {
		outputs[i] = strconv.Itoa(x)
	}
	return strings.Join(outputs, ",")
}

var r = regexp.MustCompile(`A: (\d+)\n.*B: (\d+)\n.*C: (\d+)\n\nProgram: ([\d,]*)`)

func parse(input string) computer {
	match := r.FindStringSubmatch(input)
	a, err := strconv.Atoi(match[1])
	if err != nil {
		panic(err)
	}
	b, err := strconv.Atoi(match[2])
	if err != nil {
		panic(err)
	}
	c, err := strconv.Atoi(match[3])
	if err != nil {
		panic(err)
	}

	program := make([]int, 0)
	for _, s := range strings.Split(match[4], ",") {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		program = append(program, i)
	}

	return computer{A: a, B: b, C: c, Program: program, Output: []int{}}
}

func (c *computer) readComboParam(pos int) int {
	p := c.Program[pos]
	if 0 <= p && p <= 3 {
		return p
	}
	if p == 4 {
		return c.A
	}
	if p == 5 {
		return c.B
	}
	if p == 6 {
		return c.C
	}
	panic("invalid combo param")
}

func (c *computer) step(pos int) int {
	if pos >= len(c.Program) {
		return -1
	}
	op := c.Program[pos]

	switch op {
	case 0: // adv
		operand := c.readComboParam(pos + 1)
		numerator := c.A
		denominator := math.Pow(2, float64(operand))
		c.A = int(float64(numerator) / denominator)
	case 1: // bxl
		operand := c.Program[pos+1]
		c.B = c.B ^ operand
	case 2: // bst
		operand := c.readComboParam(pos + 1)
		c.B = operand % 8
	case 3: // jnz
		if c.A != 0 {
			operand := c.Program[pos+1]
			return operand
		}
	case 4: // bxc
		c.B = c.B ^ c.C
	case 5: // out
		operand := c.readComboParam(pos + 1)
		c.Output = append(c.Output, operand%8)
	case 6: // bdv
		operand := c.readComboParam(pos + 1)
		numerator := c.A
		denominator := math.Pow(2, float64(operand))
		c.B = int(float64(numerator) / denominator)
	case 7: // cdv
		operand := c.readComboParam(pos + 1)
		numerator := c.A
		denominator := math.Pow(2, float64(operand))
		c.C = int(float64(numerator) / denominator)
	}

	return pos + 2
}

func (c *computer) run() {
	c.Output = []int{}
	pos := 0
	for {
		pos = c.step(pos)
		if pos == -1 {
			break
		}
	}
}

func solvePartA(c computer) string {
	c.run()
	return c.outputAsString()
}
