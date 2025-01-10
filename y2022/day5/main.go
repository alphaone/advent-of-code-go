package day5

import (
	"fmt"
	"strings"
)

func parse(input string) (stack, []instruction) {
	parts := strings.Split(input, "\n\n")
	return parseStack(parts[0]), parseInstructions(parts[1])
}

type stack [][]rune

func parseStack(stack string) stack {
	stackLines := strings.Split(strings.TrimRight(stack, "\n"), "\n")
	stackLines = stackLines[:len(stackLines)-1]

	lineLength := 0

	for _, line := range stackLines {
		lineLength = max(lineLength, len(line))
	}

	res := make([][]rune, lineLength/4+1)
	for i := range res {
		res[i] = make([]rune, 0)
	}

	for i := range len(res) {
		for _, line := range stackLines {
			if len(line) < i*4+1 {
				continue
			}
			r := rune(line[i*4+1])
			if r == ' ' {
				continue
			}
			res[i] = append(res[i], r)
		}
	}

	return res
}

type instruction struct {
	amount int
	from   int
	to     int
}

func parseInstructions(instructions string) []instruction {
	instructionLines := strings.Split(strings.TrimRight(instructions, "\n"), "\n")

	res := make([]instruction, len(instructionLines))
	for i, line := range instructionLines {
		var ins instruction
		fmt.Sscanf(line, "move %d from %d to %d", &ins.amount, &ins.from, &ins.to)
		ins.from -= 1
		ins.to -= 1
		res[i] = ins
	}
	return res
}

func moveA(s *stack, ins instruction) {
	for range ins.amount {
		moveB(s, instruction{amount: 1, from: ins.from, to: ins.to})
	}
}

func moveB(s *stack, ins instruction) {
	tmp := make([]rune, ins.amount)
	copy(tmp, (*s)[ins.from][:ins.amount])
	(*s)[ins.from] = (*s)[ins.from][ins.amount:]
	(*s)[ins.to] = append(tmp, (*s)[ins.to]...)
}

func (s stack) String() string {
	res := ""
	for _, runes := range s {
		res += fmt.Sprintf("%v\n", string(runes))
	}
	return res
}

func solveA(stack stack, instructions []instruction, moveFn func(*stack, instruction)) string {
	for _, ins := range instructions {
		moveFn(&stack, ins)
	}

	toprow := make([]rune, len(stack))
	for i, s := range stack {
		toprow[i] = s[0]
	}
	return string(toprow)
}
