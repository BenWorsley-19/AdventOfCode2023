package main

import (
	"adventofcode/utils"
	"fmt"
)

type point struct {
	x int
	y int
}

func main() {
	input := utils.InitInputFile("laserData.txt")
	var grid [][]tile = buildGrid(input)
	input.Close()

	var partOneResult = findNoOfEnergizedTilesFromTopLeftEastStartPoint(grid)
	printGrid(grid)
	fmt.Println("part 1: ", partOneResult)

	var partTwoResult int = findHighestEnergizedTilesFromAllStartPoints(grid)
	fmt.Println("part 2: ", partTwoResult)
}

func findNoOfEnergizedTilesFromTopLeftEastStartPoint(grid [][]tile) int {
	var startPoint = point{x: 0, y: 0}
	var energisedBeams []point = energiseBeams(grid, startPoint, East, []point{})
	return len(energisedBeams)
}

func findHighestEnergizedTilesFromAllStartPoints(grid [][]tile) int {
	var highestTilesEnergised int = 0
	for row := 0; row < len(grid); row++ {
		var noOfEnergisedTiles = countEneergisedBeams(grid, point{x: row, y: 0}, East)
		if noOfEnergisedTiles > highestTilesEnergised {
			highestTilesEnergised = noOfEnergisedTiles
		}
		noOfEnergisedTiles = countEneergisedBeams(grid, point{x: row, y: len(grid[0])}, West)
		if noOfEnergisedTiles > highestTilesEnergised {
			highestTilesEnergised = noOfEnergisedTiles
		}
	}
	for col := 0; col < len(grid[0]); col++ {
		var noOfEnergisedTiles = countEneergisedBeams(grid, point{x: 0, y: col}, South)
		if noOfEnergisedTiles > highestTilesEnergised {
			highestTilesEnergised = noOfEnergisedTiles
		}
		noOfEnergisedTiles = countEneergisedBeams(grid, point{x: len(grid), y: col}, North)
		if noOfEnergisedTiles > highestTilesEnergised {
			highestTilesEnergised = noOfEnergisedTiles
		}
	}
	return highestTilesEnergised
}

func countEneergisedBeams(grid [][]tile, startPoint point, direction Direction) int {
	resetGrid(grid)
	var energisedBeams []point = energiseBeams(grid, startPoint, direction, []point{})
	return len(energisedBeams)
}

func energiseBeams(grid [][]tile, currentPoint point, currentDirection Direction, energisedBeams []point) []point {
	if pointIsOffGrid(currentPoint, grid) {
		// we're off the grid so don't add to energised beams and return
		return energisedBeams
	}
	var currentTile tile = grid[currentPoint.x][currentPoint.y]
	if !currentTile.isEnergised() {
		currentTile.energise(currentDirection)
		energisedBeams = append(energisedBeams, currentPoint)
	} else if currentTile.hasBeamedInDirection(currentDirection) {
		return energisedBeams
	}

	var nextPoints map[Direction]point = currentTile.determineNextPoints(currentDirection, currentPoint)
	for d, p := range nextPoints {
		energisedBeams = energiseBeams(grid, p, d, energisedBeams)
	}
	return energisedBeams
}

func pointIsOffGrid(p point, grid [][]tile) bool {
	return p.x < 0 || p.x >= len(grid) ||
		p.y < 0 || p.y >= len(grid[0])
}

func buildGrid(input utils.InputFile) [][]tile {
	var grid [][]tile = [][]tile{}
	var rowCount int = 0
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		var runes []rune = []rune(line)
		var row []tile = []tile{}
		for i := 0; i < len(runes); i++ {
			var toAppend tile
			switch runes[i] {
			case '|':
				toAppend = &verticalSplinter{}
			case '-':
				toAppend = &horizontalSplinter{}
			case '.':
				toAppend = &empty{}
			case '/':
				toAppend = &forwardMirror{}
			case '\\':
				toAppend = &backwardMirror{}
			}
			row = append(row, toAppend)
		}
		rowCount++
		grid = append(grid, row)
	}
	return grid
}

func resetGrid(grid [][]tile) {
	for row := 0; row < len(grid); row++ {
		for column := 0; column < len(grid[0]); column++ {
			grid[row][column].reset()
		}
	}
}

func printGrid(arr [][]tile) {
	for row := 0; row < len(arr); row++ {
		for column := 0; column < len(arr[0]); column++ {
			if arr[row][column].isEnergised() {
				fmt.Print("#", " ")
			} else {
				fmt.Print(".", " ")

			}
		}
		fmt.Print("\n")
	}
}
