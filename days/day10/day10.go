package day10

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/littleajax/adventofcode/helpers"
)

func ProcessInputs() (inputs []int) {
	values := helpers.FetchInputs("./inputs/day10.txt")

	for _, line := range values {
		integer, _ := strconv.Atoi(line)
		inputs = append(inputs, integer)
	}

	return
}

func JoltageCalculator(integers []int) int {

	integers = quickSort(integers)
	var oneJump int
	var threeJump int

	if integers[0] == 1 {
		oneJump++
	} else {
		threeJump++
	}

	for i := 0; i < len(integers)-1; i++ {
		switch integers[i+1] - integers[i] {
		case 1:
			oneJump++
		case 3:
			threeJump++
		default:
		}
	}

	threeJump++ //Them edge cases

	fmt.Println(integers)
	return oneJump * threeJump
}

func PermutationsCalculator(integers []int) int {
	integers = quickSort(integers)
	fmt.Println(len(integers))
	integers = append([]int{0}, integers...)
	fmt.Println(len(integers))
	//integers = append(integers, 150)
	adapters := make([]adapter, len(integers))

	for i, _ := range integers {
		adapters[i] = adapter{value: integers[i]}
	}
	for i, _ := range integers {
		if i+1 == len(integers) {
			fmt.Println("EH")
			continue
		}

		if integers[i+1]-integers[i] == 1 {
			adapters[i].oneJump = &adapters[i+1]
		}

		for j := i + 1; j < len(integers) && j < i+4; j++ {
			if integers[j]-integers[i] == 3 {
				adapters[i].threeJump = &adapters[j]
			}
		}

		for k := i + 1; k < len(integers) && k < i+3; k++ {
			if integers[k]-integers[i] == 2 {
				adapters[i].twoJump = &adapters[k]
			}
		}

		fmt.Printf("%+v\n", adapters[i])
	}

	fmt.Println(adapters)

	return permustationCounter(&adapters[0])
}

func permustationCounter(adapter *adapter) (counter int) {
	fmt.Println(adapter.value)
	if adapter.oneJump == nil && adapter.threeJump == nil && adapter.twoJump == nil {
		return 1
	}

	if adapter.counter != 0 {
		return adapter.counter
	}

	if adapter.oneJump != nil {
		counter += permustationCounter(adapter.oneJump)
	}

	if adapter.twoJump != nil {
		counter += permustationCounter(adapter.twoJump)
	}

	if adapter.threeJump != nil {
		counter += permustationCounter(adapter.threeJump)
	}
	adapter.counter = counter

	return counter
}

func quickSort(a []int) []int {
	if len(a) < 2 {
		return a
	}

	left, right := 0, len(a)-1

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

type adapter struct {
	value     int
	oneJump   *adapter
	twoJump   *adapter
	threeJump *adapter
	counter   int
}
