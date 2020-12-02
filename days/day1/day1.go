package day1

import (
	"log"
	"strconv"

	"github.com/littleajax/adventofcode/helpers"
)

func ExpenseReport(report []int) int {
	for i, value := range report {
		compare := 2020 - value
		for j := i + 1; j < len(report); j++ {
			if report[j] == compare {
				return value * compare
			}
		}
	}

	return 0
}

func ThreeEntriesExpenseReport(report []int) int {
	for _, value := range report {
		for a, secondValue := range report {

			compare := 2020 - (value + secondValue)
			for j := a + 1; j < len(report); j++ {
				if report[j] == compare {
					return value * compare * secondValue
				}
			}
		}
	}

	return 0
}

func ProcessInputs() []int {
	values := helpers.FetchInputs("./inputs/day1q1.txt")

	inputs := []int{}
	for _, value := range values {
		intValue, err := strconv.Atoi(value)
		if err != nil {
			log.Fatal(err)
		}
		inputs = append(inputs, intValue)
	}

	return inputs
}
