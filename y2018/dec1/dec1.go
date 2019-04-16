package dec1

func SumUp(freqs []int) int {
	res := 0

	for _, v := range freqs {
		res += v
	}

	return res
}

func SumUpB(freqs []int) int {
	res := 0
	seen := map[int]bool{0: true}

	for {
		for _, v := range freqs {
			res += v
			if _, ok := seen[res]; !ok {
				seen[res] = true
			} else {
				return res
			}
		}
	}
}
