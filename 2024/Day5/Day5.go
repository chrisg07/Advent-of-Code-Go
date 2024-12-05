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

func isValidUpdate(rules map[int][]int, update []int) bool {
	valid := true
	for i, left := range update {
		for j, right := range update {
			if i != j && i < j {
				validUpdate := slices.Contains(rules[left], right) && !slices.Contains(rules[right], left)
				if !validUpdate {
					log.Printf("[DEBUG] Invalid update: %v", update)
					return false
				}
			}
		}
		if !valid {
			break
		}
	}

	return true
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	rules := parseOrderingRules(lines)
	pages := parsePagesToProduce(lines)

	sum := 0
	for _, update := range pages {
		if isValidUpdate(rules, update) {
			middle := len(update) / 2
			sum += update[middle]
			log.Printf("[DEBUG] Valid update: %v with middle %v", update, update[middle])
		}
	}
	return sum
}

func correctUpdate(rules map[int][]int, incorrectUpdate []int) []int {
	update := make([]int, len(incorrectUpdate))
	copy(update, incorrectUpdate)

	for i, left := range incorrectUpdate {
		for j, right := range incorrectUpdate {
			if i != j && i < j {
				validUpdate := slices.Contains(rules[left], right) && !slices.Contains(rules[right], left)
				if !validUpdate {
					tmp := update[i]
					update[i] = update[j]
					update[j] = tmp
					log.Printf("[DEBUG] Update after swaps: %v", update)
					return update
				}
			}
		}
	}
	log.Printf("[DEBUG] Update was already valid: %v", update)

	return update
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	rules := parseOrderingRules(lines)
	pages := parsePagesToProduce(lines)

	sum := 0
	for _, update := range pages {
		cpy := make([]int, len(update))
		copy(cpy, update)
		valid := isValidUpdate(rules, cpy)
		if !valid {
			validUpdate := false
			for !validUpdate {
				cpy = correctUpdate(rules, cpy)
				validUpdate = isValidUpdate(rules, cpy)
			}
			middle := len(cpy) / 2
			sum += cpy[middle]
			log.Printf("[DEBUG] Valid update: %v with middle %v", update, update[middle])

		}
	}
	return sum
}
