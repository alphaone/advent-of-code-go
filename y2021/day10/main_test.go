package day10

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"[({(<(())[]>[[{[]{<()<>>",
	"[(()[<>])]({[<{<<[]>>(",
	"{([(<{}[<>[]}>{[]{[(<()>",
	"(((({<>}<{<{<>}{[]{[]{}",
	"[[<[([]))<([[{}[[()]]]",
	"[{[{({}]{}}([{[{{{}}([]",
	"{<[[]]>}<{[{[{[]{()[[[]",
	"[<(<(<(<{}))><([]([]()",
	"<{([([[(<>()){}]>(<<{{",
	"<{([{{}}[<[[[<>{}]]]>[]]",
}

type firstIllegalCharTest struct {
	name          string
	input         string
	expected      rune
	expectedStack string
	err           error
}

func TestFirstIllegalChar(t *testing.T) {
	for _, tc := range []firstIllegalCharTest{
		{"1", "{([(<{}[<>[]}>{[]{[(<()>", '}', ">)])}", nil},
		{"2", "[[<[([]))<([[{}[[()]]]", ')', ">]]", nil},
		{"3", "[{[{({}]{}}([{[{{{}}([]", ']', "}]}]", nil},
		{"4", "[<(<(<(<{}))><([]([]()", ')', ")>)>)>]", nil},
		{"5", "<{([([[(<>()){}]>(<<{{", '>', ")])}>", nil},
	} {
		t.Run(tc.name, func(t *testing.T) {
			res, stack, err := FirstIllegalChar(tc.input)
			assert.Equal(t, tc.expected, res)
			assert.Equal(t, tc.expectedStack, string(stack))
			assert.Equal(t, tc.err, err)
		})
	}
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 26397, solveA(example))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 268845, solveA(utils.LoadStrings("input.txt")))
}

func TestToComplete(t *testing.T) {
	_, stack, _ := FirstIllegalChar("[({(<(())[]>[[{[]{<()<>>")
	assert.Equal(t, "}}]])})]", string(stack))
}

func TestAutoCompleteScore(t *testing.T) {
	assert.Equal(t, 294, autoCompleteScore("])}>"))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 288957, solveB(example))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 4038824534, solveB(utils.LoadStrings("input.txt")))
}
