package day4

import (
	"testing"

	"github.com/alphaone/advent/utils"
	"github.com/stretchr/testify/assert"
)

var exampleInput = `ecl:gry pid:860033327 eyr:2020 hcl:#fffffd
byr:1937 iyr:2017 cid:147 hgt:183cm

iyr:2013 ecl:amb cid:350 eyr:2023 pid:028048884
hcl:#cfa07d byr:1929

hcl:#ae17e1 iyr:2013
eyr:2024
ecl:brn pid:760753108 byr:1931
hgt:179cm

hcl:#cfa07d eyr:2025 pid:166559648
iyr:2011 ecl:brn hgt:59in`

func TestParse(t *testing.T) {
	assert.Equal(t, parse(exampleInput), []passport{
		{ecl: "gry", pid: "860033327", eyr: "2020", hcl: "#fffffd", byr: "1937", iyr: "2017", cid: "147", hgt: "183cm"},
		{ecl: "amb", cid: "350", eyr: "2023", pid: "028048884", hcl: "#cfa07d", byr: "1929", iyr: "2013"},
		{hcl: "#ae17e1", iyr: "2013", eyr: "2024", ecl: "brn", pid: "760753108", byr: "1931", hgt: "179cm"},
		{hcl: "#cfa07d", eyr: "2025", pid: "166559648", iyr: "2011", ecl: "brn", hgt: "59in"},
	})
}

func TestValid(t *testing.T) {
	assert.True(t, passport{ecl: "gry", pid: "860033327", eyr: "2020", hcl: "#fffffd", byr: "1937", iyr: "2017", cid: "147", hgt: "183cm"}.valid())
	assert.False(t, passport{ecl: "amb", cid: "350", eyr: "2023", pid: "028048884", hcl: "#cfa07d", byr: "1929", iyr: "2013"}.valid())
	assert.True(t, passport{hcl: "#ae17e1", iyr: "2013", eyr: "2024", ecl: "brn", pid: "760753108", byr: "1931", hgt: "179cm"}.valid())
	assert.False(t, passport{hcl: "#cfa07d", eyr: "2025", pid: "166559648", iyr: "2011", ecl: "brn", hgt: "59in"}.valid())
}

func TestSolveExampleA(t *testing.T) {
	assert.Equal(t, 2, solveA(parse(exampleInput)))
}

func TestSolveA(t *testing.T) {
	assert.Equal(t, 182, solveA(parse(utils.LoadString("input.txt"))))
}

func TestSolveExampleB(t *testing.T) {
	assert.Equal(t, 2, solveB(parse(exampleInput)))
}

func TestSolveB(t *testing.T) {
	assert.Equal(t, 109, solveB(parse(utils.LoadString("input.txt"))))
}
