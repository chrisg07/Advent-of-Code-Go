package AoC2019

import (
	_ "embed"
	"slices"
	"strconv"
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
		input = append(input, line)
	}
	return input
}

type Point struct {
	x int
	y int
}

func GetEasternPath(x int, y int, distance int) []Point {
	path := []Point{}

	for i := 1; i <= distance; i++ {
		path = append(path, Point{x + i, y})
	}

	return path
}

func GetWesternPath(x int, y int, distance int) []Point {
	path := []Point{}

	for i := 1; i <= distance; i++ {
		path = append(path, Point{x - i, y})
	}

	return path
}

func GetNorthernPath(x int, y int, distance int) []Point {
	path := []Point{}

	for i := 1; i <= distance; i++ {
		path = append(path, Point{x, y + i})
	}

	return path
}

func GetSouthernPath(x int, y int, distance int) []Point {
	path := []Point{}

	for i := 1; i <= distance; i++ {
		path = append(path, Point{x, y - i})
	}

	return path
}

func ConvertInstructionsToPoints(instructions string) []Point {
	points := []Point{}
	x := 0
	y := 0
	points = append(points, Point{x, y})

	moves := strings.Split(instructions, ",")
	for _, move := range moves {
		direction := move[:1]
		distance, _ := strconv.Atoi(move[1:])

		switch direction {
		case "R":
			path := GetEasternPath(x, y, distance)
			points = append(points, path...)
			x += distance
		case "U":
			path := GetNorthernPath(x, y, distance)
			points = append(points, path...)
			y += distance
		case "L":
			path := GetWesternPath(x, y, distance)
			points = append(points, path...)
			x -= distance
		case "D":
			path := GetSouthernPath(x, y, distance)
			points = append(points, path...)
			y -= distance
		default:
		}
	}
	return points
}

func PathsIntersect(a []Point, b []Point) bool {
	origin := Point{0, 0}

	for _, point := range a {
		if slices.Contains(b, point) && point != origin {
			return true
		}
	}
	return false
}

func GetIntersectionPoints(a []Point, b []Point) []Point {
	intersections := []Point{}
	origin := Point{0, 0}

	for _, point := range a {
		if slices.Contains(b, point) && point != origin {
			intersections = append(intersections, point)
		}
	}
	return intersections
}

func AbsInt(value int) int {
	if value < 0 {
		return value * -1
	} else {
		return value
	}
}

func GetShortestDistance(a []Point) int {
	shortestDistance := 10000000

	for _, point := range a {
		distance := AbsInt(point.x) + AbsInt(point.y)
		if distance < shortestDistance {
			shortestDistance = distance
		}
	}

	return shortestDistance
}

func LinesIntersect(a []Point, b []Point) bool {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(b); j++ {
			if a[i] == b[j] {
				return true
			}
		}
	}
	return false
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	line1 := ConvertInstructionsToPoints(input[0])
	line2 := ConvertInstructionsToPoints(input[1])
	intersections := GetIntersectionPoints(line1, line2)
	solution := GetShortestDistance(intersections)

	return solution
}

func DetermineDistancesToIntersections(intersections []Point, line1 []Point, line2 []Point) []int {
	distances := []int{}

	for _, intersection := range intersections {
		distanceToLine1Intersection := slices.Index(line1, intersection)
		distanceToLine2Intersection := slices.Index(line2, intersection)
		distances = append(distances, distanceToLine1Intersection+distanceToLine2Intersection)
	}

	return distances
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	line1 := ConvertInstructionsToPoints(input[0])
	line2 := ConvertInstructionsToPoints(input[1])
	intersections := GetIntersectionPoints(line1, line2)
	travelDistance := DetermineDistancesToIntersections(intersections, line1, line2)

	slices.Sort(travelDistance)
	// sort travel distances
	// return shortest
	return travelDistance[0]
}
