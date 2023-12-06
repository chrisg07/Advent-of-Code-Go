package AoC2021

import (
	_ "embed"
	"log"
	"strings"

	"github.com/chrisg07/Advent-of-Code-Go/Utils"
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

func solveRace(time int, distanceTraveled int, velocity int, distanceToBeat int) int {
	if distanceTraveled > distanceToBeat {
		return 1
	} else if time == 0 {
		return 0
	} else {
		return solveRace(time-1, distanceTraveled+velocity, velocity, distanceToBeat)
	}
}

func Day6PartA2023(useExample bool) int {
	lines := getInput(useExample)
	times := []int{}
	distances := []int{}
	for _, line := range lines {
		if strings.Contains(line, "Time") {
			parts := strings.Split(line, ":")
			times = Utils.ArrayFromString(parts[1], " ")
		} else if strings.Contains(line, "Distance") {
			parts := strings.Split(line, ":")
			distances = Utils.ArrayFromString(parts[1], " ")
		}
	}

	log.Printf("[WARN] Times: %v\n", times)
	log.Printf("[WARN] Distances: %v\n", distances)

	waysToBeatRecords := []int{}
	for index, time := range times {
		distance := distances[index]
		waysToBeatTheRecord := 0
		for i := 0; i < time; i++ {
			waysToBeatTheRecord += solveRace(time-i, 0, i, distance)
		}
		waysToBeatRecords = append(waysToBeatRecords, waysToBeatTheRecord)
	}

	log.Printf("[WARN] Ways to beat the record in each race: %v\n", waysToBeatRecords)

	return Utils.MultiplyArray(waysToBeatRecords)
}

func Day6PartB2023(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
