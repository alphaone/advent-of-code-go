package day12

import (
	"reflect"
	"regexp"
	"strings"

	"github.com/alphaone/advent/utils"
)

type Report struct {
	Diagnose string
	Checksum []int
}

// ???.### 1,1,3
func parseLine(s string) Report {
	res := strings.Split(s, " ")

	return Report{
		Diagnose: res[0],
		Checksum: utils.ParseNumbers(res[1]),
	}
}

func isValid(report Report) bool {
	return reflect.DeepEqual(checksum(report.Diagnose), report.Checksum)
}

func checksum(diagnose string) []int {
	re := regexp.MustCompile(`\#+|\.+`)
	res := re.FindAllStringSubmatch(diagnose, -1)
	var actualChecksum []int
	for _, match := range res {
		if match[0][0] == '#' {
			actualChecksum = append(actualChecksum, len(match[0]))
		}
	}
	return actualChecksum
}

func check(partialChecksum, fullChecksum []int) (bool, []int) {
	for i, n := range partialChecksum {
		if fullChecksum[i] != n {
			return false, []int{}
		}
	}
	return true, fullChecksum[len(partialChecksum):]
}

func countOptions(lvl int, report Report) int {
	// fmt.Println("lvl:", lvl, "counting:", report.Diagnose, report.Checksum)
	idx := strings.Index(report.Diagnose, "?")
	if idx == -1 {
		if isValid(report) {
			return 1
		}
		return 0
	}

	res := 0

	head := report.Diagnose[:idx]
	tail := report.Diagnose[idx+1:]

	newHead := head + "#"
	if success, tailChecksum := check(checksum(newHead), report.Checksum); success {
		// fmt.Println("lvl:", lvl, "Success for #:", newHead, tailChecksum)
		x := countOptions(lvl+1, Report{Diagnose: tail, Checksum: tailChecksum})
		// fmt.Println("lvl:", lvl, "x:", x)
		res += x
	} else {
		// fmt.Println("lvl:", lvl, "Failed for #:", newHead, tailChecksum)
	}

	newHead2 := head + "."
	if success, tailChecksum := check(checksum(newHead2), report.Checksum); success {
		// fmt.Println("lvl:", lvl, "Success for .:", newHead2, tailChecksum)
		x := countOptions(lvl+1, Report{Diagnose: tail, Checksum: tailChecksum})
		// fmt.Println("lvl:", lvl, "x:", x)
		res += x
	} else {
		// fmt.Println("lvl:", lvl, "Failed for .:", newHead, tailChecksum)
	}

	// fmt.Println("lvl:", lvl, " res:", res)
	return res
}

func solvePartA(input []string) int {
	res := 0
	for _, line := range input {
		report := parseLine(line)
		res += countOptions(0, report)
	}
	return res
}
