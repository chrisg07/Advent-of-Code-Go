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

func TestCalculateFuelPartA(t *testing.T) {
	answer := 71
	crabs := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	solution := CalculateFuelPartA(crabs, 10)
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
	answer := 355150
	solution := Day7PartA2021(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestCalculateFuelForDistance(t *testing.T) {
	answer := 66
	distance := 16 - 5
	solution := CalculateFuelForDistance(distance)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestCalculateFuelPartB(t *testing.T) {
	answer := 206
	crabs := []int{16, 1, 2, 0, 4, 2, 7, 1, 2, 14}
	solution := CalculateFuelPartB(crabs, 2)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay7PartB2021Example(t *testing.T) {
	answer := 168
	solution := Day7PartB2021(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay7PartB2021Complete(t *testing.T) {
	answer := 98368490
	solution := Day7PartB2021(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
