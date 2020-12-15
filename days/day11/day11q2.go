package day11

func countOccupiedSightLines(row int, col int, seating [][]rune) (adjacent int) {

	cardinalMap := make(map[string]rune)
	//Expand outward in all cardinals, stop on
	for i := 1; i < len(seating); i++ {
		back := i * -1
		forward := i

		//south
		if row+forward < len(seating) && cardinalMap["S"] == 0 {
			if seating[row+forward][col] != Empty {
				cardinalMap["S"] = seating[row+forward][col]
			}
		}
		//north
		if row+back >= 0 && cardinalMap["N"] == 0 {
			if seating[row+back][col] != Empty {
				cardinalMap["N"] = seating[row+back][col]
			}
		}
		//East
		if col+forward < len(seating[col]) && cardinalMap["E"] == 0 {
			if seating[row][col+forward] != Empty {
				cardinalMap["E"] = seating[row][col+forward]
			}
		}
		//West
		if col+back >= 0 && cardinalMap["W"] == 0 {
			if seating[row][col+back] != Empty {
				cardinalMap["W"] = seating[row][col+back]
			}
		}
		//NorthWest
		if col+back >= 0 && row+back >= 0 && cardinalMap["NW"] == 0 {
			if seating[row+back][col+back] != Empty {
				cardinalMap["NW"] = seating[row+back][col+back]
			}
		}
		//SouthEast
		if col+forward < len(seating[col]) && row+forward < len(seating) && cardinalMap["SE"] == 0 {
			if seating[row+forward][col+forward] != Empty {
				cardinalMap["SE"] = seating[row+forward][col+forward]
			}
		}

		//NorthEast
		if col+forward < len(seating[col]) && row+back >= 0 && cardinalMap["NE"] == 0 {
			if seating[row+back][col+forward] != Empty {
				cardinalMap["NE"] = seating[row+back][col+forward]
			}
		}
		//SouthWest
		if row+forward < len(seating) && col+back >= 0 && cardinalMap["SW"] == 0 {
			if seating[row+forward][col+back] != Empty {
				cardinalMap["SW"] = seating[row+forward][col+back]
			}
		}
	}

	for _, occupied := range cardinalMap {
		if occupied == Occupied {
			adjacent++
		}
	}

	return
}
