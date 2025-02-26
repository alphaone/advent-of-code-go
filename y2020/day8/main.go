package day8

import "fmt"

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
