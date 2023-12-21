package main

import (
	"adventofcode/utils"
	"fmt"
)

func partOne() {
	input := utils.InitInputFile("mapsData.txt")
	defer input.Close()
	directions := getDirections(input)
	mappings := buildMappings(input)
	result := getStepsToZ(directions, mappings)
	fmt.Println("result: ", result)
}

func getStepsToZ(directions []rune, mappings map[string][]string) int {
	var noOfSteps int = 0
	var directionIndex int = 0
	var currentPoint string = "AAA"
	for currentPoint != "ZZZ" {
		if directionIndex >= len(directions) {
			directionIndex = 0
		}
		var direction rune = directions[directionIndex]
		if direction == 'L' {
			currentPoint = mappings[currentPoint][0]
		} else {
			currentPoint = mappings[currentPoint][1]
		}
		noOfSteps++
		directionIndex++
	}
	return noOfSteps
}

func buildMappings(input utils.InputFile) map[string][]string {
	var mappings map[string][]string = map[string][]string{}
	for input.MoveToNextLine() {
		point, navigations := buildMappingPoints(input)
		mappings[point] = navigations
	}
	return mappings
}
