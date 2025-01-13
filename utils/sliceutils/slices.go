package sliceutils

import "slices"

func Insert[T any](a []T, index int, value T) []T {
	if len(a) == index { // nil or empty slice or after last element
		return append(a, value)
	}
	a = append(a[:index+1], a[index:]...) // index < len(a)
	a[index] = value
	return a
}

func RemoveIndex[T any](slice []T, index int) []T {
	ret := make([]T, 0)
	ret = append(ret, slice[:index]...)
	return append(ret, slice[index+1:]...)
}

func LastIndexFunc[S ~[]E, E any](s S, f func(E) bool) int {
	for i := len(s) - 1; i >= 0; i-- {
		if f(s[i]) {
			return i
		}
	}
	return -1
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

func IsUnique[T comparable](xs []T) bool {
	m := make(map[T]bool)
	for _, x := range xs {
		if _, ok := m[x]; ok {
			return false
		}
		m[x] = true
	}
	return true
}

func Intersection[T comparable](s1, s2 []T) []T {
	res := []T{}
	hash := make(map[T]bool)
	for _, e := range s1 {
		hash[e] = true
	}
	for _, e := range s2 {
		if hash[e] && !slices.Contains(res, e) {
			res = append(res, e)
		}
	}
	return res
}

func All[T any](xs []T, f func(T) bool) bool {
	for _, x := range xs {
		if !f(x) {
			return false
		}
	}
	return true
}

// partition the given slice into slices of given size
//
// Deprecated: Use `slices.Chunk` instead.
func ChunkBy[T any](items []T, chunkSize int) (chunks [][]T) {
	for chunkSize < len(items) {
		items, chunks = items[chunkSize:], append(chunks, items[0:chunkSize:chunkSize])
	}
	return append(chunks, items)
}
