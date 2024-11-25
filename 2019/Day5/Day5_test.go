package AoCScaffold

import (
	"bytes"
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR", "CONSOLE"},
		MinLevel: logutils.LogLevel("DEBUG"),
		Writer:   os.Stderr,
	}
	log.SetFlags(0)
	log.SetOutput(filter)
	log.Print("[CONSOLE] --------------------------\n")
	log.Print("[CONSOLE] Advent of Code 2019 Day 5:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

func TestInputInstruction(t *testing.T) {
	instructions := []int{3, 0, 99}

	// Create a pipe to mock os.Stdin
	reader, writer, _ := os.Pipe()
	defer reader.Close()
	defer writer.Close()

	// Backup the original Stdin and defer its restoration
	originalStdin := os.Stdin
	defer func() { os.Stdin = originalStdin }()

	// Replace os.Stdin with our pipe reader
	os.Stdin = reader

	// Write the mock input to the writer end of the pipe
	go func() {
		writer.Write([]byte("1337\n"))
		writer.Close()
	}()

	answer := []int{1337, 0, 99}
	solution, index := compute(instructions, 0)
	if !reflect.DeepEqual(answer, solution) || index != 2 {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestOutputInstruction(t *testing.T) {
	instructions := []int{4, 50, 99}

	var str bytes.Buffer
	log.SetOutput(&str)
	expectedOutput := "[CONSOLE] Output: 50\n"
	_, index := compute(instructions, 0)

	// Read the output
	output := str.String()
	if output != expectedOutput || index != 2 {
		t.Fatalf("Expected %q but got %q", expectedOutput, output)
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
	answer := 1
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Diagnostic code: %v", solution)
	}
}

// func TestPartBExample(t *testing.T) {
// 	answer := 1
// 	solution := PartB(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestPartBComplete(t *testing.T) {
// 	answer := 1
// 	solution := PartB(false)
// 	if solution != answer {
// 		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
// 	} else {
// 		log.Printf("[CONSOLE] Find the Elf carrying the most Calories: %v", solution)
// 	}
// }
