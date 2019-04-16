package dec1

import (
	"github.com/alphaone/advent/y2018"
	"testing"
)

func checkA(t *testing.T, input []int, expected int) {
	actual := SumUp(input)
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestA(t *testing.T) {
	checkA(t, []int{1, 1, 1}, 3)
	checkA(t, []int{1, 1, -2}, 0)
	checkA(t, []int{-1, -2, -3}, -6)
}

func TestSolvePuzzleA(t *testing.T) {
	actual := SumUp(y2018.ReadFileAsIntSlice("input.txt"))
	expected := 520
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func checkB(t *testing.T, input []int, expected int) {
	actual := SumUpB(input)
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestB(t *testing.T) {
	checkB(t, []int{1, -1}, 0)
	checkB(t, []int{+3, +3, +4, -2, -4}, 10)
	checkB(t, []int{-6, +3, +8, +5, -6}, 5)
	checkB(t, []int{+7, +7, -2, -7, -4}, 14)
}

func TestSolvePuzzleB(t *testing.T) {
	actual := SumUpB(y2018.ReadFileAsIntSlice("input.txt"))
	expected := 394
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
