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

func TestDay8PartA2023Complete(t *testing.T) {
	answer := 20569
	solution := Day8PartA2023(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestDay8PartB2023Complete(t *testing.T) {
	answer := 21366921060721
	solution := Day8PartB2023(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
