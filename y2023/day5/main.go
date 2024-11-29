package day5

import (
	"fmt"
	"regexp"
	"slices"
	"strconv"
	"strings"
	"sync"
)

type Garden struct {
	Seeds             []int
	SeedRanges        []SeedRange
	ConverionPipeline []ConversionMap
}

type ConversionMap struct {
	Name   string
	Ranges []LocationRange
}

type LocationRange struct {
	DescStart int
	SrcStart  int
	Count     int
}

type SeedRange struct {
	Start int
	Count int
}

func parseMap(input []string) ConversionMap {
	re := regexp.MustCompile(`(\d+) (\d+) (\d+)`)
	ranges := []LocationRange{}

	for _, line := range input[1:] {
		res := re.FindStringSubmatch(line)
		destStart, _ := strconv.Atoi(res[1])
		srcStart, _ := strconv.Atoi(res[2])
		count, _ := strconv.Atoi(res[3])

		ranges = append(ranges, LocationRange{
			DescStart: destStart,
			SrcStart:  srcStart,
			Count:     count,
		})
	}

	return ConversionMap{
		Name:   input[0],
		Ranges: ranges,
	}
}

func parseGarden(input []string) Garden {
	parts := strings.Split(strings.Join(input, "\n"), "\n\n")

	re := regexp.MustCompile(`seeds: (.+)`)
	r := re.FindStringSubmatch(parts[0])

	var seeds []int
	for _, seed := range strings.Split(r[1], " ") {
		s, _ := strconv.Atoi(seed)
		seeds = append(seeds, s)
	}
	var seedRanges []SeedRange
	for i := 0; i < len(seeds); i += 2 {
		seedRanges = append(seedRanges, SeedRange{
			Start: seeds[i],
			Count: seeds[i+1],
		})
	}

	var conversionPipeline []ConversionMap
	for _, part := range parts[1:] {
		conversionPipeline = append(conversionPipeline, parseMap(strings.Split(part, "\n")))
	}

	return Garden{
		Seeds:             seeds,
		SeedRanges:        seedRanges,
		ConverionPipeline: conversionPipeline,
	}
}

func solvePartA(input []string) int {
	garden := parseGarden(input)
	var res []int
	for _, seed := range garden.Seeds {
		for _, cmap := range garden.ConverionPipeline {
			seed = cmap.Convert(seed)
		}
		res = append(res, seed)
	}

	return slices.Min(res)
}

func (r LocationRange) Contains(n int) bool {
	return n >= r.SrcStart && n < r.SrcStart+r.Count
}

func (r LocationRange) Convert(n int) int {
	if !r.Contains(n) {
		return n
	}
	return n + r.DescStart - r.SrcStart
}

func (c ConversionMap) Convert(n int) int {
	for _, rng := range c.Ranges {
		if rng.Contains(n) {
			return rng.Convert(n)
		}
	}
	return n
}

func solvePartB(input []string) int {
	garden := parseGarden(input)
	ch := make(chan int, len(garden.SeedRanges))
	var wg sync.WaitGroup

	for _, seedRange := range garden.SeedRanges {
		wg.Add(1)
		go findMinForSeedRange(seedRange, garden.ConverionPipeline, ch, &wg)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	var res []int
	for num := range ch {
		res = append(res, num)
	}
	return slices.Min(res)
}

func findMinForSeedRange(seedRange SeedRange, maps []ConversionMap, ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Print(`.`)
	var res2 []int
	for i := 0; i < seedRange.Count; i++ {
		seed := seedRange.Start + i
		for _, cmap := range maps {
			seed = cmap.Convert(seed)
		}
		res2 = append(res2, seed)
	}
	fmt.Print(`#`)
	ch <- slices.Min(res2)
}
