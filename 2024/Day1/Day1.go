package AoCScaffold

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

func parseInput(lines []string) ([]int, []int) {
	leftSide := []int{}
	rightSide := []int{}

	for _, line := range lines {
		intStrings := strings.Split(line, "   ")
		a, _ := strconv.Atoi(intStrings[0])
		b, _ := strconv.Atoi(intStrings[1])
		leftSide = append(leftSide, a)
		rightSide = append(rightSide, b)
	}

	slices.Sort(leftSide)
	slices.Sort(rightSide)

	return leftSide, rightSide
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	a, b := parseInput(lines)

	deltas := []int{}
	for index := range a {
		delta := a[index] - b[index]

		if delta < 0 {
			delta *= -1
		}

		deltas = append(deltas, delta)
	}

	deltaSum := 0
	for _, delta := range deltas {
		deltaSum += delta
	}

	return deltaSum
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	a, b := parseInput(lines)

	occurrences := make(map[int]int)

	for _, value := range b {
		if occurrences[value] != 0 {
			occurrences[value]++
		} else {
			occurrences[value] = 1
		}
	}

	similarityScore := 0
	for _, value := range a {
		appearences := occurrences[value]
		similarityScore += (value * appearences)
	}

	return similarityScore
}
