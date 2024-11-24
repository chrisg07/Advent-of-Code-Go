package AoCScaffold

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

func parseInput(lines []string) []int {
	input := []int{}

	for _, line := range lines {
		parts := strings.Split(line, "-")
		lowBound, _ := strconv.Atoi(parts[0])
		highBound, _ := strconv.Atoi(parts[1])
		input = append(input, lowBound, highBound)
	}

	return input
}

func PasswordHasAdjacentDuplicates(password int) bool {
	passwordString := strconv.Itoa(password)

	for i := 0; i < len(passwordString)-1; i++ {
		if passwordString[i] == passwordString[i+1] {
			return true
		}
	}

	return false
}

func PasswordHasNoDecreasingDigits(password int) bool {
	passwordString := strconv.Itoa(password)

	for i := 0; i < len(passwordString)-1; i++ {
		leftDigit, _ := strconv.Atoi(string(passwordString[i]))
		rightDigit, _ := strconv.Atoi(string(passwordString[i+1]))
		if rightDigit < leftDigit {
			return false
		}
	}

	return true
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	potentialPasswords := []int{}

	for password := input[0]; password < input[1]; password++ {
		if IsPotentialPassword(password) {
			potentialPasswords = append(potentialPasswords, password)
		}
	}

	log.Printf("[CONSOLE] potential passwords: %v", potentialPasswords)
	return len(potentialPasswords)
}

func IsPotentialPassword(password int) bool {
	hasAdjacentDuplicates := PasswordHasAdjacentDuplicates(password)
	hasNoDecreasingDigits := PasswordHasNoDecreasingDigits(password)

	return hasAdjacentDuplicates && hasNoDecreasingDigits
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return len(input)
}
