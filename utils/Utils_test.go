package Utils

import (
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
