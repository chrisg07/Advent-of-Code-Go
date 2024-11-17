package AoC2023

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

type Node struct {
	id      string
	leftId  string
	left    *Node
	rightId string
	right   *Node
}

func createNode(id string, leftId string, rightId string) *Node {
	return &Node{id: id, leftId: leftId, rightId: rightId}
}

func Day8PartA2023(useExample bool) int {
	lines := getInput(useExample)
	nodes := make(map[string]*Node)
	instructions := ""
	for index, line := range lines {
		if index == 0 {
			instructions = line
		}
		if index > 1 {
			parts := strings.Split(line, " = (")
			id := parts[0]
			rightParts := strings.Split(parts[1], ", ")
			leftId := rightParts[0][:3]
			rightId := rightParts[1][:3]
			if nodes[id] == nil {
				nodes[id] = createNode(id, leftId, rightId)
			}
			log.Printf("[WARN] Node id: %s with left: %s and right: %s\n", id, leftId, rightId)
		}

	}
	log.Printf("[WARN] Nodes: %v\n", nodes)
	// build tree
	for _, node := range nodes {
		node.left = nodes[node.leftId]
		node.right = nodes[node.rightId]
	}

	log.Printf("[WARN] Nodes: %v\n", nodes)

	zNodeFound := false
	currentNode := nodes["AAA"]
	steps := 0
	for !zNodeFound {
		for _, instructionRune := range instructions {
			instruction := string(instructionRune)
			if instruction == "R" {
				currentNode = currentNode.right
			} else if instruction == "L" {
				currentNode = currentNode.left
			}
			steps += 1
		}
		if currentNode.id == "ZZZ" {
			zNodeFound = true
		}
	}
	return steps
}

func Day8PartB2023(useExample bool) int {
	lines := getInput(useExample)
	nodes := make(map[string]*Node)
	instructions := ""
	for index, line := range lines {
		if index == 0 {
			instructions = line
		}
		if index > 1 {
			parts := strings.Split(line, " = (")
			id := parts[0]
			rightParts := strings.Split(parts[1], ", ")
			leftId := rightParts[0][:3]
			rightId := rightParts[1][:3]
			if nodes[id] == nil {
				nodes[id] = createNode(id, leftId, rightId)
			}
		}

	}
	// build tree
	for _, node := range nodes {
		node.left = nodes[node.leftId]
		node.right = nodes[node.rightId]
	}

	// create map of nodes that end in "A"
	aNodes := make(map[string]*Node)
	for _, node := range nodes {
		if node.id[2] == 'A' {
			aNodes[node.id] = node
		}
	}

	log.Printf("[WARN] Nodes that end in A: %v\n", aNodes)

	steps := []int{}
	for _, node := range aNodes {
		steps = append(steps, 0)
		currentNode := node
		zNodeFound := false
		for !zNodeFound {
			for _, instructionRune := range instructions {
				instruction := string(instructionRune)
				if instruction == "R" {
					currentNode = currentNode.right
				} else if instruction == "L" {
					currentNode = currentNode.left
				}
				steps[len(steps)-1] += 1
			}
			if currentNode.id[2] == 'Z' {
				zNodeFound = true
				continue
			}
		}
	}

	log.Printf("[WARN] Steps required for nodes that end in A: %v\n", steps)

	lcm := Utils.LCM(steps)
	return lcm
}
