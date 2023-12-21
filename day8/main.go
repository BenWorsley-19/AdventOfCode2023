package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func main() {
	// partOne()
	input := utils.InitInputFile("mapsData.txt")
	defer input.Close()
	directions := getDirections(input)
	startingPoints, mappings := buildMappingsAndStartPoints(input)
	noOfStepsToZs := getStepsToZFromEachStartPoint(directions, startingPoints, mappings)
	result := utils.LeastCommonMultiple(noOfStepsToZs)
	fmt.Println("result: ", result)
}

func getDirections(input utils.InputFile) []rune {
	input.MoveToNextLine()
	var directionsLine string = input.ReadLine()
	directions := []rune(directionsLine)
	input.MoveToNextLine()
	return directions
}

func buildMappingsAndStartPoints(input utils.InputFile) ([]string, map[string][]string) {
	var startingPoints []string = []string{}
	var mappings map[string][]string = map[string][]string{}
	for input.MoveToNextLine() {
		point, navigations := buildMappingPoints(input)
		mappings[point] = navigations
		if strings.HasSuffix(point, "A") {
			startingPoints = append(startingPoints, point)
		}
	}
	return startingPoints, mappings
}

func getStepsToZFromEachStartPoint(directions []rune, startingPoints []string, mappings map[string][]string) []int64 {
	var noOfStepsToZs []int64 = []int64{}
	for i := 0; i < len(startingPoints); i++ {
		var noOfSteps int64 = 0
		var directionIndex int = 0
		var currentPoint string = startingPoints[i]
		for !strings.HasSuffix(currentPoint, "Z") {
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
		noOfStepsToZs = append(noOfStepsToZs, noOfSteps)
	}
	return noOfStepsToZs
}

func buildMappingPoints(input utils.InputFile) (string, []string) {
	var mappingLine string = input.ReadLine()
	var pointAndNavigations []string = strings.Split(mappingLine, " = ")
	var point string = pointAndNavigations[0]
	var navigations []string = strings.Split(pointAndNavigations[1], ", ")
	var leftNav string = strings.Replace(navigations[0], "(", "", -1)
	var rightNav string = strings.Replace(navigations[1], ")", "", -1)
	navigations[0] = leftNav
	navigations[1] = rightNav
	return point, navigations
}
