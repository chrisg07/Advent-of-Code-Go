package AoCScaffold

import (
	_ "embed"
	"log"
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

func parseInstructions(instruction string) int {
	sum := 0
	parts := strings.Split(instruction, "mul")

	for _, part := range parts {
		end := strings.Index(part, ")")
		if end != -1 && end != 0 {

			instruction := part[1:end]
			variables := strings.Split(instruction, ",")
			if len(variables) == 2 {
				left, leftOk := strconv.Atoi(variables[0])
				right, rightOk := strconv.Atoi(variables[1])
				if leftOk == nil && rightOk == nil {
					sum += (left * right)
					log.Printf("[DEBUG] Instruction: %v", instruction)
				}
			}
		}
	}
	return sum
}

func parseInput(lines []string) string {
	enabledInstructions := []string{}
	input := strings.Join(lines, "")

	log.Printf("[DEBUG] Enabled Instructions: \n%v", strings.Join(enabledInstructions, "\n"))

	return input
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return parseInstructions(input)
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	parts := strings.Split(input, "do()")
	sum := 0

	for _, part := range parts {
		end := strings.Index(part, "don't()")
		validMults := part
		if end != -1 {
			validMults = part[:end]
		}
		sum += parseInstructions(validMults)
	}
	return sum
}
