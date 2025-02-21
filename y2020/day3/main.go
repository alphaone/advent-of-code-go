package day3

type treemap [][]rune

func parse(input []string) treemap {
	tm := make(treemap, len(input))
	for i, line := range input {
		tm[i] = []rune(line)
	}
	return tm
}

func (tm treemap) countTrees(right, down int) int {
	trees := 0
	width := len(tm[0])
	for x, y := 0, 0; y < len(tm); x, y = x+right, y+down {
		if tm[y][x%width] == '#' {
			trees++
		}
	}
	return trees
}

func solveA(tm treemap) int {
	return tm.countTrees(3, 1)
}

func solveB(tm treemap) int {
	return tm.countTrees(1, 1) *
		tm.countTrees(3, 1) *
		tm.countTrees(5, 1) *
		tm.countTrees(7, 1) *
		tm.countTrees(1, 2)
}
