package day1

import "math"

func fuel(mass int) int {
	return max(int(math.Floor(float64(mass)/3)-2), 0)
}

func fuel2(mass int) int {
	res := 0
	m := mass
	for {
		f := fuel(m)
		if f == 0 {
			break
		}
		res += f
		m = f
	}
	return res
}

func solvePartA(masses []int) int {
	res := 0
	for _, m := range masses {
		res += fuel(m)
	}
	return res
}

func solvePartB(masses []int) int {
	res := 0
	for _, m := range masses {
		res += fuel2(m)
	}
	return res
}
