package day2

import (
	"strconv"
	"strings"

	"github.com/littleajax/adventofcode/helpers"
)

type passwordRules struct {
	Password string
	Char     rune
	Min      int
	Max      int
}

func ProcessInputs() []passwordRules {
	values := helpers.FetchInputs("./inputs/day2.txt")

	inputs := []passwordRules{}
	for _, value := range values {
		//strings are formatted as n-k: password
		hyphenIndex := strings.IndexRune(value, '-')
		spaceIndex := strings.IndexRune(value, ' ')
		colonIndex := strings.IndexRune(value, ':')

		min, _ := strconv.Atoi(value[:hyphenIndex])
		max, _ := strconv.Atoi(value[hyphenIndex+1 : spaceIndex])
		char := value[spaceIndex+1]
		password := value[colonIndex+2:]

		inputs = append(inputs, passwordRules{
			Password: password,
			Min:      min,
			Max:      max,
			Char:     rune(char),
		})

	}
	return inputs
}

func PasswordPositionCheck(passwords []passwordRules) (hitcount int) {

	for _, value := range passwords {
		var charCount int
		if value.Char == rune(value.Password[value.Min-1]) {
			charCount++
		}
		if value.Char == rune(value.Password[value.Max-1]) {
			charCount++
		}

		if charCount == 1 {
			hitcount++
		}
	}

	return
}

func PasswordCheck(passwords []passwordRules) (hitcount int) {

	for _, value := range passwords {
		var charCount int

		for _, char := range value.Password {
			if rune(char) == value.Char {
				charCount++
			}
		}

		if charCount >= value.Min && charCount <= value.Max {
			hitcount++
		}
	}

	return
}
