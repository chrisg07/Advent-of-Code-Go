package AoC2019

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
	log.Print("[CONSOLE] Advent of Code 2019 Day 3:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

// test that we can convert instructions to coordinates
func TestConvertInstructionsToPoints(t *testing.T) {
	answer := []Point{Point{0, 0}, Point{8, 0}}
	solution := convertInstructionsToPoints("R8")
	if !reflect.DeepEqual(answer, solution) {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

// test that we can determine whether two lines intersect

// test that we can determine where two lines intersect

// test that if there are multiple intersections they are all reported
func TestPartAExample(t *testing.T) {
	answer := 1
	solution := PartA(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartAComplete(t *testing.T) {
	answer := 1
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Find the Elf carrying the most Calories: %v", solution)
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
