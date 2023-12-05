package AoC2021

import (
	_ "embed"
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
	for index, sourceRange := range almanacMap.sources {
		if isInRange(value, sourceRange) {
			delta := value - sourceRange.start
			return almanacMap.destinations[index].start + delta
		}
	}
	return value
}

func Day5PartA2023(useExample bool) int {
	lines := getInput(useExample)

	seeds, almanac := buildState(lines)

	locations := []int{}

	for _, seed := range seeds {
		location := seed
		for _, almanacMap := range almanac {
			location = mapValue(location, almanacMap)

		}
		locations = append(locations, location)
	}

	return slices.Min(locations)
}

func buildState(lines []string) ([]int, []AlmanacMap) {
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

			sourceRange := Range{
				start:  values[1],
				length: values[2],
			}

			destinationRange := Range{
				start:  values[0],
				length: values[2],
			}

			almanac[len(almanac)-1].sources = append(almanac[len(almanac)-1].sources, sourceRange)
			almanac[len(almanac)-1].destinations = append(almanac[len(almanac)-1].destinations, destinationRange)
		}
	}
	return seeds, almanac
}

func Day5PartB2023(useExample bool) int {
	lines := getInput(useExample)

	seeds, almanac := buildState(lines)

	seedRanges := []Range{}
	lastSeed := 0
	for i := 0; i < len(seeds); i += 2 {
		seedRange := Range{
			start:  seeds[i],
			length: seeds[i+1],
		}
		seedRanges = append(seedRanges, seedRange)
		if seedRange.start+seedRange.length > lastSeed {
			lastSeed = seedRange.start + seedRange.length
		}
	}

	minLocation := 1000000000000

	for _, seedRange := range seedRanges {
		for i := seedRange.start; i < seedRange.start+seedRange.length; i++ {
			location := i
			for _, almanacMap := range almanac {
				location = mapValue(location, almanacMap)
			}
			if location < minLocation {
				minLocation = location
			}
		}
	}
	return minLocation
}
