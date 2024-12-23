package AoC2019

import (
	_ "embed"
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

func parseInput(lines []string) []int {
	input := []int{}
	for _, line := range lines {
		mass, _ := strconv.Atoi(line)
		input = append(input, mass)
	}
	return input
}

func calculateFuelCost(mass int) int {
	return (mass / 3) - 2
}

func calculateTotalFuelCost(mass int) int {
	fuelCost := (mass / 3) - 2
	if fuelCost > 8 {
		return fuelCost + calculateTotalFuelCost(fuelCost)
	} else {
		return fuelCost
	}
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	fuelRequirement := 0

	for _, mass := range input {
		fuelRequirement += calculateFuelCost(mass)
	}
	return fuelRequirement
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	fuelRequirement := 0

	for _, mass := range input {
		fuelRequirement += calculateTotalFuelCost(mass)
	}
	return fuelRequirement
}
