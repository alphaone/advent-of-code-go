package day17

import (
	"container/heap"

	"github.com/alphaone/advent/utils"
)

type Coord struct {
	Line, Col int
}
type Dir struct {
	Line, Col int
}

func (d Dir) Equals(other Dir) bool {
	return d.Line == other.Line && d.Col == other.Col
}

func (d Dir) Opposite() Dir {
	return Dir{-d.Line, -d.Col}
}

type PQItemValue struct {
	Pos    Coord
	Dir    Dir
	Dsteps int
}

type Field [][]int

func (f Field) Includes(c Coord) bool {
	return c.Line >= 0 && c.Line < len(f) && c.Col >= 0 && c.Col < len(f[0])
}

func Dijkstra(start Coord, goal Coord, field Field, minStep, maxStep int) int {
	seen := map[PQItemValue]bool{}
	pq := make(utils.PriorityQueue, 0)
	heap.Init(&pq)

	heap.Push(&pq, &utils.PQItem{
		Priority: 0,
		Value:    PQItemValue{start, Dir{0, 0}, 0},
	})

	for len(pq) > 0 {
		item := heap.Pop(&pq).(*utils.PQItem)
		itemValue := item.Value.(PQItemValue)
		if itemValue.Pos == goal && itemValue.Dsteps >= minStep {
			return item.Priority
		}
		if seen[itemValue] {
			continue
		}
		seen[itemValue] = true
		// fmt.Println("popped", item.Priority, itemValue.Pos, itemValue.Dir, itemValue.Dsteps)

		if itemValue.Dsteps < maxStep && !itemValue.Dir.Equals(Dir{0, 0}) {
			newPos := Coord{itemValue.Pos.Line + itemValue.Dir.Line, itemValue.Pos.Col + itemValue.Dir.Col}
			if field.Includes(newPos) {
				prio := item.Priority + field[newPos.Line][newPos.Col]
				// fmt.Println("pushing", newPos, itemValue.Dir, itemValue.Dsteps+1)
				heap.Push(&pq, &utils.PQItem{
					Priority: prio,
					Value:    PQItemValue{newPos, itemValue.Dir, itemValue.Dsteps + 1},
				})
			}
		}

		for _, dir := range []Dir{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
			if dir.Equals(itemValue.Dir) || dir.Equals(itemValue.Dir.Opposite()) {
				continue
			}
			if itemValue.Dsteps >= minStep || itemValue.Dir.Equals(Dir{0, 0}) {
				newPos := Coord{itemValue.Pos.Line + dir.Line, itemValue.Pos.Col + dir.Col}
				if field.Includes(newPos) {
					prio := item.Priority + field[newPos.Line][newPos.Col]
					// fmt.Println("pushing", newPos, dir, 1)
					heap.Push(&pq, &utils.PQItem{
						Priority: prio,
						Value:    PQItemValue{newPos, dir, 1},
					})
				}
			}
		}

	}
	return -1
}

func solvePartA(field Field) int {
	return Dijkstra(Coord{0, 0}, Coord{len(field) - 1, len(field[0]) - 1}, field, 1, 3)
}

func solvePartB(field Field) int {
	return Dijkstra(Coord{0, 0}, Coord{len(field) - 1, len(field[0]) - 1}, field, 4, 10)
}
