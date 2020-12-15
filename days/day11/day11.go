package day11

import (
	"github.com/littleajax/adventofcode/helpers"
)

const (
	Empty     rune = '.'
	Occupied  rune = '#'
	Available rune = 'L'
)

func GetFerryRoundCount(maxOccupied int) int {
	seating := helpers.FetchInputs("./inputs/day11.txt")
	//convert to []runes for easier swapping
	runeSeating := make([][]rune, len(seating))
	for i, line := range seating {
		runeSeating[i] = []rune(line)
	}

	//lodo loop until it settles
	//O(rows*columns*loops*~8) //
	var changeCounter int
	var occupiedCounter int
	for {

		runeSeating, changeCounter, occupiedCounter = determineSeating(runeSeating, maxOccupied)
		//runeSeating, changeCounter, occupiedCounter = determineSeating(runeSeating, maxOccupied)
		//runeSeating, changeCounter, occupiedCounter = determineSeating(runeSeating, maxOccupied)

		//return 0

		if changeCounter == 0 {
			return occupiedCounter
		}
	}
}

func determineSeating(seating [][]rune, maxOccupied int) ([][]rune, int, int) {
	//Go row by row
	var changeCounter int
	var occupiedCount int

	seatingCopy := make([][]rune, len(seating))

	for i := range seating {
		seatingCopy[i] = make([]rune, len(seating[i]))
		copy(seatingCopy[i], seating[i])
	}

	for row, seats := range seating {
		for col, _ := range seats {
			switch seatingCopy[row][col] { //. can be ignored
			case Occupied:
				if countOccupiedSightLines(row, col, seatingCopy) >= 5 {
					seating[row][col] = Available
					changeCounter++
				}

			case Available:
				if countOccupiedSightLines(row, col, seatingCopy) == 0 {
					seating[row][col] = Occupied
					changeCounter++
				}
			}
		}

	}

	for row, seats := range seating {
		for col, _ := range seats {
			if seating[row][col] == Occupied {
				occupiedCount++
			}
		}
	}

	return seating, changeCounter, occupiedCount
}

func countOccupiedAdjacency(row int, col int, seating [][]rune) (adjacent int) {

	for i := 0; i < 3; i++ {
		//Skip the first loop if we're on the first row
		if row == 0 && i == 0 {
			continue
		}
		//Skip the last loop if we're on the last row
		if row == len(seating)-1 && i == 2 {
			continue
		}

		for j := 0; j < 3; j++ {
			//Skip the first loop if we're on the first col
			if col == 0 && j == 0 {
				continue
			}
			//Skip the last loop if we're on the last col
			if col == len(seating[i])-1 && j == 2 {
				continue
			}

			//Skip if we're on the seat
			if i == 1 && j == 1 {
				continue
			}

			if seating[row-1+i][col-1+j] == Occupied {
				adjacent++
			}
		}

	}

	return
}
