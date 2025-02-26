package day8

import (
	"fmt"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"nop +0",
	"acc +1",
	"jmp +4",
	"acc +3",
	"jmp -3",
	"acc -99",
	"acc +1",
	"jmp -4",
	"acc +6",
}

type cmd struct {
	op  string
	val int
}
type machine struct {
	cmds []cmd
	acc  int
	pos  int
}

func (m *machine) run() error {
	seen := make(map[int]bool)
	for {
		_, ok := seen[m.pos]
		if ok {
			return fmt.Errorf("infinite loop")
		}
		if m.pos >= len(m.cmds) {
			return nil
		}

		nextPos := m.pos + 1
		switch m.cmds[m.pos].op {
		case "nop":
		case "acc":
			m.acc += m.cmds[m.pos].val
		case "jmp":
			nextPos = m.pos + m.cmds[m.pos].val
		}

		seen[m.pos] = true
		m.pos = nextPos
	}
}

func parse(input []string) []cmd {
	res := make([]cmd, len(input))

	for i, line := range input {
		var op string
		var val int
		fmt.Sscanf(line, "%s %d", &op, &val)
		res[i] = cmd{op, val}
	}

	return res
}

func patched(cmds []cmd) []machine {
	res := make([]machine, 0)

	for i, c := range cmds {
		switch c.op {
		case "nop":
			copyCmds := make([]cmd, len(cmds))
			copy(copyCmds, cmds)
			copyCmds[i].op = "jmp"
			res = append(res, machine{cmds: copyCmds})
		case "jmp":
			copyCmds := make([]cmd, len(cmds))
			copy(copyCmds, cmds)
			copyCmds[i].op = "nop"
			res = append(res, machine{cmds: copyCmds})
		case "acc":
		}
	}

	return res
}

func solveA(cmds []cmd) int {
	m := machine{cmds: cmds}
	m.run()
	return m.acc
}

func solveB(cmds []cmd) int {
	for _, m := range patched(cmds) {
		err := m.run()
		if err == nil {
			return m.acc
		}
	}
	panic("no solution found")
}

func TestParse(t *testing.T) {
	assert.Equal(t, []cmd{
		{op: "nop", val: 0},
		{op: "acc", val: 1},
		{op: "jmp", val: 4},
		{op: "acc", val: 3},
		{op: "jmp", val: -3},
		{op: "acc", val: -99},
		{op: "acc", val: 1},
		{op: "jmp", val: -4},
		{op: "acc", val: 6},
	}, parse(example))
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 5, solveA(parse(example)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 1586, solveA(parse(utils.LoadStrings("input.txt"))))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 8, solveB(parse(example)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 703, solveB(parse(utils.LoadStrings("input.txt"))))
}
