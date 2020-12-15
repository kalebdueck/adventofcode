package day12

var directions map[int]rune = map[int]rune{
	0:   'E',
	90:  'S',
	180: 'W',
	270: 'N',
}

type WaypointShip struct {
	x         int
	y         int
	WaypointX int
	WaypointY int
}

func (ship WaypointShip) turn(direction rune, degrees int) WaypointShip {
	turns := degrees / 90
	if direction == 'L' {

		var i int
		for i < turns {
			ship.WaypointX, ship.WaypointY = ship.WaypointY*-1, ship.WaypointX
			i++
		}

		return ship
	}

	var i int
	for i < turns {
		ship.WaypointX, ship.WaypointY = ship.WaypointY, ship.WaypointX*-1
		i++
	}

	return ship
}

func (ship WaypointShip) MoveWaypoint(direction rune, distance int) WaypointShip {

	switch direction {
	case 'E':
		ship.WaypointX += distance
	case 'W':
		ship.WaypointX -= distance
	case 'N':
		ship.WaypointY += distance
	case 'S':
		ship.WaypointY -= distance
	}

	return ship
}

func (ship WaypointShip) Move(direction rune, distance int) WaypointShip {
	yDistance := ship.WaypointY * distance
	ship.y += yDistance

	xDistance := ship.WaypointX * distance
	ship.x += xDistance

	return ship
}
