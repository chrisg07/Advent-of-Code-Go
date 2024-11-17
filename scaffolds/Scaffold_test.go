package AoCScaffold

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

func TestPartAExample(t *testing.T) {
	answer := 0
	solution := PartA(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartAComplete(t *testing.T) {
	answer := 0
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestPartBExample(t *testing.T) {
	answer := 0
	solution := PartB(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartBComplete(t *testing.T) {
	answer := 0
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
