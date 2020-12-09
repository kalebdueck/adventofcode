package day6

import (
	"github.com/littleajax/adventofcode/helpers"
)

func TallyEveryonesCustomsFields(passportLines [][]string) (tally int) {

	for _, passportGroup := range passportLines {

		tallySet := make(map[rune]int)

		for _, personalLine := range passportGroup {
			for _, char := range personalLine {
				tallySet[char] += 1
			}
		}

		for _, count := range tallySet {
			if count == len(passportGroup) {
				tally++
			}
		}

	}

	return
}

func ProcessEveryonesInputs() (passportLines [][]string) {
	inputs := helpers.FetchInputs("./inputs/day6.txt")

	var passportFields []string
	for _, passportLine := range inputs {
		if passportLine == "" {
			passportLines = append(passportLines, passportFields)
			//empty the old passport
			passportFields = []string{}
			continue
		}

		passportFields = append(passportFields, passportLine)

	}

	//Last field doesn't end...
	if len(passportFields) > 1 {
		passportLines = append(passportLines, passportFields)
	}

	return
}
