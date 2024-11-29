package day6

func Advance(ageBucket []int) []int {
	givingBirth := ageBucket[0]
	for i := range ageBucket {
		if i == 0 {
			continue
		}
		ageBucket[i-1] = ageBucket[i]
	}

	ageBucket[6] += givingBirth
	ageBucket[8] = givingBirth

	return ageBucket
}

func solvePartA(input []int, days int) int {
	ageBucket := ageBucketFromInput(input)
	for i := 0; i < days; i++ {
		ageBucket = Advance(ageBucket)
	}

	fishCount := 0
	for _, fish := range ageBucket {
		fishCount += fish
	}
	return fishCount
}

func ageBucketFromInput(input []int) []int {
	ageBucket := make([]int, 9)
	for _, fish := range input {
		ageBucket[fish]++
	}
	return ageBucket
}
