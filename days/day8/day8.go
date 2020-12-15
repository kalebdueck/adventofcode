package day8

import (
	"strconv"
	"strings"

	"github.com/littleajax/adventofcode/helpers"
)

func ProcessInputs() (instructions []instruction) {
	inputs := helpers.FetchInputs("./inputs/day8.txt")

	for _, line := range inputs {
		split := strings.Split(line, " ")
		value, err := strconv.Atoi(split[1])
		if err != nil {
			panic(err)
		}
		instructions = append(instructions, instruction{
			name:  split[0],
			value: value,
		})
	}

	return
}

type instruction struct {
	name    string
	value   int
	visited bool
}

//Wonder if i can go negative?
func GetFinalAccumulatorValue(instructions []instruction) (acc int) {
	var i int
	for {
		if instructions[i].visited == true {
			return
		}

		instructions[i].visited = true

		switch instructions[i].name {
		case "acc":
			acc += instructions[i].value
			i++
		case "jmp":
			i += instructions[i].value
		case "nop":
			i++
		}
	}
}

//Naive fix is O(n^2)
//loop once keeping a counter i of jmp/nops
//re-loop swapping the i'th jmp <-> nop
//if no loop -> return acc
//
//Better fix would be to loop through as before, updating acc
//On finding a jmp/nop, mark location, swap and run on remaining slice.
//if no loop, add result to acc -> return
//else continue the original loop.
//
//Faster because we don't need to recompute the initial loop
func GetUncorruptedAccumulatorValue(instructions []instruction) (acc int) {
	var i int

	for i != len(instructions)-1 {
		if instructions[i].visited == true {
			return
		}

		instructions[i].visited = true

		switch instructions[i].name {
		case "acc":
			acc += instructions[i].value
			i++
		case "jmp":
			newInstructions := make([]instruction, len(instructions))
			copy(newInstructions, instructions)
			newAcc, looping := CheckLoopWithSwap(newInstructions, i)
			if looping == false {
				acc += newAcc
				return
			}
			i += instructions[i].value
		case "nop":
			newInstructions := make([]instruction, len(instructions))
			copy(newInstructions, instructions)
			newAcc, looping := CheckLoopWithSwap(newInstructions, i)
			if looping == false {
				acc += newAcc
				return
			}
			i++
		}
	}

	return
}

func CheckLoopWithSwap(instructions []instruction, i int) (acc int, looped bool) {

	switch instructions[i].name {
	case "jmp":
		i++
	case "nop":
		i += instructions[i].value
	}

	for i != len(instructions)-1 {
		if instructions[i].visited == true {
			looped = true
			return
		}

		instructions[i].visited = true

		switch instructions[i].name {
		case "acc":
			acc += instructions[i].value
			i++
		case "jmp":
			i += instructions[i].value
		case "nop":
			i++
		}
	}

	return
}
