package AoC2021

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
)

//go:embed inputs/example.txt
var exampleInput string

//go:embed inputs/complete.txt
var input string

type Cubes struct {
	red   int
	blue  int
	green int
}

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

func Day2PartA2021(useExample bool) int {
	lines := getInput(useExample)

	availableCubes := Cubes{
		red:   12,
		green: 13,
		blue:  14,
	}

	sumOfIDs := 0
	for _, line := range lines {
		parts := strings.Split(line, ": ")
		sets := strings.Split(parts[1], "; ")

		id := strings.Split(parts[0], " ")
		idValue, _ := strconv.Atoi(id[1])
		log.Printf("[WARN] Sets: %v\n", sets)
		setHadBeenAdded := false
		for _, set := range sets {
			cubes := strings.Split(set, ", ")
			if setHadBeenAdded {
				continue
			}
			for _, cube := range cubes {
				cubeParts := strings.Split(cube, " ")
				value, _ := strconv.Atoi(cubeParts[0])
				switch cubeParts[1] {
				case "red":
					if value > availableCubes.red {
						setHadBeenAdded = true
					}
				case "blue":
					if value > availableCubes.blue {
						setHadBeenAdded = true
					}
				case "green":
					if value > availableCubes.green {
						setHadBeenAdded = true
					}
				}
				if setHadBeenAdded {
					break
				}
				log.Printf("[WARN] Cube parts: %v\n", cubeParts)
			}
		}

		if !setHadBeenAdded {
			sumOfIDs += idValue
			log.Printf("[WARN] Adding %d to sum\n", idValue)
		}
		setHadBeenAdded = false

	}

	return sumOfIDs
}

func Day2PartB2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
