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
		t.Fatalf(`Example solution = %q, want match for %#q, nil`, solution, answer)
	}
}

// TestHelloEmpty calls greetings.Hello with an empty string,
// checking for an error.
// func TestDay1PartA(t *testing.T) {
//     msg, err := Hello("")
//     if msg != "" || err == nil {
//         t.Fatalf(`Hello("") = %q, %v, want "", error`, msg, err)
//     }
// }
