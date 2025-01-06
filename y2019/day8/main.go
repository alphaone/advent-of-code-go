package day8

import (
	"fmt"
	"strings"

	"github.com/alphaone/advent/utils"
)

func parseInput(input string, width, height int) [][]int {
	layers := make([][]int, 0)
	layerSize := width * height

	for range len(input) / layerSize {
		layer := make([]int, 0)
		layers = append(layers, layer)
	}

	for i, c := range strings.TrimSpace(input) {
		layerIdx := i / layerSize
		layers[layerIdx] = append(layers[layerIdx], int(c-'0'))
	}

	return layers
}

func solveA(layers [][]int) int {
	freqsWithFewestZeros := map[int]int{0: 999}

	for _, layer := range layers {
		freqs := utils.Frequencies(layer)
		if freqs[0] < freqsWithFewestZeros[0] {
			freqsWithFewestZeros = freqs
		}
	}

	return freqsWithFewestZeros[1] * freqsWithFewestZeros[2]
}

func stack(layers [][]int) []int {
	res := layers[0]

	for _, layer := range layers[1:] {
		for i, p := range res {
			if p != 2 {
				continue
			}

			res[i] = layer[i]
		}
	}

	return res
}

func solveB(layers [][]int, width int) string {
	strs := make([]string, 0)

	for i, v := range stack(layers) {
		switch v {
		case 0:
			strs = append(strs, " ")
		case 1:
			strs = append(strs, "X")
		}
		if i%width == width-1 {
			strs = append(strs, "\n")
		}

	}
	fmt.Println(strings.Join(strs, ""))

	return strings.Join(strs, "")
}
