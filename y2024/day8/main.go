package day8

type (
	Coord struct{ l, c int }
	Grid  [][]rune
)

func (c Coord) add(other Coord) Coord {
	return Coord{c.l + other.l, c.c + other.c}
}

func (c Coord) subtract(other Coord) Coord {
	return Coord{c.l - other.l, c.c - other.c}
}

func (c Coord) inside(boxStart Coord, boxEnd Coord) bool {
	return c.l >= boxStart.l && c.c >= boxStart.c && c.l < boxEnd.l && c.c < boxEnd.c
}

func parseInput(input []string) Grid {
	res := make([][]rune, len(input))
	for i, line := range input {
		res[i] = []rune(line)
	}
	return res
}

func (g Grid) groupFrequencies() map[rune][]Coord {
	res := make(map[rune][]Coord)
	for l, line := range g {
		for r, rune := range line {
			if rune == '.' {
				continue
			}

			_, ok := res[rune]
			if !ok {
				res[rune] = make([]Coord, 0)
			}

			res[rune] = append(res[rune], Coord{l, r})
		}
	}
	return res
}

func pairs(nodes []Coord) [][2]Coord {
	res := make([][2]Coord, 0)
	for _, n := range nodes {
		for _, other := range nodes {
			if n == other {
				continue
			}

			res = append(res, [2]Coord{n, other})
		}
	}
	return res
}

func antinodes(nodes []Coord, limit Coord) []Coord {
	res := make([]Coord, 0)
	for _, pair := range pairs(nodes) {
		n := pair[0]
		other := pair[1]

		resonance := n.add(n.subtract(other))
		if resonance.inside(Coord{0, 0}, limit) {
			res = append(res, resonance)
		}
	}
	return res
}

func antinodes2(nodes []Coord, limit Coord) []Coord {
	res := make([]Coord, 0)
	for _, pair := range pairs(nodes) {
		n := pair[0]
		other := pair[1]

		diff := Coord{0, 0}
		for {
			resonance := n.add(diff)
			if resonance.inside(Coord{0, 0}, limit) {
				res = append(res, resonance)
			} else {
				break
			}
			diff = diff.add(n.subtract(other))
		}
	}
	return res
}

func solve(g Grid, antinodeFn func([]Coord, Coord) []Coord) int {
	limit := Coord{len(g), len(g[0])}
	res := make(map[Coord]bool)
	for _, coords := range g.groupFrequencies() {
		for _, n := range antinodeFn(coords, limit) {
			res[n] = true
		}
	}
	return len(res)
}
