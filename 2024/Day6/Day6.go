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

type Coordinates struct {
	x int
	y int
}

func coordinatesAreInBounds(width, height int, coords Coordinates) bool {
	xInBounds := coords.x >= 0 && coords.x <= width
	yInBounds := coords.y >= 0 && coords.y <= height

	return xInBounds && yInBounds
}

var MAX_GUARD_STEPS = 75000
var MIN_VISITS_FOR_LOOP = 20

func simulateGuard(width, height int, guard Coordinates, obstructions map[Coordinates]int) map[Coordinates]int {
	steps := 0
	visited := make(map[Coordinates]int)
	guardIsInBounds := coordinatesAreInBounds(width, height, guard)
	currentDirection := 'N'

	for guardIsInBounds && steps < MAX_GUARD_STEPS {
		steps++
		visited[guard]++

		switch currentDirection {
		case 'N':
			// moving North is the negative Y direction due to map orientation
			nextPosition := Coordinates{guard.x, guard.y - 1}
			if obstructions[nextPosition] != 0 {
				log.Printf("[DEBUG] Encountered obstacle at: %v, %v", nextPosition.x, nextPosition.y)
				currentDirection = 'E'
			} else {
				guard = nextPosition
			}
		case 'S':
			// moving South is the positive Y direction due to map orientation
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

	return visited
}

func loopOccurred(visited map[Coordinates]int) bool {
	for _, visits := range visited {
		if visits > MIN_VISITS_FOR_LOOP {
			return true
		}
	}

	return false
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	guard := parseGuardLocation(lines)
	obstructions := parseObstructionLocations(lines)
	width := len(lines[0])
	height := len(lines)
	visited := simulateGuard(width, height, guard, obstructions)

	return len(visited) - 1
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	guard := parseGuardLocation(lines)
	width := len(lines[0])
	height := len(lines)

	loops := 0
	for x := 0; x < width; x++ {
		for y := 0; y < height; y++ {
			cell := lines[y][x]
			if cell != '#' && cell != '^' {
				sovereignGlueTank := Coordinates{x, y}
				obstructions := parseObstructionLocations(lines)
				obstructions[sovereignGlueTank]++
				visited := simulateGuard(width, height, guard, obstructions)

				if loopOccurred(visited) {
					loops++
				}
			}
		}
	}

	return loops
}
