package AoC2021

import (
	_ "embed"
	"fmt"
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

func PartA(useExample bool) int {
	lines := getInput(useExample)

	rate := []int{}

	for _, line := range lines {
		for y, char := range line {
			if y > len(rate)-1 {
				rate = append(rate, 0)
			}

			bit, _ := strconv.Atoi(string(char))

			if bit == 0 {
				rate[y] -= 1
			} else {
				rate[y] += 1
			}
		}
	}

	for index, value := range rate {
		if value > 0 {
			rate[index] = 1
		} else {
			rate[index] = 0
		}
	}

	fmt.Printf("Rate: %v\n", rate)

	gammaRateValue := Utils.BitArrayToDecimal(rate)
	epsilonRate := Utils.FlipBitArray(rate)

	fmt.Printf("Epsilon rate: %v\n", epsilonRate)
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

func PartB(useExample bool) int {
	lines := getInput(useExample)

	rate := []int{}
	readings := [][]int{}

	for x, line := range lines {
		readings = append(readings, []int{})
		for y, char := range line {
			if y > len(rate)-1 {
				rate = append(rate, 0)
			}

			bit, _ := strconv.Atoi(string(char))
			readings[x] = append(readings[x], bit)

			if bit == 0 {
				rate[y] -= 1
			} else {
				rate[y] += 1
			}
		}
	}

	fmt.Printf("Rate: %v\n", rate)
	for index, value := range rate {
		if value >= 0 {
			rate[index] = 1
		} else {
			rate[index] = 0
		}
	}

	oxygenReadings := Utils.DuplicateMatrix[int](readings)
	for index := range rate {
		if len(oxygenReadings) != 1 {
			currentRate := CalculateRate(oxygenReadings)
			oxygenReadings = ReduceReadings(oxygenReadings, currentRate, index)
			fmt.Printf("Oxygen readings: %v\n", oxygenReadings)
		}
	}
	oxygenRating := Utils.BitArrayToDecimal(oxygenReadings[0])
	fmt.Printf("Oxygen rating: %d\n", oxygenRating)

	scrubberReadings := Utils.DuplicateMatrix[int](readings)
	flippedRate := Utils.FlipBitArray(rate)
	for index := range flippedRate {
		if len(scrubberReadings) != 1 {
			currentRate := CalculateRate(scrubberReadings)
			currentFlippedRate := Utils.FlipBitArray(currentRate)
			scrubberReadings = ReduceReadings(scrubberReadings, currentFlippedRate, index)
			fmt.Printf("Scrubber readings: %v\n", scrubberReadings)
		}
	}
	scrubberRating := Utils.BitArrayToDecimal(scrubberReadings[0])
	fmt.Printf("Scrubber rating: %d\n", scrubberRating)

	return oxygenRating * scrubberRating
}
