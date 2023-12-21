package main

import (
	"adventofcode/utils"
	"container/heap"
	"fmt"
)

type point struct {
	x int
	y int
}

func main() {
	input := utils.InitInputFile("cityBlockData.txt")
	var grid [][]rune = input.ToRuneGrid()
	input.Close()

	var partOneResult = djikstra(grid, 0, 0, 3)
	fmt.Println("part 1: ", partOneResult)

	var partTwoResult = djikstra(grid, 4, 4, 10)
	fmt.Println("part 2: ", partTwoResult)
}

func djikstra(grid [][]rune, minStepsInDirBeforeFin int, minStepsInDirBeforeTurn int, maxStepsInOneDir int) int {
	pq := make(PriorityQueue, 0)
	heap.Init(&pq)

	var start cityBlock = cityBlock{
		heatLoss:       0,
		itemPoint:      point{x: 0, y: 0},
		dir:            Stationary,
		noOfTimesInDir: 0,
		index:          0,
	}
	heap.Push(&pq, &start)
	var seen []cityBlock = []cityBlock{}
	for len(pq) > 0 {
		var curr *cityBlock = heap.Pop(&pq).(*cityBlock)
		if curr.itemPoint.x == len(grid)-1 && curr.itemPoint.y == len(grid[0])-1 && curr.noOfTimesInDir >= minStepsInDirBeforeFin {
			return curr.heatLoss
		}
		if seenBefore(seen, *curr) {
			continue
		}
		seen = append(seen, *curr)

		// check same dir
		if curr.noOfTimesInDir < maxStepsInOneDir && curr.dir != Stationary {
			var nextPoint point = curr.dir.nextPoint(curr.itemPoint)
			if !pointIsOffGrid(nextPoint, grid) {
				var heatLoss = utils.RuneToInt(grid[nextPoint.x][nextPoint.y])
				var next cityBlock = cityBlock{
					heatLoss:       curr.heatLoss + heatLoss,
					itemPoint:      nextPoint,
					dir:            curr.dir,
					noOfTimesInDir: curr.noOfTimesInDir + 1,
				}
				heap.Push(&pq, &next)
			}
		}

		//check turns
		if curr.noOfTimesInDir >= minStepsInDirBeforeTurn || curr.dir == Stationary {
			for _, d := range directions {
				// can't go opposite direction
				if d == curr.dir.opposite() {
					continue
				}
				// already handled current direction above
				if d == curr.dir {
					continue
				}

				var nextPoint point = d.nextPoint(curr.itemPoint)
				if !pointIsOffGrid(nextPoint, grid) {
					var heatLoss = utils.RuneToInt(grid[nextPoint.x][nextPoint.y])
					var next cityBlock = cityBlock{
						heatLoss:       curr.heatLoss + heatLoss,
						itemPoint:      nextPoint,
						dir:            d,
						noOfTimesInDir: 1,
					}
					heap.Push(&pq, &next)
				}
			}
		}
	}
	panic("We should have found the end")
}

// TODO optimise
func seenBefore(seen []cityBlock, curr cityBlock) bool {
	for i := 0; i < len(seen); i++ {
		if seen[i].itemPoint == curr.itemPoint &&
			seen[i].dir == curr.dir &&
			seen[i].noOfTimesInDir == curr.noOfTimesInDir {
			return true
		}
	}
	return false
}

func pointIsOffGrid(p point, grid [][]rune) bool {
	return p.x < 0 || p.x >= len(grid) ||
		p.y < 0 || p.y >= len(grid[0])
}
