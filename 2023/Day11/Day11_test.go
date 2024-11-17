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

func TestCalculateLengthOfLine(t *testing.T) {
	answer := 17
	solution := calculateLengthOfLine(Line{Point{0, 2}, Point{12, 7}})
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}

	answer = 9
	solution = calculateLengthOfLine(Line{Point{1, 6}, Point{5, 11}})
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}

	answer = 15
	solution = calculateLengthOfLine(Line{Point{4, 0}, Point{9, 10}})
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}

	answer = 5
	solution = calculateLengthOfLine(Line{Point{0, 11}, Point{5, 11}})
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay11PartA2023Example(t *testing.T) {
	answer := 374
	solution := Day112023(true, 1)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay11PartA2023Complete(t *testing.T) {
	answer := 9177603
	solution := Day112023(false, 2)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestDay11PartB2023Example(t *testing.T) {
	answer := 1850
	solution := Day112023(true, 10)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay11PartB2023Complete(t *testing.T) {
	answer := 632003913611
	solution := Day112023(false, 1000000)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
