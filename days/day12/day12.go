package day12

import (
	"fmt"
	"strconv"

	"github.com/littleajax/adventofcode/helpers"
)

type orders struct {
	operation rune
	amount    int
}

func ProcessInputs() []orders {
	values := helpers.FetchInputs("./inputs/day12.txt")
	ordersList := make([]orders, len(values))

	for i, value := range values {
		order := orders{}
		r := []rune(value)
		order.operation = r[0]
		order.amount, _ = strconv.Atoi(string(r[1:]))

		ordersList[i] = order
	}

	return ordersList
}

func ExecuteOrders(ship Ship, orders []orders) int {

	for _, order := range orders {

		switch order.operation {
		case 'L', 'R':
			ship = ship.turn(order.operation, order.amount)
		default:
			ship = ship.Move(order.operation, order.amount)
		}
	}
	fmt.Printf("%v", ship)

	return abs(ship.x) + abs(ship.y)
}

func ExecuteWaypointOrders(ship WaypointShip, orders []orders) int {
	for _, order := range orders {
		switch order.operation {
		case 'L', 'R':
			ship = ship.turn(order.operation, order.amount)
		case 'F':
			ship = ship.Move(order.operation, order.amount)
		default:
			ship = ship.MoveWaypoint(order.operation, order.amount)
		}
	}
	fmt.Printf("%v", ship)

	return abs(ship.x) + abs(ship.y)

}

func abs(x int) int {
	if x < 0 {
		return x * -1
	}
	return x
}
