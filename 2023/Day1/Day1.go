package AoC2021

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

func ParseDigitFromString(str string) string {
	for _, char := range str {
		_, err := strconv.Atoi(string(char))
		if err == nil {
			return string(char)
		}
	}
	return "false"
}

func ParseKeyFromString(str string) int {
	reverseStr := Utils.ReverseString(str)
	key, _ := strconv.Atoi(ParseDigitFromString(str) + ParseDigitFromString(reverseStr))
	return key
}

func Day1PartA2023(useExample bool) int {
	lines := getInput(useExample)
	sum := 0
	for _, line := range lines {
		value := ParseKeyFromString(line)
		sum += value
	}

	return sum
}

func GetDigitWordFromString(str string) string {
	log.Printf("[WARN] Getting digit from string: %s", str)
	if strings.Contains(str, "one") {
		return "1"
	} else if strings.Contains(str, "two") {
		return "2"
	} else if strings.Contains(str, "three") {
		return "3"
	} else if strings.Contains(str, "four") {
		return "4"
	} else if strings.Contains(str, "five") {
		return "5"
	} else if strings.Contains(str, "six") {
		return "6"
	} else if strings.Contains(str, "seven") {
		return "7"
	} else if strings.Contains(str, "eight") {
		return "8"
	} else if strings.Contains(str, "nine") {
		return "9"
	}
	return "false"
}

func ParseDigitAndWordsFromString(str string) string {
	for index, char := range str {
		charStr := string(char)
		slice := string(str[:index])
		wordFromString := GetDigitWordFromString(slice)
		log.Printf("[WARN] Digit in word from from string: %s", wordFromString)
		_, part1Err := strconv.Atoi(charStr)
		_, part2Err := strconv.Atoi(wordFromString)
		if part2Err == nil {
			return wordFromString
		}
		if part1Err == nil {
			return charStr
		}
	}
	return "false"
}
func ParseDigitAndWordsFromReverseString(str string) string {
	for i := len(str) - 1; i >= 0; i-- {
		charStr := string(str[i])
		slice := string(str[i:])
		wordFromString := GetDigitWordFromString(slice)
		log.Printf("[WARN] Digit in word from string: %s", wordFromString)
		_, part1Err := strconv.Atoi(charStr)
		_, part2Err := strconv.Atoi(wordFromString)
		if part2Err == nil {
			return wordFromString
		}
		if part1Err == nil {
			return charStr
		}
	}
	return "false"
}

func ParseKeyFromStringPartB(str string) int {
	key, _ := strconv.Atoi(ParseDigitAndWordsFromString(str) + ParseDigitAndWordsFromReverseString(str))
	return key
}

func Day1PartB2023(useExample bool) int {
	lines := getInput(useExample)
	sum := 0
	for _, line := range lines {
		value := ParseKeyFromStringPartB(line)
		sum += value
	}

	return sum
}
