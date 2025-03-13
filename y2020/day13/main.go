package day13

import (
	"strconv"
	"strings"
)

func parse(input string) (int, []int) {
	parts := strings.Split(input, "\n")
	ts, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	buses := []int{}
	for _, bus := range strings.Split(parts[1], ",") {
		if bus == "x" {
			continue
		}
		n, err := strconv.Atoi(bus)
		if err != nil {
			panic(err)
		}
		buses = append(buses, n)
	}

	return ts, buses
}

func parseB(input string) []bus {
	parts := strings.Split(input, "\n")
	buses := []bus{}
	for i, nrstr := range strings.Split(parts[1], ",") {
		if nrstr == "x" {
			continue
		}
		nr, err := strconv.Atoi(nrstr)
		if err != nil {
			panic(err)
		}
		buses = append(buses, bus{nr, i})
	}
	return buses
}

func solveA(ts int, buses []int) int {
	for i := ts; ; i++ {
		for _, bus := range buses {
			if i%bus == 0 {
				return bus * (i - ts)
			}
		}
	}
}

type bus struct {
	nr, offset int
}

// ; BusA(7| 0)
// ; BusB(5| 1)
// ;
// ; 0---------1---------1---------2---------3---------4
// ; 0----5----0----5----0----5----0----5----0----5----0
// ; A......A......A......A......A......A......A......A
// ; B....B....B....B....B....B....B....B....B....B....B
// ;               ^^                                 ^^
// ; combined BusAB(35| 14)
func combineBus(busA, busB bus) bus {
	offset := busA.offset
	for {
		offset += busA.nr
		if (offset+busB.offset)%busB.nr == 0 {
			return bus{busA.nr * busB.nr, offset}
		}
	}
}

func solveB(buses []bus) int {
	combined := buses[0]
	for _, bus := range buses[1:] {
		combined = combineBus(combined, bus)
	}

	return combined.offset
}
