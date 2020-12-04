package day3

import (
	"fmt"

	"github.com/littleajax/adventofcode/helpers"
)

func ProcessInputs() (mountainSide [][]bool) {
	inputs := helpers.FetchInputs("./inputs/day3.txt")

	for _, treeLine := range inputs {

		var level []bool
		for _, char := range treeLine {
			level = append(level, char == '#')
		}

		mountainSide = append(mountainSide, level)
	}

	return
}

type path struct {
	over     int
	down     int
	latitude int
}

func TreesSmashed(mountainSide [][]bool) (total int) {
	//drop 1
	paths := []path{
		{
			over: 1,
			down: 1,
		},
		{
			over: 3,
			down: 1,
		},
		{
			over: 5,
			down: 1,
		},
		{
			over: 7,
			down: 1,
		},
		{
			over: 1,
			down: 2,
		},
	}

	for _, run := range paths {
		var hits int

		for level, treeLine := range mountainSide {

			if level%run.down != 0 {
				continue
			}

			patternLength := len(treeLine)

			if treeLine[run.latitude] == true {
				hits++
			}

			run.latitude += run.over

			if run.latitude > patternLength-1 {
				run.latitude -= patternLength
			}

		}
		fmt.Println(hits)

		if total == 0 {
			total = hits
		} else {
			total *= hits
		}
		hits = 0

	}

	return
}
