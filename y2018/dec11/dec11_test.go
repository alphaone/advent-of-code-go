package dec11

import "testing"

func TestHundredsDigit(t *testing.T) {
	actual := hundredsDigit(1234)
	expected := 2
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}
func TestHundredsDigit2(t *testing.T) {
	actual := hundredsDigit(949)
	expected := 9
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func checkA(t *testing.T, serialNumber int, cell Coords, expected int) {
	actual := powerLevel(serialNumber, cell)
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestA(t *testing.T) {
	checkA(t, 8, Coords{3, 5}, 4)
	checkA(t, 57, Coords{122, 79}, -5)
	checkA(t, 39, Coords{217, 196}, 0)
	checkA(t, 71, Coords{101, 153}, 4)
}

func TestSquarePowerLevel(t *testing.T) {
	actual := squarePowerLevel(18, 3, Coords{33, 45})
	expected := 29
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func checkBestSquareOfSize(t *testing.T, serialNumber, size int, expected Coords) {
	actual, _ := bestSquareOfSize(serialNumber, size)
	if actual != expected {
		t.Errorf("Expected %v but got %v", expected, actual)
	}
}

func TestBestSquareOfSize(t *testing.T) {
	checkBestSquareOfSize(t, 18, 3, Coords{33, 45})
	checkBestSquareOfSize(t, 18, 16, Coords{90, 269})
}

func TestSolveA(t *testing.T) {
	checkBestSquareOfSize(t, 4151, 3, Coords{20, 46})
}

func checkBestSquare(t *testing.T, serialNumber int, expectedCoords Coords, expectedSize int) {
	actualCoords, actualSize := bestSquare(serialNumber)
	if actualCoords != expectedCoords {
		t.Errorf("Expected %v but got %v", expectedCoords, actualCoords)
	}
	if actualSize != expectedSize {
		t.Errorf("Expected %v but got %v", expectedSize, actualSize)
	}
}

func TestBestSquare(t *testing.T) {
	checkBestSquare(t, 18, Coords{90, 269}, 16)
}

func TestSolveB(t *testing.T) {
	checkBestSquare(t, 4151, Coords{231, 65}, 14)
}
