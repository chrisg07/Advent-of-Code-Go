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

func ParseDigitFromString(str string) string {
	for _, char := range str {
		charStr := string(char)
		_, err := strconv.Atoi(charStr)
		if err == nil {
			return charStr
		}
	}
	return "false"
}

func ParseKeyFromString(str string) int {
	reverseStr := ""
	for i := len(str) - 1; i >= 0; i-- {
		reverseStr += string(str[i])
	}
	key, _ := strconv.Atoi(ParseDigitFromString(str) + ParseDigitFromString(reverseStr))
	return key
}

func Day1PartA2021(useExample bool) int {
	lines := getInput(useExample)
	sum := 0
	for _, line := range lines {
		value := ParseKeyFromString(line)
		sum += value
	}

	return sum
}

func Day1PartB2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
