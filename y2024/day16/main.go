package day16

import (
	"container/heap"

	"github.com/alphaone/advent/utils"
)

type (
	Grid  [][]rune
	Coord struct{ l, r int }
)

func (c Coord) add(other Coord) Coord {
	return Coord{c.l + other.l, c.r + other.r}
}

func (d Coord) equals(other Coord) bool {
	return d.l == other.l && d.r == other.r
}

func (d Coord) opposite() Coord {
	return Coord{-d.l, -d.r}
}

func parse(input []string) Grid {
	grid := make(Grid, 0)
	for _, line := range input {
		grid = append(grid, []rune(line))
	}

	return grid
}

func (g *Grid) findCoord(lookingfor rune) Coord {
	for l, line := range *g {
		for r, rune := range line {
			if rune == lookingfor {
				return Coord{l, r}
			}
		}
	}
	panic("not found")
}

type PQItemValue struct {
	Pos    Coord
	Dir    Coord
	Dsteps int
}

func Dijkstra(start Coord, goal Coord, field Grid) int {
	seen := map[PQItemValue]bool{}
	pq := make(utils.PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &utils.PQItem{
		Priority: 0,
		Value:    PQItemValue{start, Coord{0, 1}, 0},
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
			if dir.equals(itemValue.Dir.opposite()) {
				continue
			}

			newPos := itemValue.Pos.add(dir)
			if field[newPos.l][newPos.r] == '.' || field[newPos.l][newPos.r] == 'E' {
				addedPrio := 1
				if dir != itemValue.Dir {
					addedPrio += 1000
				}
				prio := item.Priority + addedPrio
				heap.Push(&pq, &utils.PQItem{
					Priority: prio,
					Value:    PQItemValue{newPos, dir, 1},
				})
			}
		}

	}
	return -1
}

func solvePartA(g Grid) int {
	start := g.findCoord('S')
	goal := g.findCoord('E')
	return Dijkstra(start, goal, g)
}
