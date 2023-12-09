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

func TestDay9PartA2023Example(t *testing.T) {
	answer := 114
	solution := Day9PartA2023(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay9PartA2023Complete(t *testing.T) {
	answer := 1887980197
	solution := Day9PartA2023(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

// func TestDay9PartB2023Example(t *testing.T) {
// 	answer := 0
// 	solution := Day9PartB2023(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDay9PartB2023Complete(t *testing.T) {
// 	answer := 0
// 	solution := Day9PartB2023(false)
// 	if solution != answer {
// 		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
// 	}
// }
