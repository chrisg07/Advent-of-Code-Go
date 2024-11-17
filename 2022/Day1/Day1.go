package AoC2022

import (
	_ "embed"
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

func parseInput(lines []string) []string {
	input := []string{}
	for _, line := range lines {
		// for _, char := range line {
		// 	log.Print(string(char))
		// }

		// log.Printf("[CONSOLE] %v", line)
		input = append(input, line)
	}
	return input
}

type Elf struct {
	food     []int
	calories int
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	elves := []Elf{}
	maxCalories := 0
	currentElf := Elf{food: []int{}, calories: 0}
	for _, line := range input {
		if line == "" {
			elves = append(elves, currentElf)
			if currentElf.calories > maxCalories {
				maxCalories = currentElf.calories
			}
			currentElf = Elf{food: []int{}, calories: 0}
		} else {
			food, _ := strconv.Atoi(line)
			currentElf.calories += food
		}
	}
	elves = append(elves, currentElf)
	if currentElf.calories > maxCalories {
		maxCalories = currentElf.calories
	}

	return maxCalories
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	elves := []Elf{}
	maxCalories := 0
	currentElf := Elf{food: []int{}, calories: 0}
	for _, line := range input {
		if line == "" {
			elves = append(elves, currentElf)
			if currentElf.calories > maxCalories {
				maxCalories = currentElf.calories
			}
			currentElf = Elf{food: []int{}, calories: 0}
		} else {
			food, _ := strconv.Atoi(line)
			currentElf.calories += food
		}
	}

	elves = append(elves, currentElf)

	sort.Slice(elves, func(i, j int) bool {
		return elves[i].calories > elves[j].calories
	})

	return elves[0].calories + elves[1].calories + elves[2].calories
}
