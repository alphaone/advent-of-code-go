package day23

import (
	"fmt"
	"slices"
)

type Pos struct {
	Line, Col int
}

func (p Pos) add(other Pos) Pos {
	return Pos{Line: p.Line + other.Line, Col: p.Col + other.Col}
}

func FindStart(input []string) Pos {
	line := input[0]
	for col, char := range line {
		if char == '.' {
			return Pos{Line: 0, Col: col}
		}
	}
	panic("No start found")
}

func FindEnd(input []string) Pos {
	line := input[len(input)-1]
	for col, char := range line {
		if char == '.' {
			return Pos{Line: len(input) - 1, Col: col}
		}
	}
	panic("No end found")
}

func FindPointsOfInterest(input []string) []Pos {
	res := []Pos{FindStart(input), FindEnd(input)}

	for l, line := range input {
		for c, char := range line {
			if char == '#' {
				continue
			}
			neighbors := 0
			for _, dir := range []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}} {
				np := Pos{l, c}.add(dir)
				if inbounds(input, np) && input[np.Line][np.Col] != '#' {
					neighbors++
				}
			}
			if neighbors >= 3 {
				res = append(res, Pos{l, c})
			}
		}
	}
	return res
}

func inbounds(input []string, pos Pos) bool {
	return pos.Line >= 0 && pos.Line < len(input) &&
		pos.Col >= 0 && pos.Col < len(input[pos.Line])
}

type StackNode struct {
	Pos      Pos
	Distance int
}

type Graph map[Pos]map[Pos]int

func BuildGraph(pois []Pos, input []string, possibleDirections func(rune) []Pos) Graph {
	graph := make(Graph)
	for _, p := range pois {
		graph[p] = make(map[Pos]int)
	}

	for _, p := range pois {
		stack := []StackNode{{Pos: p, Distance: 0}}
		seen := make(map[Pos]bool)
		seen[p] = true

		for len(stack) > 0 {
			popped := stack[0]
			stack = stack[1:]

			if popped.Distance != 0 && slices.Contains(pois, popped.Pos) {
				graph[p][popped.Pos] = popped.Distance
				continue
			}

			for _, dir := range possibleDirections(rune(input[popped.Pos.Line][popped.Pos.Col])) {
				np := popped.Pos.add(dir)
				if inbounds(input, np) && input[np.Line][np.Col] != '#' && !seen[np] {
					seen[np] = true
					stack = append(stack, StackNode{Pos: np, Distance: popped.Distance + 1})
				}
			}
		}
	}

	return graph
}

func possibleDirectionsA(char rune) []Pos {
	switch char {
	case '^':
		return []Pos{{-1, 0}}
	case 'v':
		return []Pos{{1, 0}}
	case '<':
		return []Pos{{0, -1}}
	case '>':
		return []Pos{{0, 1}}
	case '.':
		return []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	default:
		panic("Unknown direction")
	}
}

func possibleDirectionsB(char rune) []Pos {
	return []Pos{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
}

func FindLongestRoute(graph Graph, start, end Pos, seen []Pos) (int, error) {
	if start == end {
		return 0, nil
	}

	maxDistance := -1
	seen = append(seen, start)
	for next, nodeDistance := range graph[start] {
		if !slices.Contains(seen, next) {
			d, err := FindLongestRoute(graph, next, end, seen)
			if err != nil {
				continue
			}
			distance := nodeDistance + d
			maxDistance = max(maxDistance, distance)
		}
	}
	if maxDistance == -1 {
		return 0, fmt.Errorf("sackgasse")
	}
	return maxDistance, nil
}

func solvePartA(input []string) int {
	pois := FindPointsOfInterest(input)
	graph := BuildGraph(pois, input, possibleDirectionsA)
	d, err := FindLongestRoute(graph, FindStart(input), FindEnd(input), []Pos{})
	if err != nil {
		panic(err)
	}
	return d
}

func solvePartB(input []string) int {
	pois := FindPointsOfInterest(input)
	graph := BuildGraph(pois, input, possibleDirectionsB)
	d, err := FindLongestRoute(graph, FindStart(input), FindEnd(input), []Pos{})
	if err != nil {
		panic(err)
	}
	return d
}
