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

func isValidOccurrencePartB(wordsearch [][]rune, coords []Coordinate) bool {
	if len(coords) != 5 {
		return false
	}

	for _, coord := range coords {
		if coord.x < 0 || coord.x > len(wordsearch)-1 || coord.y < 0 || coord.y > len(wordsearch[0])-1 {
			return false
		}

		if wordsearch[coord.x][coord.y] == 'X' {
			return false
		}
	}

	centerA := wordsearch[coords[0].x][coords[0].y] == 'A'

	northEastX := wordsearch[coords[1].x][coords[1].y] == 'M'
	southEastX := wordsearch[coords[2].x][coords[2].y] == 'M'
	northWestX := wordsearch[coords[3].x][coords[3].y] == 'M'
	southWestX := wordsearch[coords[4].x][coords[4].y] == 'M'

	Ms := make(map[bool]int)
	Ms[northEastX]++
	Ms[northWestX]++
	Ms[southEastX]++
	Ms[southWestX]++

	if Ms[true] != 2 {
		return false
	}

	validXsToEast := northEastX && southEastX
	validXsToNorth := northEastX && northWestX
	validXsToWest := northWestX && southWestX
	validXsToSouth := southEastX && southWestX

	northEastS := wordsearch[coords[1].x][coords[1].y] == 'S'
	southEastS := wordsearch[coords[2].x][coords[2].y] == 'S'
	northWestS := wordsearch[coords[3].x][coords[3].y] == 'S'
	southWestS := wordsearch[coords[4].x][coords[4].y] == 'S'

	validSsToEast := northEastS && southEastS
	validSsToNorth := northEastS && northWestS
	validSsToWest := northWestS && southWestS
	validSsToSouth := southEastS && southWestS

	Ss := make(map[bool]int)
	Ss[northEastS]++
	Ss[northWestS]++
	Ss[southEastS]++
	Ss[southWestS]++

	if Ss[true] != 2 {
		return false
	}
	valid := centerA && ((validXsToEast && validSsToWest) || (validXsToNorth && validSsToSouth) || (validXsToWest && validSsToEast) || (validXsToSouth || validSsToNorth))

	if valid {
		for _, coord := range coords {
			log.Printf("[DEBUG] %v", string(wordsearch[coord.x][coord.y]))
		}
	}

	return valid
}

func checkCellPartB(wordsearch [][]rune, x, y int) int {
	offsets := []Coordinate{
		{x + 0, y + 0}, // center
		{x + 1, y + 1}, // northeast
		{x + 1, y - 1}, // southeast
		{x - 1, y + 1}, // northwest
		{x - 1, y - 1}, // southwest
	}

	valid := isValidOccurrencePartB(wordsearch, offsets)

	if valid {
		log.Printf("[DEBUG] Valid occurence located at: (%d, %d)", x, y)
		return 1
	} else {
		return 0
	}
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	wordsearch := parseInput(lines)

	totalValidOccurrences := 0
	for x, row := range wordsearch {
		for y := range row {
			validOccurrences := checkCellPartB(wordsearch, x, y)
			totalValidOccurrences += validOccurrences
		}
	}
	return totalValidOccurrences
}
