package day2

import (
	"regexp"
	"strconv"
	"strings"
)

type Game struct {
	Id    int
	Draws []Draw
}
type Draw map[string]int

// Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
var (
	reLine  = regexp.MustCompile(`Game (\d+): (.+)`)
	reDraws = regexp.MustCompile(`(\d+) (\w+)`)
)

func parseLine(line string) Game {
	res := reLine.FindStringSubmatch(line)
	id, _ := strconv.Atoi(res[1])
	drawsString := res[2]

	var draws []Draw
	for _, drawString := range strings.Split(drawsString, ";") {
		var draw Draw = make(Draw)
		for _, colorString := range strings.Split(drawString, ",") {
			res2 := reDraws.FindStringSubmatch(colorString)
			count, _ := strconv.Atoi(res2[1])
			draw[res2[2]] = count
		}
		draws = append(draws, draw)
	}
	return Game{
		Id:    id,
		Draws: draws,
	}
}

func (g Game) fewestDice() map[string]int {
	max := make(map[string]int)
	for _, color := range []string{"red", "green", "blue"} {
		for _, draw := range g.Draws {
			c := draw[color]
			if c > max[color] {
				max[color] = c
			}
		}
	}
	return max
}

func solvePartA(input []string) int {
	res := 0
	for _, line := range input {
		game := parseLine(line)

		if isValid(game.fewestDice()) {
			res += game.Id
		}
	}
	return res
}

// only 12 red cubes, 13 green cubes, and 14 blue cubes
func isValid(max map[string]int) bool {
	return max["red"] <= 12 && max["green"] <= 13 && max["blue"] <= 14
}

func solvePartB(input []string) int {
	res := 0
	for _, line := range input {
		game := parseLine(line)
		max := game.fewestDice()
		res += max["red"] * max["green"] * max["blue"]
	}
	return res
}
