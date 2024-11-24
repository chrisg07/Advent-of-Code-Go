package AoCScaffold

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR", "CONSOLE"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	log.SetFlags(0)
	log.SetOutput(filter)
	log.Print("[CONSOLE] --------------------------\n")
	log.Print("[CONSOLE] Advent of Code 2019 Day 5:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

func TestPartASmallPrograms(t *testing.T) {
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
	solution := compute(instructions, 0)
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

func TestPartBExample(t *testing.T) {
	answer := 1
	solution := PartB(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartBComplete(t *testing.T) {
	answer := 1
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Find the Elf carrying the most Calories: %v", solution)
	}
}
