package AoC2021

import (
	_ "embed"
	"log"
	"slices"
	"strings"
	"unicode"
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

type node struct {
	value       string
	neighbors   []*node
	isSmallCave bool
	visits      int
}

func hasUppercaseCharacter(value string) bool {
	hasUpper := false
	for _, r := range value {
		if unicode.IsUpper(r) {
			hasUpper = true
			break
		}
	}
	return hasUpper
}

func createNode(value string) *node {
	return &node{
		value:       value,
		neighbors:   []*node{},
		visits:      0,
		isSmallCave: !hasUppercaseCharacter(value),
	}
}

func listContainsNode(nodes []*node, value string) int {
	for index := range nodes {
		if nodes[index].value == value {
			return index
		}
	}
	return -1
}

func linkNodes(a *node, b *node) {
	a.neighbors = append(a.neighbors, b)
	b.neighbors = append(b.neighbors, a)
}

var paths = []string{}

func dfsRecursivePartA(nodes []*node, vertex *node, path []*node) {
	path = append(path, vertex)
	if vertex.value == "end" {
		readablePath := ""

		for _, step := range path {
			readablePath += step.value + "-"
		}
		readablePath = readablePath[:len(readablePath)-2]
		paths = append(paths, readablePath)
		return
	}

	for _, neighbor := range vertex.neighbors {
		pathContainsNeighbor := slices.Contains(path, neighbor)
		if !neighbor.isSmallCave || (neighbor.isSmallCave && !pathContainsNeighbor) || neighbor.value == "end" {
			dfsRecursivePartA(nodes[:], neighbor, path)
		}
	}

}

func Day12PartA2021(useExample bool) int {
	lines := getInput(useExample)

	nodes := buildNodes(lines)

	startNodeIndex := listContainsNode(nodes, "start")
	startNode := nodes[startNodeIndex]

	dfsRecursivePartA(nodes[:], startNode, []*node{})

	for index, readablePath := range paths {
		log.Printf("[DEBUG] Path %2d: %s", index, readablePath)
	}

	numberOfPaths := len(paths)
	paths = []string{}
	return numberOfPaths
}

func buildNodes(lines []string) []*node {
	nodes := []*node{}

	for _, line := range lines {
		parts := strings.Split(line, "-")

		var nodeA *node
		nodeAIndex := listContainsNode(nodes, parts[0])
		if nodeAIndex >= 0 {
			nodeA = nodes[nodeAIndex]
		} else {
			nodeA = createNode(parts[0])
			nodes = append(nodes, nodeA)
		}

		var nodeB *node
		nodeBIndex := listContainsNode(nodes, parts[1])
		if nodeBIndex >= 0 {
			nodeB = nodes[nodeBIndex]
		} else {
			nodeB = createNode(parts[1])
			nodes = append(nodes, nodeB)
		}

		linkNodes(nodeA, nodeB)
	}
	return nodes
}

func dfsRecursivePartB(nodes []*node, vertex *node, path []*node, smallCaveVisitedTwice bool) {
	path = append(path, vertex)
	if vertex.value == "end" {
		readablePath := ""

		for _, step := range path {
			readablePath += step.value + "-"
		}
		readablePath = readablePath[:len(readablePath)-2]
		paths = append(paths, readablePath)
		return
	}

	for _, neighbor := range vertex.neighbors {
		if neighbor.value != "start" {
			pathContainsNeighbor := slices.Contains(path, neighbor)
			if !neighbor.isSmallCave || neighbor.value == "end" {
				dfsRecursivePartB(nodes[:], neighbor, path, smallCaveVisitedTwice)
			} else if neighbor.isSmallCave && !pathContainsNeighbor {
				dfsRecursivePartB(nodes[:], neighbor, path, smallCaveVisitedTwice)
			} else if !smallCaveVisitedTwice && neighbor.isSmallCave && pathContainsNeighbor {
				dfsRecursivePartB(nodes[:], neighbor, path, true)
			}
		}
	}

}

func Day12PartB2021(useExample bool) int {
	lines := getInput(useExample)

	nodes := buildNodes(lines)

	startNodeIndex := listContainsNode(nodes, "start")
	startNode := nodes[startNodeIndex]

	dfsRecursivePartB(nodes[:], startNode, []*node{}, false)

	for index, readablePath := range paths {
		log.Printf("[DEBUG] Path %2d: %s", index, readablePath)
	}

	numberOfPaths := len(paths)
	paths = []string{}
	return numberOfPaths
}
