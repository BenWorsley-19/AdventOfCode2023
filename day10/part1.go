package main

// import (
// 	"adventofcode/utils"
// 	"fmt"
// )

// type point struct {
// 	x int
// 	y int
// }

// func main() {
// 	input := utils.InitInputFile("pipesData.txt")
// 	grid, startPoint := buildGridAndStartPoint(input)
// 	input.Close()

// 	_, noOfSteps := searchForPipeLoop(grid, startPoint, North, 0)
// 	fmt.Println("steps: ", noOfSteps)

// 	result := noOfSteps/2 + noOfSteps%2
// 	fmt.Println("result: ", result)
// }

// func searchForPipeLoop(grid [][]tile, currentPoint point, currentDirection Direction, steps int) (bool, int) {
// 	if pointIsOffGrid(currentPoint, grid) {
// 		return false, steps
// 	}
// 	var currentTile tile = grid[currentPoint.x][currentPoint.y]
// 	fmt.Printf("curr: %T\n", currentTile)
// 	if !connectsInDirection(currentTile, currentDirection) {
// 		return false, steps
// 	}
// 	_, isStart := currentTile.(startPipe)
// 	if steps != 0 && isStart {
// 		return true, steps
// 	}
// 	// steps = append(steps, currentPoint)
// 	steps++
// 	for _, direction := range currentTile.getConnectionPoints() {
// 		// don't go backwards
// 		if direction == currentDirection.opposite() {
// 			continue
// 		}

// 		var nextPoint point = direction.nextPoint(currentPoint)

// 		fmt.Println("steps ", steps)
// 		found, steps := searchForPipeLoop(grid, nextPoint, direction, steps)
// 		fmt.Println("f ", found)
// 		if found {
// 			return true, steps
// 		}
// 		fmt.Println("steps ", steps)
// 	}
// 	steps--
// 	return false, steps
// }

// func pointIsOffGrid(p point, grid [][]tile) bool {
// 	return p.x < 0 || p.x >= len(grid[0]) ||
// 		p.y < 0 || p.y > len(grid)

// }

// func buildGridAndStartPoint(input utils.InputFile) ([][]tile, point) {
// 	var startPoint point
// 	var grid [][]tile = [][]tile{}
// 	var rowCount int = 0
// 	for input.MoveToNextLine() {
// 		var line string = input.ReadLine()
// 		var runes []rune = []rune(line)
// 		var row []tile = []tile{}
// 		for i := 0; i < len(runes); i++ {
// 			var toAppend tile
// 			switch runes[i] {
// 			case '|':
// 				toAppend = verticalPipe{}
// 			case '-':
// 				toAppend = horizontalPipe{}
// 			case 'L':
// 				toAppend = lBend{}
// 			case 'J':
// 				toAppend = jBend{}
// 			case '7':
// 				toAppend = sevenBend{}
// 			case 'F':
// 				toAppend = fBend{}
// 			case '.':
// 				toAppend = ground{}
// 			case 'S':
// 				toAppend = startPipe{}
// 				startPoint = point{x: rowCount, y: i}
// 			}
// 			row = append(row, toAppend)
// 		}
// 		rowCount++
// 		grid = append(grid, row)
// 	}
// 	return grid, startPoint
// }
