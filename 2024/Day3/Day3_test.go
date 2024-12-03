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
	log.Print("[CONSOLE] Advent of Code 2024 Day 3:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

func TestPartAComplete(t *testing.T) {
	answer := 190604937
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Result of all multiplications: %v", solution)
	}
}

func TestPartBComplete(t *testing.T) {
	answer := 82857512
	solution := PartB(false)
	if solution != answer || solution == 87020895 {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Result of all enabled multiplications: %v", solution)
	}
}
