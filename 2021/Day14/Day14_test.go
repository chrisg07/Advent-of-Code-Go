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
		MinLevel: logutils.LogLevel("ERROR"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
}

// func TestDay14PartA2021Example(t *testing.T) {
// 	answer := 1588
// 	solution := Day142021(true, 10)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestDay14PartA2021Complete(t *testing.T) {
	answer := 2621
	solution := Day142021(false, 10)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

// func TestDay14PartB2021Example(t *testing.T) {
// 	answer := 2188189693529
// 	solution := Day142021(true, 40)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestDay14PartB2021Complete(t *testing.T) {
	answer := 2843834241366
	solution := Day142021(false, 40)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
