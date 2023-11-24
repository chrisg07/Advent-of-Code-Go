package AoC2021

import (
	_ "embed"
	"log"
	"sort"
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

func isLowPoint(matrix [][]int, x int, y int) bool {
	point := matrix[x][y]

	xLen := len(matrix)
	yLen := len(matrix[0])

	// north
	north := 10
	if y != yLen-1 {
		north = matrix[x][y+1]
	}
	// east
	east := 10
	if x != xLen-1 {
		east = matrix[x+1][y]
	}
	// south
	south := 10
	if y != 0 {
		south = matrix[x][y-1]
	}
	// west
	west := 10
	if x != 0 {
		west = matrix[x-1][y]
	}

	return point < north && point < east && point < south && point < west
}

func Day9PartA2021(useExample bool) int {
	lines := getInput(useExample)

	matrix := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, char := range line {
			value, _ := strconv.Atoi(string(char))
			row = append(row, value)
		}
		matrix = append(matrix, row)
	}

	Utils.PrettyPrint(matrix)

	risk := 0
	for x := range matrix {
		for y := range matrix[x] {
			if isLowPoint(matrix, x, y) {
				risk += matrix[x][y] + 1
			}
		}
	}

	return risk
}

func calculateBasinSize(matrix [100][100]int) int {
	size := 0
	for x := range matrix {
		for y := range matrix[x] {
			if matrix[x][y] == 1 {
				size += 1
			}
		}
	}
	return size
}

func calculateBasin(matrix [][]int, visited [100][100]int, x int, y int) [100][100]int {
	point := matrix[x][y]

	xLen := len(matrix)
	yLen := len(matrix[0])

	// log.Printf("[WARN] Location: %d,%d height: %d", x, y, point)

	if visited[x][y] == 1 || visited[x][y] == 9 {
		panic("basin had already been visited or is too large")
	}

	if point <= 8 {
		visited[x][y] = 1
		// north
		if (y != yLen-1 && point <= matrix[x][y+1]-1) && visited[x][y+1] == 0 {
			visited = calculateBasin(matrix, visited, x, y+1)
		}
		// east
		if x != xLen-1 && (point <= matrix[x+1][y]-1) && visited[x+1][y] == 0 {
			visited = calculateBasin(matrix, visited, x+1, y)
		}
		// south
		if y != 0 && (point <= matrix[x][y-1]-1) && visited[x][y-1] == 0 {
			visited = calculateBasin(matrix, visited, x, y-1)
		}
		// west
		if x != 0 && (point <= matrix[x-1][y]-1) && visited[x-1][y] == 0 {
			visited = calculateBasin(matrix, visited, x-1, y)
		}
	}

	return visited
}

func printBasin(matrix [][]int, visited [100][100]int) {
	visitedXMin := 100
	visitedXMax := 0
	visitedYMin := 100
	visitedYMax := 0

	for x := range visited {
		for y := range visited[x] {
			if visited[x][y] == 1 {
				if x < visitedXMin {
					visitedXMin = x
				}
				if x > visitedXMax {
					visitedXMax = x
				}
				if y < visitedYMin {
					visitedYMin = y
				}
				if y > visitedYMax {
					visitedYMax = y
				}
			}
		}
	}

	for x := visitedXMin; x <= visitedXMax; x++ {
		printStr := ""
		for y := visitedYMin; y <= visitedYMax; y++ {
			if visited[x][y] == 1 {
				str := strconv.Itoa(matrix[x][y])
				printStr += str
			} else {
				printStr += "-"
			}
		}
		log.Printf("[WARN] x: %2d | "+printStr, x)
	}
}

func Day9PartB2021(useExample bool) int {
	lines := getInput(useExample)

	matrix := [][]int{}
	for _, line := range lines {
		row := []int{}
		for _, char := range line {
			value, _ := strconv.Atoi(string(char))
			row = append(row, value)
		}
		matrix = append(matrix, row)
	}

	// Utils.PrettyPrint(matrix)

	basins := []int{}
	visited := [100][100]int{}
	for x := range matrix {
		for y := range matrix[x] {
			if isLowPoint(matrix, x, y) {
				basin := calculateBasin(matrix, visited, x, y)
				basinSize := calculateBasinSize(basin)
				printBasin(matrix, basin)
				basins = append(basins, basinSize)
				log.Printf("[WARN] Discovered basin of size: %d at %d,%d", basinSize, x, y)
			}
		}
	}

	sort.Ints(basins)

	Utils.PrettyPrint(basins)

	// log.Printf("[WARN] Visited: %v", visited)

	numBasins := len(basins)
	basinSizes := basins[numBasins-3] * basins[numBasins-2] * basins[numBasins-1]
	return basinSizes
}
