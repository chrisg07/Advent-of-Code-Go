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

func convertInstructionsToPoints(instructions string) []Point {
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
