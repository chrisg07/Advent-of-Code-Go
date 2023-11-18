package AoC2021

import (
	"testing"
)

// TestHelloName calls greetings.Hello with a name, checking
// for a valid return value.
func TestDay1PartAExample(t *testing.T) {
	answer := 7
	solution := Day1PartA(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
func TestDay1PartA(t *testing.T) {
	answer := 1342
	solution := Day1PartA(false)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}
