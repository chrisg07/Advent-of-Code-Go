package AoC2021

import (
	_ "embed"
	"log"
	"slices"
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

type Mirror struct {
	matrix                          [][]string
	verticalReflection              bool
	columnsLeftOfVerticalReflection int
	horizontalReflection            bool
	rowsAboveHorizontalReflection   int
}

func createMirrorFromMatrix(matrix [][]string) Mirror {
	// determine if there is a vertical reflection
	isVerticalReflection := false
	verticalReflectionAxis := 0
	for y, row := range matrix {
		for x := 1; x < len(matrix[y]); x++ {
			left := row[:x]
			leftReflection := Utils.ReverseArray(left)
			right := row[x:]
			if len(leftReflection) >= len(right) {
				for index := range right {
					if leftReflection[index] != right[index] {
						isVerticalReflection = false
					}
				}
			}
			if len(leftReflection) < len(right) {
				if slices.Equal[[]string, string](leftReflection, right[:len(leftReflection)]) {
					isVerticalReflection = true
					verticalReflectionAxis = x - 1
				}
			}
			if isVerticalReflection {
				break
			}
		}
		if isVerticalReflection {
			break
		}
	}
	// determine if there is a horizontal reflections
	isHorizontalReflection := true
	horizontalReflectionAxis := 0
	// for y := range matrix[0] {
	// 	for x := range matrix[y] {
	// 		// which values to compare?
	// 		left := matrix[x]
	// 	}
	// }
	return Mirror{
		matrix:                          matrix,
		verticalReflection:              isVerticalReflection,
		columnsLeftOfVerticalReflection: verticalReflectionAxis,
		horizontalReflection:            isHorizontalReflection,
		rowsAboveHorizontalReflection:   horizontalReflectionAxis,
	}
}

func Day13PartA2023(useExample bool) int {
	lines := getInput(useExample)
	mirrorIndex := 0
	matrices := [][][]string{[][]string{}}
	for _, line := range lines {

		if len(line) == 0 {
			matrices = append(matrices, [][]string{})
			mirrorIndex += 1
		} else {
			chars := strings.Split(line, "")
			matrices[mirrorIndex] = append(matrices[mirrorIndex], chars)
		}
	}

	mirrors := []Mirror{}
	for index, matrix := range matrices {
		log.Printf("[WARN] Mirror %d: \n%v\n", index, matrix)
		mirrors = append(mirrors, createMirrorFromMatrix(matrix))

	}

	sumOfColumnsLeftOfVerticalReflections := 0
	sumOfRowsAboveHorizontalReflections := 0
	// for each mirror
	for _, mirror := range mirrors {
		if mirror.verticalReflection {
			sumOfColumnsLeftOfVerticalReflections += mirror.columnsLeftOfVerticalReflection
		}
		if mirror.horizontalReflection {
			sumOfRowsAboveHorizontalReflections += mirror.rowsAboveHorizontalReflection
		}
	}

	for index, mirror := range mirrors {
		log.Printf("[WARN] Mirror %d: \n%v\n", index, mirror)
	}

	return sumOfColumnsLeftOfVerticalReflections + (100 * sumOfRowsAboveHorizontalReflections)
}

func Day13PartB2023(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
