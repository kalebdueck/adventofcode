package main

import (
	"fmt"

	"github.com/littleajax/adventofcode/days/day2"

	"github.com/littleajax/adventofcode/days/day1"
)

func main() {
	fmt.Println("Day 1: -----------")
	day1Result := day1.ExpenseReport(day1.ProcessInputs())
	fmt.Println(day1Result)
	day1Result2 := day1.ThreeEntriesExpenseReport(day1.ProcessInputs())
	fmt.Println(day1Result2)

	fmt.Println("Day 2: -----------")
	day2Result := day2.PasswordCheck(day2.ProcessInputs())
	fmt.Println(day2Result)
	day2Result2 := day2.PasswordPositionCheck(day2.ProcessInputs())
	fmt.Println(day2Result2)
}
