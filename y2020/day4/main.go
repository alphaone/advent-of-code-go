package day4

import (
	"regexp"
	"strings"
)

type passport struct {
	byr string
	iyr string
	eyr string
	hgt string
	hcl string
	ecl string
	pid string
	cid string
}

func parse(input string) []passport {
	parts := strings.Split(input, "\n\n")

	passports := []passport{}
	for _, part := range parts {
		part = strings.ReplaceAll(part, "\n", " ")
		passport := passport{}
		fields := strings.Split(part, " ")
		for _, field := range fields {
			keyvalue := strings.Split(field, ":")
			switch keyvalue[0] {
			case "byr":
				passport.byr = keyvalue[1]
			case "iyr":
				passport.iyr = keyvalue[1]
			case "eyr":
				passport.eyr = keyvalue[1]
			case "hgt":
				passport.hgt = keyvalue[1]
			case "hcl":
				passport.hcl = keyvalue[1]
			case "ecl":
				passport.ecl = keyvalue[1]
			case "pid":
				passport.pid = keyvalue[1]
			case "cid":
				passport.cid = keyvalue[1]
			}
		}
		passports = append(passports, passport)
	}

	return passports
}

func (p passport) valid() bool {
	return p.byr != "" &&
		p.iyr != "" &&
		p.eyr != "" &&
		p.hgt != "" &&
		p.hcl != "" &&
		p.ecl != "" &&
		p.pid != ""
}

func (p passport) validB() bool {
	return p.valid() &&
		p.byrValid() &&
		p.iyrValid() &&
		p.eyrValid() &&
		p.hgtValid() &&
		p.hclValid() &&
		p.eclValid() &&
		p.pidValid()
}

func (p passport) byrValid() bool {
	return p.byr >= "1920" && p.byr <= "2002"
}

func (p passport) iyrValid() bool {
	return p.iyr >= "2010" && p.iyr <= "2020"
}

func (p passport) eyrValid() bool {
	return p.eyr >= "2020" && p.eyr <= "2030"
}

func (p passport) hgtValid() bool {
	if strings.HasSuffix(p.hgt, "cm") {
		return p.hgt >= "150cm" && p.hgt <= "193cm"
	}
	if strings.HasSuffix(p.hgt, "in") {
		return p.hgt >= "59in" && p.hgt <= "76in"
	}
	return false
}

func (p passport) hclValid() bool {
	re := regexp.MustCompile(`^#[0-9a-f]{6}$`)
	return re.MatchString(p.hcl)
}

func (p passport) eclValid() bool {
	re := regexp.MustCompile(`^(amb|blu|brn|gry|grn|hzl|oth)$`)
	return re.MatchString(p.ecl)
}

func (p passport) pidValid() bool {
	re := regexp.MustCompile(`^\d{9}$`)
	return re.MatchString(p.pid)
}

func solveA(passports []passport) int {
	valid := 0
	for _, p := range passports {
		if p.valid() {
			valid++
		}
	}
	return valid
}

func solveB(passports []passport) int {
	valid := 0
	for _, p := range passports {
		if p.validB() {
			valid++
		}
	}
	return valid
}
