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

func TestCalculateFueld(t *testing.T) {
	answer := 71
	crabs := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	solution := CalculateFuel(crabs, 10)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay7PartA2021Example(t *testing.T) {
	answer := 37
	solution := Day7PartA2021(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay7PartA2021Complete(t *testing.T) {
	answer := 37
	solution := Day7PartA2021(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

// func TestDay7PartB2021Example(t *testing.T) {
// 	answer := 0
// 	solution := Day7PartB2021(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDay7PartBComplete(t *testing.T) {
// 	answer := 0
// 	solution := Day7PartB2021(false)
// 	if solution != answer {
// 		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
// 	}
// }
