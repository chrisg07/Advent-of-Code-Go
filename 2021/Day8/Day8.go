package AoC2021

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

func Day8PartA2021(useExample bool) int {
	lines := getInput(useExample)
	occurences := 0
	for _, line := range lines {
		lineParts := strings.Split(line, " | ")
		// signalPatterns := strings.Split(lineParts[0], " ")
		outputs := strings.Split(lineParts[1], " ")
		for _, output := range outputs {
			switch len(output) {
			case 2:
				occurences += 1
			case 4:
				occurences += 1
			case 3:
				occurences += 1
			case 7:
				occurences += 1

			}
		}
	}

	return occurences
}

func StringContainsOtherStringsChars(a string, b string) bool {
	for _, char := range b {
		subStr := string(char)
		if !strings.Contains(a, subStr) {
			return false
		}
	}
	return true
}

func DecodeLine(line string) int {
	lineParts := strings.Split(line, " | ")
	signalPatterns := strings.Split(lineParts[0], " ")
	outputs := strings.Split(lineParts[1], " ")

	patterns := append(signalPatterns, outputs...)
	encodings := [10]string{}

	for _, pattern := range patterns {
		switch len(pattern) {
		case 2:
			encodings[1] = pattern
		case 4:
			encodings[4] = pattern
		case 3:
			encodings[7] = pattern
		case 7:
			encodings[8] = pattern
		}
	}

	for _, pattern := range patterns {
		// is 9 if
		// pattern length is 6
		// pattern contains 4's pattern
		if len(pattern) == 6 && StringContainsOtherStringsChars(pattern, encodings[4]) {
			encodings[9] = pattern
		}
	}

	for _, pattern := range patterns {
		// is 3 if
		// pattern length is 5
		// pattern contains 1's pattern
		if len(pattern) == 5 && StringContainsOtherStringsChars(pattern, encodings[1]) {
			encodings[3] = pattern
		}
	}

	for _, pattern := range patterns {
		// is 5 if
		// pattern length is 5
		// pattern is contained in 9's pattern and not 3's pattern
		if len(pattern) == 5 && StringContainsOtherStringsChars(encodings[9], pattern) && !StringContainsOtherStringsChars(pattern, encodings[3]) {
			encodings[5] = pattern
		}
	}

	for _, pattern := range patterns {
		// is 2 if
		// pattern length is 5
		// and is not equal to the pattern for 3 or 5
		if len(pattern) == 5 && !StringContainsOtherStringsChars(pattern, encodings[3]) && !StringContainsOtherStringsChars(pattern, encodings[5]) {
			encodings[2] = pattern
		}
	}

	for _, pattern := range patterns {
		// is 6 if
		// pattern length is 6
		// pattern contains 5's pattern and not 9's pattern
		if len(pattern) == 6 && StringContainsOtherStringsChars(pattern, encodings[5]) && !StringContainsOtherStringsChars(pattern, encodings[9]) {
			encodings[6] = pattern
		}
	}

	for _, pattern := range patterns {
		// is 0 if
		// pattern length is 6
		// and is not equal to the pattern for 6 or 9
		if len(pattern) == 6 && !StringContainsOtherStringsChars(pattern, encodings[6]) && !StringContainsOtherStringsChars(pattern, encodings[9]) {
			encodings[0] = pattern
		}
	}

	Utils.PrettyPrint(encodings)

	// sum := 0
	outputValues := []int{}
	for _, output := range outputs {
		for encodedValue, encoding := range encodings {
			if StringContainsOtherStringsChars(output, encoding) && StringContainsOtherStringsChars(encoding, output) {
				outputValues = append(outputValues, encodedValue)
			}
		}
	}

	Utils.PrettyPrint(outputValues)

	return (outputValues[0] * 1000) + (outputValues[1] * 100) + (outputValues[2] * 10) + outputValues[3]
}

func Day8PartB2021(useExample bool) int {
	lines := getInput(useExample)

	sum := 0
	for _, line := range lines {
		sum += DecodeLine(line)
	}

	return sum
}
