package day16

import (
	"fmt"
	"slices"
)

type Node struct {
	Mirror        rune
	EntryBeamDirs []rune
}
type Field [][]Node

func parseInput(input []string) Field {
	result := make(Field, len(input))
	for i, line := range input {
		result[i] = make([]Node, len(line))
		for j, c := range line {
			result[i][j] = Node{Mirror: c, EntryBeamDirs: []rune{}}
		}
	}
	return result
}

func newCoords(line, col int, dir rune) (int, int) {
	switch dir {
	case 'N':
		line--
	case 'S':
		line++
	case 'E':
		col++
	case 'W':
		col--
	}
	return line, col
}

func advance(field Field, line, col int, dir rune) Field {
	// fmt.Println(line, col, dir)
	node := field[line][col]
	node.EntryBeamDirs = append(node.EntryBeamDirs, dir)
	field[line][col] = node

	dirs := []rune{}
	switch node.Mirror {
	case '.':
		dirs = []rune{dir}
	case '/':
		switch dir {
		case 'N':
			dirs = []rune{'E'}
		case 'E':
			dirs = []rune{'N'}
		case 'S':
			dirs = []rune{'W'}
		case 'W':
			dirs = []rune{'S'}
		}
	case '\\':
		switch dir {
		case 'N':
			dirs = []rune{'W'}
		case 'E':
			dirs = []rune{'S'}
		case 'S':
			dirs = []rune{'E'}
		case 'W':
			dirs = []rune{'N'}
		}
	case '-':
		switch dir {
		case 'N', 'S':
			dirs = []rune{'E', 'W'}
		case 'E', 'W':
			dirs = []rune{dir}

		}
	case '|':
		switch dir {
		case 'N', 'S':
			dirs = []rune{dir}
		case 'E', 'W':
			dirs = []rune{'N', 'S'}
		}
	}

	for _, dir := range dirs {
		// been here already?
		nl, nc := newCoords(line, col, dir)
		if nl < 0 || nl >= len(field) || nc < 0 || nc >= len(field[0]) {
			continue
		}
		if !slices.Contains(field[nl][nc].EntryBeamDirs, dir) {
			field = advance(field, nl, nc, dir)
		}
	}

	return field
}

func solvePartA(input []string) int {
	field := parseInput(input)
	field = advance(field, 0, 0, 'E')
	return field.Energy()
}

func (f Field) String() string {
	result := ""
	for _, line := range f {
		var lineResult string
		for _, node := range line {
			switch node.Mirror {
			case '/', '\\', '-', '|':
				lineResult += string(node.Mirror)
			case '.':
				if len(node.EntryBeamDirs) == 0 {
					lineResult += "."
				} else if len(node.EntryBeamDirs) == 1 {
					switch node.EntryBeamDirs[0] {
					case 'N':
						lineResult += "^"
					case 'S':
						lineResult += "v"
					case 'E':
						lineResult += ">"
					case 'W':
						lineResult += "<"
					}
				} else {
					lineResult += fmt.Sprint(len(node.EntryBeamDirs))
				}
			}
		}
		result += lineResult + "\n"
	}
	return result
}

func (f Field) Energy() int {
	result := 0
	for _, line := range f {
		for _, node := range line {
			if len(node.EntryBeamDirs) > 0 {
				result++
			}
		}
	}
	return result
}

func solvePartB(input []string) int {
	field := parseInput(input)

	energies := []int{}

	for col := 0; col < len(field[0]); col++ {
		field := parseInput(input)
		field = advance(field, 0, col, 'S')
		energies = append(energies, field.Energy())

		field = parseInput(input)
		field = advance(field, len(field)-1, col, 'N')
		energies = append(energies, field.Energy())
	}
	for line := 0; line < len(field); line++ {
		field = parseInput(input)
		field = advance(field, line, 0, 'E')
		energies = append(energies, field.Energy())

		field = parseInput(input)
		field = advance(field, line, len(field[0])-1, 'W')
		energies = append(energies, field.Energy())
	}

	return slices.Max(energies)
}
