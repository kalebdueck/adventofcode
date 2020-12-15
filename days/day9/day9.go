package day9

import (
	"strconv"

	"github.com/littleajax/adventofcode/helpers"
)

func ProcessInputs() (inputs []int) {
	values := helpers.FetchInputs("./inputs/day9.txt")

	for _, line := range values {
		integer, _ := strconv.Atoi(line)
		inputs = append(inputs, integer)
	}

	return
}

func FirstInvalidNumber(integers []int) int {

	for i, integer := range integers {
		//Skip the preamble
		if i <= 25 {
			continue
		}

		//Quick little 2 sum on the previous
		compliments := make(map[int]int)
		var found bool
		for _, val := range integers[i-25 : i] {

			compliment := integer - val

			if _, exists := compliments[compliment]; exists {
				found = true
				break
			}

			compliments[val] = compliment
		}

		if found == false {

			return integer
		}
	}

	return 0
}

func SmallestAndLargestOfAContiguousRange(integers []int, invalid int) (low int, high int) {
	var overrun int
	for i, val := range integers {
		if val == invalid {
			continue
		}
		if val < overrun {
			continue
		}

		overrun = 0
		var tally int
		var j = i
		high = 0
		low = 0
		for {

			if low == 0 || integers[j] < low {
				low = integers[j]
			}
			if integers[j] > high {
				high = integers[j]
			}

			tally += integers[j] //Look into ensuring I gets in there

			if tally == invalid {
				return
			}

			if tally > invalid {
				overrun = tally - invalid
				tally = 0
				break
			}

			j++
		}
		//fmt.Println("wentOver")
		//fmt.Println(i)
		//fmt.Println(tally)
	}

	return low, high
}
