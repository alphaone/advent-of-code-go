package utils

func Insert[T any](a []T, index int, value T) []T {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func Transpose[T any](input [][]T) [][]T {
	result := make([][]T, len(input[0]))
	for i := 0; i < len(result); i++ {
		for j := 0; j < len(input); j++ {
			result[i] = append(result[i], input[j][i])
		}
	}
	return result
}

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

func Join[T any](input [][]T, sep T) []T {
	result := []T{}
	for i, line := range input {
		result = append(result, line...)
		if i < len(input)-1 {
			result = append(result, sep)
		}
	}
	return result
}

func MakeSequence(start, length int, backwards bool) []int {
	a := make([]int, length)
	step := 1
	if backwards {
		step = -1
	}
	for i := range a {
		a[i] = start + step*i
	}
	return a
}

func Frequencies[T comparable](xs []T) map[T]int {
	res := make(map[T]int)
	for _, x := range xs {
		res[x] = res[x] + 1
	}
	return res
}
