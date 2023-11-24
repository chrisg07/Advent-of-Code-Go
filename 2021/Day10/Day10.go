package AoC2021

import (
	_ "embed"
	"log"
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

func enqueue(queue []string, value string) []string {
	return append(queue, value)
}

func dequeue(queue []string) []string {
	if len(queue) > 0 {
		return queue[:len(queue)-1]
	} else {
		return queue
	}
}

func getIllegalCharScore(char string) int {
	switch char {
	case ")":
		return 3
	case "]":
		return 57
	case "}":
		return 1197
	case ">":
		return 25137
	default:
		return 0
	}
}

func Day10PartA2021(useExample bool) int {
	lines := getInput(useExample)
	score := 0
	for _, line := range lines {
		queue := []string{}
		for _, char := range line {
			chunk := string(char)
			// log.Printf("[WARN] Queue: %v\n", queue)
			isLeftChunk := chunk == "(" || chunk == "[" || chunk == "{" || chunk == "<"
			if isLeftChunk {
				queue = enqueue(queue, chunk)
			} else {
				// right chunk
				lastLeftChunk := queue[len(queue)-1]
				updatedQueue := dequeue(queue)
				queue = updatedQueue
				switch lastLeftChunk {
				case "(":
					if chunk != ")" {
						score += getIllegalCharScore(chunk)
						log.Printf("[WARN] %s - Expected ), but found %s instead.\n", line, chunk)
					}
				case "{":
					if chunk != "}" {
						score += getIllegalCharScore(chunk)
						log.Printf("[WARN] %s - Expected }, but found %s instead.\n", line, chunk)
					}
				case "[":
					if chunk != "]" {
						score += getIllegalCharScore(chunk)
						log.Printf("[WARN] %s - Expected ], but found %s instead.\n", line, chunk)
					}
				case "<":
					if chunk != ">" {
						score += getIllegalCharScore(chunk)
						log.Printf("[WARN] %s - Expected >, but found %s instead.\n", line, chunk)
					}
				}
			}

		}
	}

	return score
}

func Day10PartB2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
