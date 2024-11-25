package AoC2019

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"strings"
	"testing"

	Utils "github.com/chrisg07/Advent-of-Code-Go/Utils"
	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR", "CONSOLE"},
		MinLevel: logutils.LogLevel("ERROR"),
		Writer:   os.Stderr,
	}
	log.SetFlags(0)
	log.SetOutput(filter)
	log.Print("[CONSOLE] --------------------------\n")
	log.Print("[CONSOLE] Advent of Code 2019 Day 5:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

func TestInputOutputInstruction(t *testing.T) {
	instructions := []int{3, 0, 4, 0, 99}

	cleanup, _ := Utils.MockStdin("1337\n")
	defer cleanup()

	var str bytes.Buffer
	log.SetOutput(&str)
	expectedOutput := "[CONSOLE] Output: 1337"
	parseInstructions(instructions)
	output := str.String()
	if strings.Contains(output, expectedOutput) {
		t.Fatalf(`Expected output "%v", actual = "%v"`, expectedOutput, output)
	}
}

func TestOpcodeParsing(t *testing.T) {
	opcode := 1002
	actualOpcode, mode1, mode2, mode3 := ParseOpcode(opcode)
	if actualOpcode != 2 || mode1 != 0 || mode2 != 1 || mode3 != 0 {
		t.Fatalf(`Opcode was not parsed correctly`)
	}
}
func TestImmediateMode(t *testing.T) {
	instructions := []int{1002, 4, 3, 4, 33}
	answer := []int{1002, 4, 3, 4, 99}
	solution, _ := compute(instructions, 0)
	if !reflect.DeepEqual(answer, solution) {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
	instructions = []int{1101, 100, -1, 4, 0}
	answer = []int{1101, 100, -1, 4, 99}
	solution, _ = compute(instructions, 0)
	if !reflect.DeepEqual(answer, solution) {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartAComplete(t *testing.T) {
	answer := 13547311
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] the diagnostic code for system ID 1: %v", solution)
	}
}

func TestJumpInstructionUsingPosition(t *testing.T) {
	// output should be 0 if 0 is provided as input
	instructions := []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	cleanup, _ := Utils.MockStdin("0\n")
	defer cleanup()
	solution := parseInstructions(instructions)
	if solution != 0 {
		t.Fatalf(`Example output to be 0`)
	}

	// output should be 1 if the input is non-zero
	instructions = []int{3, 12, 6, 12, 15, 1, 13, 14, 13, 4, 13, 99, -1, 0, 1, 9}
	cleanup, _ = Utils.MockStdin("1\n")
	defer cleanup()
	solution = parseInstructions(instructions)
	if solution != 1 {
		t.Fatalf(`Example output to be 1`)
	}
}

func TestJumpInstructionUsingImmediate(t *testing.T) {
	// output should be 0 if 0 is provided as input
	instructions := []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	cleanup, _ := Utils.MockStdin("0\n")
	defer cleanup()
	solution := parseInstructions(instructions)
	if solution != 0 {
		t.Fatalf(`Expected example output to be 0`)
	}

	// output should be the 1 if the provided input is non-zero
	instructions = []int{3, 3, 1105, -1, 9, 1101, 0, 0, 12, 4, 12, 99, 1}
	cleanup, _ = Utils.MockStdin("1\n")
	defer cleanup()
	solution = parseInstructions(instructions)
	if solution != 1 {
		t.Fatalf(`Expected example output to be 1`)
	}
}

func TestLessThanInstructionUsingPosition(t *testing.T) {
	// consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not)
	instructions := []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	cleanup, _ := Utils.MockStdin("0\n")
	defer cleanup()
	solution := parseInstructions(instructions)
	if solution != 1 {
		t.Fatalf(`Example output to be 1 was %v`, solution)
	}

	instructions = []int{3, 9, 7, 9, 10, 9, 4, 9, 99, -1, 8}
	cleanup, _ = Utils.MockStdin("8\n")
	defer cleanup()
	solution = parseInstructions(instructions)
	if solution != 0 {
		t.Fatalf(`Example output to be 0`)
	}
}

func TestLessThanInstructionUsingImmediate(t *testing.T) {
	// consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not)
	instructions := []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	cleanup, _ := Utils.MockStdin("0\n")
	defer cleanup()
	solution := parseInstructions(instructions)
	if solution != 1 {
		t.Fatalf(`Example output to be 1 was %v`, solution)
	}

	instructions = []int{3, 3, 1107, -1, 8, 3, 4, 3, 99}
	cleanup, _ = Utils.MockStdin("8\n")
	defer cleanup()
	solution = parseInstructions(instructions)
	if solution != 0 {
		t.Fatalf(`Example output to be 0`)
	}
}

func TestEqualityInstructionUsingPosition(t *testing.T) {
	// consider whether the input is less than 8; output 1 (if it is) or 0 (if it is not)
	instructions := []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	cleanup, _ := Utils.MockStdin("0\n")
	defer cleanup()
	solution := parseInstructions(instructions)
	if solution != 0 {
		t.Fatalf(`Example output to be 0 was %v`, solution)
	}

	instructions = []int{3, 9, 8, 9, 10, 9, 4, 9, 99, -1, 8}
	cleanup, _ = Utils.MockStdin("8\n")
	defer cleanup()
	solution = parseInstructions(instructions)
	if solution != 1 {
		t.Fatalf(`Example output to be 1`)
	}
}

func TestPartBComplete(t *testing.T) {
	answer := 236453
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] The diagnostic code for system ID 5: %v", solution)
	}
}
