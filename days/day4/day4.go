package day4

import (
	"regexp"
	"strconv"
	"strings"

	"github.com/littleajax/adventofcode/helpers"
)

func checkField(name string, value string) bool {
	switch name {
	case "cid":
		return false
	case "byr":
		return stringToIntRange(value, 1920, 2002)
	case "iyr":
		return stringToIntRange(value, 2010, 2020)
	case "eyr":
		return stringToIntRange(value, 2020, 2030)
	case "hgt":
		if len(value) < 3 {
			return false
		}
		if value[len(value)-2:] == "cm" {
			return stringToIntRange(value[:len(value)-2], 150, 193)
		}
		if value[len(value)-2:] == "in" {
			return stringToIntRange(value[:len(value)-2], 59, 76)
		}
	case "hcl":
		if len(value) != 7 {
			return false
		}
		regex := regexp.MustCompile(`^\#[0-9a-z]{6}$`)
		return regex.MatchString(value)

	case "ecl":
		eclValues := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}
		for _, val := range eclValues {
			if value == val {
				return true
			}
		}

	case "pid": //A nine-digit number, including leading zeros
		if len(value) != 9 {
			return false
		}
		regex := regexp.MustCompile(`^[0-9]{9}$`)
		return regex.MatchString(value)
	}
	return false
}
func stringToIntRange(s string, min int, max int) bool {
	number, err := strconv.Atoi(s)
	if err != nil {

		return false
	}
	if number >= min && number <= max {
		return true
	}

	return false
}

func ValidPassportCount(passportLines [][]string) (valids int) {
	var counter int
	for _, passportFields := range passportLines {
		counter = 0

		for _, field := range passportFields {
			if checkField(field[:3], field[4:]) {
				counter++
			}
		}

		if counter >= 7 {
			valids++
		}
	}

	return
}

func ProcessInputs() (passportLines [][]string) {
	inputs := helpers.FetchInputs("./inputs/day4.txt")

	var passportFields []string
	for _, passportLine := range inputs {
		if passportLine == "" {
			passportLines = append(passportLines, passportFields)
			//empty the old passport
			passportFields = []string{}
			continue
		}

		//Split on space to get each line,
		for _, field := range strings.Split(passportLine, " ") {
			passportFields = append(passportFields, field)
		}
	}

	//Last field doesn't end...
	if len(passportFields) > 1 {
		passportLines = append(passportLines, passportFields)
	}

	return
}
