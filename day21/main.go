package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
)

func main() {
	input := utils.InitInputFile("gardensData.txt")
	defer input.Close()
	var grid [][]rune = input.ToRuneGrid()
	var start point = findStart(grid)
	var partOneResult = findReachablePoints(grid, start, 64)
	fmt.Println("Part one result: ", partOneResult)
	var partTwoResult = findReachablePointsOnInfiniteGrid(grid, start, 26501365)
	fmt.Println("Part two result: ", partTwoResult)
}

func findReachablePoints(grid [][]rune, start point, noOfSteps int64) int64 {
	var reachablePoints []point = []point{start}
	for i := 0; int64(i) < noOfSteps; i++ {
		var nextReachablePoints []point = []point{}
		for _, currReachablePoint := range reachablePoints {
			for _, dir := range directions {
				var next point = dir.nextPoint(currReachablePoint, 1)
				if utils.ArrayContains(nextReachablePoints, next) {
					continue
				}
				if !pointIsOffGrid(next, grid) && grid[next.x][next.y] != '#' {
					nextReachablePoints = append(nextReachablePoints, next)
				}
			}
		}
		reachablePoints = nextReachablePoints
	}
	return int64(len(reachablePoints))
}

func findReachablePointsOnInfiniteGrid(grid [][]rune, start point, noOfSteps int64) int64 {
	var gridHeight int64 = int64(len(grid))
	var gridWidth int64 = noOfSteps/gridHeight - 1
	var odd int64 = int64(math.Pow(float64(gridWidth/2*2+1), 2))
	var even int64 = int64(math.Pow((float64(gridWidth+1))/2*2, 2))
	var oddPoints int64 = findReachablePoints(grid, start, gridHeight*2+1)
	var evenPoints int64 = findReachablePoints(grid, start, gridHeight*2)
	var topCorner int64 = findReachablePoints(grid, point{x: gridHeight - 1, y: start.y}, gridHeight-1)
	var rightCorner int64 = findReachablePoints(grid, point{x: start.x, y: 0}, gridHeight-1)
	var bottomCorner int64 = findReachablePoints(grid, point{x: 0, y: start.y}, gridHeight-1)
	var leftCorner int64 = findReachablePoints(grid, point{x: start.x, y: gridHeight - 1}, gridHeight-1)
	var smallTopRight int64 = findReachablePoints(grid, point{x: gridHeight - 1, y: 0}, gridHeight/2-1)
	var smallTopLeft int64 = findReachablePoints(grid, point{x: gridHeight - 1, y: gridHeight - 1}, gridHeight/2-1)
	var smallBottomRight int64 = findReachablePoints(grid, point{x: 0, y: 0}, gridHeight/2-1)
	var smallBottomLeft int64 = findReachablePoints(grid, point{x: 0, y: gridHeight - 1}, gridHeight/2-1)
	var largeTopRight int64 = findReachablePoints(grid, point{x: gridHeight - 1, y: 0}, gridHeight*3/2-1)
	var largeTopLeft int64 = findReachablePoints(grid, point{x: gridHeight - 1, y: gridHeight - 1}, gridHeight*3/2-1)
	var largeBottomRight int64 = findReachablePoints(grid, point{x: 0, y: 0}, gridHeight*3/2-1)
	var largeBottomLeft int64 = findReachablePoints(grid, point{x: 0, y: gridHeight - 1}, gridHeight*3/2-1)
	return odd*oddPoints + even*evenPoints + topCorner + rightCorner + bottomCorner + leftCorner +
		(gridWidth+1)*(smallTopRight+smallTopLeft+smallBottomRight+smallBottomLeft) +
		gridWidth*(largeTopRight+largeTopLeft+largeBottomRight+largeBottomLeft)
}

func findStart(grid [][]rune) point {
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 'S' {
				return point{x: int64(row), y: int64(col)}
			}
		}
	}
	panic("A start is required")
}

func pointIsOffGrid(p point, grid [][]rune) bool {
	return p.x < 0 || p.x >= int64(len(grid)) ||
		p.y < 0 || p.y >= int64(len(grid[0]))
}
