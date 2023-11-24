package AoC2021

import (
	_ "embed"
	"log"
	"sort"
	"strconv"
	"strings"
)

//go:embed inputs/example.txt
var exampleInput string

//go:embed inputs/complete.txt
var input string

func getInput(useExample bool) []string {
	var lines []string
	var unsplitLines string
	if useExample {
		unsplitLines = strings.TrimRight(exampleInput, "\n")
	} else {
		unsplitLines = strings.TrimRight(input, "\n")
	}
	lines = strings.Split(unsplitLines, "\n")
	return lines
}

func CalculateFuelPartA(crabs []int, xPos int) int {
	fuelCost := 0
	for _, crab := range crabs {
		if crab-xPos < 0 {
			fuelCost += (crab - xPos) * -1
		} else {
			fuelCost += (crab - xPos)
		}
	}
	return fuelCost
}

func Day7PartA2021(useExample bool) int {
	lines := getInput(useExample)
	crabs := []int{}
	for _, line := range lines {
		crabsStr := strings.Split(line, ",")
		for _, str := range crabsStr {
			log.Print(string(str))
			value, _ := strconv.Atoi(string(str))
			crabs = append(crabs, value)
		}
		log.Println("")
	}

	sort.Ints(crabs)

	fuelCosts := []int{}
	for x := 0; x < crabs[len(crabs)-1]; x++ {
		fuelCost := CalculateFuelPartA(crabs, x)
		fuelCosts = append(fuelCosts, fuelCost)
	}

	sort.Ints(fuelCosts)

	log.Printf("[WARN] %v", crabs)
	return fuelCosts[0]
}

func CalculateFuelForDistance(distance int) int {
	fuelCost := 0
	for x := 0; x < distance; x++ {
		fuelCost += (x + 1)
	}
	return fuelCost
}

func CalculateFuelPartB(crabs []int, xPos int) int {
	fuelCost := 0
	for _, crab := range crabs {
		distance := 0
		if crab-xPos < 0 {
			distance += (crab - xPos) * -1
		} else {
			distance += (crab - xPos)
		}
		fuelCost += CalculateFuelForDistance(distance)
	}
	return fuelCost
}

func Day7PartB2021(useExample bool) int {
	lines := getInput(useExample)
	crabs := []int{}
	for _, line := range lines {
		crabsStr := strings.Split(line, ",")
		for _, str := range crabsStr {
			log.Print(string(str))
			value, _ := strconv.Atoi(string(str))
			crabs = append(crabs, value)
		}
		log.Println("")
	}

	sort.Ints(crabs)

	fuelCosts := []int{}
	for x := 0; x < crabs[len(crabs)-1]; x++ {
		fuelCost := CalculateFuelPartB(crabs, x)
		fuelCosts = append(fuelCosts, fuelCost)
	}

	sort.Ints(fuelCosts)

	log.Printf("[WARN] %v", crabs)
	return fuelCosts[0]
}
