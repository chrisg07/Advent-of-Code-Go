package AoC2021

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
}

func TestDay4PartA2023Example(t *testing.T) {
	answer := 13
	solution := Day4PartA2023(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay4PartA2023Complete(t *testing.T) {
	answer := 33950
	solution := Day4PartA2023(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

// func TestDay4PartB2023Example(t *testing.T) {
// 	answer := 0
// 	solution := Day4PartB2023(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDay4PartB2023Complete(t *testing.T) {
// 	answer := 0
// 	solution := Day4PartB2023(false)
// 	if solution != answer {
// 		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
// 	}
// }
