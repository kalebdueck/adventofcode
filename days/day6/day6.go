package day6

import (
	"github.com/littleajax/adventofcode/helpers"
)

func TallyCustomsFields(passportLines [][]rune) (tally int) {

	for _, passportGroup := range passportLines {
		tallySet := make(map[rune]bool)

		for _, char := range passportGroup {
			tallySet[char] = true
		}

		tally += len(tallySet)
	}

	return
}

func ProcessInputs() (passportLines [][]rune) {
	inputs := helpers.FetchInputs("./inputs/day6.txt")

	var passportFields []rune
	for _, passportLine := range inputs {
		if passportLine == "" {
			passportLines = append(passportLines, passportFields)
			//empty the old passport
			passportFields = []rune{}
			continue
		}

		//Split on space to get each line,
		for _, field := range passportLine {
			passportFields = append(passportFields, field)
		}
	}

	//Last field doesn't end...
	if len(passportFields) > 1 {
		passportLines = append(passportLines, passportFields)
	}

	return
}
