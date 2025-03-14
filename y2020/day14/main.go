package day14

import (
	"fmt"
	"math"
)

type (
	mask        string
	instruction struct {
		op    string
		mask  mask
		addr  int
		value int
	}
)

func (m mask) apply(value int) int {
	maskLen := len(m)
	orMask := 0
	andMask := int(math.Pow(2, float64(maskLen)) - 1)
	for i, c := range m {
		switch c {
		case '1':
			orMask |= 1 << (maskLen - 1 - i)
		case '0':
			andMask &= ^(1 << (maskLen - 1 - i))
		case 'X':
			andMask |= 1 << (maskLen - 1 - i)
			orMask &= ^(1 << (maskLen - 1 - i))
		}
	}

	return value&andMask | orMask
}

func parse(input []string) []instruction {
	res := make([]instruction, len(input))

	for i, line := range input {
		if line[:4] == "mask" {
			res[i].op = "mask"
			res[i].mask = mask(line[7:])
		} else {
			res[i].op = "mem"
			fmt.Sscanf(line, "mem[%d] = %d", &res[i].addr, &res[i].value)
		}
	}

	return res
}

func solveA(instructions []instruction) int {
	mem := map[int]int{}
	var mask mask

	for _, instr := range instructions {
		switch instr.op {
		case "mask":
			mask = instr.mask
		case "mem":
			mem[instr.addr] = mask.apply(instr.value)
		}
	}

	sum := 0
	for _, v := range mem {
		sum += v
	}

	return sum
}
