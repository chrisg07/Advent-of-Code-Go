package AoCScaffold

import (
	_ "embed"
	"log"
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

func parseOrderingRules(lines []string) map[int][]int {
	rules := make(map[int][]int)
	for _, line := range lines {
		if strings.Contains(line, "|") {
			nums := strings.Split(line, "|")
			left, _ := strconv.Atoi(nums[0])
			right, _ := strconv.Atoi(nums[1])
			rules[left] = append(rules[left], right)
		}
	}
	return rules
}

func parsePagesToProduce(lines []string) [][]int {
	pages := [][]int{}
	for _, line := range lines {
		if strings.Contains(line, ",") {
			nums := strings.Split(line, ",")
			ints := []int{}
			for _, num := range nums {
				val, _ := strconv.Atoi(num)
				ints = append(ints, val)
			}
			pages = append(pages, ints)
		}
	}
	return pages
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	rules := parseOrderingRules(lines)
	pages := parsePagesToProduce(lines)
	log.Printf("[DEBUG] Rules: %v", rules)
	log.Printf("[DEBUG] Updates: %v", pages)

	sum := 0
	for _, update := range pages {
		valid := true
		for i, left := range update {
			for j, right := range update {
				if i != j && i < j {
					log.Printf("[DEBUG] Checking: %v against %v", left, right)
					log.Printf("[DEBUG] Rules left: %v", rules[left])
					log.Printf("[DEBUG] Rules right: %v", rules[right])

					// if rules[right] != 0 && rules[right] == left {
					// 	valid = false
					// 	break
					// }
					validUpdate := slices.Contains(rules[left], right) && !slices.Contains(rules[right], left)
					if !validUpdate {
						log.Printf("[DEBUG] ----> Invalid update: %v", update)
						valid = false
						break
					}
				}
			}
			if !valid {
				break
			}
		}
		if valid {
			middle := len(update) / 2
			sum += update[middle]
			log.Printf("[DEBUG] Valid update: %v with middle %v", update, update[middle])
		}
	}
	return sum
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	// input := parseInput(lines)

	return len(lines)
}
