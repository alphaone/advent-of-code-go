package day10

import (
	"fmt"
	"log"
	"slices"
)

type cpu struct {
	registerX   int
	currentTask *task
}

type task struct {
	instruction instruction
	cyclesLeft  int
}

type instruction struct {
	op  string
	arg int
}

var cycleCount = map[string]int{
	"noop": 1,
	"addx": 2,
}

func (c *cpu) tick(instructions *[]instruction) {
	if c.currentTask == nil {
		instr := (*instructions)[0]
		*instructions = (*instructions)[1:]
		c.currentTask = &task{instruction: instr, cyclesLeft: cycleCount[instr.op]}
	}

	// after
	c.currentTask.cyclesLeft--
	if c.currentTask.cyclesLeft == 0 {
		if c.currentTask.instruction.op == "addx" {
			c.registerX += c.currentTask.instruction.arg
		}
		c.currentTask = nil
	}
}

func parse(input []string) []instruction {
	res := make([]instruction, 0)
	for _, l := range input {
		in := instruction{}
		fmt.Sscanf(l, "%s %d", &in.op, &in.arg)
		res = append(res, in)
	}
	return res
}

func solve(input []string) int {
	instructions := parse(input)
	c := cpu{registerX: 1}

	interestingCycles := []int{20, 60, 100, 140, 180, 220}
	res := 0
	for i := 1; i <= 220; i++ {
		if slices.Contains(interestingCycles, i) {
			log.Printf("cycle %d: %v", i, c.registerX)
			res += i * c.registerX
		}
		c.tick(&instructions)
	}

	return res
}

func solveB(input []string) string {
	instructions := parse(input)
	c := cpu{registerX: 1}

	res := ""

	for i := 1; i <= 240; i++ {
		// fmt.Printf("cycle %d: %v\n", i, c.registerX)
		x := (i-1)%40 - c.registerX
		if x >= -1 && x <= 1 {
			res += "#"
		} else {
			res += "."
		}
		if i%40 == 0 {
			res += fmt.Sprintf("%d\n", i)
		}

		c.tick(&instructions)
	}
	return res
}
