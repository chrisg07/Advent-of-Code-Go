package AoC2021

import (
	_ "embed"
	"log"
	"sort"
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

func Day14PartA2021(useExample bool, steps int) int {
	lines := getInput(useExample)
	insertionRules := make(map[string]string)
	pairQuantities := make(map[string]int)
	elementQuantities := make(map[string]int)

	for _, line := range lines {
		if strings.Contains(line, "->") {
			// add pair rule
			parts := strings.Split(line, " -> ")
			insertionRules[parts[0]] = parts[1]
		} else if len(line) > 0 {
			// count element quanities
			for _, char := range line {
				element := string(char)
				elementQuantities[element] += 1
			}
			// populate pairs
			for i := 0; i < len(line)-1; i++ {
				pair := line[i : i+2]
				pairQuantities[pair] += 1
			}
		}
	}

	log.Printf("[WARN] Pair mappings: %v", insertionRules)
	// take 10 steps
	for i := 0; i < steps; i++ {
		log.Printf("[WARN] Step %d\n", i)
		log.Printf("[WARN] Element quantities: %v\n", elementQuantities)
		log.Printf("[DEBUG] Pair quantities: %v\n", pairQuantities)
		pairQuantities, elementQuantities = step(pairQuantities, elementQuantities, insertionRules)
	}

	quantities := []int{}
	for _, value := range elementQuantities {
		quantities = append(quantities, value)
	}

	sort.Ints(quantities)

	log.Printf("[WARN] Element quantities after insertions: %v\n", quantities)

	quanityOfLeastCommonElement := quantities[0]
	quanityOfMostCommonElement := quantities[len(quantities)-1]
	return quanityOfMostCommonElement - quanityOfLeastCommonElement
}

func copyMap(mapToCopy map[string]int) map[string]int {
	copy := make(map[string]int)
	for k, v := range mapToCopy {
		copy[k] = v
	}
	return copy
}

func step(pairQuantities map[string]int, elementQuantities map[string]int, insertionRules map[string]string) (map[string]int, map[string]int) {
	updatedPairQuantities := make(map[string]int)
	updatedElementQuantities := copyMap(elementQuantities)
	for key, value := range pairQuantities {
		elementToInsert := insertionRules[key]
		log.Printf("[DEBUG] Inserting element: %s for pair: %s %d times\n", elementToInsert, key, value)

		// update pair quantities
		leftPair := string(key[0]) + elementToInsert
		rightPair := elementToInsert + string(key[1])
		log.Printf("[DEBUG] Adding pairs: %s and %s\n", leftPair, rightPair)

		updatedPairQuantities[leftPair] += value
		updatedPairQuantities[rightPair] += value

		// update element quantities
		updatedElementQuantities[elementToInsert] += value
	}

	return updatedPairQuantities, updatedElementQuantities
}

func Day14PartB2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
