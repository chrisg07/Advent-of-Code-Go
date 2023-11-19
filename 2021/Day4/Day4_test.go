package AoC2021

import (
	"testing"
)

func TestDay4PartA2021Example(t *testing.T) {
	answer := 4512
	solution := Day4PartA2021(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay4PartA2021Complete(t *testing.T) {
	answer := 4512
	solution := Day4PartA2021(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

// func TestDay4PartB2021Example(t *testing.T) {
// 	answer := 0
// 	solution := Day4PartB2021(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDay4PartBComplete(t *testing.T) {
// 	answer := 0
// 	solution := Day4PartB2021(false)
// 	if solution != answer {
// 		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
// 	}
// }
