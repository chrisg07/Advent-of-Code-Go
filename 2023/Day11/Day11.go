package AoC2021

import (
	_ "embed"
	"log"
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

func Day11PartA2023(useExample bool) int {
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
		log.Println("")
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
	log.Printf("[WARN] Galaxies before expansion:\n")
	for _, galaxy := range galaxies {
		log.Printf("[WARN] -- [%d, %d]\n", galaxy.x, galaxy.y)
	}
	log.Printf("[WARN] Empty rows: %v\n", emptyRows)
	log.Printf("[WARN] Empty columns: %v\n", emptyColumns)

	for _, row := range Utils.ReverseArray[int](emptyRows) {
		for index, galaxy := range galaxies {
			if row < galaxy.y {
				galaxies[index].y += 1
			}
		}
	}

	for _, column := range Utils.ReverseArray[int](emptyColumns) {
		for index, galaxy := range galaxies {
			if column < galaxy.x {
				galaxies[index].x += 1
			}
		}
	}

	// create pairs
	pairs := []Line{}
	for i := 0; i < len(galaxies)-1; i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pairs = append(pairs, Line{*galaxies[i], *galaxies[j]})
		}
	}

	log.Printf("[WARN] Galaxies after expansion:\n")
	for _, galaxy := range galaxies {
		log.Printf("[WARN] -- [%d, %d]\n", galaxy.x, galaxy.y)
	}
	// determine shortest distance between each point
	sumOfShortestPaths := 0

	log.Printf("[WARN] %d pairs:\n", len(pairs))
	for _, pair := range pairs {
		log.Printf("[WARN] -- %v\n", pair)
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

func Day11PartB2023(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
