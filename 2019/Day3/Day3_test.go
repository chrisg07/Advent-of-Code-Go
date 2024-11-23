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

func TestConvertPointsToCoordinateTrail(t *testing.T) {
	answer := []Point{{0, 0}, {1, 0}, {2, 0}, {3, 0}, {4, 0}, {5, 0}, {6, 0}, {7, 0}, {8, 0}, {8, 1}, {8, 2}, {8, 3}, {8, 4}, {8, 5}}
	points := ConvertInstructionsToPoints("R8,U5")
	if !reflect.DeepEqual(answer, points) {
		t.Fatalf(`Example solution = %d, should = %d`, points, answer)
	}
}

func TestTwoPathsDoNotIntersect(t *testing.T) {
	line1 := ConvertInstructionsToPoints("R8")
	line2 := ConvertInstructionsToPoints("L8")
	solution := PathsIntersect(line1, line2)
	if solution == true {
		t.Fatalf(`Example solution = %v, should = false`, solution)
	}
}

func TestTwoPathsDoIntersect(t *testing.T) {
	line1 := ConvertInstructionsToPoints("R8,U5")
	line2 := ConvertInstructionsToPoints("U3, R10")
	solution := PathsIntersect(line1, line2)
	if solution == true {
		t.Fatalf(`Example solution = %v, should = false`, solution)
	}
}

func TestPathIntersections(t *testing.T) {
	answer := []Point{{8, 3}}
	line1 := ConvertInstructionsToPoints("R8,U5")
	line2 := ConvertInstructionsToPoints("U3, R10")
	solution := GetIntersectionPoints(line1, line2)
	if reflect.DeepEqual(solution, answer) {
		t.Fatalf(`Example solution = %v, should = false`, solution)
	}
}

func TestShortPathIntersections(t *testing.T) {
	answer := []Point{{3, 3}, {6, 5}}
	line1 := ConvertInstructionsToPoints("R8,U5,L5,D3")
	line2 := ConvertInstructionsToPoints("U7,R6,D4,L4")
	solution := GetIntersectionPoints(line1, line2)
	if reflect.DeepEqual(solution, answer) {
		t.Fatalf(`Example solution = %v, should = false`, solution)
	}
}

func TestShortestManhattanDistance(t *testing.T) {
	answer := 6
	line1 := ConvertInstructionsToPoints("R8,U5,L5,D3")
	line2 := ConvertInstructionsToPoints("U7,R6,D4,L4")
	intersections := GetIntersectionPoints(line1, line2)
	solution := GetShortestDistance(intersections)
	if solution != answer {
		t.Fatalf(`Example solution = %v, should = false`, solution)
	}
}

func TestPartAComplete(t *testing.T) {
	answer := 2129
	solution := PartA(false)
	if solution != answer {
		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
	} else {
		log.Printf("[CONSOLE] Find the Elf carrying the most Calories: %v", solution)
	}
}

// func TestPartBExample(t *testing.T) {
// 	answer := 1
// 	solution := PartB(true)
// 	if solution != answer {
// 		t.Fatalf(`Example solution = %d, should = %d`, solution, answer)
// 	}
// }

// func TestPartBComplete(t *testing.T) {
// 	answer := 1
// 	solution := PartB(false)
// 	if solution != answer {
// 		t.Fatalf(`Complete solution = %d, should = %d`, solution, answer)
// 	} else {
// 		log.Printf("[CONSOLE] Find the Elf carrying the most Calories: %v", solution)
// 	}
// }
