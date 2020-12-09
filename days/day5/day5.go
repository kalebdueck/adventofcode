package day5

import (
	"math/rand"

	"github.com/littleajax/adventofcode/helpers"
)

func GetHighestSeatNumber(seatings []string) (mySeatNumber int) {
	var seatNumbers []int
	for _, seating := range seatings {
		var seatedRow int
		var seatedCol int
		//Build a slice of 128 for row
		rows := make([]int, 128)
		for i := range rows {
			rows[i] = i
		}
		//Build a slice of 8 for columns
		cols := make([]int, 8)
		for i := range cols {
			cols[i] = i
		}

		for _, direction := range seating[:7] {
			halfwayPoint := len(rows) / 2 //64
			if direction == 'F' {
				rows = rows[:halfwayPoint]
			} else {
				rows = rows[halfwayPoint:]
			}
		}
		seatedRow = rows[0]

		for _, direction := range seating[7:] {
			halfwayPoint := len(cols) / 2 //64
			if direction == 'L' {
				cols = cols[:halfwayPoint]
			} else {
				cols = cols[halfwayPoint:]
			}

		}
		seatedCol = cols[0]

		seatNumber := seatedRow*8 + seatedCol

		seatNumbers = append(seatNumbers, seatNumber)
	}

	//sort the seatNumbers
	seatNumbers = quickSort(seatNumbers)

	for i, _ := range seatNumbers {
		if i == len(seatNumbers)-1 {
			continue
		}
		if seatNumbers[i+1]-seatNumbers[i] != 1 {
			return seatNumbers[i] + 1
		}
	}

	return
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}
	left, right := 0, len(a)-1

	//pivot := len(a) / 2
	pivot := rand.Int() % len(a)

	a[pivot], a[right] = a[right], a[pivot]

	for i, _ := range a {
		if a[i] < a[right] {
			a[left], a[i] = a[i], a[left]
			left++
		}
	}
	a[left], a[right] = a[right], a[left]

	quickSort(a[:left])
	quickSort(a[left+1:])
	return a
}

func ProcessInputs() []string {
	values := helpers.FetchInputs("./inputs/day5.txt")
	return values
}
