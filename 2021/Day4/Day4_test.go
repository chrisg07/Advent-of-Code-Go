package AoC2021

import (
	"testing"
)

func TestCheckBoardWinConditionTrue(t *testing.T) {
	expected := true
	draws := []int{7, 4, 9, 5, 11, 17, 23, 2, 0, 14, 21, 24}
	board := BingoBoard{
		rows: [][]int{{14, 21, 17, 24, 4}},
	}
	actual := CheckBoardWinCondition(board, draws)

	if actual != expected {
		t.Fatalf(`Example solution = %t, should = %t`, actual, expected)
	}
}

func TestDay4PartA2021Example(t *testing.T) {
	answer := 4512
	solution := Day4PartA2021(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay4PartA2021Complete(t *testing.T) {
	answer := 33462
	solution := Day4PartA2021(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestDay4PartB2021Example(t *testing.T) {
	answer := 1924
	solution := Day4PartB2021(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay4PartBComplete(t *testing.T) {
	answer := 30070
	solution := Day4PartB2021(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
