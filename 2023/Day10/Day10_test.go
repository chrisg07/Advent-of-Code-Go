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

func TestDay10PartA2023Complete(t *testing.T) {
	answer := 6909
	solution := Day10PartA2023(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

// func TestDay10PartB2023Example(t *testing.T) {
// 	answer := 4
// 	solution := Day10PartB2023(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDay10PartB2023Complete(t *testing.T) {
// 	answer := 0
// 	solution := Day10PartB2023(false)
// 	if solution >= 721 {
// 		t.Fatalf(`Answer is too high`)
// 	}
// 	if solution != answer {
// 		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
// 	}
// }
