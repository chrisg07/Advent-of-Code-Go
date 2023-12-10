package AoC2021

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

func validateLoop(tiles [][]string, position Point, direction rune, path []Point) ([]Point, bool) {
	// determine next position according to current position
	switch tiles[position.x][position.y] {
	case "S": // starting location
		switch direction {
		case 'N':
			return moveFromSouthToNorth(position, tiles, path)
		case 'E':
			return moveFromWestToEast(position, tiles, path)
		case 'S':
			return moveFromNorthToSouth(position, tiles, path)
		case 'W':
			return moveFromEastToWest(position, tiles, path)
		}
	case "|": // vertical pipe
		switch direction {
		case 'N':
			return moveFromSouthToNorth(position, tiles, path)
		case 'S':
			return moveFromNorthToSouth(position, tiles, path)
		}
	case "-": // horizontal pipe
		switch direction {
		case 'N':
			return path, false
		case 'E':
			return moveFromWestToEast(position, tiles, path)
		case 'S':
			return path, false
		case 'W':
			return moveFromEastToWest(position, tiles, path)
		}
	case "L": // 90 degree bend North -> East
		switch direction {
		case 'N':
			return moveFromSouthToNorth(position, tiles, path)
		case 'E':
			return moveFromWestToEast(position, tiles, path)
		}
	case "J": // 90 degree bend North -> West
		switch direction {
		case 'N':
			return moveFromSouthToNorth(position, tiles, path)
		case 'W':
			return moveFromEastToWest(position, tiles, path)
		}
	case "7": // 90 degree bend South -> West
		switch direction {
		case 'S':
			return moveFromNorthToSouth(position, tiles, path)
		case 'W':
			return moveFromEastToWest(position, tiles, path)
		}
	case "F": // 90 degree bend South -> East
		switch direction {
		case 'E':
			return moveFromWestToEast(position, tiles, path)
		case 'S':
			return moveFromNorthToSouth(position, tiles, path)
		}
	case ".": // ground
		return append(path, position), false
	default:
		return path, false
	}

	return path, false
}

func moveFromNorthToSouth(position Point, tiles [][]string, path []Point) ([]Point, bool) {
	nextPosition := getSouth(position)
	southTile := tiles[nextPosition.x][nextPosition.y]

	switch southTile {
	case "|":
		return validateLoop(tiles, nextPosition, 'S', append(path, nextPosition))
	case "L":
		return validateLoop(tiles, nextPosition, 'E', append(path, nextPosition))
	case "J":
		return validateLoop(tiles, nextPosition, 'W', append(path, nextPosition))
	case "S":
		return append(path, nextPosition), true
	default:
		return path, false
	}
}

func moveFromSouthToNorth(position Point, tiles [][]string, path []Point) ([]Point, bool) {
	nextPosition := getNorth(position)
	northTile := tiles[nextPosition.x][nextPosition.y]

	switch northTile {
	case "|":
		return validateLoop(tiles, nextPosition, 'N', append(path, nextPosition))
	case "7":
		return validateLoop(tiles, nextPosition, 'W', append(path, nextPosition))
	case "F":
		return validateLoop(tiles, nextPosition, 'E', append(path, nextPosition))
	case "S":
		return append(path, nextPosition), true
	default:
		return path, false
	}
}

func moveFromWestToEast(position Point, tiles [][]string, path []Point) ([]Point, bool) {
	nextPosition := getEast(position)
	eastTile := tiles[nextPosition.x][nextPosition.y]
	switch eastTile {
	case "-":
		return validateLoop(tiles, nextPosition, 'E', append(path, nextPosition))
	case "J":
		return validateLoop(tiles, nextPosition, 'N', append(path, nextPosition))
	case "7":
		return validateLoop(tiles, nextPosition, 'S', append(path, nextPosition))
	case "S":
		return append(path, nextPosition), true
	default:
		return path, false
	}
}

func moveFromEastToWest(position Point, tiles [][]string, path []Point) ([]Point, bool) {
	nextPosition := getWest(position)
	westTile := tiles[nextPosition.x][nextPosition.y]
	switch westTile {
	case "-":
		return validateLoop(tiles, nextPosition, 'W', append(path, nextPosition))
	case "L":
		return validateLoop(tiles, nextPosition, 'N', append(path, nextPosition))
	case "F":
		return validateLoop(tiles, nextPosition, 'S', append(path, nextPosition))
	case "S":
		return append(path, nextPosition), true
	default:
		return path, false
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
	}

	northPath, northIsValidLoop := validateLoop(tiles, animalStartingPos, 'N', []Point{})
	if northIsValidLoop {
		return (len(northPath) + 1) / 2
	}
	eastPath, eastIsValidLoop := validateLoop(tiles, animalStartingPos, 'E', []Point{})
	if eastIsValidLoop {
		return (len(eastPath) + 1) / 2
	}
	southPath, southIsValidLoop := validateLoop(tiles, animalStartingPos, 'S', []Point{})
	if southIsValidLoop {
		return (len(southPath) + 1) / 2
	}
	westPath, westIsValidLoop := validateLoop(tiles, animalStartingPos, 'W', []Point{})
	if westIsValidLoop {
		return (len(westPath) + 1) / 2
	}

	return 0
}

func positionIsInPath(position Point, path []Point) bool {
	for _, tile := range path {
		if tile.x == position.x && tile.y == position.y {
			return true
		}
	}
	return false
}

func rangesNorthAndSouth(position Point, path []Point) (int, int) {
	rangesNorth := 0
	for x := 0; x < position.x-1; x++ {
		if positionIsInPath(Point{x, position.y}, path) && !positionIsInPath(Point{x + 1, position.y}, path) {
			rangesNorth += 1
		}
	}
	rangesSouth := 0
	for x := position.x + 1; x < 145; x++ {
		if positionIsInPath(Point{x, position.y}, path) && !positionIsInPath(Point{x + 1, position.y}, path) {
			rangesSouth += 1
		}
	}
	return rangesNorth, rangesSouth
}

func rangesEastAndWest(position Point, path []Point) (int, int) {
	rangesEast := 0
	for y := 0; y < position.y-1; y++ {
		if positionIsInPath(Point{position.x, y}, path) && !positionIsInPath(Point{position.x, y + 1}, path) {
			rangesEast += 1
		}
	}
	rangesWest := 0
	for y := position.y + 1; y < 145; y++ {
		if positionIsInPath(Point{position.x, y}, path) && !positionIsInPath(Point{position.x, y + 1}, path) {
			rangesWest += 1
		}
	}
	return rangesEast, rangesWest
}

func postionIsContainedInPath(position Point, path []Point) bool {
	pointsToNorth, pointsToSouth := rangesNorthAndSouth(position, path)
	pointsToEast, pointsToWest := rangesEastAndWest(position, path)
	positionIsContained := pointsToNorth%2 == 1 && pointsToSouth%2 == 1 && pointsToEast%2 == 1 && pointsToWest%2 == 1
	if positionIsContained {
		log.Printf("[WARN] Position: %v is contained by path\n", position)
	}
	return positionIsContained
}

func countTilesContainedInPath(tilesWidth int, tilesHeight int, path []Point) int {
	tilesContained := 0
	for x := 0; x < tilesWidth; x++ {
		for y := 0; y < tilesHeight; y++ {
			if postionIsContainedInPath(Point{x, y}, path) {
				tilesContained += 1
			}
		}
	}
	return tilesContained
}
func Day10PartB2023(useExample bool) int {
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
	}

	northPath, northIsValidLoop := validateLoop(tiles, animalStartingPos, 'N', []Point{})
	if northIsValidLoop {
		return countTilesContainedInPath(len(tiles), len(tiles[0]), northPath)
	}
	eastPath, eastIsValidLoop := validateLoop(tiles, animalStartingPos, 'E', []Point{})
	if eastIsValidLoop {
		return countTilesContainedInPath(len(tiles), len(tiles[0]), eastPath)
	}
	southPath, southIsValidLoop := validateLoop(tiles, animalStartingPos, 'S', []Point{})
	if southIsValidLoop {
		return countTilesContainedInPath(len(tiles), len(tiles[0]), southPath)
	}
	westPath, westIsValidLoop := validateLoop(tiles, animalStartingPos, 'W', []Point{})
	if westIsValidLoop {
		return countTilesContainedInPath(len(tiles), len(tiles[0]), westPath)
	}

	return 0
}
