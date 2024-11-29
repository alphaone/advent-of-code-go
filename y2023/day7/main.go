package day7

import (
	"sort"
	"strconv"
	"strings"
)

var (
	cardValues          = []rune("23456789TJQKA")
	cardValuesWithJoker = []rune("J23456789TQKA")
)

type Game struct {
	Hand Hand
	Bid  int
}

type Hand []rune

func parseLine(line string) Game {
	x := strings.Split(line, " ")
	bid, _ := strconv.Atoi(x[1])

	return Game{
		Bid:  bid,
		Hand: []rune(x[0]),
	}
}

// Define a collection type that implements sort.Interface
type GamesByValue []Game

func (u GamesByValue) Len() int {
	return len(u)
}

func (u GamesByValue) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u GamesByValue) Less(i, j int) bool {
	return u[i].Hand.LessThan(u[j].Hand, false)
}

// Define a collection type that implements sort.Interface
type GamesByValueWithJoker []Game

func (u GamesByValueWithJoker) Len() int {
	return len(u)
}

func (u GamesByValueWithJoker) Swap(i, j int) {
	u[i], u[j] = u[j], u[i]
}

func (u GamesByValueWithJoker) Less(i, j int) bool {
	return u[i].Hand.LessThan(u[j].Hand, true)
}

func (h Hand) LessThan(other Hand, enableJokers bool) bool {
	if h.Type(enableJokers) == other.Type(enableJokers) {
		for i := range h {
			if h[i] != other[i] {
				if enableJokers {
					return strings.Index(string(cardValuesWithJoker), string(h[i])) < strings.Index(string(cardValuesWithJoker), string(other[i]))
				}
				return strings.Index(string(cardValues), string(h[i])) < strings.Index(string(cardValues), string(other[i]))
			}
		}
	}
	return h.Type(enableJokers) < other.Type(enableJokers)
}

func (h Hand) String() string {
	return string(h)
}

func (h Hand) Type(enableJokers bool) int {
	freq := h.Frequencies(enableJokers)
	switch freq[0].Count {
	case 5:
		return 6
	case 4:
		return 5
	case 3:
		if freq[1].Count == 2 {
			return 4
		}
		return 3
	case 2:
		if freq[1].Count == 2 {
			return 2
		}
		return 1
	default:
		return 0
	}
}

type Frequency struct {
	Rune  rune
	Count int
}

func (h Hand) Frequencies(enableJokers bool) []Frequency {
	freq := make(map[rune]int)
	for _, c := range h {
		freq[c]++
	}

	jokerCount := freq['J']
	if enableJokers {
		delete(freq, 'J')
	}

	if len(freq) == 0 {
		return []Frequency{{'J', 5}}
	}

	res := sortMap(freq)

	if enableJokers {
		res[0].Count += jokerCount
	}
	return res
}

// sorting map by value is rediculously verbose in go
func sortMap(freq map[rune]int) []Frequency {
	keys := make([]rune, 0, len(freq))
	for k := range freq {
		keys = append(keys, k)
	}

	sort.SliceStable(keys, func(i, j int) bool {
		return freq[keys[i]] > freq[keys[j]]
	})

	var result []Frequency
	for _, k := range keys {
		result = append(result, Frequency{k, freq[k]})
	}
	return result
}

func solvePartA(input []string) int {
	var games []Game
	for _, line := range input {
		games = append(games, parseLine(line))
	}
	sort.Sort(GamesByValue(games))
	result := 0
	for idx, g := range games {
		result += g.Bid * (idx + 1)
	}
	return result
}

func solvePartB(input []string) int {
	var games []Game
	for _, line := range input {
		games = append(games, parseLine(line))
	}
	sort.Sort(GamesByValueWithJoker(games))
	result := 0
	for idx, g := range games {
		result += g.Bid * (idx + 1)
	}
	return result
}
