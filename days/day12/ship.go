package day12

type Ship struct {
	x      int
	y      int
	Facing int
}

func (ship Ship) turn(direction rune, degrees int) Ship {
	if direction == 'L' {
		ship.Facing -= degrees

		if ship.Facing < 0 {
			ship.Facing += 360
		}

		return ship
	}

	ship.Facing += degrees

	if ship.Facing >= 360 {
		ship.Facing -= 360
	}

	return ship
}

func (ship Ship) Move(direction rune, distance int) Ship {
	if direction == 'F' {
		direction = directions[ship.Facing]
	}

	switch direction {
	case 'E':
		ship.x += distance
	case 'W':
		ship.x -= distance
	case 'N':
		ship.y += distance
	case 'S':
		ship.y -= distance
	}

	return ship
}
