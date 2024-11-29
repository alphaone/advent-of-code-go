package day8

import (
	"regexp"
	"strings"

	"github.com/alphaone/advent/utils"
)

type Node struct {
	Name  string
	Left  string
	Right string
}

var reNode = regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)

func parseNodes(input []string) map[string]Node {
	result := make(map[string]Node)
	for _, line := range input {
		r := reNode.FindStringSubmatch(line)
		result[r[1]] = Node{r[1], r[2], r[3]}
	}
	return result
}

func solvePartA(input []string) int {
	directions := input[0]
	nodes := parseNodes(input[2:])

	current := "AAA"
	step := 0
	for {
		if current == "ZZZ" {
			break
		}

		direction := rune(directions[step%len(directions)])
		switch direction {
		case 'L':
			current = nodes[current].Left
		case 'R':
			current = nodes[current].Right

		}
		step++
	}

	return step
}

func solvePartB(input []string) int {
	directions := input[0]
	nodes := parseNodes(input[2:])

	startingNodes := findAllEndingWith(nodes, "A")

	var steps []int

	for _, nodeName := range startingNodes {
		step := 0

		for {
			if strings.HasSuffix(nodeName, "Z") {
				break
			}

			direction := rune(directions[step%len(directions)])
			switch direction {
			case 'L':
				nodeName = nodes[nodeName].Left
			case 'R':
				nodeName = nodes[nodeName].Right

			}
			step++
		}
		steps = append(steps, step)
	}

	lcm := steps[0]
	for _, x := range steps[1:] {
		lcm = utils.LCM(lcm, x)
	}

	return lcm
}

func findAllEndingWith(nodes map[string]Node, suffix string) []string {
	var result []string
	for nodeName := range nodes {
		if strings.HasSuffix(nodeName, suffix) {
			result = append(result, nodeName)
		}
	}
	return result
}

func allEndingWith(nodes []string, suffix string) bool {
	for _, nodeName := range nodes {
		if !strings.HasSuffix(nodeName, suffix) {
			return false
		}
	}
	return true
}

func countEndingWith(nodes []string, suffix string) int {
	result := 0
	for _, nodeName := range nodes {
		if strings.HasSuffix(nodeName, suffix) {
			result += 1
		}
	}
	return result
}
