package AoC2021

import (
	_ "embed"
	"fmt"
	"strings"
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

func Day4PartA2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			fmt.Print(char)
		}
	}

	return 0
}

func Day4PartB2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			fmt.Print(char)
		}
	}

	return 0
}
