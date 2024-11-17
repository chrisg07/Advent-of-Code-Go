package AoC2023

import (
	"log"
	"os"

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

// func TestDay13PartA2023Example(t *testing.T) {
// 	answer := 405
// 	solution := Day13PartA2023(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDay13PartA2023Complete(t *testing.T) {
// 	answer := 405
// 	solution := Day13PartA2023(false)
// 	if solution != answer {
// 		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDay13PartB2023Example(t *testing.T) {
// 	answer := 0
// 	solution := Day13PartB2023(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDay13PartB2023Complete(t *testing.T) {
// 	answer := 0
// 	solution := Day13PartB2023(false)
// 	if solution != answer {
// 		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
// 	}
// }
