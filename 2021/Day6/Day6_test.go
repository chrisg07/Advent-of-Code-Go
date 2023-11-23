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
		MinLevel: logutils.LogLevel("DEBUG"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
}

func TestDay6PartA2021ExampleSmall(t *testing.T) {
	answer := 26
	solution := Day62021(true, 18)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay6PartA2021Example(t *testing.T) {
	answer := 5934
	solution := Day62021(true, 80)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay6PartA2021Complete(t *testing.T) {
	answer := 375482
	solution := Day62021(false, 80)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestDay6PartB2021Example(t *testing.T) {
	answer := 26984457539
	solution := Day62021(true, 256)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay6PartBComplete(t *testing.T) {
	answer := 1689540415957
	solution := Day62021(false, 256)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
