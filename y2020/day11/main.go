package day11

type Grid [][]rune

func parseInput(input []string) Grid {
	result := make([][]rune, len(input))
	for i, line := range input {
		result[i] = []rune(line)
	}
	return result
}

func (g Grid) copy() Grid {
	result := make([][]rune, len(g))
	for i, row := range g {
		result[i] = make([]rune, len(row))
		copy(result[i], row)
	}
	return result
}

func (g Grid) String() string {
	result := ""
	for _, line := range g {
		result += string(line) + "\n"
	}
	return result
}

func (g Grid) step() Grid {
	result := g.copy()
	for l, line := range g {
		for c, seat := range line {
			if seat == '.' {
				continue
			}
			occupied := g.countOccupiedNeighbors(l, c)
			if seat == 'L' && occupied == 0 {
				result[l][c] = '#'
			} else if seat == '#' && occupied >= 4 {
				result[l][c] = 'L'
			}
		}
	}
	return result
}

func (g Grid) stepB() Grid {
	result := g.copy()
	for l, line := range g {
		for c, seat := range line {
			if seat == '.' {
				continue
			}
			occupied := g.countVisibleOccupiedNeighbors(l, c)
			if seat == 'L' && occupied == 0 {
				result[l][c] = '#'
			} else if seat == '#' && occupied >= 5 {
				result[l][c] = 'L'
			}
		}
	}
	return result
}

func (g Grid) countOccupiedNeighbors(l, c int) int {
	count := 0
	for _, dl := range []int{-1, 0, 1} {
		for _, dc := range []int{-1, 0, 1} {
			if dl == 0 && dc == 0 {
				continue
			}
			if g.at(l+dl, c+dc) == '#' {
				count++
			}
		}
	}
	return count
}

func (g Grid) countVisibleOccupiedNeighbors(l, c int) int {
	count := 0
	for _, dl := range []int{-1, 0, 1} {
		for _, dc := range []int{-1, 0, 1} {
			if dl == 0 && dc == 0 {
				continue
			}
			nl := l
			nc := c
		look:
			for {
				nl += dl
				nc += dc
				switch g.at(nl, nc) {
				case 'L':
					break look
				case '#':
					count++
					break look
				case ' ':
					break look
				}
			}
		}
	}
	return count
}

func (g Grid) at(l, c int) rune {
	if l < 0 || l >= len(g) || c < 0 || c >= len(g[l]) {
		return ' '
	}
	return g[l][c]
}

func solveA(g Grid) int {
	for {
		next := g.step()
		if g.String() == next.String() {
			break
		}
		g = next
	}
	count := 0
	for _, line := range g {
		for _, seat := range line {
			if seat == '#' {
				count++
			}
		}
	}
	return count
}

func solveB(g Grid) int {
	for {
		next := g.stepB()
		if g.String() == next.String() {
			break
		}
		g = next
	}
	count := 0
	for _, line := range g {
		for _, seat := range line {
			if seat == '#' {
				count++
			}
		}
	}
	return count
}
