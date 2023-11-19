package AoC2021

import (
	_ "embed"
	"log"
	"strconv"
	"strings"

	Utils "github.com/chrisg07/Advent-of-Code-Go/Utils"
)

//go:embed inputs/example.txt
var exampleInput string

//go:embed inputs/complete.txt
var input string

func getInput(useExample bool) []string {
	var lines []string
	if useExample {
		exampleInput = strings.TrimRight(exampleInput, "\n")
		lines = strings.Split(exampleInput, "\n")
	} else {
		input = strings.TrimRight(input, "\n")
		lines = strings.Split(input, "\n")
	}
	return lines
}

func PartA(useExample bool) int {
	lines := getInput(useExample)

	previousDepth := 10000
	descents := 0

	for _, line := range lines {
		depth, parseErr := strconv.Atoi(line)

		if parseErr != nil {
			log.Fatal(parseErr)
		}

		didDescend := depth > previousDepth

		if didDescend {
			descents += 1
		}

		previousDepth = depth
	}

	return descents
}

func PartB(useExample bool) int {
	lines := getInput(useExample)

	descents := 0
	measurements := []int{}

	for _, line := range lines {
		currentDepth, parseErr := strconv.Atoi(line)

		if parseErr != nil {
			log.Fatal(parseErr)
		}

		var previousDepthWindow int
		var currentDepthWindow int
		measurements = append(measurements, currentDepth)
		var measurementsLength = len(measurements)
		if measurementsLength >= 4 {
			previousDepthWindow = Utils.SumArray(measurements[measurementsLength-4 : measurementsLength-1])
			currentDepthWindow = Utils.SumArray(measurements[measurementsLength-3 : measurementsLength])
			if currentDepthWindow > previousDepthWindow {
				descents += 1
			}
		}
	}

	return descents
}
