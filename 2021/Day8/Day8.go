package AoC2021

import (
	_ "embed"
	"log"
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

func Day8PartA2021(useExample bool) int {
	lines := getInput(useExample)
	occurences := 0
	for _, line := range lines {
		lineParts := strings.Split(line, " | ")
		// signalPatterns := strings.Split(lineParts[0], " ")
		outputs := strings.Split(lineParts[1], " ")
		for _, output := range outputs {
			switch len(output) {
			case 2:
				occurences += 1
			case 4:
				occurences += 1
			case 3:
				occurences += 1
			case 7:
				occurences += 1

			}
		}
	}

	return occurences
}

func Day8PartB2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
