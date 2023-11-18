package AoC2021

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed inputs/Day1/PartAExample.txt
var exampleInput string

//go:embed inputs/Day1/PartA.txt
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

func Day1PartA(useExample bool) int {
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
