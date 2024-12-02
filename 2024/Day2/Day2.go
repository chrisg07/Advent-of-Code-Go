package AoCScaffold

import (
	_ "embed"
	"log"
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

func parseInput(lines []string) [][]int {
	reports := [][]int{}
	for _, line := range lines {
		parts := strings.Split(line, " ")
		levels := []int{}
		for _, part := range parts {
			level, _ := strconv.Atoi(part)
			levels = append(levels, level)
		}
		reports = append(reports, levels)
	}
	return reports
}

func isSafe(levels []int) bool {
	for i := 1; i < len(levels); i++ {
		previousLevel := levels[i-1]
		level := levels[i]
		delta := (previousLevel - level)
		if !(previousLevel <= level && (delta <= -1 && delta >= -3)) {
			return false
		}
	}
	return true
}

func determineSafety(levels []int) bool {
	levelsReversed := make([]int, len(levels))
	copy(levelsReversed, levels)
	Utils.ReverseArray(levelsReversed)

	safeFromLeftToRight := isSafe(levels)
	safeFromRightToLeft := isSafe(levelsReversed)

	safe := safeFromLeftToRight || safeFromRightToLeft
	if safe {
		log.Printf("[DEBUG] Safe report: %v", levels)
	}

	return safe
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	safeReports := 0
	for _, report := range input {
		reportCopy := make([]int, len(report))
		copy(reportCopy, report)

		Utils.ReverseArray(reportCopy)
		if determineSafety(report) {
			safeReports++
		}
	}
	return safeReports
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	safeReports := 0
	for _, report := range input {
		if determineSafety(report) {
			safeReports++
			continue
		}

		for index := range report {
			leftSlice := []int{}
			rightSlice := []int{}
			if index > 0 {
				leftSlice = report[:index]
			}
			if index <= len(report) {
				rightSlice = report[index+1:]
			}
			slice := []int{}
			slice = append(slice, leftSlice...)
			slice = append(slice, rightSlice...)

			log.Printf("[DEBUG] Checking safety of slice %v from report %v", slice, report)
			if determineSafety(slice) {
				safeReports++
				break
			}
		}
	}
	return safeReports
}
