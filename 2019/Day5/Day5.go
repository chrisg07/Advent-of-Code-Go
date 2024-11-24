package AoCScaffold

import (
	_ "embed"
	"log"
	"slices"
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
	instructions := []int{}
	for _, line := range lines {
		opcodes := strings.Split(line, ",")
		for _, opcode := range opcodes {
			instruction, _ := strconv.Atoi(opcode)
			instructions = append(instructions, instruction)
		}
	}
	return instructions
}

func compute(instructions []int, index int) []int {
	switch instructions[index] {
	case 1:
		instructions[instructions[index+3]] = instructions[instructions[index+1]] + instructions[instructions[index+2]]
		break
	case 2:
		instructions[instructions[index+3]] = instructions[instructions[index+1]] * instructions[instructions[index+2]]
		break
	case 99:
		// Halt instruction
		break
	default:
		log.Printf("[ERROR] Unsupported instruction encountered: %v", instructions[index])
	}
	return instructions
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	// provide 1 as the only input
	// support op codes with 5 digits
	// support immediate mode
	// support op code 3: input
	// support op code 4: output

	input = parseInstructions(input)

	return input[0]
}

func parseInstructions(input []int) []int {
	for i := 0; i < len(input); i += 4 {
		if input[i] == 99 {
			break
		}
		if i%4 == 0 {
			input = compute(input, i)
		}
	}
	return input
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	for noun := 0; noun < len(input); noun++ {
		for verb := 0; verb < len(input); verb++ {
			temporaryInstructions := slices.Clone(input)
			temporaryInstructions[1] = noun
			temporaryInstructions[2] = verb
			parseInstructions(temporaryInstructions)
			if temporaryInstructions[0] == 19690720 {
				log.Print("[CONSOLE] The noun and verb that cause the program to produce the output 19690720:\n")
				log.Printf("[CONSOLE] Noun: %v\n", noun)
				log.Printf("[CONSOLE] Verb: %v\n", verb)

				return 100*noun + verb
			}
		}
	}

	return -1
}
