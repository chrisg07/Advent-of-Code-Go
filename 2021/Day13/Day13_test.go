package AoC2021

import (
	"log"
	"os"
	"reflect"
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

func TestFoldX(t *testing.T) {
	answer := [][]string{
		[]string{"#", "#", "#", "#", "#"},
		[]string{"#", ".", ".", ".", "#"},
		[]string{"#", ".", ".", ".", "#"},
		[]string{"#", ".", ".", ".", "#"},
		[]string{"#", "#", "#", "#", "#"},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", "#", "."},
	}

	paper := [][]string{
		[]string{"#", ".", "#", "#", ".", "|", "#", ".", ".", "#", "."},
		[]string{"#", ".", ".", ".", "#", "|", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", "|", "#", ".", ".", ".", "#"},
		[]string{"#", ".", ".", ".", "#", "|", ".", ".", ".", ".", "."},
		[]string{".", "#", ".", "#", ".", "|", "#", ".", "#", "#", "#"},
		[]string{".", ".", ".", ".", ".", "|", ".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", ".", "|", ".", "#", ".", ".", "."},
	}

	solution := foldX(paper, 5)
	if !reflect.DeepEqual(solution, answer) {
		t.Fatalf(`Example solution = %v, should = %v`, solution, answer)
	}
}

func TestFoldY(t *testing.T) {
	answer := [][]string{
		[]string{"#", ".", "#", ".", "#"},
		[]string{"#", ".", ".", ".", "."},
		[]string{".", "#", ".", "#", "."},
	}

	paper := [][]string{
		[]string{"#", ".", "#", ".", "#"},
		[]string{"#", ".", ".", ".", "."},
		[]string{".", "#", ".", "#", "."},
		[]string{"-", "-", "-", "-", "-"},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
		[]string{".", ".", ".", ".", "."},
	}

	solution := foldY(paper, 3)
	if !reflect.DeepEqual(solution, answer) {
		t.Fatalf(`Example solution = %v, should = %v`, solution, answer)
	}
}

func TestDay132021Example(t *testing.T) {
	answer := 17
	solution := Day132021(true, 15, 11, 1)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay132021Complete(t *testing.T) {
	answer := 785
	solution := Day132021(false, 895, 1311, 1)
	if solution != answer {
		t.Fatalf(`Comeplete solution = %d, should = %d`, solution, answer)
	}
}

func TestDay13PartB2021Example(t *testing.T) {
	answer := 16
	solution := Day132021(true, 15, 11, 0)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestDay13PartB2021Complete(t *testing.T) {
	answer := 98
	solution := Day132021(false, 895, 1311, 0)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	}
}
