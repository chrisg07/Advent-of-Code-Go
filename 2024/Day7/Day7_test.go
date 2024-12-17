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
		MinLevel: logutils.LogLevel("DEBUG"),
		Writer:   os.Stderr,
	}
	log.SetFlags(0)
	log.SetOutput(filter)
	log.Print("[CONSOLE] --------------------------\n")
	log.Print("[CONSOLE] Advent of Code 2024 Day 7:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

func TestInputParsing(t *testing.T) {
	example := PartA(true)
	complete := PartA(false)
	if reflect.DeepEqual(example, complete) {
		t.Fatalf(`Example solution = %d, should = %d`, example, complete)
	}
}

func TestPartAExample(t *testing.T) {
	answer := uint64(3749)
	solution := PartA(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartAComplete(t *testing.T) {
	answer := uint64(1)
	solution := PartA(false)
	if solution != answer || solution <= 661987571 || solution == 28919153446 {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Total calibration result: %v", solution)
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
