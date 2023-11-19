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

func PartB(useExample bool) int {
	// lines := getInput(useExample)

	// for _, line := range lines {
	// }

	return 0
}
