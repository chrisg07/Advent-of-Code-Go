package AoC2023

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("ERROR"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
}

// func TestDay3PartA2023Example(t *testing.T) {
// 	answer := 4361
// 	solution := Day3PartA2023(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestDay3PartA2023Complete(t *testing.T) {
	answer := 540212
	solution := Day3PartA2023(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

// func TestDay3PartB2023Example(t *testing.T) {
// 	answer := 467835
// 	solution := Day3PartB2023(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestDay3PartB2023Complete(t *testing.T) {
	answer := 87605697
	solution := Day3PartB2023(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
