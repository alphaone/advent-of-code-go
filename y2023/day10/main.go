package day10

type Point struct {
	Line int
	Col  int
	Char rune
}

func FindStart(input []string) Point {
	for y, line := range input {
		for x, char := range line {
			if char == 'S' {
				return Point{Line: y, Col: x, Char: 'S'}
			}
		}
	}
	panic("no start found")
}

func OptionsFromStart(input []string, start Point) (rune, rune) {
	options := []rune{}
	if start.Line > 0 {
		north := input[start.Line-1][start.Col]
		if north == '|' || north == 'F' || north == '7' {
			options = append(options, 'N')
		}
	}
	if start.Col < len(input[0])-1 {
		east := input[start.Line][start.Col+1]
		if east == '-' || east == 'J' || east == '7' {
			options = append(options, 'E')
		}
	}
	if start.Line < len(input)-1 {
		south := input[start.Line+1][start.Col]
		if south == '|' || south == 'L' || south == 'J' {
			options = append(options, 'S')
		}
	}
	if start.Col > 0 {
		west := input[start.Line][start.Line-1]
		if west == '-' || west == 'F' || west == 'L' {
			options = append(options, 'W')
		}
	}

	switch string(options) {
	case "NS":
		return 'N', '|'
	case "EW":
		return 'E', '-'
	case "NE":
		return 'N', 'L'
	case "NW":
		return 'N', 'J'
	case "ES":
		return 'E', 'F'
	case "SW":
		return 'S', '7'
	}
	panic("no options found")
}

func solvePartA(input []string) int {
	return len(Route(input)) / 2
}

func Route(input []string) []Point {
	start := FindStart(input)
	startDir, startRune := OptionsFromStart(input, start)
	route := []Point{}

	dir := startDir
	current := Point{Col: start.Col, Line: start.Line, Char: startRune}
	for {
		route = append(route, current)

		current = Walk(current, dir)
		current.Char = rune(input[current.Line][current.Col])

		if current.Char == 'S' {
			break
		}
		dir = Direction(current.Char, dir)
	}
	return route
}

func Walk(current Point, dir rune) Point {
	switch dir {
	case 'N':
		return Point{Col: current.Col, Line: current.Line - 1}
	case 'E':
		return Point{Col: current.Col + 1, Line: current.Line}
	case 'S':
		return Point{Col: current.Col, Line: current.Line + 1}
	case 'W':
		return Point{Col: current.Col - 1, Line: current.Line}
	}
	panic("invalid direction")
}

func Direction(x rune, lastDir rune) rune {
	switch x {
	case '|':
		if lastDir == 'N' {
			return 'N'
		} else {
			return 'S'
		}
	case '-':
		if lastDir == 'E' {
			return 'E'
		} else {
			return 'W'
		}
	case 'L':
		if lastDir == 'S' {
			return 'E'
		} else {
			return 'N'
		}
	case 'J':
		if lastDir == 'S' {
			return 'W'
		} else {
			return 'N'
		}
	case '7':
		if lastDir == 'E' {
			return 'S'
		} else {
			return 'W'
		}
	case 'F':
		if lastDir == 'W' {
			return 'S'
		} else {
			return 'E'
		}
	}

	panic("invalid direction")
}

func solvePartB(input []string) int {
	blank := make([][]rune, len(input))
	for i := range blank {
		blank[i] = make([]rune, len(input[0]))
		for j := range input[0] {
			blank[i][j] = '.'
		}
	}
	for _, p := range Route(input) {
		blank[p.Line][p.Col] = p.Char
	}

	in := "out"
	enclosed := 0
	for y, line := range blank {
		for x, char := range line {
			switch char {
			case '|':
				if in == "out" {
					in = "in"
				} else {
					in = "out"
				}
			case 'F':
				if in == "out" {
					in = "half-in-from-s"
				} else {
					in = "half-in-from-n"
				}
			case 'L':
				if in == "out" {
					in = "half-in-from-n"
				} else {
					in = "half-in-from-s"
				}
			case '7':
				if in == "half-in-from-n" {
					in = "in"
				} else {
					in = "out"
				}
			case 'J':
				if in == "half-in-from-s" {
					in = "in"
				} else {
					in = "out"
				}
			}
			if in == "in" && char == '.' {
				enclosed++
				blank[y][x] = 'X'
			}
		}
	}

	return enclosed
}
