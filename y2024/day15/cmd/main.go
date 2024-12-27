package main

import (
	"fmt"

	"github.com/alphaone/advent/y2024/day15"
	"github.com/eiannone/keyboard"
)

var largeExample = `##########
#..O..O.O#
#......O.#
#.OO..O.O#
#..O@..O.#
#O#..O...#
#O..O..O.#
#.OO.O.OO#
#....O...#
##########

<vv>^<v^>v>^vv^v>v<>v^v<v<^vv<<<^><<><>>v<vvv<>^v^>^<<<><<v<<<v^vv^v>^
vvv<<^>^v^^><<>>><>^<<><^vv^^<>vvv<>><^^v>^>vv<>v<<<<v<^v>^<^^>>>^<v<v
><>vv>v^v^<>><>>>><^^>vv>v<^^^>>v^v^<^^>v^^>v^<^v>v<>>v^v^<v>v^^<^^vv<
<<v<^>>^^^^>>>v^<>vvv^><v<<<>^^^vv^<vvv>^>v<^^^^v<>^>vvvv><>>v^<<^^^^^
^><^><>>><>^^<<^^v>>><^<v>^<vv>>v>>>^v><>^v><<<<v>>v<v<v>vvv>^<><<>^><
^>><>^v<><^vvv<^^<><v<<<<<><^v<<<><<<^^<v<^^^><^>>^<v^><<<^>>^v<v^v<v^
>^>>^v>vv>^<<^v<>><<><<v<<v><>v<^vv<<<>^^v^>^^>>><<^v>>v^v><^^>>^<>vv^
<><^^>^^^<><vvvvv^v<v<<>^v<v>v<<^><<><<><<<^^<<<^<<>><<><^^^>^^<>^>v<>
^^>vv<^v^v<vv>^<><v<^v>^^^>>>^^vvv^>vvv<>>>^<^>>>>>^<<^v>^vvv<>^<><<v>
v^^>>><<^^<>>^v^<v^vv<>v^<<>^<^v^v><^<<<><<^<v><v<>vv>>v><v^<vv<>v^<<^
`

var keyboardDirections = map[rune]day15.Coord{
	'h': {L: 0, C: -1},
	'j': {L: 1, C: 0},
	'k': {L: -1, C: 0},
	'l': {L: 0, C: 1},
}

func main() {
	g, movements := day15.ParseB(largeExample)

	pos := g.FindCoord('@')[0]
	g[pos.L][pos.C] = '.'

	for i, m := range movements {
		fmt.Print("\033[H\033[2J")
		fmt.Printf("move %d of %d: %s\n", i, len(movements), string(m))
		fmt.Print("\n", g.Visual(pos))
		char, key, err := keyboard.GetSingleKey()
		if err != nil {
			panic(err)
		}
		if key == keyboard.KeyCtrlC {
			panic("Ctrl+C pressed")
		}

		dir, ok := keyboardDirections[char]
		fmt.Println(key, string(char), ok, dir)
		if ok {
			pos = g.Move(pos, dir)
		} else {
			pos = g.Move(pos, day15.Directions[m])
		}

	}
}
