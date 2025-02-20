package day1

func solveA(input []int) int {
	a, b := findTwo(input)
	return a * b
}

func solveB(input []int) int {
	a, b, c := findThree(input)
	return a * b * c
}

func findTwo(input []int) (int, int) {
	for i, a := range input {
		for _, b := range input[i+1:] {
			if a+b == 2020 {
				return a, b
			}
		}
	}
	return 0, 0
}

func findThree(input []int) (int, int, int) {
	for i, a := range input {
		for j, b := range input[i+1:] {
			for _, c := range input[j+1:] {
				if a+b+c == 2020 {
					return a, b, c
				}
			}
		}
	}
	return 0, 0, 0
}
