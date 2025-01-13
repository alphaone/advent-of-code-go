package rotate

func RotateClockwise[T any](input [][]T) [][]T {
	result := make([][]T, len(input))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			result[i] = append(result[i], input[len(input)-j-1][i])
		}
	}
	return result
}

func RotateCounterClockwise[T any](input [][]T) [][]T {
	result := make([][]T, len(input))
	for i := 0; i < len(input); i++ {
		for j := 0; j < len(input); j++ {
			result[i] = append(result[i], input[j][len(input)-i-1])
		}
	}
	return result
}
