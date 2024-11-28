package AoC2019

import (
	_ "embed"
	"log"
	"math"
	"strconv"
	"strings"

	Utils "github.com/chrisg07/Advent-of-Code-Go/Utils"
)

//go:embed inputs/example.txt
var exampleInput string

//go:embed inputs/complete.txt
var input string

var inputQueue = make([]int, 0)
var outputQueue = make([]int, 0)

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

func ParseOpcode(fullCode int) int {
	return fullCode % 100
}

func ParseModes(fullCode int) (mode1 int, mode2 int, mode3 int) {
	fullCode /= 100

	// extract modes
	mode1 = fullCode % 10
	fullCode /= 10
	mode2 = fullCode % 10
	fullCode /= 10
	mode3 = fullCode % 10

	return mode1, mode2, mode3
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
	opcode := ParseOpcode(instructions[index])

	switch opcode {
	case 1:
		add(instructions, index)
		index += 4
	case 2:
		multiply(instructions, index)
		index += 4
	case 3:
		recieveInput(instructions, index)
		index += 2
	case 4:
		log.Printf("[DEBUG] Output: %v\n", instructions[instructions[index+1]])
		outputQueue = append(outputQueue, instructions[instructions[index+1]])
		index += 2
	case 5:
		index = jumpIfTrue(instructions, index)
	case 6:
		index = jumpIfFalse(instructions, index)
	case 7:
		lessThan(instructions, index)
		index += 4
	case 8:
		equals(instructions, index)
		index += 4
	case 99:
		// Halt instruction
	default:
		log.Fatalf("[ERROR] Unsupported instruction encountered: %v", instructions[index])
	}
	return instructions, index
}

func getParameter(instructions []int, instructionIndex int, parameterIndex int) int {
	mode := (instructions[instructionIndex] / int(math.Pow(10, float64(parameterIndex)+1)))
	mode %= 10

	if mode == 0 || parameterIndex == 3 {
		return instructions[instructions[instructionIndex+parameterIndex]]
	} else {
		return instructions[instructionIndex+parameterIndex]
	}
}

func equals(instructions []int, index int) {
	parameter1 := getParameter(instructions, index, 1)
	parameter2 := getParameter(instructions, index, 2)
	parameter3 := getParameter(instructions, index, 3)

	if parameter1 == parameter2 {
		instructions[instructions[index+3]] = 1
		log.Printf("[DEBUG] Parameter 1 (%v) is equal to parameter 2 (%v). Storing 1 at %v", parameter1, parameter2, parameter3)
	} else {
		instructions[instructions[index+3]] = 0
		log.Printf("[DEBUG] Parameter 1 (%v) is not equal to parameter 2 (%v). Storing 0 at %v", parameter1, parameter2, parameter3)
	}
}

func recieveInput(instructions []int, index int) {
	input := inputQueue[0]
	inputQueue = inputQueue[1:]

	log.Printf("[DEBUG] Input %v and stored it at address %v", input, instructions[index+1])
	instructions[instructions[index+1]] = input
}

func add(instructions []int, index int) {
	parameter1 := getParameter(instructions, index, 1)
	parameter2 := getParameter(instructions, index, 2)

	log.Printf("[DEBUG] Adding %v + %v and storing it at address %v", parameter1, parameter2, instructions[index+3])
	instructions[instructions[index+3]] = parameter1 + parameter2
}

func multiply(instructions []int, index int) {
	parameter1 := getParameter(instructions, index, 1)
	parameter2 := getParameter(instructions, index, 2)

	log.Printf("[DEBUG] Multiplying %v * %v and storing it at address %v", parameter1, parameter2, instructions[index+3])
	instructions[instructions[index+3]] = parameter1 * parameter2
}

func jumpIfTrue(instructions []int, index int) int {
	parameter1 := getParameter(instructions, index, 1)
	parameter2 := getParameter(instructions, index, 2)

	if parameter1 != 0 {
		index = parameter2
		log.Printf("[DEBUG] Moved instruction pointer to address %v", index)
	} else {
		index += 3
	}
	return index
}

func jumpIfFalse(instructions []int, index int) int {
	parameter1 := getParameter(instructions, index, 1)
	parameter2 := getParameter(instructions, index, 2)

	if parameter1 == 0 {
		index = parameter2
		log.Printf("[DEBUG] Moved instruction pointer to address %v", index)
	} else {
		index += 3
	}
	return index
}

func lessThan(instructions []int, index int) {
	parameter1 := getParameter(instructions, index, 1)
	parameter2 := getParameter(instructions, index, 2)
	parameter3 := getParameter(instructions, index, 3)

	if parameter1 < parameter2 {
		instructions[instructions[index+3]] = 1
		log.Printf("[DEBUG] Parameter 1 (%v) is less than parameter 2 (%v). Storing 1 at %v", parameter1, parameter2, parameter3)
	} else {
		instructions[instructions[index+3]] = 0
		log.Printf("[DEBUG] Parameter 1 (%v) is not less than parameter 2 (%v). Storing 0 at %v", parameter1, parameter2, parameter3)
	}
}

func calculateAmplifierOutput(program []int) int {
	log.Printf("[CONSOLE] Amplifier input: %v, %v", inputQueue[0], inputQueue[1])

	parseInstructions(program)
	output := outputQueue[0]
	outputQueue = outputQueue[1:]

	log.Printf("[CONSOLE] Amplifier output: %v", output)
	return output
}

func calculateThrusterSignal(program []int, phaseSetting []int) int {
	inputQueue = append(inputQueue, phaseSetting[0])
	inputQueue = append(inputQueue, 0)
	amplifierA := calculateAmplifierOutput(program)

	inputQueue = append(inputQueue, phaseSetting[1])
	inputQueue = append(inputQueue, amplifierA)
	amplifierB := calculateAmplifierOutput(program)

	inputQueue = append(inputQueue, phaseSetting[2])
	inputQueue = append(inputQueue, amplifierB)
	amplifierC := calculateAmplifierOutput(program)

	inputQueue = append(inputQueue, phaseSetting[3])
	inputQueue = append(inputQueue, amplifierC)
	amplifierD := calculateAmplifierOutput(program)

	inputQueue = append(inputQueue, phaseSetting[4])
	inputQueue = append(inputQueue, amplifierD)
	amplifierE := calculateAmplifierOutput(program)

	return amplifierE
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	maxThrusterSignal := 0
	for a := 0; a < 5; a++ {
		for b := 0; b < 5; b++ {
			for c := 0; c < 5; c++ {
				for d := 0; d < 5; d++ {
					for e := 0; e < 5; e++ {
						permutation := make(map[int]bool)
						permutation[a] = true
						permutation[b] = true
						permutation[c] = true
						permutation[d] = true
						permutation[e] = true
						log.Printf("[CONSOLE] Permutation: %v with length %v", permutation, len(permutation))

						if len(permutation) == 5 {
							thrusterSignal := calculateThrusterSignal(input, []int{a, b, c, d, e})
							if thrusterSignal > maxThrusterSignal {
								maxThrusterSignal = thrusterSignal
							}
						}
					}
				}
			}
		}
	}

	return maxThrusterSignal
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	cleanup, _ := Utils.MockStdin("5\n")
	defer cleanup()

	diagnosticCode := parseInstructions(input)
	return diagnosticCode
}
