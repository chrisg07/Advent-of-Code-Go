package AoCScaffold

import (
	_ "embed"
	"log"
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

func parseInput(lines []string) [][]rune {
	wordsearch := [][]rune{}
	for _, line := range lines {
		row := []rune{}
		for _, char := range line {
			row = append(row, char)
		}

		wordsearch = append(wordsearch, row)
		log.Printf("[DEBUG] %v", row)
	}

	return wordsearch
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	wordsearch := parseInput(lines)

	totalValidOccurrences := 0
	for x, row := range wordsearch {
		for y := range row {
			validOccurrences := checkCell(wordsearch, x, y)
			totalValidOccurrences += validOccurrences
		}
	}
	return totalValidOccurrences
}

type Coordinate struct {
	x int
	y int
}

func isValidOccurrence(wordsearch [][]rune, coords []Coordinate) bool {
	if len(coords) != 4 {
		return false
	}

	for _, coord := range coords {
		if coord.x < 0 || coord.x > len(wordsearch)-1 || coord.y < 0 || coord.y > len(wordsearch[0])-1 {
			return false
		}
	}

	validX := wordsearch[coords[0].x][coords[0].y] == 'X'
	validM := wordsearch[coords[1].x][coords[1].y] == 'M'
	validA := wordsearch[coords[2].x][coords[2].y] == 'A'
	validS := wordsearch[coords[3].x][coords[3].y] == 'S'

	return (validX && validM && validA && validS)
}

func checkCell(wordsearch [][]rune, x, y int) int {
	occurrences := 0
	// check east
	offsets := []Coordinate{
		{x + 0, y},
		{x + 1, y},
		{x + 2, y},
		{x + 3, y},
	}

	valid := isValidOccurrence(wordsearch, offsets)

	if valid {
		occurrences++
	}

	// check west
	offsets = []Coordinate{
		{x + 0, y},
		{x - 1, y},
		{x - 2, y},
		{x - 3, y},
	}

	valid = isValidOccurrence(wordsearch, offsets)

	if valid {
		occurrences++
	}
	// check north
	offsets = []Coordinate{
		{x, y + 0},
		{x, y + 1},
		{x, y + 2},
		{x, y + 3},
	}

	valid = isValidOccurrence(wordsearch, offsets)

	if valid {
		occurrences++
	}
	// check south
	offsets = []Coordinate{
		{x, y + 0},
		{x, y - 1},
		{x, y - 2},
		{x, y - 3},
	}

	valid = isValidOccurrence(wordsearch, offsets)

	if valid {
		occurrences++
	}
	// check SW
	offsets = []Coordinate{
		{x + 0, y + 0},
		{x - 1, y - 1},
		{x - 2, y - 2},
		{x - 3, y - 3},
	}

	valid = isValidOccurrence(wordsearch, offsets)

	if valid {
		occurrences++
	}
	// check SE
	offsets = []Coordinate{
		{x + 0, y + 0},
		{x + 1, y - 1},
		{x + 2, y - 2},
		{x + 3, y - 3},
	}

	valid = isValidOccurrence(wordsearch, offsets)

	if valid {
		occurrences++
	}
	// check NE
	offsets = []Coordinate{
		{x + 0, y + 0},
		{x + 1, y + 1},
		{x + 2, y + 2},
		{x + 3, y + 3},
	}

	valid = isValidOccurrence(wordsearch, offsets)

	if valid {
		occurrences++
	}
	// check NW
	offsets = []Coordinate{
		{x + 0, y + 0},
		{x - 1, y + 1},
		{x - 2, y + 2},
		{x - 3, y + 3},
	}

	valid = isValidOccurrence(wordsearch, offsets)

	if valid {
		occurrences++
	}
	return occurrences
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return len(input)
}
