package AoCScaffold

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

func parseInput(lines []string) ([]int, []int) {
	left := []int{}
	right := []int{}
	for _, line := range lines {
		nums := strings.Split(line, "|")
		log.Printf("[DEBUG] %v", line)
		input = append(input, line)
	}
	return input
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return len(input)
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return len(input)
}
