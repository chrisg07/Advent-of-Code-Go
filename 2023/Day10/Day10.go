package AoC2021

import (
	_ "embed"
	"log"
	"slices"
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

type Point struct {
	x int
	y int
}

func getNorth(point Point) Point {
	return Point{point.x - 1, point.y}
}

func getEast(point Point) Point {
	return Point{point.x, point.y + 1}
}

func getSouth(point Point) Point {
	return Point{point.x + 1, point.y}
}

func getWest(point Point) Point {
	return Point{point.x, point.y - 1}
}

func isValidTile(tile string, validTiles []string) bool {
	return slices.Index(validTiles, tile) >= 0
}

func validateLoop(tiles [][]string, position Point, direction rune, length int) (int, bool) {
	// determine next position according to current position
	switch tiles[position.x][position.y] {
	case "S": // starting location
		switch direction {
		case 'N':
			return moveFromSouthToNorth(position, tiles, length)
		case 'E':
			return moveFromWestToEast(position, tiles, length)
		case 'S':
			return moveFromNorthToSouth(position, tiles, length)
		case 'W':
			return moveFromEastToWest(position, tiles, length)
		}
	case "|": // vertical pipe
		switch direction {
		case 'N':
			return moveFromSouthToNorth(position, tiles, length)
		case 'E':
			return length, false
		case 'S':
			return moveFromNorthToSouth(position, tiles, length)
		case 'W':
			return length, false
		}
	case "-": // horizontal pipe
		switch direction {
		case 'N':
			return length, false
		case 'E':
			return moveFromWestToEast(position, tiles, length)
		case 'S':
			return length, false
		case 'W':
			return moveFromEastToWest(position, tiles, length)
		}
	case "L": // 90 degree bend North -> East
		switch direction {
		case 'N':
			return moveFromSouthToNorth(position, tiles, length)
		case 'E':
			return moveFromWestToEast(position, tiles, length)
		case 'S':
			return length, false
		case 'W':
			return length, false
		}
	case "J": // 90 degree bend North -> West
		switch direction {
		case 'N':
			return moveFromSouthToNorth(position, tiles, length)
		case 'E':
			return length, false
		case 'S':
			return length, false
		case 'W':
			return moveFromEastToWest(position, tiles, length)
		}
	case "7": // 90 degree bend South -> West
		switch direction {
		case 'N':
			return length, false
		case 'E':
			return length, false
		case 'S':
			return moveFromNorthToSouth(position, tiles, length)
		case 'W':
			return moveFromEastToWest(position, tiles, length)
		}
	case "F": // 90 degree bend South -> East
		switch direction {
		case 'N':
			return length, false
		case 'E':
			return moveFromWestToEast(position, tiles, length)
		case 'S':
			return moveFromNorthToSouth(position, tiles, length)
		case 'W':
			return length, false
		}
	case ".": // ground
		return length + 1, false
	default:
		return length, false
	}

	return length, false
}

func moveFromNorthToSouth(position Point, tiles [][]string, length int) (int, bool) {
	nextPosition := getSouth(position)
	southTile := tiles[nextPosition.x][nextPosition.y]

	switch southTile {
	case "|":
		return validateLoop(tiles, nextPosition, 'S', length+1)
	case "L":
		return validateLoop(tiles, nextPosition, 'E', length+1)
	case "J":
		return validateLoop(tiles, nextPosition, 'W', length+1)
	case "S":
		return length + 1, true
	default:
		return length, false
	}
}

func moveFromSouthToNorth(position Point, tiles [][]string, length int) (int, bool) {
	nextPosition := getNorth(position)
	northTile := tiles[nextPosition.x][nextPosition.y]

	switch northTile {
	case "|":
		return validateLoop(tiles, nextPosition, 'N', length+1)
	case "7":
		return validateLoop(tiles, nextPosition, 'W', length+1)
	case "F":
		return validateLoop(tiles, nextPosition, 'E', length+1)
	case "S":
		return length + 1, true
	default:
		return length, false
	}
}

func moveFromWestToEast(position Point, tiles [][]string, length int) (int, bool) {
	nextPosition := getEast(position)
	eastTile := tiles[nextPosition.x][nextPosition.y]
	switch eastTile {
	case "-":
		return validateLoop(tiles, nextPosition, 'E', length+1)
	case "J":
		return validateLoop(tiles, nextPosition, 'N', length+1)
	case "7":
		return validateLoop(tiles, nextPosition, 'S', length+1)
	case "S":
		return length + 1, true
	default:
		return length, false
	}
}

func moveFromEastToWest(position Point, tiles [][]string, length int) (int, bool) {
	nextPosition := getWest(position)
	westTile := tiles[nextPosition.x][nextPosition.y]
	switch westTile {
	case "-":
		return validateLoop(tiles, nextPosition, 'W', length+1)
	case "L":
		return validateLoop(tiles, nextPosition, 'N', length+1)
	case "F":
		return validateLoop(tiles, nextPosition, 'S', length+1)
	case "S":
		return length + 1, true
	default:
		return length, false
	}
}

func Day10PartA2023(useExample bool) int {
	lines := getInput(useExample)
	tiles := [][]string{}
	animalStartingPos := Point{0, 0}
	for x, line := range lines {
		tiles = append(tiles, []string{})
		for y, char := range line {
			if string(char) == "S" {
				animalStartingPos = Point{x, y}
			}
			tiles[x] = append(tiles[x], string(char))
		}

		log.Printf("[WARN] %v\n", tiles[x])
	}

	log.Printf("[WARN] Animal starting position: %v", animalStartingPos)

	// only need to find valid pipes starting from North, East, South, and West of starting position
	// A loop means two of these paths will be valid, we can return on the first valid loop
	validLoopLengthNorth, northIsValidLoop := validateLoop(tiles, animalStartingPos, 'N', 0)
	if northIsValidLoop {
		return (validLoopLengthNorth + 1) / 2
	}
	validLoopLengthEast, eastIsValidLoop := validateLoop(tiles, animalStartingPos, 'E', 0)
	if eastIsValidLoop {
		return (validLoopLengthEast + 1) / 2
	}
	validLoopLengthSouth, southIsValidLoop := validateLoop(tiles, animalStartingPos, 'S', 0)
	if southIsValidLoop {
		return (validLoopLengthSouth + 1) / 2
	}
	validLoopLengthWest, westIsValidLoop := validateLoop(tiles, animalStartingPos, 'W', 0)
	if westIsValidLoop {
		return (validLoopLengthWest + 1) / 2
	}

	return 0
}

func Day10PartB2023(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
