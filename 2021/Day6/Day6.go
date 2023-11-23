package AoC2021

import (
	_ "embed"
	"log"
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

func Day62021(useExample bool, days int) int {
	lines := getInput(useExample)
	lanternfish := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, line := range lines {
		splitStrs := strings.Split(line, ",")
		for _, lanternfishStr := range splitStrs {
			lanternfishValue, _ := strconv.Atoi(lanternfishStr)
			lanternfish[lanternfishValue] += 1
		}
		log.Printf("[DEBUG] Initial state: %v", lanternfish)
	}

	for x := 0; x < days; x += 1 {
		lanternfish = [9]int{
			lanternfish[1],
			lanternfish[2],
			lanternfish[3],
			lanternfish[4],
			lanternfish[5],
			lanternfish[6],
			lanternfish[7] + lanternfish[0],
			lanternfish[8],
			lanternfish[0],
		}
	}

	log.Printf("[DEBUG] After %d days: %v", days, lanternfish)

	return lanternfish[0] + lanternfish[1] + lanternfish[2] + lanternfish[3] + lanternfish[4] + lanternfish[5] + lanternfish[6] + lanternfish[7] + lanternfish[8]
}
