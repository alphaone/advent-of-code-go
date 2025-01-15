package day9

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = []string{
	"R 4",
	"U 4",
	"L 3",
	"D 1",
	"R 4",
	"D 1",
	"L 5",
	"R 2",
}

func TestParse(t *testing.T) {
	exp := []instruction{
		{coord{0, 1}, 4},
		{coord{-1, 0}, 4},
		{coord{0, -1}, 3},
		{coord{1, 0}, 1},
	}
	assert.Equal(t, exp, parse(example)[:4])
}

type moveTowardsTest struct {
	name string
	tail coord
	head coord
	exp  coord
}

func TestMoveTowards(t *testing.T) {
	for _, tc := range []moveTowardsTest{
		{name: "stay1", tail: coord{0, 0}, head: coord{0, 0}, exp: coord{0, 0}},
		{name: "stay2", tail: coord{0, 1}, head: coord{0, 0}, exp: coord{0, 1}},
		{name: "move -y", tail: coord{0, 2}, head: coord{0, 0}, exp: coord{0, 1}},
		{name: "move -y 2", tail: coord{0, 5}, head: coord{0, 3}, exp: coord{0, 4}},
		{name: "move -y 0", tail: coord{0, 1}, head: coord{0, -1}, exp: coord{0, 0}},
		{name: "move y", tail: coord{0, -2}, head: coord{0, 0}, exp: coord{0, -1}},
		{name: "move y 2", tail: coord{0, -5}, head: coord{0, -3}, exp: coord{0, -4}},
		{name: "move y 0", tail: coord{0, -1}, head: coord{0, 1}, exp: coord{0, 0}},
		{name: "move -x", tail: coord{2, 0}, head: coord{0, 0}, exp: coord{1, 0}},
		{name: "move -x 2", tail: coord{5, 0}, head: coord{3, 0}, exp: coord{4, 0}},
		{name: "move -x 0", tail: coord{1, 0}, head: coord{-1, 0}, exp: coord{0, 0}},
		{name: "move y", tail: coord{-2, 0}, head: coord{0, 0}, exp: coord{-1, 0}},
		{name: "move y 2", tail: coord{-5, 0}, head: coord{-3, 0}, exp: coord{-4, 0}},
		{name: "move y 0", tail: coord{-1, 0}, head: coord{1, 0}, exp: coord{0, 0}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			assert.Equal(t, tc.exp, tc.tail.moveTowards(tc.head))
		})
	}
}

func TestExampleA(t *testing.T) {
	assert.Equal(t, 13, solveA(example))
}

func TestA(t *testing.T) {
	assert.Equal(t, 5960, solveA(utils.LoadStrings("input.txt")))
}

var exampleB = []string{
	"R 5",
	"U 8",
	"L 8",
	"D 3",
	"R 17",
	"D 10",
	"L 25",
	"U 20",
}

func TestExampleB(t *testing.T) {
	assert.Equal(t, 36, solveB(exampleB))
}

func TestB(t *testing.T) {
	assert.Equal(t, 2327, solveB(utils.LoadStrings("input.txt")))
}
