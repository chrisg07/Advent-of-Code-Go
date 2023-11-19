package AoC2021

import (
	_ "embed"
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
	if useExample {
		exampleInput = strings.TrimRight(exampleInput, "\n")
		lines = strings.Split(exampleInput, "\n")
	} else {
		input = strings.TrimRight(input, "\n")
		lines = strings.Split(input, "\n")
	}
	return lines
}

func convertInputToReadings(input []string) [][]int {
	readings := [][]int{}

	for x, line := range input {
		readings = append(readings, []int{})
		for _, char := range line {
			bit, _ := strconv.Atoi(string(char))
			readings[x] = append(readings[x], bit)
		}
	}
	return readings
}

func PartA(useExample bool) int {
	lines := getInput(useExample)

	readings := convertInputToReadings(lines)

	rate := CalculateRate(readings)

	gammaRateValue := Utils.BitArrayToDecimal(rate)
	epsilonRate := Utils.FlipBitArray(rate)

	epsilonRateValue := Utils.BitArrayToDecimal(epsilonRate)
	return gammaRateValue * epsilonRateValue
}

func CalculateRate(readings [][]int) []int {
	rate := []int{}

	for _, reading := range readings {
		for y, bit := range reading {
			if y > len(rate)-1 {
				rate = append(rate, 0)
			}

			if bit == 0 {
				rate[y] -= 1
			} else {
				rate[y] += 1
			}
		}
	}

	for index, value := range rate {
		if value >= 0 {
			rate[index] = 1
		} else {
			rate[index] = 0
		}
	}
	return rate
}

func ReduceReadings(readings [][]int, bits []int, index int) [][]int {
	matchingReadings := [][]int{}

	for _, reading := range readings {
		if reading[index] == bits[index] {
			matchingReadings = append(matchingReadings, reading)
		}
	}

	return matchingReadings
}

func calculateOxygenRating(readings [][]int) int {
	oxygenReadings := Utils.DuplicateMatrix[int](readings)
	rate := CalculateRate(oxygenReadings)
	for index := range rate {
		if len(oxygenReadings) != 1 {
			currentRate := CalculateRate(oxygenReadings)
			oxygenReadings = ReduceReadings(oxygenReadings, currentRate, index)
		}
	}
	oxygenRating := Utils.BitArrayToDecimal(oxygenReadings[0])
	return oxygenRating
}

func calculateScrubberRating(readings [][]int) int {
	scrubberReadings := Utils.DuplicateMatrix[int](readings)
	rate := CalculateRate(scrubberReadings)
	flippedRate := Utils.FlipBitArray(rate)
	for index := range flippedRate {
		if len(scrubberReadings) != 1 {
			currentRate := CalculateRate(scrubberReadings)
			currentFlippedRate := Utils.FlipBitArray(currentRate)
			scrubberReadings = ReduceReadings(scrubberReadings, currentFlippedRate, index)

		}
	}
	scrubberRating := Utils.BitArrayToDecimal(scrubberReadings[0])
	return scrubberRating
}

func PartB(useExample bool) int {
	lines := getInput(useExample)

	readings := convertInputToReadings(lines)

	oxygenRating := calculateOxygenRating(readings)

	scrubberRating := calculateScrubberRating(readings)

	return oxygenRating * scrubberRating
}
