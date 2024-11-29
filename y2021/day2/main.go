package day2

import (
	"regexp"
	"strconv"
)

type Coord struct {
	Horizontal int
	Depth      int
	Aim        int
}
type Command struct {
	Direction string
	Amount    int
}

func (c Coord) Score() int {
	return c.Horizontal * c.Depth
}

/*
 * "forward 8",
 * "up 3",
 * "down 8"
 */
func parseLine(line string) Command {
	re := regexp.MustCompile(`(forward|up|down) (\d+)`)
	result := re.FindStringSubmatch(line)
	distance, _ := strconv.Atoi(result[2])
	return Command{
		Direction: result[1],
		Amount:    distance,
	}
}

func solvePartA(input []string) Coord {
	var result Coord
	for _, line := range input {
		command := parseLine(line)
		switch command.Direction {
		case "forward":
			result.Horizontal += command.Amount
		case "up":
			result.Depth -= command.Amount
		case "down":
			result.Depth += command.Amount

		}
	}
	return result
}

func solvePartB(input []string) Coord {
	var result Coord
	for _, line := range input {
		command := parseLine(line)
		switch command.Direction {
		case "forward":
			result.Horizontal += command.Amount
			result.Depth += command.Amount * result.Aim
		case "up":
			result.Aim -= command.Amount
		case "down":
			result.Aim += command.Amount
		}
	}
	return result
}
