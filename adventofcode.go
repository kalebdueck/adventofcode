package main

import (
	"fmt"

	"github.com/littleajax/adventofcode/days/day7"

	"github.com/littleajax/adventofcode/days/day5"
	"github.com/littleajax/adventofcode/days/day6"

	"github.com/littleajax/adventofcode/days/day4"

	"github.com/littleajax/adventofcode/days/day2"
	"github.com/littleajax/adventofcode/days/day3"

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

	fmt.Println("Day 3: -----------")
	day3Result := day3.TreesSmashed(day3.ProcessInputs())
	fmt.Println(day3Result)

	fmt.Println("Day 4: -----------")
	day4Result := day4.ValidPassportCount(day4.ProcessInputs())
	fmt.Println(day4Result)

	fmt.Println("Day 5: -----------")
	day5Results := day5.GetHighestSeatNumber(day5.ProcessInputs())
	fmt.Println(day5Results)

	fmt.Println("Day 6: -----------")
	day6Results := day6.TallyCustomsFields(day6.ProcessInputs())
	fmt.Println(day6Results)
	day6q2Results := day6.TallyEveryonesCustomsFields(day6.ProcessEveryonesInputs())
	fmt.Println(day6q2Results)

	fmt.Println("Day 7: -----------")
	day7Results := day7.ShinyGoldBagContainers(day7.ProcessInputs())
	fmt.Println(day7Results)
	day7q2Results := day7.ShinyGoldBagChildren(day7.ProcessWithChildrenInputs())
	fmt.Println(day7q2Results)
}
