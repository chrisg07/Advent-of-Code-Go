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

type octopus struct {
	energy     int
	neighbors  []*octopus
	hasFlashed bool
	x          int
	y          int
}

func newOctopus(energy int, x int, y int) *octopus {
	cephalopod := octopus{
		energy:     energy,
		neighbors:  []*octopus{},
		hasFlashed: false,
		x:          x,
		y:          y,
	}
	return &cephalopod
}

func printOctopi(octopi []*octopus) {
	for row := 0; row < 100; row += 10 {
		printStr := ""
		for column := 0; column < 10; column++ {
			value := strconv.Itoa(octopi[row+column].energy)
			printStr += value
		}
		log.Printf("[WARN] |  %s", printStr)
	}

	log.Print("[WARN]\n")
}

func flashOctopus(cephalopod *octopus) *octopus {
	if cephalopod.energy > 9 && !cephalopod.hasFlashed {
		cephalopod.hasFlashed = true
		for x := range cephalopod.neighbors {
			if !cephalopod.neighbors[x].hasFlashed {
				cephalopod.neighbors[x].energy += 1
			}

			if cephalopod.neighbors[x].energy > 9 && !cephalopod.neighbors[x].hasFlashed {
				flashOctopus(cephalopod.neighbors[x])
			}
		}
	}
	return cephalopod
}

func Day11PartA2021(useExample bool) int {
	lines := getInput(useExample)
	octopi := []*octopus{}
	flashes := 0

	octopi = setState(lines, octopi)

	buildNeighbors(octopi)

	for step := 0; step < 100; step++ {
		flashes = takeStep(step, octopi, flashes)
	}
	return flashes
}

func takeStep(step int, octopi []*octopus, flashes int) int {
	log.Printf("[WARN] |  After step %d:\n\n", step)
	printOctopi(octopi)

	// step 1, increase value of all octopi by 1
	for index, _ := range octopi {
		octopi[index].energy += 1
	}

	// step 2, any octopus with an energy level of 9 or above flashes
	for index, _ := range octopi {
		if octopi[index].energy > 9 && !octopi[index].hasFlashed {
			flashOctopus(octopi[index])
		}
	}

	// step 3, any octopus that has flashed has it's energy set to 0
	for index, _ := range octopi {
		if octopi[index].hasFlashed {
			flashes += 1
			octopi[index].hasFlashed = false
			octopi[index].energy = 0
		}
	}
	return flashes
}

func setState(lines []string, octopi []*octopus) []*octopus {
	for x, line := range lines {
		for y, char := range line {
			energy, _ := strconv.Atoi(string(char))
			cephalopod := newOctopus(energy, x, y)

			octopi = append(octopi, cephalopod)
		}
	}
	return octopi
}

func buildNeighbors(octopi []*octopus) {
	for index, cephalopod := range octopi {
		neighborIndexOffsets := []int{}
		if index%10 == 0 {
			neighborIndexOffsets = []int{-10, -9, 1, 10, 11}
		} else if (index+1)%10 == 0 {
			neighborIndexOffsets = []int{-11, -10, -1, 9, 10}
		} else {
			neighborIndexOffsets = []int{-11, -10, -9, -1, 1, 9, 10, 11}
		}

		for _, offset := range neighborIndexOffsets {
			if index+offset >= 0 && index+offset < len(octopi) {
				octopi[index].neighbors = append(cephalopod.neighbors, octopi[index+offset])
			}
		}
	}
}

func Day11PartB2021(useExample bool) int {
	lines := getInput(useExample)
	octopi := []*octopus{}

	octopi = setState(lines, octopi)
	buildNeighbors(octopi)

	step := 0
	notAllOctopiHaveFlashed := true
	for notAllOctopiHaveFlashed {
		flashes := takeStep(step, octopi, 0)
		step += 1

		if flashes == 100 {
			notAllOctopiHaveFlashed = false
		}
	}

	return step
}
