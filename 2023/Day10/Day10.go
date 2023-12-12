package AoC2021

import (
	_ "embed"
	"fmt"
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

type Tile struct {
	pos   Point
	value string
}

func getNorth(tiles [][]string, point Point) Tile {
	return Tile{Point{point.x - 1, point.y}, tiles[point.x-1][point.y]}
}

func getEast(tiles [][]string, point Point) Tile {
	return Tile{Point{point.x, point.y + 1}, tiles[point.x][point.y+1]}
}

func getSouth(tiles [][]string, point Point) Tile {
	return Tile{Point{point.x + 1, point.y}, tiles[point.x+1][point.y]}
}

func getWest(tiles [][]string, point Point) Tile {
	return Tile{Point{point.x, point.y - 1}, tiles[point.x][point.y-1]}
}

func validateLoop(tiles [][]string, position Point, direction rune, path []Tile) ([]Tile, bool) {
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
		return append(path, Tile{position, "."}), false
	default:
		return path, false
	}

	return path, false
}

func moveFromNorthToSouth(position Point, tiles [][]string, path []Tile) ([]Tile, bool) {
	nextPosition := getSouth(tiles, position)
	southTile := tiles[nextPosition.pos.x][nextPosition.pos.y]

	switch southTile {
	case "|":
		return validateLoop(tiles, nextPosition.pos, 'S', append(path, nextPosition))
	case "L":
		return validateLoop(tiles, nextPosition.pos, 'E', append(path, nextPosition))
	case "J":
		return validateLoop(tiles, nextPosition.pos, 'W', append(path, nextPosition))
	case "S":
		return append(path, nextPosition), true
	default:
		return path, false
	}
}

func moveFromSouthToNorth(position Point, tiles [][]string, path []Tile) ([]Tile, bool) {
	nextPosition := getNorth(tiles, position)
	northTile := tiles[nextPosition.pos.x][nextPosition.pos.y]

	switch northTile {
	case "|":
		return validateLoop(tiles, nextPosition.pos, 'N', append(path, nextPosition))
	case "7":
		return validateLoop(tiles, nextPosition.pos, 'W', append(path, nextPosition))
	case "F":
		return validateLoop(tiles, nextPosition.pos, 'E', append(path, nextPosition))
	case "S":
		return append(path, nextPosition), true
	default:
		return path, false
	}
}

func moveFromWestToEast(position Point, tiles [][]string, path []Tile) ([]Tile, bool) {
	nextPosition := getEast(tiles, position)
	eastTile := tiles[nextPosition.pos.x][nextPosition.pos.y]
	switch eastTile {
	case "-":
		return validateLoop(tiles, nextPosition.pos, 'E', append(path, nextPosition))
	case "J":
		return validateLoop(tiles, nextPosition.pos, 'N', append(path, nextPosition))
	case "7":
		return validateLoop(tiles, nextPosition.pos, 'S', append(path, nextPosition))
	case "S":
		return append(path, nextPosition), true
	default:
		return path, false
	}
}

func moveFromEastToWest(position Point, tiles [][]string, path []Tile) ([]Tile, bool) {
	nextPosition := getWest(tiles, position)
	westTile := tiles[nextPosition.pos.x][nextPosition.pos.y]
	switch westTile {
	case "-":
		return validateLoop(tiles, nextPosition.pos, 'W', append(path, nextPosition))
	case "L":
		return validateLoop(tiles, nextPosition.pos, 'N', append(path, nextPosition))
	case "F":
		return validateLoop(tiles, nextPosition.pos, 'S', append(path, nextPosition))
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

	northPath, northIsValidLoop := validateLoop(tiles, animalStartingPos, 'N', []Tile{})
	if northIsValidLoop {
		return (len(northPath) + 1) / 2
	}
	eastPath, eastIsValidLoop := validateLoop(tiles, animalStartingPos, 'E', []Tile{})
	if eastIsValidLoop {
		return (len(eastPath) + 1) / 2
	}
	southPath, southIsValidLoop := validateLoop(tiles, animalStartingPos, 'S', []Tile{})
	if southIsValidLoop {
		return (len(southPath) + 1) / 2
	}
	westPath, westIsValidLoop := validateLoop(tiles, animalStartingPos, 'W', []Tile{})
	if westIsValidLoop {
		return (len(westPath) + 1) / 2
	}

	return 0
}

func positionIsInPath(position Point, path []Tile) bool {
	for _, tile := range path {
		if tile.pos.x == position.x && tile.pos.y == position.y {
			return true
		}
	}
	return false
}

func rangesNorthAndSouth(position Point, path []Tile) (int, int) {
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

// if the right eye sees the position two times more then the left eye?
func rangesEastAndWest(position Point, path []Tile) (int, int) {
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

func positionIsContainedInPath(position Point, path []Tile) bool {
	pointsToNorth, pointsToSouth := rangesNorthAndSouth(position, path)
	pointsToEast, pointsToWest := rangesEastAndWest(position, path)
	positionIsContained := pointsToNorth%2 == 1 && pointsToSouth%2 == 1 && pointsToEast%2 == 1 && pointsToWest%2 == 1
	if positionIsContained {
		log.Printf("[WARN] Position: %v is contained by path\n", position)
	}
	return positionIsContained
}

func fillTilesLeftToRight(tiles [][]string, position Point, path []Tile) [][]string {
	for x := 0; x < len(tiles[0]); x++ {
		crossedPath := false
		for y := 0; y < len(tiles); y++ {
			if positionIsInPath(Point{x, y}, path) && tiles[x][y] == "|" {
				crossedPath = true
			}
			if !crossedPath && tiles[x][y] != "■" && tiles[x][y] != "|" && tiles[x][y] != "-" {
				tiles = fillTilesRightwardUntilPathEncounter(tiles, Point{x, y}, path)
			}
		}
	}
	return tiles
}

func fillTilesTopToBottom(tiles [][]string, position Point, path []Tile) [][]string {
	for x := 0; x < len(tiles[0]); x++ {
		crossedPath := false
		for y := 0; y < len(tiles); y++ {
			if positionIsInPath(Point{y, x}, path) && tiles[y][x] == "-" {
				crossedPath = true
			}
			if !crossedPath && tiles[y][x] != "■" && tiles[y][x] != "|" && tiles[y][x] != "-" {
				tiles = fillTilesDownwardUntilPathEncounter(tiles, Point{y, x}, path)
			}
		}
	}
	return tiles
}

func fillTilesRightToLeft(tiles [][]string, position Point, path []Tile) [][]string {
	for x := len(tiles) - 1; x > 0; x-- {
		crossedPath := false
		for y := len(tiles[x]) - 1; y > 0; y-- {
			if positionIsInPath(Point{x, y}, path) && tiles[x][y] == "|" {
				crossedPath = true
			}
			if !crossedPath && tiles[x][y] != "■" && tiles[x][y] != "|" && tiles[x][y] != "-" {
				tiles = fillTilesLeftwardUntilPathEncounter(tiles, Point{x, y}, path)
			}
		}
	}
	return tiles
}

func fillTilesBottomToTop(tiles [][]string, position Point, path []Tile) [][]string {
	for x := len(tiles[0]) - 1; x > 0; x-- {
		crossedPath := false
		for y := len(tiles) - 1; y > 0; y-- {
			if positionIsInPath(Point{y, x}, path) && tiles[y][x] == "-" {
				crossedPath = true
			}
			if !crossedPath && tiles[y][x] != "■" && tiles[y][x] != "|" && tiles[y][x] != "-" {
				tiles = fillTilesUpwardUntilPathEncounter(tiles, Point{y, x}, path)
			}
		}
	}
	return tiles
}

func fillTilesRightwardUntilPathEncounter(tiles [][]string, position Point, path []Tile) [][]string {
	if !positionIsInPath(position, path) {
		tiles[position.x][position.y] = "■"
	}
	if position.x > 0 && tiles[position.x-1][position.y] != "■" && !positionIsInPath(Point{position.x - 1, position.y}, path) {
		tiles = fillTilesRightwardUntilPathEncounter(tiles, Point{position.x - 1, position.y}, path)
	}
	if position.x < len(tiles)-1 && tiles[position.x+1][position.y] != "■" && !positionIsInPath(Point{position.x + 1, position.y}, path) {
		tiles = fillTilesRightwardUntilPathEncounter(tiles, Point{position.x + 1, position.y}, path)
	}
	if position.x > 0 && position.y < len(tiles)-1 && tiles[position.x-1][position.y+1] != "■" && !positionIsInPath(Point{position.x - 1, position.y + 1}, path) {
		tiles = fillTilesRightwardUntilPathEncounter(tiles, Point{position.x - 1, position.y + 1}, path)
	}
	if position.y < len(tiles)-1 && position.x < len(tiles)-1 && tiles[position.x+1][position.y+1] != "■" && !positionIsInPath(Point{position.x + 1, position.y + 1}, path) {
		tiles = fillTilesRightwardUntilPathEncounter(tiles, Point{position.x + 1, position.y + 1}, path)
	}
	return tiles
}

func fillTilesLeftwardUntilPathEncounter(tiles [][]string, position Point, path []Tile) [][]string {
	if !positionIsInPath(position, path) {
		tiles[position.x][position.y] = "■"
	}
	if position.x > 0 && tiles[position.x-1][position.y] != "■" && !positionIsInPath(Point{position.x - 1, position.y}, path) {
		tiles = fillTilesLeftwardUntilPathEncounter(tiles, Point{position.x - 1, position.y}, path)
	}
	if position.x < len(tiles)-1 && tiles[position.x+1][position.y] != "■" && !positionIsInPath(Point{position.x + 1, position.y}, path) {
		tiles = fillTilesLeftwardUntilPathEncounter(tiles, Point{position.x + 1, position.y}, path)
	}
	if position.x > 0 && position.y > 0 && tiles[position.x-1][position.y-1] != "■" && !positionIsInPath(Point{position.x - 1, position.y - 1}, path) {
		tiles = fillTilesLeftwardUntilPathEncounter(tiles, Point{position.x - 1, position.y - 1}, path)
	}
	if position.y > 0 && position.x < len(tiles)-1 && tiles[position.x+1][position.y-1] != "■" && !positionIsInPath(Point{position.x + 1, position.y - 1}, path) {
		tiles = fillTilesLeftwardUntilPathEncounter(tiles, Point{position.x + 1, position.y - 1}, path)
	}
	return tiles
}

func fillTilesDownwardUntilPathEncounter(tiles [][]string, position Point, path []Tile) [][]string {
	if !positionIsInPath(position, path) {
		tiles[position.x][position.y] = "■"
	}
	if position.y > 0 && tiles[position.x][position.y-1] != "■" && !positionIsInPath(Point{position.x, position.y - 1}, path) {
		tiles = fillTilesDownwardUntilPathEncounter(tiles, Point{position.x, position.y - 1}, path)
	}
	if position.y < len(tiles)-1 && tiles[position.x][position.y+1] != "■" && !positionIsInPath(Point{position.x, position.y + 1}, path) {
		tiles = fillTilesDownwardUntilPathEncounter(tiles, Point{position.x, position.y + 1}, path)
	}
	if position.y > 0 && position.x > 0 && tiles[position.x-1][position.y-1] != "■" && !positionIsInPath(Point{position.x - 1, position.y - 1}, path) {
		tiles = fillTilesDownwardUntilPathEncounter(tiles, Point{position.x - 1, position.y - 1}, path)
	}
	if position.x > 0 && position.y < len(tiles)-1 && tiles[position.x-1][position.y+1] != "■" && !positionIsInPath(Point{position.x - 1, position.y + 1}, path) {
		tiles = fillTilesDownwardUntilPathEncounter(tiles, Point{position.x - 1, position.y + 1}, path)
	}
	return tiles
}

func fillTilesUpwardUntilPathEncounter(tiles [][]string, position Point, path []Tile) [][]string {
	if !positionIsInPath(position, path) {
		tiles[position.x][position.y] = "■"
	}
	if position.y > 0 && tiles[position.x][position.y-1] != "■" && !positionIsInPath(Point{position.x, position.y - 1}, path) {
		tiles = fillTilesUpwardUntilPathEncounter(tiles, Point{position.x, position.y - 1}, path)
	}
	if position.y < len(tiles)-1 && tiles[position.x][position.y+1] != "■" && !positionIsInPath(Point{position.x, position.y + 1}, path) {
		tiles = fillTilesUpwardUntilPathEncounter(tiles, Point{position.x, position.y + 1}, path)
	}
	if position.y > 0 && position.x < len(tiles)-1 && tiles[position.x+1][position.y-1] != "■" && !positionIsInPath(Point{position.x + 1, position.y - 1}, path) {
		tiles = fillTilesUpwardUntilPathEncounter(tiles, Point{position.x + 1, position.y - 1}, path)
	}
	if position.x < len(tiles)-1 && position.y < len(tiles)-1 && tiles[position.x+1][position.y+1] != "■" && !positionIsInPath(Point{position.x + 1, position.y + 1}, path) {
		tiles = fillTilesUpwardUntilPathEncounter(tiles, Point{position.x + 1, position.y + 1}, path)
	}
	return tiles
}

func countTilesNotContainedInPath(tiles [][]string, path []Tile) int {
	tilesNotContained := 0
	for x := 0; x < len(tiles); x++ {
		for y := 0; y < len(tiles[0]); y++ {
			if tiles[x][y] != "■" {
				tilesNotContained += 1
			}
		}
		fmt.Printf("%v\n", tiles[x])
	}
	return tilesNotContained
}

func updateTilesNotContainedInPath(tiles [][]string, path []Tile) [][]string {
	// for x := 0; x < len(tiles); x++ {
	// 	for y := 0; y < len(tiles[0]); y++ {
	// 		if positionIsInPath(Point{x, y}, path) && tiles[x][y] != "|" && tiles[x][y] != "-" {
	// 			tiles[x][y] = "■"
	// 		}
	// 	}
	// }
	// tiles = fillTilesLeftToRight(tiles, Point{0, 0}, path)
	// for _, tile := range path {
	// 	tiles[tile.pos.x][tile.pos.y] = tile.value
	// }

	// for x := 0; x < len(tiles); x++ {
	// 	for y := 0; y < len(tiles[0]); y++ {
	// 		if positionIsInPath(Point{x, y}, path) && tiles[x][y] != "|" && tiles[x][y] != "-" {
	// 			tiles[x][y] = "■"
	// 		}
	// 	}
	// }
	// tiles = fillTilesRightToLeft(tiles, Point{0, 0}, path)
	// for _, tile := range path {
	// 	tiles[tile.pos.x][tile.pos.y] = tile.value
	// }

	for x := 0; x < len(tiles); x++ {
		for y := 0; y < len(tiles[0]); y++ {
			if positionIsInPath(Point{x, y}, path) && tiles[x][y] != "|" && tiles[x][y] != "-" {
				tiles[x][y] = "■"
			}
		}
	}
	tiles = fillTilesTopToBottom(tiles, Point{0, 0}, path)
	for _, tile := range path {
		tiles[tile.pos.x][tile.pos.y] = tile.value
	}

	// for x := 0; x < len(tiles); x++ {
	// 	for y := 0; y < len(tiles[0]); y++ {
	// 		if positionIsInPath(Point{x, y}, path) && tiles[x][y] != "|" && tiles[x][y] != "-" {
	// 			tiles[x][y] = "■"
	// 		}
	// 	}
	// }
	// tiles = fillTilesBottomToTop(tiles, Point{0, 0}, path)
	// for _, tile := range path {
	// 	tiles[tile.pos.x][tile.pos.y] = tile.value
	// }
	for x := 0; x < len(tiles); x++ {
		fmt.Printf("%v\n", tiles[x])
	}

	return tiles
}

func findAnimalStartingPos(tiles [][]string) Point {
	for x := range tiles {
		for y := range tiles[x] {
			if tiles[x][y] == "S" {
				return Point{x, y}
			}
		}
	}
	return Point{0, 0}
}

func updateRotatedTiles(tiles [][]string) [][]string {
	for x, row := range tiles {
		for y, tile := range row {
			switch tile {
			case "|":
				tiles[x][y] = "-"
			case "-":
				tiles[x][y] = "|"
			case "L":
				tiles[x][y] = "F"
			case "J":
				tiles[x][y] = "L"
			case "7":
				tiles[x][y] = "J"
			case "F":
				tiles[x][y] = "7"
			default:
				tiles[x][y] = tiles[x][y]
			}
		}
	}
	return tiles
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
	northPath, northIsValidLoop := validateLoop(tiles, animalStartingPos, 'N', []Tile{})
	if northIsValidLoop {
		path := northPath[:]
		tiles = updateTilesNotContainedInPath(tiles, path)
		for _, tile := range path {
			tiles[tile.pos.x][tile.pos.y] = "■"
		}
		for x := 0; x < len(tiles); x++ {
			fmt.Printf("%v\n", tiles[x])
		}
		return countTilesNotContainedInPath(tiles, path)
	}
	eastPath, eastIsValidLoop := validateLoop(tiles, animalStartingPos, 'E', []Tile{})
	if eastIsValidLoop {
		path := eastPath[:]
		tiles = updateTilesNotContainedInPath(tiles, path)
		for _, tile := range path {
			tiles[tile.pos.x][tile.pos.y] = "■"
		}
		for x := 0; x < len(tiles); x++ {
			fmt.Printf("%v\n", tiles[x])
		}
		return countTilesNotContainedInPath(tiles, path)
	}
	southPath, southIsValidLoop := validateLoop(tiles, animalStartingPos, 'S', []Tile{})
	if southIsValidLoop {
		path := southPath[:]
		tiles = updateTilesNotContainedInPath(tiles, path)
		for _, tile := range path {
			tiles[tile.pos.x][tile.pos.y] = "■"
		}
		for x := 0; x < len(tiles); x++ {
			fmt.Printf("%v\n", tiles[x])
		}
		return countTilesNotContainedInPath(tiles, path)
	}
	westPath, westIsValidLoop := validateLoop(tiles, animalStartingPos, 'W', []Tile{})
	if westIsValidLoop {
		path := westPath[:]
		tiles = updateTilesNotContainedInPath(tiles, path)
		for _, tile := range path {
			tiles[tile.pos.x][tile.pos.y] = "■"
		}
		for x := 0; x < len(tiles); x++ {
			fmt.Printf("%v\n", tiles[x])
		}
		return countTilesNotContainedInPath(tiles, path)
	}
	return 0
}
