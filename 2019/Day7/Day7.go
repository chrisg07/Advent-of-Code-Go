package AoC2019

import (
	_ "embed"
	"log"
	"math"
	"strconv"
	"strings"
)

//go:embed inputs/example.txt
var exampleInput string

//go:embed inputs/complete.txt
var input string

var inputQueue = make([]int, 0)

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

func parseInstructions(input []int, inputs []int, outputs *[]int) int {
	index := 0
	for index < len(input) {
		if input[index] == 99 {
			return input[input[index-1]]
		}
		input, index = compute(input, index, &inputs, outputs)
	}
	return -1
}

func compute(instructions []int, index int, inputs *[]int, outputs *[]int) ([]int, int) {
	opcode := ParseOpcode(instructions[index])

	switch opcode {
	case 1:
		add(instructions, index)
		index += 4
	case 2:
		multiply(instructions, index)
		index += 4
	case 3:
		recieveInput(instructions, index, inputs)
		index += 2
	case 4:
		output := instructions[instructions[index+1]]
		log.Printf("[DEBUG] Output: %v\n", output)
		*outputs = append(*outputs, output)
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

func recieveInput(instructions []int, index int, inputs *[]int) {
	for len(*inputs) == 0 {
		// log.Fatalf("[ERROR] No items in input queue at instruction index %v\n", index)
	}

	log.Printf("[DEBUG] Inputs before consumption: %v\n", *inputs)
	input := (*inputs)[0]
	*inputs = (*inputs)[1:]
	log.Printf("[DEBUG] Inputs after consumption: %v\n", *inputs)
	log.Printf("[DEBUG] Input %v and stored it at address %v\n", input, instructions[index+1])
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
		log.Printf("[DEBUG] Moved instruction pointer to address %v\n", index)
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
		log.Printf("[DEBUG] Moved instruction pointer to address %v\n", index)
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
		log.Printf("[DEBUG] Parameter 1 (%v) is less than parameter 2 (%v). Storing 1 at %v\n", parameter1, parameter2, parameter3)
	} else {
		instructions[instructions[index+3]] = 0
		log.Printf("[DEBUG] Parameter 1 (%v) is not less than parameter 2 (%v). Storing 0 at %v\n", parameter1, parameter2, parameter3)
	}
}

func calculateAmplifierOutput(program []int, inputs []int, outputs *[]int) int {
	log.Printf("[DEBUG] Amplifier started with inputs: %v\n", inputs)
	parseInstructions(program[:], inputs[:], outputs)
	if len(*outputs) == 0 {
		log.Fatalf("[ERROR] No output generated by amplifier\n")
	}
	output := (*outputs)[0]
	*outputs = (*outputs)[1:]
	log.Printf("[DEBUG] Amplifier finished with output: %v\n", output)
	return output
}

func cloneArray(array []int) []int {
	clone := make([]int, len(array))
	copy(clone, array)
	return clone
}

func calculateThrusterSignal(program []int, phaseSetting []int, amplifierAInput int) int {
	outputQueueE := []int{phaseSetting[0], amplifierAInput}
	outputQueueA := []int{phaseSetting[1]}
	programA := cloneArray(program)
	calculateAmplifierOutput(programA, outputQueueE, &outputQueueA)

	outputQueueB := []int{phaseSetting[2]}
	programB := cloneArray(program)
	calculateAmplifierOutput(programB, outputQueueA, &outputQueueB)

	outputQueueC := []int{phaseSetting[3]}
	programC := cloneArray(program)
	calculateAmplifierOutput(programC, outputQueueB, &outputQueueC)

	outputQueueD := []int{phaseSetting[4]}
	programD := cloneArray(program)
	calculateAmplifierOutput(programD, outputQueueC, &outputQueueD)

	programE := cloneArray(program)
	calculateAmplifierOutput(programE, outputQueueD, &outputQueueE)

	// last output from E is the most recently added input for A
	return outputQueueE[0]
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
						log.Printf("[DEBUG] Permutation: %v with length %v", permutation, len(permutation))

						if len(permutation) == 5 {
							thrusterSignal := calculateThrusterSignal(input, []int{a, b, c, d, e}, 0)
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

	maxThrusterSignal := 0
	for f := 5; f < 10; f++ {
		for g := 5; g < 10; g++ {
			for h := 5; h < 10; h++ {
				for i := 5; i < 10; i++ {
					for j := 5; j < 10; j++ {
						permutation := make(map[int]bool)
						permutation[f] = true
						permutation[g] = true
						permutation[h] = true
						permutation[i] = true
						permutation[j] = true
						log.Printf("[DEBUG] Permutation: %v with length %v", permutation, len(permutation))

						if len(permutation) == 5 {
							// can't stop calculation
							// each amplifier needs it's own input and output queues
							thrusterSignal := calculateThrusterSignal(input, []int{f, g, h, i, j}, 0)
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
