package AoC2021

import (
	"log"
	"os"
	"testing"

	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR"},
		MinLevel: logutils.LogLevel("DEBUG"),
		Writer:   os.Stderr,
	}
	log.SetOutput(filter)
}

// func TestSimulateLanterfish1Day(t *testing.T) {
// 	answer := []int{2, 3, 2, 0, 1}
// 	solution := SimulateLanternfish([]int{3, 4, 3, 1, 2}, 1)
// 	if !reflect.DeepEqual(solution, answer) {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestSimulateLanterfish2Day(t *testing.T) {
// 	answer := []int{1, 2, 1, 6, 0, 8}
// 	solution := SimulateLanternfish([]int{3, 4, 3, 1, 2}, 2)
// 	if !reflect.DeepEqual(solution, answer) {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestSimulateLanterfishSingleFish3Days(t *testing.T) {
// 	answer := []int{6, 1}
// 	solution := SimulateLanternfish([]int{0}, 7)
// 	if !reflect.DeepEqual(solution, answer) {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestSimulateLanterfish18Days(t *testing.T) {
// 	answer := []int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8}
// 	solution := SimulateLanternfish([]int{3, 4, 3, 1, 2}, 18)
// 	if !reflect.DeepEqual(solution, answer) {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDay62Days(t *testing.T) {
// 	answer := 6
// 	solution := Day62021(true, 2, 1)
// 	if solution == answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestDay618Days(t *testing.T) {
// 	answer := 26
// 	solution := Day62021(true, 18, 1)
// 	if solution == answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestSimulateLanternfishInChunks(t *testing.T) {
// 	answer := []int{6, 0, 6, 4, 5, 6, 0, 1, 1, 2, 6, 0, 1, 1, 1, 2, 2, 3, 3, 4, 6, 7, 8, 8, 8, 8}
// 	sort.Ints(answer)
// 	solution := SimulateLanternfishInChunks(18, 3, []int{3, 4, 3, 1, 2})
// 	sort.Ints(solution)
// 	if !reflect.DeepEqual(solution, answer) {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

func TestDay6PartA2021ExampleSmall(t *testing.T) {
	answer := 26
	solution := Day62021(true, 18, 4)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}
func TestDay6PartA2021Example(t *testing.T) {
	answer := 5934
	solution := Day62021(true, 80, 4)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay6PartA2021Complete(t *testing.T) {
	answer := 375482
	solution := Day62021(false, 80, 4)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestDay6PartB2021Example(t *testing.T) {
	answer := 26984457539
	solution := Day62021(true, 256, 8)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay6PartBComplete(t *testing.T) {
	answer := 269844575390
	solution := Day62021(false, 256, 0)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
