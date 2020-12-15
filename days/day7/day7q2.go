package day7

import (
	"strconv"
	"strings"

	"github.com/littleajax/adventofcode/helpers"
)

//Build a distributed graph
//Going the other direction now, not parents, but children, with multiplicative counts
func ShinyGoldBagChildren(shinyGold *BagWithChildren) (totalChildren int) {
	for _, child := range shinyGold.children {
		if child.bag.visited == true {
			//continue
		}

		childCount := child.count
		child.bag.visited = true

		subCount := ShinyGoldBagChildren(child.bag)

		if subCount != 0 {
			subCount = childCount * subCount
		}

		childCount += subCount

		totalChildren += childCount
	}

	return
}

func ProcessWithChildrenInputs() *BagWithChildren {
	inputs := helpers.FetchInputs("./inputs/day7.txt")

	bags := make(map[string]*BagWithChildren)

	//Create all the different bags
	for _, rule := range inputs {
		var bag BagWithChildren
		splits := strings.Split(rule, " ")
		name := splits[0] + splits[1]
		bag.name = name
		bags[name] = &bag
	}

	//Now set containers
	for _, rule := range inputs {
		var bagGroup BagGroup
		splits := strings.Split(rule, " ")
		parentName := splits[0] + splits[1]

		if splits[4] == "no" {
			continue //this bag is no ones container
		}

		//Jump through bag line
		for i := 4; i <= len(splits)-1; i += 4 {
			bagName := splits[i+1] + splits[i+2]
			bagGroup.bag = bags[bagName]
			bagGroup.count, _ = strconv.Atoi(splits[i]) //bad me, ignoring err
			bags[parentName].children = append(bags[parentName].children, bagGroup)
		}
	}

	return bags["shinygold"]
}

type BagWithChildren struct {
	name     string
	children []BagGroup
	visited  bool //So we don't enter infinite bag loops
}

type BagGroup struct {
	count int
	bag   *BagWithChildren
}
