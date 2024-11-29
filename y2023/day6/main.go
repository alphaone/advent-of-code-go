package day6

type Race struct {
	time     int
	distance int
}

func outcomes(raceTime int) []int {
	var result []int
	for i := 0; i <= raceTime; i++ {
		result = append(result, accelaration(i, raceTime))
	}
	return result
}

func accelaration(accelarationTime int, raceTime int) int {
	speed := accelarationTime
	return (raceTime - accelarationTime) * speed
}

func solve(input []Race) int {
	result := 1
	for _, race := range input {
		wins := 0
		for _, out := range outcomes(race.time) {
			if out > race.distance {
				wins++
			}
		}
		result *= wins
	}
	return result
}
