package AoC2021

import (
	_ "embed"
	"fmt"
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

func drawLine(matrix [1000][1000]int, x1 int, y1 int, x2 int, y2 int) [1000][1000]int {
	xMin := x1
	xMax := x2
	yMin := y1
	yMax := y2
	if x2 < x1 || y2 < y1 {
		xMin = x2
		xMax = x1
		yMin = y2
		yMax = y1
	}

	// log.Printf("[DEBUG] Drawing line from %d,%d -> %d,%d\n", xMin, yMin, xMax, yMax)

	for x := xMin; x <= xMax; x += 1 {
		for y := yMin; y <= yMax; y += 1 {
			log.Printf("[DEBUG] Drawing vent at %d,%d\n", x, y)
			matrix[x][y] += 1
		}
	}

	return matrix
}

func prettyPrintMatrix(matrix [1000][1000]int) {
	// for _, row := range matrix {
	// 	// log.Printf("[DEBUG] %v\n", row)
	// }
}

func countOverlaps(matrix [1000][1000]int) int {
	overlaps := 0
	for _, row := range matrix {
		// log.Printf("[DEBUG] %v\n", row)
		for _, column := range row {
			if column >= 2 {
				overlaps += 1
			}
		}
	}
	return overlaps
}

func Day5PartA2021(useExample bool) int {
	lines := getInput(useExample)
	matrix := [1000][1000]int{}
	for _, line := range lines {
		// log.Printf("[DEBUG] %s\n", line)
		positions := strings.Split(line, " -> ")
		startCoords := strings.Split(positions[0], ",")
		endCoords := strings.Split(positions[1], ",")
		x1, _ := strconv.Atoi(startCoords[0])
		y1, _ := strconv.Atoi(startCoords[1])
		x2, _ := strconv.Atoi(endCoords[0])
		y2, _ := strconv.Atoi(endCoords[1])
		if x1 == x2 || y1 == y2 {
			matrix = drawLine(matrix, x1, y1, x2, y2)
		}

		prettyPrintMatrix(matrix)
	}

	return countOverlaps(matrix)
}

func Day5PartB2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println("")
	}

	return 0
}
