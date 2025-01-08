package day2

import "fmt"

type game struct {
	opp    string
	resp   string
	intent string
}

func parse(input []string) []game {
	var games []game
	for _, i := range input {
		var g game
		fmt.Sscanf(i, "%s %s", &g.opp, &g.resp)
		g.intent = g.resp
		games = append(games, g)
	}
	return games
}

// A Rock, B Paper, C Scissors
// X Rock, Y Paper, Z Scissors (part 1)
// X Lose, Y Draw, Z Win (part 2)

var scoreSelected = map[string]int{
	"X": 1,
	"Y": 2,
	"Z": 3,
}

var scoreOutcome = map[[2]string]int{
	{"X", "A"}: 3,
	{"X", "B"}: 0,
	{"X", "C"}: 6,
	{"Y", "A"}: 6,
	{"Y", "B"}: 3,
	{"Y", "C"}: 0,
	{"Z", "A"}: 0,
	{"Z", "B"}: 6,
	{"Z", "C"}: 3,
}

var selectForIntent = map[[2]string]string{
	{"A", "X"}: "Z",
	{"A", "Y"}: "X",
	{"A", "Z"}: "Y",
	{"B", "X"}: "X",
	{"B", "Y"}: "Y",
	{"B", "Z"}: "Z",
	{"C", "X"}: "Y",
	{"C", "Y"}: "Z",
	{"C", "Z"}: "X",
}

func (g game) score() int {
	return scoreOutcome[[2]string{g.resp, g.opp}] + scoreSelected[g.resp]
}

func (g *game) selectResp() {
	g.resp = selectForIntent[[2]string{g.opp, g.intent}]
}

func solve(games []game) int {
	sum := 0
	for _, g := range games {
		sum += g.score()
	}
	return sum
}

func solveB(games []game) int {
	sum := 0
	for _, g := range games {
		g.selectResp()
		sum += g.score()
	}
	return sum
}
