package AoC2021

import (
	_ "embed"
	"log"
	"strconv"
	"strings"

	"github.com/chrisg07/Advent-of-Code-Go/Utils"
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

func drawStraightLine(matrix [1000][1000]int, x1 int, y1 int, x2 int, y2 int) [1000][1000]int {
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

	log.Printf("[DEBUG] Drawing straight line from %d,%d -> %d,%d\n", x1, y1, x2, y2)
	for x := xMin; x <= xMax; x += 1 {
		for y := yMin; y <= yMax; y += 1 {
			matrix[x][y] += 1
		}
	}

	return matrix
}

func drawDiagonalLine(matrix [1000][1000]int, x1 int, y1 int, x2 int, y2 int) [1000][1000]int {
	northEast := x1 < x2 && y1 < y2
	southEast := x1 > x2 && y1 < y2
	southWest := x1 > x2 && y1 > y2
	northWest := x1 < x2 && y1 > y2

	if northEast {
		log.Printf("[DEBUG] Drawing line northeast from %d,%d -> %d,%d\n", x1, y1, x2, y2)
		for x := x1; x <= x2; x += 1 {
			for y := y1; y <= y2; y += 1 {
				matrix[x][y] += 1
				x += 1
			}
		}
	} else if southEast {
		log.Printf("[DEBUG] Drawing line southeast from %d,%d -> %d,%d\n", x1, y1, x2, y2)
		for x := x1; x >= x2; x -= 1 {
			for y := y1; y <= y2; y += 1 {
				matrix[x][y] += 1
				x -= 1
			}
		}
	} else if southWest {
		log.Printf("[DEBUG] Drawing line southwest from %d,%d -> %d,%d\n", x1, y1, x2, y2)
		for x := x1; x >= x2; x -= 1 {
			for y := y1; y >= y2; y -= 1 {
				matrix[x][y] += 1
				x -= 1
			}
		}
	} else if northWest {
		log.Printf("[DEBUG] Drawing line northwest from %d,%d -> %d,%d\n", x1, y1, x2, y2)
		for x := x1; x <= x2; x += 1 {
			for y := y1; y >= y2; y -= 1 {
				matrix[x][y] += 1
				x += 1
			}
		}
	}

	return matrix
}

func prettyPrintMatrix(matrix [1000][1000]int) {
	log.Printf("[DEBUG] --- Displaying map of sea vents ---\n\n")

	matrixToTranspose := [][]int{}
	for x := 9; x >= 0; x -= 1 {
		matrixToTranspose = append(matrixToTranspose, matrix[x][:10])
	}

	transposedMatrix := Utils.Transpose(matrixToTranspose)
	for x := 0; x < 10; x += 1 {
		arrayForDisplay := Utils.ReverseArray[int](transposedMatrix[x][:10])
		log.Printf("[DEBUG]       %v\n", arrayForDisplay)
	}
	log.Printf("[DEBUG] -----------------------------------\n")
}

func countOverlaps(matrix [1000][1000]int) int {
	overlaps := 0
	for _, row := range matrix {
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
		x1, y1, x2, y2 := parseCoordinates(line)
		if x1 == x2 || y1 == y2 {
			matrix = drawStraightLine(matrix, x1, y1, x2, y2)
		}

		if useExample {
			prettyPrintMatrix(matrix)
		}
	}

	return countOverlaps(matrix)
}

func Day5PartB2021(useExample bool) int {
	lines := getInput(useExample)
	matrix := [1000][1000]int{}
	for _, line := range lines {
		x1, y1, x2, y2 := parseCoordinates(line)
		if x1 == x2 || y1 == y2 {
			matrix = drawStraightLine(matrix, x1, y1, x2, y2)
		} else {
			matrix = drawDiagonalLine(matrix, x1, y1, x2, y2)
		}

		if useExample {
			prettyPrintMatrix(matrix)
		}
	}

	return countOverlaps(matrix)
}

func parseCoordinates(line string) (int, int, int, int) {
	positions := strings.Split(line, " -> ")
	startCoords := strings.Split(positions[0], ",")
	endCoords := strings.Split(positions[1], ",")
	x1, _ := strconv.Atoi(startCoords[0])
	y1, _ := strconv.Atoi(startCoords[1])
	x2, _ := strconv.Atoi(endCoords[0])
	y2, _ := strconv.Atoi(endCoords[1])
	return x1, y1, x2, y2
}
