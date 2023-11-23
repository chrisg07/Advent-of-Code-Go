package AoC2021

import (
	_ "embed"
	"fmt"
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

func Day6PartA2021(useExample bool) int {
	lines := getInput(useExample)
	lanternfish := []int{}
	for _, line := range lines {
		splitStrs := strings.Split(line, ",")
		for _, lanternfishStr := range splitStrs {
			lanternfishValue, _ := strconv.Atoi(lanternfishStr)
			lanternfish = append(lanternfish, lanternfishValue)
		}
		log.Printf("[DEBUG] Initial state: %v", lanternfish)
	}

	days := 80
	for x := 0; x < days; x += 1 {
		for index, fish := range lanternfish {
			if fish == 0 {
				lanternfish[index] = 7
				lanternfish = append(lanternfish, 8)
			}
			lanternfish[index] += -1
		}
		log.Printf("[DEBUG] After %d days: %v", x+1, lanternfish)

	}

	return len(lanternfish)
}

func Day6PartB2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			fmt.Print(string(char))
		}
		fmt.Println("")
	}

	return 0
}
