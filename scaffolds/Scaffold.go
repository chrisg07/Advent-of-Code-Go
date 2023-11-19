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
		input = strings.TrimRight(exampleInput, "\n")
	} else {
		input = strings.TrimRight(input, "\n")
	}
	lines = strings.Split(input, "\n")
	return lines
}

func DayXPartA2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println("")
	}

	return 0
}

func DayXPartB2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println("")
	}

	return 0
}
