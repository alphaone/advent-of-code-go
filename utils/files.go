package utils

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func LoadString(filename string) string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result += scanner.Text() + "\n"
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func LoadStrings(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var result []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		result = append(result, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return result
}

func SplitIntoSection(input []string) [][]string {
	sections := [][]string{{}}

	for _, line := range input {
		if len(line) == 0 {
			sections = append(sections, []string{})
		} else {
			sections[len(sections)-1] = append(sections[len(sections)-1], line)
		}
	}

	return sections
}

func LoadInts(filename string) []int {
	var result []int
	for _, line := range LoadStrings(filename) {
		lineInt, _ := strconv.Atoi(line)
		result = append(result, lineInt)
	}
	return result
}

func LoadIntCols(filename string) [][]int {
	var result [][]int
	for _, line := range LoadStrings(filename) {
		fields := strings.Fields(line)
		var lineInts []int
		for _, f := range fields {
			i, _ := strconv.Atoi(f)
			lineInts = append(lineInts, i)
		}
		result = append(result, lineInts)
	}
	return result
}

func AsIntField(input []string) [][]int {
	var result [][]int
	for _, line := range input {
		lineInts := []int{}
		for _, c := range line {
			i := int(c - '0')
			lineInts = append(lineInts, i)
		}
		result = append(result, lineInts)
	}
	return result
}

func AsRunes(input []string) [][]rune {
	result := make([][]rune, len(input))
	for i, line := range input {
		result[i] = []rune(line)
	}
	return result
}

func AsStrings(input [][]rune) []string {
	result := make([]string, len(input))
	for i, line := range input {
		result[i] = string(line)
	}
	return result
}
