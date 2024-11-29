package day4

import (
	"regexp"
	"strconv"

	"github.com/alphaone/advent/utils"
)

type Card struct {
	Id            int
	Winning       []int
	Played        []int
	instanceCount int
}

// Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
var reCard = regexp.MustCompile(`Card\s+(\d+): (.+) \| (.+)`)

func parseLine(line string) Card {
	res := reCard.FindStringSubmatch(line)
	id, _ := strconv.Atoi(res[1])

	winning := utils.ParseNumbers(res[2])
	played := utils.ParseNumbers(res[3])

	return Card{
		Id:            id,
		Winning:       winning,
		Played:        played,
		instanceCount: 1,
	}
}

func (c Card) Matching() []int {
	var matching []int
	for _, w := range c.Winning {
		for _, p := range c.Played {
			if w == p {
				matching = append(matching, w)
			}
		}
	}
	return matching
}

func (c Card) Score() int {
	matching := c.Matching()
	if len(matching) == 0 {
		return 0
	}
	return utils.PowInts(2, len(matching)-1)
}

func solvePartA(input []string) int {
	var cards []Card
	for _, line := range input {
		cards = append(cards, parseLine(line))
	}
	var score int
	for _, card := range cards {
		score += card.Score()
	}
	return score
}

func solvePartB(input []string) int {
	var cards []Card
	for _, line := range input {
		cards = append(cards, parseLine(line))
	}

	var result int
	for idx, card := range cards {
		noOfMatching := len(card.Matching())

		for i := 0; i < noOfMatching; i++ {
			cards[idx+i+1].instanceCount += card.instanceCount
		}
		result += card.instanceCount
	}
	return result
}
