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

func TestDay12PartA2021Example(t *testing.T) {
	answer := 19
	solution := Day12PartA2021(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay12PartA2021Complete(t *testing.T) {
	answer := 4912
	solution := Day12PartA2021(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestDay12PartB2021Example(t *testing.T) {
	answer := 103
	solution := Day12PartB2021(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay12PartB2021Complete(t *testing.T) {
	answer := 150004
	solution := Day12PartB2021(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
