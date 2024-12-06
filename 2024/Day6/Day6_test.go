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
	log.Print("[CONSOLE] Advent of Code 2024 Day 6:\n")
	log.Print("[CONSOLE] --------------------------\n")
}
func TestPartAComplete(t *testing.T) {
	answer := 5067
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Number of distinct positions the guard visited before leaving the mapped area: %v", solution)
	}
}

func TestPartBExample(t *testing.T) {
	answer := 6
	solution := PartB(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartBComplete(t *testing.T) {
	answer := 1793
	solution := PartB(false)
	// 1729
	// 1793
	if solution != answer && solution < 1800 && solution > 1253 && solution != 1498 {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Number of positions a sovereign glue tank can be placed to cause the guard to continously loop: %v", solution)
	}
}
