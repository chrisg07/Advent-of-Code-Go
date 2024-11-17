package AoC2022

import (
	"log"
	"os"
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
	log.Print("[CONSOLE] Advent of Code 2022 Day 1:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

// func TestPartAExample(t *testing.T) {
// 	answer := 24000
// 	solution := PartA(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestPartAComplete(t *testing.T) {
	answer := 69912
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Find the Elf carrying the most Calories: %v", solution)
	}
}

// func TestPartBExample(t *testing.T) {
// 	answer := 45000
// 	solution := PartB(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestPartBComplete(t *testing.T) {
	answer := 208180
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Find the top three Elves carrying the most Calories: %v", solution)
	}
}
