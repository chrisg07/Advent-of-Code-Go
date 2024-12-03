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

func parseInput(lines []string) int {
	input := []string{}
	sum := 0
	for _, line := range lines {
		// for _, char := range line {
		// 	log.Print(string(char))
		// }

		// log.Printf("[DEBUG] Line: %v", line)
		parts := strings.Split(line, "mul")

		for _, part := range parts {
			// log.Printf("[DEBUG] Part: %v", part)
			end := strings.Index(part, ")")
			if end != -1 {

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
		input = append(input, line)
	}
	return sum
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return input
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return input
}
