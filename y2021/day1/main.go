package day1

func solvePartA(input []int) int {
	var result int
	for i, depth := range input {
		if i == 0 {
			continue
		}

		oldDepth := input[i-1]
		if depth > oldDepth {
			result += 1
		}
	}
	return result
}

func solvePartB(input []int) int {
	var windowSums []int
	for i := range input {
		if i < 2 {
			continue
		}
		window := input[i-2 : i+1]
		windowSums = append(windowSums, window[0]+window[1]+window[2])
	}
	return solvePartA(windowSums)
}
