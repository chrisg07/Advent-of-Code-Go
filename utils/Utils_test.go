package Utils

import (
	"reflect"
	"testing"
)

func TestSumArray(t *testing.T) {
	answer := 3
	arr := []int{1, 1, 1}
	solution := SumArray(arr)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestBitArrayToDecimal(t *testing.T) {
	answer := 9
	arr := []int{0, 1, 0, 0, 1}
	solution := BitArrayToDecimal(arr)
	if solution != answer {
		t.Fatalf(`Decimal = %d, 01001 in binary is = %d`, solution, answer)
	}
}

func TestReverseArrayWithInts(t *testing.T) {
	answer := []int{0, 1, 2, 3, 4}
	arr := []int{4, 3, 2, 1, 0}
	solution := ReverseArray(arr)
	if !reflect.DeepEqual(solution, answer) {
		t.Fatalf(`Reversed array = %v, {0, 1, 2, 3, 4} reversed is = %v`, solution, answer)
	}
}

func TestReverseArrayWithStrings(t *testing.T) {
	answer := []string{"hip", "hop", "drink", "soda", "pop"}
	arr := []string{"pop", "soda", "drink", "hop", "hip"}
	solution := ReverseArray(arr)
	if !reflect.DeepEqual(solution, answer) {
		t.Fatalf(`Reversed array = %v, {"hip", "hop", "drink", "soda", "pop"} reversed is = %v`, solution, answer)
	}
}

func TestToPower(t *testing.T) {
	answer := 100
	solution := ToPower(10, 2)
	if answer != solution {
		t.Fatalf(`Solution = %v, 10^2 is = %v`, solution, answer)
	}

	answer = 1
	solution = ToPower(2, 0)
	if answer != solution {
		t.Fatalf(`Solution = %v, 2^0 is = %v`, solution, answer)
	}

	answer = 2
	solution = ToPower(2, 1)
	if answer != solution {
		t.Fatalf(`Solution = %v, 2^0 is = %v`, solution, answer)
	}

	answer = 4
	solution = ToPower(2, 2)
	if answer != solution {
		t.Fatalf(`Solution = %v, 2^0 is = %v`, solution, answer)
	}
}

func TestFlipBitArray(t *testing.T) {
	answer := []int{0, 1, 0, 1, 0}
	arr := []int{1, 0, 1, 0, 1}
	solution := FlipBitArray(arr)
	if !reflect.DeepEqual(solution, answer) {
		t.Fatalf(`Flipped array = %v, {0, 1, 0, 1, 0} flipped is = %v`, solution, answer)
	}
}

func TestRemoveIndex(t *testing.T) {
	answer := []int{0, 1, 1, 0}
	arr := []int{0, 1, 0, 1, 0}
	solution := RemoveIndex(arr, 2)
	if !reflect.DeepEqual(solution, answer) {
		t.Fatalf(`Actual = %v, expected = %v`, solution, answer)
	}
}

func TestDuplicateMatrix(t *testing.T) {
	answer := [][]int{{0, 0, 0, 0, 0}, {1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}}
	arr := [][]int{{0, 0, 0, 0, 0}, {1, 0, 1, 0, 1}, {1, 1, 1, 1, 1}}
	solution := DuplicateMatrix(arr)
	if !reflect.DeepEqual(solution, answer) && &answer != &solution {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestSplitStringToInts(t *testing.T) {
	answer := []int{8, 2, 23, 4, 24}
	str := " 8  2 23  4 24"
	solution := SplitStringToInts(str, " ")
	if !reflect.DeepEqual(solution, answer) {
		t.Fatalf(`Actual = %v, expected = %v`, solution, answer)
	}
}

func TestArrayContainsArrayTrue(t *testing.T) {
	expected := true
	haystack := []int{14, 21, 17, 24, 4}
	needles := []int{4, 24, 17, 21, 14}
	actual := ArrayContainsArray(haystack, needles)
	if actual != expected {
		t.Fatalf(`Actual = %v, expected = %v`, actual, expected)
	}
}

func TestArrayContainsArrayFalse(t *testing.T) {
	expected := false
	haystack := []int{14, 21, 17, 24, 4}
	needles := []int{4, 1337}
	actual := ArrayContainsArray(haystack, needles)
	if actual != expected {
		t.Fatalf(`Actual = %v, expected = %v`, actual, expected)
	}
}

func TestReverseString(t *testing.T) {
	expected := "cba"
	actual := ReverseString("abc")
	if actual != expected {
		t.Fatalf(`Actual = %s, expected = %s`, actual, expected)
	}
}
