package AoC2021

import (
	_ "embed"
	"log"
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
	distanceAbleToBeTraveled := time * velocity
	if distanceAbleToBeTraveled > distanceToBeat {
		return 1
	} else {
		return 0
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

	marginOfError := findMarginOfError(times, distances)

	return marginOfError
}

func findMarginOfError(times []int, distances []int) int {
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
	marginOfError := Utils.MultiplyArray(waysToBeatRecords)
	return marginOfError
}

func Day6PartB2023(useExample bool) int {
	lines := getInput(useExample)
	times := []int{}
	distances := []int{}

	for _, line := range lines {
		if strings.Contains(line, "Time") {
			parts := strings.Split(line, ":")
			splitTimes := strings.Split(parts[1], " ")
			singleTime := strings.Join(splitTimes, "")
			time, _ := strconv.Atoi(singleTime)
			times = append(times, time)
		} else if strings.Contains(line, "Distance") {
			parts := strings.Split(line, ":")
			splitRecords := strings.Split(parts[1], " ")
			singleRecord := strings.Join(splitRecords, "")
			distance, _ := strconv.Atoi(singleRecord)
			distances = append(distances, distance)
		}
	}

	marginOfError := findMarginOfError(times, distances)

	return marginOfError
}
