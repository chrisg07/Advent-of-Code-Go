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

func TestParseDigitFromString(t *testing.T) {
	answer := "8"
	solution := ParseDigitFromString("8ndowt")
	if solution != answer {
		t.Fatalf(`Example solution = %s, should = %s`, solution, answer)
	}
}

func TestParseKeyFromString(t *testing.T) {
	answer := 15
	solution := ParseKeyFromString("15195one")
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay1PartA2021Example(t *testing.T) {
	answer := 142
	solution := Day1PartA2021(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay1PartA2021Complete(t *testing.T) {
	answer := 142
	solution := Day1PartA2021(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestParseKeyFromStringPartB(t *testing.T) {
	answer := 29
	solution := ParseKeyFromStringPartB("two1nine")
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay1PartB2021Example(t *testing.T) {
	answer := 281
	solution := Day1PartB2021(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay1PartB2021Complete(t *testing.T) {
	answer := 281
	solution := Day1PartB2021(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
