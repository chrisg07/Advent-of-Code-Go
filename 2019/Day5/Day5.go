package AoC2019

import (
	"bufio"
	_ "embed"
	"fmt"
	"log"
	"os"
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

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	cleanup, _ := Utils.MockStdin("1\n")
	defer cleanup()

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

func compute(instructions []int, index int) ([]int, int) {
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

		input = strings.TrimSpace(input)
		inputValue, _ := strconv.Atoi(input)

		log.Printf("[DEBUG] Input %v and stored it at address %v", input, instructions[index+1])
		instructions[instructions[index+1]] = inputValue
		index += 2
	case 4:
		log.Printf("[DEBUG] Output: %v\n", instructions[instructions[index+1]])
		index += 2
	case 5:
		// if the first parameter is non-zero, it sets the instruction pointer to the value from the second parameter. Otherwise, it does nothing.
		parameter1 := getParameter(mode1, instructions, index+1)
		parameter2 := getParameter(mode2, instructions, index+2)

		if parameter1 != 0 {
			index = parameter2
			log.Printf("[DEBUG] Moved instruction pointer to address %v", index)
		} else {
			index += 3
		}
	case 6:
		parameter1 := getParameter(mode1, instructions, index+1)
		parameter2 := getParameter(mode2, instructions, index+2)

		if parameter1 == 0 {
			index = parameter2
			log.Printf("[DEBUG] Moved instruction pointer to address %v", index)
		} else {
			index += 3
		}
	case 7:
		parameter1 := getParameter(mode1, instructions, index+1)
		parameter2 := getParameter(mode2, instructions, index+2)
		parameter3 := getParameter(0, instructions, index+3)

		if parameter1 < parameter2 {
			instructions[instructions[index+3]] = 1
			log.Printf("[DEBUG] Parameter 1 (%v) is less than parameter 2 (%v). Storing 1 at %v", parameter1, parameter2, parameter3)
		} else {
			instructions[instructions[index+3]] = 0
			log.Printf("[DEBUG] Parameter 1 (%v) is not less than parameter 2 (%v). Storing 0 at %v", parameter1, parameter2, parameter3)
		}
		index += 4
	case 8:
		parameter1 := getParameter(mode1, instructions, index+1)
		parameter2 := getParameter(mode2, instructions, index+2)
		parameter3 := getParameter(0, instructions, index+3)

		if parameter1 == parameter2 {
			instructions[instructions[index+3]] = 1
			log.Printf("[DEBUG] Parameter 1 (%v) is equal to parameter 2 (%v). Storing 1 at %v", parameter1, parameter2, parameter3)
		} else {
			instructions[instructions[index+3]] = 0
			log.Printf("[DEBUG] Parameter 1 (%v) is not equal to parameter 2 (%v). Storing 0 at %v", parameter1, parameter2, parameter3)
		}
		index += 4
	case 99:
		// Halt instruction
	default:
		log.Fatalf("[ERROR] Unsupported instruction encountered: %v", instructions[index])
	}
	return instructions, index
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	cleanup, _ := Utils.MockStdin("5\n")
	defer cleanup()

	diagnosticCode := parseInstructions(input)

	return diagnosticCode
}
