package AoCScaffold

import (
	_ "embed"
	"log"
	"slices"
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

type Planet struct {
	identifier string
	orbits     []*Planet
}

func createPlanet(identifier string) *Planet {
	return &Planet{identifier: identifier}
}

func parseInput(lines []string) map[string]*Planet {
	planets := make(map[string]*Planet)

	for _, line := range lines {
		planetIdentifiers := strings.Split(line, ")")

		mainPlanet, mainPlanetExists := planets[planetIdentifiers[0]]
		orbittingPlanet, orbittingPlanetExists := planets[planetIdentifiers[1]]

		if !mainPlanetExists {
			mainPlanet = createPlanet(planetIdentifiers[0])
			planets[planetIdentifiers[0]] = mainPlanet
		}

		if !orbittingPlanetExists {
			orbittingPlanet = createPlanet(planetIdentifiers[1])
			planets[planetIdentifiers[1]] = orbittingPlanet
		}

		mainPlanet.orbits = append(mainPlanet.orbits, orbittingPlanet)
	}
	return planets
}

func DFS(node *Planet, depth int) int {
	orbits := 0
	for _, planet := range node.orbits {
		indirectOrbits := depth + DFS(planet, depth+1) + 1
		orbits += indirectOrbits
	}
	return orbits
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	planets := parseInput(lines)

	root := planets["COM"]
	orbits := DFS(root, 0)

	return orbits
}

func Search(node *Planet, goal *Planet, depth int) int {
	log.Printf("[DEBUG] Searching planet %v for node %v", node.identifier, goal.identifier)
	if slices.Contains(node.orbits, goal) {
		log.Printf("[DEBUG] Found %v at a depth of %v", goal.identifier, depth)
		return depth
	} else {
		for _, planet := range node.orbits {
			distance := Search(planet, goal, depth+1)
			if distance > 0 {
				return distance
			}
		}
	}
	return -1
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	planets := parseInput(lines)

	root := planets["YOU"]
	goal := planets["SAN"]

	minTransfer := 10000000
	for _, planet := range planets {
		log.Printf("[DEBUG] Began search for path from root node %v", planet.identifier)
		distanceToYou := Search(planet, root, 0)
		distanceToSan := Search(planet, goal, 0)
		log.Printf("[DEBUG] Distance from root to YOU: %v", distanceToYou)
		log.Printf("[DEBUG] Distance from root to SAN: %v", distanceToSan)
		if distanceToSan > 0 && distanceToYou > 0 {
			transferDistance := distanceToSan + distanceToYou
			if transferDistance < minTransfer {
				minTransfer = transferDistance
			}
		}
	}

	return minTransfer
}
