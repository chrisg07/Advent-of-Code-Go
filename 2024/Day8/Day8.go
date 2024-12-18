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

func parseInput(lines []string) []string {
	input := []string{}
	for _, line := range lines {
		// for _, char := range line {
		// 	log.Print(string(char))
		// }

		log.Printf("[DEBUG] %v", line)
		input = append(input, line)
	}
	return input
}

type Cell struct {
	value rune
	x     int
	y     int
}

type Coordinates struct {
	x int
	y int
}

func parseGrid(lines []string) []Cell {
	cells := []Cell{}

	for y, line := range lines {
		for x, char := range line {
			switch char {
			case '\r':
				fallthrough
			case '.':
				continue
			default:
				cell := Cell{char, x, y}
				cells = append(cells, cell)
			}
		}
	}

	log.Printf("[DEBUG] Parsed grid with cells: %v", cells)
	return cells
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	cells := parseGrid(lines)

	cellMap := make(map[rune][]Cell)

	for _, cell := range cells {
		cellMap[cell.value] = append(cellMap[cell.value], cell)
	}

	log.Printf("[DEBUG] Parsed cells to map: %v", cellMap)

	antinodes := []Coordinates{}
	for _, group := range cellMap {
		for i, cell := range group {
			for j, sibling := range group {
				if i < j {
					deltaX := cell.x - sibling.x
					deltaY := cell.y - sibling.y
					log.Printf("[DEBUG] Comparing: %v and %v", cell, sibling)

					var antinodeA Coordinates
					var antinodeB Coordinates
					if deltaX < 0 && deltaY < 0 {
						antinodeA = Coordinates{cell.x - (deltaX * -1), cell.y - (deltaY * -1)}
						antinodeB = Coordinates{sibling.x + (deltaX * -1), sibling.y + (deltaY * -1)}
					} else if deltaX >= 0 && deltaY >= 0 {
						antinodeA = Coordinates{cell.x - deltaX, cell.y - deltaY}
						antinodeB = Coordinates{sibling.x + deltaX, sibling.y + deltaY}
					} else if deltaX < 0 && deltaY >= 0 {
						antinodeA = Coordinates{cell.x - (deltaX * -1), cell.y - deltaY}
						antinodeB = Coordinates{cell.x - (deltaX * -1), sibling.y + deltaY}
					} else if deltaX >= 0 && deltaY < 0 {
						antinodeA = Coordinates{cell.x + deltaX, cell.y - (deltaY * -1)}
						antinodeB = Coordinates{sibling.x - deltaX, sibling.y + (deltaY * -1)}
					}

					log.Printf("[DEBUG] Creating antinodes: %v and %v", antinodeA, antinodeB)
					antinodes = append(antinodes, antinodeA, antinodeB)
				}
			}
		}
	}

	log.Printf("[DEBUG] Antinodes before filtering for bounds: %v", antinodes)
	filteredAntinodes := make(map[Coordinates]int)

	width := len(lines[0])
	height := len(lines)
	for _, coord := range antinodes {
		if coord.x >= 0 && coord.x < width && coord.y >= 0 && coord.y < height {
			filteredAntinodes[coord]++
		}
	}

	return len(filteredAntinodes)
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return len(input)
}
