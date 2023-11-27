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

func prettyPrintMatrix(matrix [][]string) {
	log.Printf("[DEBUG] -----------------------------------\n")

	for x := range matrix {
		log.Printf("[DEBUG] y:%2d %v", x, matrix[x])
		// for y := range matrix[x] {

		// }
	}
	log.Printf("[DEBUG] -----------------------------------\n")
}

func foldX(matrix [][]string, x int) [][]string {
	foldedMatrix := makeMatrix(len(matrix), x)

	prettyPrintMatrix(matrix)
	log.Printf("[DEBUG] fold along x=%d\n", x)
	for index := 0; index < len(matrix); index++ {
		rowToCompare := matrix[index][x:]
		rowToCompare = Utils.ReverseArray[string](rowToCompare)

		for y := range foldedMatrix[index] {
			if matrix[index][y] == "#" || rowToCompare[y] == "#" {
				foldedMatrix[index][y] = "#"
			} else {
				foldedMatrix[index][y] = "."
			}
		}
	}

	prettyPrintMatrix(foldedMatrix)

	return foldedMatrix
}

func foldY(matrix [][]string, y int) [][]string {
	foldedMatrix := makeMatrix(y, len(matrix[0]))

	prettyPrintMatrix(matrix)

	log.Printf("[DEBUG] fold along y=%d\n", y)
	for x := range foldedMatrix[0] {
		for index := 0; index < y; index++ {
			if matrix[index][x] == "#" || matrix[len(matrix)-index-1][x] == "#" {
				foldedMatrix[index][x] = "#"
			} else {
				foldedMatrix[index][x] = "."
			}
		}
	}

	prettyPrintMatrix(foldedMatrix)

	return foldedMatrix
}

func makeMatrix(x int, y int) [][]string {
	matrix := make([][]string, x)
	for i := range matrix {
		matrix[i] = make([]string, y)
		for j := range matrix[i] {
			matrix[i][j] = "."
		}
	}
	return matrix
}

func Day132021(useExample bool, width int, height int, maxFolds int) int {
	lines := getInput(useExample)
	paper := makeMatrix(width, height)
	folds := 0
	for _, line := range lines {
		if strings.Contains(line, "fold") {
			if folds < maxFolds || maxFolds == 0 {
				parts := strings.Split(line, "=")
				axis, _ := strconv.Atoi(parts[1])
				if strings.Contains(parts[0], "x") {
					paper = foldX(paper, axis)
				} else if strings.Contains(parts[0], "y") {
					paper = foldY(paper, axis)
				}
				folds += 1
			}
		} else if len(line) > 0 {
			coords := strings.Split(line, ",")
			xStr, yStr := coords[0], coords[1]
			x, _ := strconv.Atoi(xStr)
			y, _ := strconv.Atoi(yStr)
			paper[y][x] = "#"
		}
	}

	// count dots: "#"
	dots := 0

	for x := range paper {
		for y := range paper[x] {
			if paper[x][y] == "#" {
				dots += 1
			}
		}
	}

	return dots
}
