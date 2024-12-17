package AoCScaffold

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
		unsplitLines = exampleInput
	} else {
		unsplitLines = input
	}

	// Normalize line endings: replace all \r\n with \n
	unsplitLines = strings.ReplaceAll(unsplitLines, "\r\n", "\n")

	// Trim any trailing newlines (whether \n or \r\n)
	unsplitLines = strings.TrimRight(unsplitLines, "\n")

	log.Printf("[DEBUG] Raw input lines: %q", unsplitLines)

	// Split into lines by \n
	lines = strings.Split(unsplitLines, "\n")

	return lines
}

func evaluate(left uint64, nums []uint64, results chan uint64) {
	// log.Printf("[DEBUG] Left: %v Nums: %v Results: %v", left, nums, len(results))
	sum := left
	product := left

	if len(nums) > 1 {
		sum += nums[0]
		evaluate(sum, nums[1:], results)

		product *= nums[0]
		evaluate(product, nums[1:], results)
	}

	if len(nums) == 1 {
		sum += nums[0]
		results <- sum

		product *= nums[0]
		results <- product
	}
}

func parseInput(lines []string) uint64 {
	var sum uint64 = 0

	for _, line := range lines {
		log.Printf("[DEBUG] Evaluating: %v", line)
		sides := strings.Split(line, ": ")
		left, err := strconv.ParseUint(sides[0], 10, 64)
		if err != nil {
			log.Fatalf("[ERROR] Failed to parse left side: %v", err)
		}
		rightNumbers := strings.Fields(sides[1])

		log.Printf("[DEBUG] Right numbers: %v", rightNumbers)
		nums := []uint64{}

		for _, str := range rightNumbers {
			num, _ := strconv.Atoi(str)
			nums = append(nums, uint64(num))
		}

		// log.Printf("[DEBUG] Evaluating: %v", left)
		// first operation
		sumLeft := nums[0] + nums[1]
		// log.Printf("[DEBUG] First sum: %v", sumLeft)
		productLeft := nums[0] * nums[1]
		// log.Printf("[DEBUG] First product: %v", productLeft)

		calibrated := false
		if len(nums) > 2 {
			nums := nums[2:]
			results := make(chan uint64, 1000000)
			evaluate(sumLeft, nums, results)
			evaluate(productLeft, nums, results)

			// time.Sleep(5 * time.Millisecond)
			close(results)

			for result := range results {
				if result == uint64(left) {
					calibrated = true
					break
				}
			}

		} else {
			calibrated = sumLeft == uint64(left) || productLeft == uint64(left)
		}

		if calibrated {
			log.Printf("[DEBUG] Calibrated result: %v", line)
			log.Printf("[DEBUG] Previous sum: %v", sum)
			sum += uint64(left)
			log.Printf("[DEBUG] Updated sum : %v", sum)
		}

	}

	return sum
}

func PartA(useExample bool) uint64 {
	lines := getInput(useExample)
	result := parseInput(lines)

	return result
}

func PartB(useExample bool) uint64 {
	lines := getInput(useExample)
	input := parseInput(lines)

	return input
}
