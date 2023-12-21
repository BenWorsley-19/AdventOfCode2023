package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
)

type point struct {
	x int64
	y int64
}

func main() {
	input := utils.InitInputFile("galaxyData.txt")
	var universe [][]rune = input.ToRuneGrid()
	input.Close()

	identifyEmptySpace(universe)
	// printGrid(universe) //  TODO move to utils
	var galaxyCoordinates []point = getGalaxyCoordinates(universe, 2)
	var result = addUpShortestDistances(galaxyCoordinates)
	fmt.Println("Part 1 result: ", int64(result))

	var p2GalaxyCoordinates []point = getGalaxyCoordinates(universe, 1000000)
	var p2Result = addUpShortestDistances(p2GalaxyCoordinates)
	fmt.Println("Part 2 result: ", int64(p2Result))
}

// func printGrid(arr [][]rune) {
// 	for row := 0; row < len(arr); row++ {
// 		for column := 0; column < len(arr[0]); column++ {
// 			fmt.Print(string(arr[row][column]), " ")
// 		}
// 		fmt.Print("\n")
// 	}
// }

func identifyEmptySpace(universe [][]rune) {
	for i := 0; i < len(universe); i++ {
		var row []rune = universe[i]
		if !rowContains(row, '#') {
			for j := 0; j < len(row); j++ {
				row[j] = 'E'
			}
		}

		if !colContains(universe, i, '#') {
			for j := 0; j < len(universe[i]); j++ {
				universe[j][i] = 'E'
			}
		}
	}
}

func getGalaxyCoordinates(universe [][]rune, expansionFactor int64) []point {
	var coordinates []point = []point{}
	var x int64 = 0
	for i := 0; i < len(universe); i++ {
		var row []rune = universe[i]
		if !rowContains(row, '.') && !rowContains(row, '#') {
			x += expansionFactor
			continue
		}
		var y int64 = 0
		for j := 0; j < len(universe[i]); j++ {
			var currentPoint rune = universe[i][j]
			if currentPoint == 'E' {
				y += expansionFactor
				continue
			}
			if currentPoint == '#' {
				coordinates = append(coordinates, point{x: x, y: y})
			}
			y += 1
		}
		x += 1
	}
	return coordinates
}

func addUpShortestDistances(coordiantes []point) float64 {
	var total float64 = 0
	var count int = 0
	var calculatedPoints []point = []point{}
	for i := 0; i < len(coordiantes); i++ {
		var galaxyCoord point = coordiantes[i]
		for j := 0; j < len(coordiantes); j++ {
			var otherGalaxyCoord point = coordiantes[j]
			if galaxyCoord == otherGalaxyCoord || containsPoint(calculatedPoints, otherGalaxyCoord) {
				continue
			}

			count++
			total += manhattanDistance(galaxyCoord, otherGalaxyCoord)
		}
		calculatedPoints = append(calculatedPoints, galaxyCoord)
	}
	return total
}

func containsPoint(arr []point, p point) bool {
	for i := 0; i < len(arr); i++ {
		if arr[i] == p {
			return true
		}
	}
	return false
}

func manhattanDistance(pointOne point, pointTwo point) float64 {
	return math.Abs(float64(pointOne.x)-float64(pointTwo.x)) + math.Abs(float64(pointOne.y)-float64(pointTwo.y))
}

func colContains(grid [][]rune, x int, r rune) bool {
	for y := 0; y < len(grid[x]); y++ {
		if grid[y][x] == r {
			return true
		}
	}
	return false
}

func rowContains(arr []rune, r rune) bool {
	for _, i := range arr {
		if i == r {
			return true
		}
	}
	return false
}
