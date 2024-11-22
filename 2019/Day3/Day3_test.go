package AoC2019

import (
	"log"
	"os"
	"reflect"
	"testing"

	"github.com/hashicorp/logutils"
)

func init() {
	filter := &logutils.LevelFilter{
		Levels:   []logutils.LogLevel{"DEBUG", "WARN", "ERROR", "CONSOLE"},
		MinLevel: logutils.LogLevel("WARN"),
		Writer:   os.Stderr,
	}
	log.SetFlags(0)
	log.SetOutput(filter)
	log.Print("[CONSOLE] --------------------------\n")
	log.Print("[CONSOLE] Advent of Code 2019 Day 3:\n")
	log.Print("[CONSOLE] --------------------------\n")
}

func TestConvertInstructionsToPoints(t *testing.T) {
	answer := []Point{{0, 0}, {8, 0}, {8, 5}, {3, 5}, {3, 2}}
	solution := ConvertInstructionsToPoints("R8,U5,L5,D3")
	if !reflect.DeepEqual(answer, solution) {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestTwoLinesDoNotIntersect(t *testing.T) {
	line1 := []Point{{0, 0}, {8, 0}}
	line2 := []Point{{0, 1}, {8, 1}}
	solution := LinesIntersect(line1, line2)
	if solution == true {
		t.Fatalf(`Example solution = %v, should = false`, solution)
	}
}

func TestTwoPairsOfCoordinatesIntersect(t *testing.T) {
	p1 := Point{0, 0}
	p2 := Point{8, 0}
	p3 := Point{4, -4}
	p4 := Point{4, 4}
	solution := CoordinatesIntersect(p1, p2, p3, p4)
	if solution == false {
		t.Fatalf(`Example solution = %v, should = false`, solution)
	}
}
func TestLinesIntersect(t *testing.T) {
	line1 := []Point{{0, 0}, {8, 0}}
	line2 := []Point{{4, -4}, {4, 4}}
	solution := LinesIntersect(line1, line2)
	if solution == false {
		t.Fatalf(`Example solution = %v, should = false`, solution)
	}
}

// test that if there are multiple intersections they are all reported
func TestPartAExample(t *testing.T) {
	answer := 1
	solution := PartA(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartAComplete(t *testing.T) {
	answer := 1
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Find the Elf carrying the most Calories: %v", solution)
	}
}

func TestPartBExample(t *testing.T) {
	answer := 1
	solution := PartB(true)
	if solution != answer {
		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
	}
}

func TestPartBComplete(t *testing.T) {
	answer := 1
	solution := PartB(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Find the Elf carrying the most Calories: %v", solution)
	}
}
