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

type Node struct {
	id       int
	value    int
	adjacent []*Node
	visited  bool
}

func createNode(id int, value int) *Node {
	return &Node{
		id:       id,
		value:    value,
		adjacent: []*Node{},
		visited:  false,
	}
}

var minRiskLevel = 1000000
var width = 0

func copySlice(sliceToCopy []*Node) []*Node {
	copy := make([]*Node, len(sliceToCopy))
	for k, v := range sliceToCopy {
		copy[k] = v
	}
	return copy
}

func dfsRecursivePartA(nodes []*Node, vertex *Node, path []*Node, totalRisk int) {
	if totalRisk > minRiskLevel || len(path) > minRiskLevel {
		return
	}
	if vertex.id != 0 {
		totalRisk += vertex.value
	}
	path = append(path, vertex)
	vertex.visited = true
	if vertex.id == len(nodes)-1 {
		if totalRisk < minRiskLevel {
			minRiskLevel = totalRisk
		}
		return
	}

	nodesCopy := copySlice(nodes)
	for _, neighbor := range vertex.adjacent {
		pathContainsNeighbor := false
		// need to use map for quicker look ups
		for _, visited := range path {
			if visited.id == neighbor.id {
				pathContainsNeighbor = true
			}
		}
		if !pathContainsNeighbor {
			if totalRisk < minRiskLevel {
				dfsRecursivePartA(nodesCopy, neighbor, path, totalRisk)
			}
		}
	}

}

func Day15PartA2021(useExample bool) int {
	lines := getInput(useExample)

	// build array of nodes
	nodes := []*Node{}

	for _, line := range lines {
		width = len(line)
		for _, char := range line {
			value, _ := strconv.Atoi(string(char))
			node := createNode(len(nodes), value)
			nodes = append(nodes, node)
		}
	}

	log.Printf("[DEBUG] Nodes: %v", nodes)
	// build neighbors based on indexes in the array

	for index, node := range nodes {
		adjacentIndexOffsets := []int{}
		if index%width == 0 {
			adjacentIndexOffsets = []int{-width, 1, width}
		} else if (index+1)%width == 0 {
			adjacentIndexOffsets = []int{-width, -1, width}
		} else {
			adjacentIndexOffsets = []int{-width, -1, 1, width}
		}

		for _, offset := range adjacentIndexOffsets {
			if index+offset >= 0 && index+offset < len(nodes) {
				nodes[index].adjacent = append(node.adjacent, nodes[index+offset])
			}
		}
	}
	// dfs, push path cost to array
	nodesCopy := copySlice(nodes)
	dfsRecursivePartA(nodesCopy, nodesCopy[0], []*Node{}, 0)

	return minRiskLevel
}

func Day15PartB2021(useExample bool) int {
	lines := getInput(useExample)
	for _, line := range lines {
		for _, char := range line {
			log.Print(string(char))
		}
		log.Println("")
	}

	return 0
}
