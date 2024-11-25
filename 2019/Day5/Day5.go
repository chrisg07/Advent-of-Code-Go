package AoC2019

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
	"slices"
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

func getParameter(mode1 int, instructions []int, index int) int {
	parameter := 0
	if mode1 == 0 {
		parameter = instructions[instructions[index]]
	} else {
		parameter = instructions[index]
	}
	return parameter
}

func ParseOpcode(fullCode int) (opcode int, mode1 int, mode2 int, mode3 int) {
	// extract opcode
	opcode = fullCode % 100
	fullCode /= 100

	// extract modes
	mode1 = fullCode % 10
	fullCode /= 10
	mode2 = fullCode % 10
	fullCode /= 10
	mode3 = fullCode % 10

	return opcode, mode1, mode2, mode3
}

func compute(instructions []int, index int) ([]int, int) {
	// take instruction at index and extract op code and parameter modes, if any
	opcode, mode1, mode2, _ := ParseOpcode(instructions[index])

	switch opcode {
	case 1:
		parameter1 := getParameter(mode1, instructions, index+1)
		parameter2 := getParameter(mode2, instructions, index+2)

		log.Printf("[DEBUG] Adding %v + %v and storing it at address %v", parameter1, parameter2, instructions[index+3])
		instructions[instructions[index+3]] = parameter1 + parameter2
		index += 4
	case 2:
		parameter1 := getParameter(mode1, instructions, index+1)
		parameter2 := getParameter(mode2, instructions, index+2)

		log.Printf("[DEBUG] Multiplying %v * %v and storing it at address %v", parameter1, parameter2, instructions[index+3])

		instructions[instructions[index+3]] = parameter1 * parameter2
		index += 4
	case 3:
		// Create a new reader to read input from the standard input
		reader := bufio.NewReader(os.Stdin)
		log.Println("Enter the input instruction: ")

		// Read input until the user presses Enter
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Println("Error reading input:", err)
		}

		// Trim the newline character from the input
		input = strings.TrimSpace(input)
		inputValue, _ := strconv.Atoi(input)

		log.Printf("[DEBUG] Input %v and stored it at address %v", input, instructions[index+1])
		instructions[instructions[index+1]] = inputValue
		index += 2
	case 4:
		log.Printf("[DEBUG] Output: %v\n", instructions[instructions[index+1]])
		index += 2
	case 99:
		// Halt instruction
	default:
		log.Fatalf("[ERROR] Unsupported instruction encountered: %v", instructions[index])
	}
	return instructions, index
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	cleanup, _ := Utils.MockStdin("1\n")
	defer cleanup() // Ensure cleanup is called

	diagnosticCode := parseInstructions(input)

	return diagnosticCode
}

func parseInstructions(input []int) int {
	index := 0
	for index < len(input) {
		if input[index] == 99 {
			return input[input[index-1]]
		}
		input, index = compute(input, index)
	}
	return -1
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
