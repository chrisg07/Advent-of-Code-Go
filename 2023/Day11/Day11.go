package AoC2023

import (
	_ "embed"
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

type Point struct {
	x int
	y int
}

type Line struct {
	a Point
	b Point
}

func Day112023(useExample bool, expansionFactor int) int {
	lines := getInput(useExample)
	galaxies := []*Point{}
	emptyRows := []int{}
	emptyColumns := []int{}
	for y, line := range lines {
		if !strings.Contains(line, "#") {
			emptyRows = append(emptyRows, y)
		}
		for x, char := range line {
			if char == '#' {
				galaxies = append(galaxies, &Point{x, y})
			}
		}
	}

	for x := range lines[0] {
		containsGalaxy := false
		for y := range lines {
			if lines[y][x] == '#' {
				containsGalaxy = true
			}
		}
		if !containsGalaxy {
			emptyColumns = append(emptyColumns, x)
		}
	}

	for _, row := range Utils.ReverseArray[int](emptyRows) {
		for index, galaxy := range galaxies {
			if row < galaxy.y {
				galaxies[index].y += (expansionFactor - 1)
			}
		}
	}

	for _, column := range Utils.ReverseArray[int](emptyColumns) {
		for index, galaxy := range galaxies {
			if column < galaxy.x {
				galaxies[index].x += (expansionFactor - 1)
			}
		}
	}

	pairs := []Line{}
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pairs = append(pairs, Line{*galaxies[i], *galaxies[j]})
		}
	}

	sumOfShortestPaths := 0
	for _, pair := range pairs {
		distance := calculateLengthOfLine(pair)
		sumOfShortestPaths += distance
	}

	return sumOfShortestPaths
}

func calculateLengthOfLine(pair Line) int {
	xDistance := pair.a.x - pair.b.x
	if xDistance < 0 {
		xDistance *= -1
	}
	yDistance := pair.a.y - pair.b.y
	if yDistance < 0 {
		yDistance *= -1
	}
	return xDistance + yDistance
}
