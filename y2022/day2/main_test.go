package day2

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{"A Y", "B X", "C Z"}

func TestParse(t *testing.T) {
	assert.Equal(t,
		[]game{
			{opp: "A", resp: "Y", intent: "Y"},
			{opp: "B", resp: "X", intent: "X"},
			{opp: "C", resp: "Z", intent: "Z"},
		}, parse(example))
}

type scoreTestCases struct {
	name string
	game game
	exp  int
}

func TestScoreA(t *testing.T) {
	for _, tc := range []scoreTestCases{
		{"Rock vs my Paper", game{opp: "A", resp: "Y"}, 8},
		{"Paper vs my Rock", game{opp: "B", resp: "X"}, 1},
		{"Scissors vs my Scissors", game{opp: "C", resp: "Z"}, 6},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp, tc.game.score())
		})
	}
}

type selectRespTestCases struct {
	name string
	game game
	resp string
}

func TestSelectResp(t *testing.T) {
	for _, tc := range []selectRespTestCases{
		{"Rock vs my wish to draw", game{opp: "A", intent: "Y"}, "X"},
		{"Paper vs my with to lose", game{opp: "B", intent: "X"}, "X"},
		{"Scissors vs my wish to win", game{opp: "C", intent: "Z"}, "X"},
	} {
		t.Run(tc.name, func(t *testing.T) {
			tc.game.selectResp()
			assert.Equal(t, tc.resp, tc.game.resp)
		})
	}
}

func TestSolveExample(t *testing.T) {
	assert.Equal(t, 15, solve(parse(example)))
}

func TestSolve(t *testing.T) {
	assert.Equal(t, 10718, solve(parse(utils.LoadStrings("input.txt"))))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 12, solveB(parse(example)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 14652, solveB(parse(utils.LoadStrings("input.txt"))))
}
