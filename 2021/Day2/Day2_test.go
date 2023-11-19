package AoC2021

import (
	"testing"
)

func TestPartAExample(t *testing.T) {
	answer := 150
	solution := PartA(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartAComplete(t *testing.T) {
	answer := 2150351
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestPartBExample(t *testing.T) {
	answer := 900
	solution := PartB(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartBComplete(t *testing.T) {
	answer := 1842742223
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
