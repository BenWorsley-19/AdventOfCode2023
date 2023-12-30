package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	input := utils.InitInputFile("hikeMapData.txt")
	defer input.Close()
	var grid [][]rune = input.ToRuneGrid()
	start, end := findStartAndEnd(grid)
	var compressedPoints []point = compressPoints(start, end, grid)

	var partOneResult int = longestPathFollowingSlopes(start, end, compressedPoints, grid)
	fmt.Println("Part one result: ", partOneResult)
	var partTwoResult int = longestPathIgnoringSlopes(start, end, compressedPoints, grid)
	fmt.Println("Part Two result: ", partTwoResult)
}

func longestPathFollowingSlopes(start, end point, compressedPoints []point, grid [][]rune) int {
	var tileDirs map[rune][]direction = map[rune][]direction{
		'^': {North},
		'v': {South},
		'<': {West},
		'>': {East},
		'.': directions,
	}
	var hikeGraph graph = initGraph(start, end, compressedPoints, grid, tileDirs)
	return hikeGraph.dfs()
}

func longestPathIgnoringSlopes(start, end point, compressedPoints []point, grid [][]rune) int {
	var tileDirs map[rune][]direction = map[rune][]direction{
		'^': directions,
		'v': directions,
		'<': directions,
		'>': directions,
		'.': directions,
	}
	var hikeGraph graph = initGraph(start, end, compressedPoints, grid, tileDirs)
	return hikeGraph.dfs()
}

func findStartAndEnd(grid [][]rune) (point, point) {
	var start point
	for col := 0; col < len(grid[0]); col++ {
		if grid[0][col] == '.' {
			start = point{x: 0, y: col}
		}
	}
	var end point
	for col := 0; col < len(grid[0]); col++ {
		if grid[len(grid)-1][col] == '.' {
			end = point{x: len(grid) - 1, y: col}
		}
	}
	return start, end
}

func compressPoints(start point, end point, grid [][]rune) []point {
	var compressedPoints []point = []point{start, end}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == '#' {
				continue
			}
			var currPoint point = point{x: row, y: col}
			var noOfNeighbours int = 0
			for _, dir := range directions {
				var np point = dir.nextPoint(currPoint)
				if !pointIsOffGrid(np, grid) && grid[np.x][np.y] != '#' {
					noOfNeighbours++
				}
			}
			if noOfNeighbours >= 3 {
				compressedPoints = append(compressedPoints, currPoint)
			}
		}
	}
	return compressedPoints
}

func pointIsOffGrid(p point, grid [][]rune) bool {
	return p.x < 0 || p.x >= len(grid) ||
		p.y < 0 || p.y >= len(grid[0])
}
