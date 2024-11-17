package AoC2021

import (
	"testing"
)

// func TestPartAExample(t *testing.T) {
// 	answer := 7
// 	solution := PartA(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestDay1PartA2021Complete(t *testing.T) {
	answer := 1342
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

// func TestPartBExample(t *testing.T) {
// 	answer := 5
// 	solution := PartB(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestDay1PartB2021Complete(t *testing.T) {
	answer := 1378
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}
