package main

import (
	"fmt"
	"strings"
)

const TOTAL_BRIDGES = 7

var BRIDGES = []string{"AaB", "AbB", "AcC", "AdC", "AeD", "BfD", "CgD"}
var PATHS []string

func main() {
	allAreas := "ABCD"
	var numberOfSteps, numberOfSolutions int

	for _, area := range allAreas {
		startWalking(string(area))
	}

	for _, path := range PATHS {
		fmt.Println(path)
	}

	for _, currentWalk := range PATHS {
		if listContains([]string{"a", "b", "c", "d", "e", "f", "g"}, currentWalk) {
			numberOfSolutions++
		}
		for _, letter := range currentWalk {
			if letter >= 'a' && letter <= 'z' {
				numberOfSteps++
			}
		}
	}

	fmt.Println("Number of bridges crossed: ", numberOfSteps)
	fmt.Println("Number of solutions: ", numberOfSolutions)
}

func startWalking(area string) []string {
	createPath(area, "", []string{})
	return PATHS
}

func createPath(area, lastpath string, bridgesCrossed []string) {
	if lastpath == "" {
		lastpath = "NONE"
	}

	if lastpath == "NONE" {
		lastpath = area
	}

	var availableBridges []string

	for i := 0; i < TOTAL_BRIDGES; i++ {
		if strings.Contains(BRIDGES[i], area) && !contains(bridgesCrossed, string(BRIDGES[i][1])) {
			availableBridges = append(availableBridges, BRIDGES[i])
		}
	}

	if len(availableBridges) == 0 {
		PATHS = append(PATHS, lastpath)
		lastpath = "NONE"
		return
	}

	for _, bridge := range availableBridges {
		currentlyCrossing := reverse(bridge[0:2])
		if bridge[0] == area[0] {
			currentlyCrossing = bridge[1:]
		}
		bridgesCrossed = append(bridgesCrossed, string(bridge[1]))
		currentBack := currentlyCrossing[len(currentlyCrossing)-1]
		createPath(string(currentBack), lastpath+currentlyCrossing, bridgesCrossed)
	}
}

func listContains(list []string, item string) bool {
	var counter int
	for i := range list {
		if strings.Contains(item, list[i]) {
			counter++
		}
	}
	return counter == len(list)
}

func contains(list []string, item string) bool {
	for i := range list {
		if strings.Contains(list[i], item) {
			return true
		}
	}
	return false
}

func reverse(s string) string {
	rns := []rune(s)
	for i, j := 0, len(rns)-1; i < j; i, j = i+1, j-1 {
		rns[i], rns[j] = rns[j], rns[i]
	}
	return string(rns)
}
