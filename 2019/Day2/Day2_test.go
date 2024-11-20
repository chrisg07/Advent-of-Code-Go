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
	log.Print("[CONSOLE] Advent of Code 2019 Day 2:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

func TestPartASmallPrograms(t *testing.T) {
	instructions := []int{1, 0, 0, 0, 99}
	answer := []int{2, 0, 0, 0, 99}
	solution := compute(instructions, 0)
	if !reflect.DeepEqual(answer, solution) {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
	instructions = []int{2, 3, 0, 3, 99}
	answer = []int{2, 3, 0, 6, 99}
	solution = compute(instructions, 0)
	if !reflect.DeepEqual(answer, solution) {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
	instructions = []int{2, 4, 4, 5, 99, 0}
	answer = []int{2, 4, 4, 5, 99, 9801}
	solution = compute(instructions, 0)
	if !reflect.DeepEqual(answer, solution) {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
	instructions = []int{1, 1, 1, 4, 99, 5, 6, 0, 99}
	answer = []int{30, 1, 1, 4, 2, 5, 6, 0, 99}
	solution = compute(instructions, 0)
	solution = compute(solution, 4)
	if !reflect.DeepEqual(answer, solution) {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartAComplete(t *testing.T) {
	answer := 3101878
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] The value left at position 0 after the program halts is: %v", solution)
	}
}

func TestPartBComplete(t *testing.T) {
	answer := 8444
	solution := PartB(false)
	if solution != answer || solution == 2722500 || solution == 369600 {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
