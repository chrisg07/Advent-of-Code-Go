package AoC2021

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

type Range struct {
	start  int
	length int
}

func isInRange(value int, rangeToCheck Range) bool {
	return value >= rangeToCheck.start && value < (rangeToCheck.start+rangeToCheck.length)
}

type AlmanacMap struct {
	sources      []Range
	destinations []Range
}

func mapValue(value int, almanacMap AlmanacMap) int {
	if len(almanacMap.destinations) != len(almanacMap.sources) {
		panic("Map lengths don't match")
	}
	for index, sourceRange := range almanacMap.sources {
		if isInRange(value, sourceRange) {
			log.Printf("[WARN] Seed was in range\n")
			delta := value - sourceRange.start
			return almanacMap.destinations[index].start + delta
		}
	}
	return value
}

func Day5PartA2023(useExample bool) int {
	lines := getInput(useExample)

	seeds := []int{}

	almanac := []AlmanacMap{}

	for _, line := range lines {
		if strings.Contains(line, "seeds") {
			parts := strings.Split(line, ": ")
			numbers := strings.Split(parts[1], " ")
			for _, number := range numbers {
				value, _ := strconv.Atoi(number)
				seeds = append(seeds, value)
			}
		} else if strings.Contains(line, "map") {
			almanacMap := AlmanacMap{
				sources:      []Range{},
				destinations: []Range{},
			}
			almanac = append(almanac, almanacMap)
		} else if len(line) > 0 {
			valueStrings := strings.Split(line, " ")
			values := []int{}
			for _, str := range valueStrings {
				value, _ := strconv.Atoi(str)
				values = append(values, value)
			}
			// create source range
			sourceRange := Range{
				start:  values[1],
				length: values[2],
			}
			// create destination range
			destinationRange := Range{
				start:  values[0],
				length: values[2],
			}

			// append ranges to current map
			almanac[len(almanac)-1].sources = append(almanac[len(almanac)-1].sources, sourceRange)
			almanac[len(almanac)-1].destinations = append(almanac[len(almanac)-1].destinations, destinationRange)
		}
	}

	log.Printf("[WARN] Seeds: %v", seeds)
	log.Printf("[WARN] Almanac: %v", almanac)

	locations := []int{}

	for _, seed := range seeds {
		// map that seed through all almanac maps
		location := seed
		for _, almanacMap := range almanac {
			newValue := mapValue(location, almanacMap)
			log.Printf("[WARN] Seed mapped from %d to %d\n", location, newValue)
			location = newValue
		}
		locations = append(locations, location)
		log.Printf("[WARN] Seed mapped from %d to %d\n", seed, location)
	}

	return slices.Min(locations)
}

func Day5PartB2023(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
