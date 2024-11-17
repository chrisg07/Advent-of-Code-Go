package AoC2021

import (
	"reflect"
	"testing"
)

// func TestPartAExample(t *testing.T) {
// 	answer := 198
// 	solution := PartA(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestDay3PartA2021Complete(t *testing.T) {
	answer := 1307354
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestReduceReadings(t *testing.T) {
	answer := [][]int{{1, 1, 1, 1, 1}}
	arr := [][]int{{0, 0, 0, 0, 0}, {1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}}
	solution := ReduceReadings(arr, []int{0, 1, 0, 0, 0}, 1)
	if !reflect.DeepEqual(solution, answer) {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestCalculateRate(t *testing.T) {
	answer := []int{1, 0, 1, 1, 0}
	arr := [][]int{
		{0, 0, 1, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 1, 1, 0},
		{1, 0, 1, 1, 1},
		{1, 0, 1, 0, 1},
		{0, 1, 1, 1, 1},
		{0, 0, 1, 1, 1},
		{1, 1, 1, 0, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 1},
		{0, 0, 0, 1, 0},
		{0, 1, 0, 1, 0},
	}
	solution := CalculateRate(arr)
	if !reflect.DeepEqual(solution, answer) {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

// func TestPartBExample(t *testing.T) {
// 	answer := 230
// 	solution := PartB(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestDay3PartB2021Complete(t *testing.T) {
	answer := 482500
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
