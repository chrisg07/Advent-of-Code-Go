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

func SimulateLanternfish(lanternfish []int, days int) []int {
	fishes := append([]int{}, lanternfish...)
	chunkSize := 8
	for x := 0; x < days; x += chunkSize {
		for index, fish := range fishes {
			if fish == 0 {
				fishes[index] = 0
				fishes = append(fishes, 1)
			}
			fishes[index] += -1
		}
		log.Printf("[DEBUG] After %d days: %v", x+1, fishes)
	}

	return fishes
}

func Day62021(useExample bool, days int, chunkSize int) int {
	lines := getInput(useExample)
	lanternfish := [9]int{0, 0, 0, 0, 0, 0, 0, 0, 0}

	for _, line := range lines {
		splitStrs := strings.Split(line, ",")
		for _, lanternfishStr := range splitStrs {
			lanternfishValue, _ := strconv.Atoi(lanternfishStr)
			lanternfish[lanternfishValue] += 1
			// lanternfish = append(lanternfish, lanternfishValue)
		}
		log.Printf("[DEBUG] Initial state: %v", lanternfish)
	}

	// lanternfish = SimulateLanternfishInChunks(days, chunkSize, lanternfish)

	for x := 0; x < days; x += 1 {
		lanternfish = [9]int{
			lanternfish[1],
			lanternfish[2],
			lanternfish[3],
			lanternfish[4],
			lanternfish[5],
			lanternfish[6],
			lanternfish[7] + lanternfish[0],
			lanternfish[8],
			lanternfish[0],
		}
	}

	log.Printf("[DEBUG] After %d days: %v", days, lanternfish)

	return lanternfish[0] + lanternfish[1] + lanternfish[2] + lanternfish[3] + lanternfish[4] + lanternfish[5] + lanternfish[6] + lanternfish[7] + lanternfish[8]
}

func SimulateLanternfishInChunks(days int, chunkSize int, lanternfish []int) []int {
	lanternfishChunk0 := SimulateLanternfish([]int{0}, days/chunkSize)
	lanternfishChunk1 := SimulateLanternfish([]int{1}, days/chunkSize)
	lanternfishChunk2 := SimulateLanternfish([]int{2}, days/chunkSize)
	lanternfishChunk3 := SimulateLanternfish([]int{3}, days/chunkSize)
	lanternfishChunk4 := SimulateLanternfish([]int{4}, days/chunkSize)
	lanternfishChunk5 := SimulateLanternfish([]int{5}, days/chunkSize)
	lanternfishChunk6 := SimulateLanternfish([]int{6}, days/chunkSize)
	lanternfishChunk7 := SimulateLanternfish([]int{7}, days/chunkSize)
	lanternfishChunk8 := SimulateLanternfish([]int{8}, days/chunkSize)

	// length := 0
	for x := 0; x < chunkSize; x += 1 {
		fishes := []int{}
		for _, fish := range lanternfish {
			switch fish {
			case 0:
				fishes = append(fishes, lanternfishChunk0...)
			case 1:
				fishes = append(fishes, lanternfishChunk1...)
			case 2:
				fishes = append(fishes, lanternfishChunk2...)
			case 3:
				fishes = append(fishes, lanternfishChunk3...)
			case 4:
				fishes = append(fishes, lanternfishChunk4...)
			case 5:
				fishes = append(fishes, lanternfishChunk5...)
			case 6:
				fishes = append(fishes, lanternfishChunk6...)
			case 7:
				fishes = append(fishes, lanternfishChunk7...)
			case 8:
				fishes = append(fishes, lanternfishChunk8...)
			}
		}
		log.Printf("[DEBUG] After %d days: %v", (x+1)*(days/chunkSize), fishes)
		lanternfish = fishes
	}
	return lanternfish
}

// func Day6PartB2021(useExample bool) int {
// 	lines := getInput(useExample)
// 	for _, line := range lines {
// 		for _, char := range line {
// 			fmt.Print(string(char))
// 		}
// 		fmt.Println("")
// 	}

// 	return 0
// }
