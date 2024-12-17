package AoCScaffold

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR", "CONSOLE"},
		MinLevel: logutils.LogLevel("CONSOLE"),
		Writer:   os.Stderr,
	}
	log.SetFlags(0)
	log.SetOutput(filter)
	log.Print("[CONSOLE] --------------------------\n")
	log.Print("[CONSOLE] Advent of Code 2024 Day 7:\n")
	log.Print("[CONSOLE] --------------------------\n")
}
func TestConcatenateInts(t *testing.T) {
	expected := uint64(111222)
	actual := concatenateInts(111, 222)
	if actual != expected {
		t.Fatalf(`Actual solution = %d, expected = %d`, actual, expected)
	}
}

func TestConcatenateIntsExample(t *testing.T) {
	expected := uint64(156)
	actual := concatenateInts(15, 6)
	if actual != expected {
		t.Fatalf(`Actual solution = %d, expected = %d`, actual, expected)
	}
}

func TestPartBComplete(t *testing.T) {
	answer := uint64(328790210468594)
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Total calibration result: %v", solution)
	}
}
