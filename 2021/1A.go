package AoC2021

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"strconv"
)

func Day1PartA() {
	utils.FormatMessage("Day 1 Part A:")

	lines := utils.ReadAoCInput("./2021/inputs/Day1/PartAExample.txt")

	previousDepth := 10000
	descents := 0

	for _, line := range lines {
		depth, parseErr := strconv.Atoi(line)

		if parseErr != nil {
			log.Fatal(parseErr)
		}

		didDescend := depth > previousDepth
		status := "decreased"

		if didDescend {
			descents += 1
			status = "increased"
		}

		previousDepth = depth
		fmt.Printf("%s (%s)\n", line, status)
	}

	fmt.Printf("Number of descents: %d\n", descents)

}
