package AoC2021

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

	xPos := 0
	yPos := 0

	for _, line := range lines {
		command := strings.Split(line, " ")

		distance, _ := strconv.Atoi(command[1])

		switch command[0] {
		case "forward":
			xPos += distance
		case "up":
			yPos -= distance
		case "down":
			yPos += distance
		}
	}

	return xPos * yPos
}

func PartB(useExample bool) int {
	// lines := getInput(useExample)

	answer := 0

	// for _, line := range lines {
	// }

	return answer
}
