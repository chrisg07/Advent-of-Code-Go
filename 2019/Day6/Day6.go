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

func addIndirectOrbits(planets map[string]*Planet, mainPlanet *Planet, orbittingPlanet *Planet) map[string]*Planet {
	for _, planet := range planets {
		if slices.Contains(planet.orbits, mainPlanet) {
			log.Printf("[CONSOLE] Adding indirect orbits for planet %v for all planets that %v orbits", orbittingPlanet.identifier, mainPlanet.identifier)

			planet.orbits = append(planet.orbits, orbittingPlanet)
			addIndirectOrbits(planets, planet, orbittingPlanet)
		}
	}

	return planets
}

func printPlanet(planet Planet) {
	orbits := ""
	for _, orbittingPlanet := range planet.orbits {
		orbits += orbittingPlanet.identifier + ", "
	}
	log.Printf("[CONSOLE] Planet %v has orbits: %v", planet.identifier, orbits)
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
		log.Printf("[CONSOLE] Adding indirect orbits (%v) between planet %v and %v at depth %v", indirectOrbits, node.identifier, planet.identifier, depth)
		orbits += indirectOrbits
	}
	return orbits
}

func PartA(useExample bool) int {
	lines := getInput(useExample)
	planets := parseInput(lines)

	// for _, planet := range planets {
	// 	for _, orbittingPlanet := range planet.orbits {
	// 		planets = addIndirectOrbits(planets, planet, orbittingPlanet)
	// 	}
	// }

	// depth first search

	root, _ := planets["COM"]
	orbits := DFS(root, 0)

	for _, planet := range planets {
		printPlanet(*planet)
	}

	return orbits
}

func PartB(useExample bool) int {
	lines := getInput(useExample)
	input := parseInput(lines)

	return len(input)
}
