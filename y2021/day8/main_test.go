package day8

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe",
	"edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc",
	"fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg",
	"fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb",
	"aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea",
	"fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb",
	"dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe",
	"bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef",
	"egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb",
	"gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce",
}

func TestParse(t *testing.T) {
	assert.Equal(t, Line{
		numbers: []string{"be", "cfbegad", "cbdgef", "fgaecd", "cgeb", "fdcge", "agebfd", "fecdb", "fabcd", "edb", "fdgacbe", "cefdb", "cefbgd", "gcbe"},
		output:  []string{"fdgacbe", "cefdb", "cefbgd", "gcbe"},
	}, parse(example)[0])
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 26, solveA(parse(example)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 303, solveA(parse(utils.LoadStrings("input.txt"))))
}

func TestDiff(t *testing.T) {
	assert.Equal(t, 6, Line{}.diff([]byte("abc"), []byte("def")))
	// diff 5 - 4
	assert.Equal(t, 3, Line{}.diff([]byte("cdfbe"), []byte("eafb")))
	// diff 2 - 4
	assert.Equal(t, 5, Line{}.diff([]byte("gcdfa"), []byte("eafb")))
}

func TestDigit(t *testing.T) {
	line := parse([]string{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"})[0]
	assert.Equal(t, 1, line.toDigit([]byte("ab")))
	assert.Equal(t, 7, line.toDigit([]byte("dab")))
	assert.Equal(t, 4, line.toDigit([]byte("eafb")))
	assert.Equal(t, 8, line.toDigit([]byte("abcdefg")))
	assert.Equal(t, 2, line.toDigit([]byte("gcdfa")))
	assert.Equal(t, 5, line.toDigit([]byte("cdfbe")))
	assert.Equal(t, 3, line.toDigit([]byte("fbcad")))
	assert.Equal(t, 9, line.toDigit([]byte("cefabd")))
	assert.Equal(t, 6, line.toDigit([]byte("cdfgeb")))
	assert.Equal(t, 0, line.toDigit([]byte("cagedb")))
}

func TestOutput(t *testing.T) {
	line := parse([]string{"acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf"})[0]
	assert.Equal(t, 5353, line.Output())
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 61229, solveB(parse(example)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 961734, solveB(parse(utils.LoadStrings("input.txt"))))
}
