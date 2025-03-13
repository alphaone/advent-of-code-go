package day13

import (
	"strconv"
	"strings"
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var example = `939
7,13,x,x,59,x,31,19`

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 295, solveA(parse(example)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 2845, solveA(parse(utils.LoadString("input.txt"))))
}

func parse(input string) (int, []int) {
	parts := strings.Split(input, "\n")
	ts, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	buses := []int{}
	for _, bus := range strings.Split(parts[1], ",") {
		if bus == "x" {
			continue
		}
		n, err := strconv.Atoi(bus)
		if err != nil {
			panic(err)
		}
		buses = append(buses, n)
	}

	return ts, buses
}

func solveA(ts int, buses []int) int {
	for i := ts; ; i++ {
		for _, bus := range buses {
			if i%bus == 0 {
				return bus * (i - ts)
			}
		}
	}
}
