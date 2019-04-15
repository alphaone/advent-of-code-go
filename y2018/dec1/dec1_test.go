package dec1

import (
	"testing"
)

func test(t *testing.T, input []int, expected int) {
	actual := SumUp(input)
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func Test(t *testing.T) {
	test(t, []int{1, 1, 1}, 3)
	test(t, []int{1, 1, -2}, 0)
	test(t, []int{-1, -2, -3}, -6)
}

func TestSolvePuzzleA(t *testing.T) {
	actual := SumUp(readInput())
	expected := 520
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
