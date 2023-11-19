package AoC2021

import (
	"testing"
)

func TestDayXPartA2021Example(t *testing.T) {
	answer := 0
	solution := DayXPartA2021(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDayXPartA2021Complete(t *testing.T) {
	answer := 0
	solution := DayXPartA2021(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestDayXPartB2021Example(t *testing.T) {
	answer := 0
	solution := DayXPartB2021(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDayXPartBComplete(t *testing.T) {
	answer := 0
	solution := DayXPartB2021(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
