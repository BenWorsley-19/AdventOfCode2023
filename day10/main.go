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
	input := utils.InitInputFile("pipesData.txt")
	grid, startPoint := buildGridAndStartPoint(input)
	input.Close()

	_, loop := searchForPipeLoop(grid, startPoint, North, []point{})
	noOfSteps := len(loop)

	partOneResult := noOfSteps/2 + noOfSteps%2
	fmt.Println("part 1: ", partOneResult)

	// replaceSInGrid(grid, startPoint)
	// TODO
	grid[startPoint.x][startPoint.y] = lBend{}

	partTwoResult := countTilesInLoop(grid, loop)
	fmt.Println("part 2: ", partTwoResult)
}

func searchForPipeLoop(grid [][]tile, currentPoint point, currentDirection Direction, loop []point) (bool, []point) {
	if pointIsOffGrid(currentPoint, grid) {
		return false, loop
	}
	var currentTile tile = grid[currentPoint.x][currentPoint.y]
	if !connectsInDirection(currentTile, currentDirection) {
		return false, loop
	}
	_, isStart := currentTile.(startPipe)
	if len(loop) != 0 && isStart {
		return true, loop
	}
	loop = append(loop, currentPoint)
	for _, direction := range currentTile.getConnectionPoints() {
		// don't go backwards
		if direction == currentDirection.opposite() {
			continue
		}

		var nextPoint point = direction.nextPoint(currentPoint)
		found, steps := searchForPipeLoop(grid, nextPoint, direction, loop)
		if found {
			return true, steps
		}
	}
	deleteElement(loop, len(loop))
	return false, loop
}

func countTilesInLoop(grid [][]tile, loop []point) int {
	var itemsInside int = 0
	for i := 0; i < len(grid); i++ {
		var linesCrossed int = 0
		var previousTile tile
		for j := 0; j < len(grid[0]); j++ {
			currentPoint := point{x: i, y: j}
			currentTile := grid[i][j]
			if loopContains(loop, currentPoint) {
				_, currIsHorizontalPipe := currentTile.(horizontalPipe)
				if currIsHorizontalPipe {
					continue
				}
				_, currIsVerticalPipe := currentTile.(verticalPipe)
				if currIsVerticalPipe {
					linesCrossed++
				}
				_, prevIsFBend := previousTile.(fBend)
				_, currIsJBend := currentTile.(jBend)
				if prevIsFBend && currIsJBend {
					linesCrossed++
				}
				_, prevIsLBend := previousTile.(lBend)
				_, currIsSevenBend := currentTile.(sevenBend)
				if prevIsLBend && currIsSevenBend {
					linesCrossed++
				}
				previousTile = currentTile
			} else if linesCrossed%2 != 0 {
				itemsInside++
			}
		}
	}
	return itemsInside
}

func pointIsOffGrid(p point, grid [][]tile) bool {
	return p.x < 0 || p.x >= len(grid) ||
		p.y < 0 || p.y > len(grid[0])
}

func buildGridAndStartPoint(input utils.InputFile) ([][]tile, point) {
	var startPoint point
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
				toAppend = verticalPipe{}
			case '-':
				toAppend = horizontalPipe{}
			case 'L':
				toAppend = lBend{}
			case 'J':
				toAppend = jBend{}
			case '7':
				toAppend = sevenBend{}
			case 'F':
				toAppend = fBend{}
			case '.':
				toAppend = ground{}
			case 'S':
				toAppend = startPipe{}
				startPoint = point{x: rowCount, y: i}
			}
			row = append(row, toAppend)
		}
		rowCount++
		grid = append(grid, row)
	}
	return grid, startPoint
}

func deleteElement(slice []point, index int) []point {
	return append(slice[:index], slice[index+1:]...)
}

func loopContains(loop []point, currentPoint point) bool {
	for i := 0; i < len(loop); i++ {
		if loop[i] == currentPoint {
			return true
		}
	}
	return false
}
