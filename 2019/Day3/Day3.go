package AoC2019

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed inputs/example.txt
var exampleInput string

//go:embed inputs/complete.txt
var input string

func getInput(useExample bool) []string {
	var lines []string
	var unsplitLines string
	if useExample {
		unsplitLines = strings.TrimRight(exampleInput, "\n")
	} else {
		unsplitLines = strings.TrimRight(input, "\n")
	}
	lines = strings.Split(unsplitLines, "\n")
	return lines
}

func parseInput(lines []string) []string {
	input := []string{}
	for _, line := range lines {
		// for _, char := range line {
		// 	log.Print(string(char))
		// }

		log.Printf("[CONSOLE] %v", line)
		input = append(input, line)
	}
	return input
}

type Point struct {
	x int
	y int
}

func ConvertInstructionsToPoints(instructions string) []Point {
	points := []Point{}
	x := 0
	y := 0
	points = append(points, Point{x, y})

	moves := strings.Split(instructions, ",")
	for _, move := range moves {
		direction := move[:1]
		distance, _ := strconv.Atoi(move[1:])

		switch direction {
		case "R":
			x += distance
		case "U":
			y += distance
		case "L":
			x -= distance
		case "D":
			y -= distance
		default:
		}

		points = append(points, Point{x, y})
	}
	return points
}

func CoordinatesIntersect(p1 Point, p2 Point, p3 Point, p4 Point) bool {
	xIntersects := (p1.x > p3.x && p1.x < p4.x) && (p2.x > p3.x && p2.x < p4.x) && (p3.x > p1.x) && (p3.x > p2.x && p3.x < p2.x)
	yIntersects := (p1.y > p3.y && p1.y < p4.y) && (p2.y > p3.y && p2.y < p4.y) && (p3.y > p1.y) && (p3.y > p2.y && p3.y < p2.y)
	return xIntersects && yIntersects
}

func LinesIntersect(a []Point, b []Point) bool {
	for i := 0; i < len(a)-2; i++ {
		for j := 0; j < len(b)-2; j++ {

		}
	}
	return false
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return len(input)
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return len(input)
}
