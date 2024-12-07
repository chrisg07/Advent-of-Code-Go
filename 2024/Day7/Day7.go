package AoCScaffold

import (
	_ "embed"
	"log"
	"strconv"
	"strings"
	"time"
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

func evaluate(left int, nums []int, results chan int) {
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

func parseInput(lines []string) int {
	sum := 0

	for _, line := range lines {
		log.Printf("[DEBUG] Evaluating: %v", line)
		sides := strings.Split(line, ": ")
		left, _ := strconv.Atoi(sides[0])
		rightNumbers := strings.Split(sides[1][:len(sides[1])-1], " ")
		nums := []int{}

		for _, str := range rightNumbers {
			num, _ := strconv.Atoi(str)
			nums = append(nums, num)
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
			results := make(chan int, 10000)
			evaluate(sumLeft, nums, results)
			evaluate(productLeft, nums, results)

			time.Sleep(10 * time.Millisecond)
			close(results)

			for result := range results {
				if result == left {
					calibrated = true
					break
				}
			}

		} else {
			calibrated = sumLeft == left || productLeft == left
		}

		if calibrated {
			log.Printf("[DEBUG] Calibrated result: %v", line)
			log.Printf("[DEBUG] Previous sum: %v", sum)
			sum += left
			log.Printf("[DEBUG] Updated sum : %v", sum)
		}

	}

	return sum
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	result := parseInput(lines)

	return result
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return input
}
