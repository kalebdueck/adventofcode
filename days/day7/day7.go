package day7

import (
	"strings"

	"github.com/littleajax/adventofcode/helpers"
)

//Build a distributed graph
func ShinyGoldBagContainers(shinyGold *Bag) (possibleContainers int) {
	for _, container := range shinyGold.containedIn {
		if container.visited == true {
			continue
		}

		container.visited = true
		possibleContainers++
		possibleContainers += ShinyGoldBagContainers(container)
	}

	return
}

func ProcessInputs() *Bag {
	inputs := helpers.FetchInputs("./inputs/day7.txt")

	bags := make(map[string]*Bag)

	//Create all the different bags
	for _, rule := range inputs {
		var bag Bag
		splits := strings.Split(rule, " ")
		name := splits[0] + splits[1]
		bag.name = name
		bags[name] = &bag
	}

	//Now set containers
	for _, rule := range inputs {
		splits := strings.Split(rule, " ")
		container := splits[0] + splits[1]

		if splits[4] == "no" {
			continue //this bag is no ones container
		}

		//Jump through bag line
		for i := 4; i <= len(splits)-1; i += 4 {
			bagName := splits[i+1] + splits[i+2]
			bags[bagName].containedIn = append(bags[bagName].containedIn, bags[container])
		}
	}

	return bags["shinygold"]
}

type Bag struct {
	name        string
	containedIn []*Bag
	visited     bool //So we don't enter infinite bag loops
}
