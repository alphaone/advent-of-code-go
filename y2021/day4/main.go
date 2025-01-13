package day4

import (
	"slices"
	"strconv"
	"strings"

	"github.com/alphaone/advent/utils"
	"github.com/alphaone/advent/utils/sliceutils"
)

type Game struct {
	Numbers []int
	Boards  []Board
}
type Board struct {
	Numbers [][]int
	Crosses []Coord
}
type Coord struct {
	Line, Col int
}

func parseInput(input []string) Game {
	result := Game{Numbers: []int{}}

	res := strings.Split(input[0], ",")
	for _, r := range res {
		i, _ := strconv.Atoi(r)
		result.Numbers = append(result.Numbers, i)
	}

	boardStrings := strings.Split(strings.Join(input[2:], "\n"), "\n\n")
	for _, boardString := range boardStrings {
		board := Board{
			Numbers: [][]int{},
			Crosses: []Coord{},
		}
		for _, line := range strings.Split(boardString, "\n") {
			board.Numbers = append(board.Numbers, utils.ParseNumbers(line))
		}
		result.Boards = append(result.Boards, board)
	}

	return result
}

func (b Board) Completed() bool {
	for l, line := range b.Numbers {
		crossed := 0
		for c := range line {
			if slices.Contains(b.Crosses, Coord{Line: l, Col: c}) {
				crossed++
			}
		}
		if crossed == 5 {
			return true
		}
	}
	for l, line := range sliceutils.Transpose(b.Numbers) {
		crossed := 0
		for c := range line {
			if slices.Contains(b.Crosses, Coord{Line: c, Col: l}) {
				crossed++
			}
		}
		if crossed == 5 {
			return true
		}
	}
	return false
}

func (b Board) Find(n int) []Coord {
	var result []Coord
	for l, line := range b.Numbers {
		for c, num := range line {
			if num == n {
				result = append(result, Coord{Line: l, Col: c})
			}
		}
	}
	return result
}

func (g *Game) Play() int {
	if len(g.Numbers) == 0 {
		return -1
	}

	n := g.Numbers[0]
	g.Numbers = g.Numbers[1:]

	for i, board := range g.Boards {
		coords := board.Find(n)
		board.Crosses = append(board.Crosses, coords...)
		g.Boards[i] = board
	}
	return n
}

func (b Board) Score(n int) int {
	result := 0
	for l, line := range b.Numbers {
		for c, num := range line {
			if !slices.Contains(b.Crosses, Coord{Line: l, Col: c}) {
				result += num
			}
		}
	}
	return result * n
}

func solvePartA(input []string) int {
	game := parseInput(input)
	for {
		n := game.Play()
		if n == -1 {
			return -1
		}

		for _, board := range game.Boards {
			if board.Completed() {
				return board.Score(n)
			}
		}
	}
}

func solvePartB(input []string) int {
	game := parseInput(input)
	for {
		n := game.Play()
		if n == -1 {
			return -1
		}
		keepBoards := []Board{}
		for _, board := range game.Boards {
			if !board.Completed() {
				keepBoards = append(keepBoards, board)
			}
		}
		if len(keepBoards) == 0 {
			return game.Boards[0].Score(n)
		}
		game.Boards = keepBoards

	}
}
