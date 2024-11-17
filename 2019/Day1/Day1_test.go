package AoC2019

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
	log.Print("[CONSOLE] Advent of Code 2019 Day 1:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

func TestFuelCalculation(t *testing.T) {
	answer := 2
	solution := calculateFuelCost(12)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
	answer = 2
	solution = calculateFuelCost(14)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
	answer = 654
	solution = calculateFuelCost(1969)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
	answer = 33583
	solution = calculateFuelCost(100756)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartAComplete(t *testing.T) {
	answer := 3262991
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] The sum of the fuel requirements for all the modules on the spacecraft: %v", solution)
	}
}

func TestRecursiveFuelCalculation(t *testing.T) {
	answer := 2
	solution := calculateTotalFuelCost(14)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
	answer = 966
	solution = calculateTotalFuelCost(1969)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
	answer = 50346
	solution = calculateTotalFuelCost(100756)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

// func TestPartBExample(t *testing.T) {
// 	answer := 1
// 	solution := PartB(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestPartBComplete(t *testing.T) {
	answer := 4891620
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] The sum of the fuel requirements for all of the modules on the spacecraft when also taking into account the mass of the added fuel: %v", solution)
	}
}
