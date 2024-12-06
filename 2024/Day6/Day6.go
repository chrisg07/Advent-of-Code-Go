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

func parseGuardLocation(lines []string) Coordinates {
	for y, line := range lines {
		for x, char := range line {
			if char == '^' {
				return Coordinates{x, y}
			}
		}
	}

	return Coordinates{-1, -1}
}

func parseObstructionLocations(lines []string) map[Coordinates]int {
	coords := make(map[Coordinates]int)
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				position := Coordinates{x, y}
				coords[position]++
			}
		}
	}

	return coords
}

type State interface {
	Move() State
	IsFacingObstacle() bool
	Turn() State
	InBounds() bool
}

type Coordinates struct {
	x int
	y int
}

func coordinatesAreInBounds(width, height int, coords Coordinates) bool {
	xInBounds := coords.x >= 0 && coords.x <= width
	yInBounds := coords.y >= 0 && coords.y <= height

	return xInBounds && yInBounds
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	guard := parseGuardLocation(lines)
	obstructions := parseObstructionLocations(lines)
	visited := make(map[Coordinates]int)
	width := len(lines[0])
	height := len(lines)
	guardIsInBounds := coordinatesAreInBounds(width, height, guard)
	currentDirection := 'N'

	for guardIsInBounds {
		// update guard position
		visited[guard]++

		switch currentDirection {
		case 'N':
			nextPosition := Coordinates{guard.x, guard.y - 1}
			if obstructions[nextPosition] != 0 {
				log.Printf("[DEBUG] Encountered obstacle at: %v, %v", nextPosition.x, nextPosition.y)
				currentDirection = 'E'
			} else {
				guard = nextPosition
			}
		case 'S':
			nextPosition := Coordinates{guard.x, guard.y + 1}
			if obstructions[nextPosition] != 0 {
				log.Printf("[DEBUG] Encountered obstacle at: %v, %v", nextPosition.x, nextPosition.y)
				currentDirection = 'W'
			} else {
				guard = nextPosition
			}
		case 'E':
			nextPosition := Coordinates{guard.x + 1, guard.y}
			if obstructions[nextPosition] != 0 {
				log.Printf("[DEBUG] Encountered obstacle at: %v, %v", nextPosition.x, nextPosition.y)
				currentDirection = 'S'
			} else {
				guard = nextPosition
			}
		case 'W':
			nextPosition := Coordinates{guard.x - 1, guard.y}
			if obstructions[nextPosition] != 0 {
				log.Printf("[DEBUG] Encountered obstacle at: %v, %v", nextPosition.x, nextPosition.y)
				currentDirection = 'N'
			} else {
				guard = nextPosition
			}
		default:
		}

		guardIsInBounds = coordinatesAreInBounds(width, height, guard)
	}

	for _, line := range lines {
		log.Printf("[DEBUG] %v", line)
	}

	log.Printf("[DEBUG] Obstruction positions: %v", obstructions)
	log.Printf("[DEBUG] Visited: %v", visited)
	return len(visited) - 1
}

func PartB(useExample bool) int {
	lines := getInput(useExample)

	return len(lines)
}
