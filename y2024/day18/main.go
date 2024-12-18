package day18

import (
	"container/heap"
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils"
)

type (
	Coord struct{ l, c int }
	Grid  [][]rune
)

func (c Coord) add(other Coord) Coord {
	return Coord{c.l + other.l, c.c + other.c}
}

func (c Coord) insideOf(upperLeft, lowerRight Coord) bool {
	return upperLeft.l <= c.l && c.l <= lowerRight.l &&
		upperLeft.c <= c.c && c.c <= lowerRight.c
}

func (g Grid) String() string {
	res := ""
	for _, line := range g {
		res += string(line) + "\n"
	}
	return res
}

func parse(input []string) []Coord {
	res := make([]Coord, 0)
	for _, line := range input {
		parts := strings.Split(strings.TrimSpace(line), ",")
		c, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		l, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		res = append(res, Coord{l, c})
	}
	return res
}

func solvePartA(coords []Coord, goal Coord) int {
	grid := make([][]rune, goal.l+1)
	for i := range goal.l + 1 {
		grid[i] = make([]rune, goal.c+1)
		for j := range goal.c + 1 {
			grid[i][j] = '.'
		}
	}

	for _, c := range coords {
		grid[c.l][c.c] = '#'
	}
	return Dijkstra(Coord{0, 0}, goal, grid)
}

func solvePartB(coords []Coord, goal Coord) Coord {
	grid := make([][]rune, goal.l+1)
	for i := range goal.l + 1 {
		grid[i] = make([]rune, goal.c+1)
		for j := range goal.c + 1 {
			grid[i][j] = '.'
		}
	}

	for _, c := range coords {
		grid[c.l][c.c] = '#'

		if Dijkstra(Coord{0, 0}, goal, grid) == -1 {
			return c
		}
	}
	return Coord{-1, -1}
}

type PQItemValue struct {
	Pos Coord
}

func Dijkstra(start Coord, goal Coord, grid Grid) int {
	seen := map[PQItemValue]bool{}
	pq := make(utils.PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &utils.PQItem{
		Priority: 0,
		Value:    PQItemValue{start},
	})

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*utils.PQItem)
		itemValue := item.Value.(PQItemValue)
		if itemValue.Pos == goal {
			return item.Priority
		}
		if seen[itemValue] {
			continue
		}
		seen[itemValue] = true

		for _, dir := range []Coord{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			newPos := itemValue.Pos.add(dir)
			if !newPos.insideOf(start, goal) {
				continue
			}
			if grid[newPos.l][newPos.c] == '.' {
				prio := item.Priority + 1
				heap.Push(&pq, &utils.PQItem{
					Priority: prio,
					Value:    PQItemValue{newPos},
				})
			}
		}

	}
	return -1
}
