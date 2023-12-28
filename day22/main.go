package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

type sandBrick struct {
	from               position
	to                 position
	providesSupportFor []*sandBrick
	sitsOn             []*sandBrick
}

func (b *sandBrick) addToSupports(brick *sandBrick) {
	b.providesSupportFor = append(b.providesSupportFor, brick)
}

func (b *sandBrick) addToSitsOn(brick *sandBrick) {
	b.sitsOn = append(b.sitsOn, brick)
}

type position struct {
	x int
	y int
	z int
}

func main() {
	input := utils.InitInputFile("sandData.txt")
	defer input.Close()
	var bricks []*sandBrick = parseInput(input)
	quickSort(bricks, 0, len(bricks)-1)
	landBricks(bricks)
	determineSupportingBricks(bricks)
	var partOneResult int = countDisintegratableSandBlocks(bricks)
	fmt.Println("Part one result: ", partOneResult)
	var partTwoResult int = calcTotalOfChainReactions(bricks)
	fmt.Println("Part two result: ", partTwoResult)
}

func landBricks(bricks []*sandBrick) {
	var noOfFalls int = 0
	for i := 0; i < len(bricks); i++ {
		var highestPointCanLand int = 1
		var brickA *sandBrick = bricks[i]
		for j := 0; j < i; j++ {
			var brickB *sandBrick = bricks[j]
			if intersects(brickA, brickB) {
				if brickB.to.z+1 > highestPointCanLand {
					highestPointCanLand = brickB.to.z + 1
				}
			}
		}
		var priorTo int = brickA.to.z
		var priorFrom int = brickA.from.z
		brickA.to.z -= brickA.from.z - highestPointCanLand
		brickA.from.z = highestPointCanLand
		if priorTo != brickA.to.z || priorFrom != brickA.from.z {
			noOfFalls++
		}
	}
	quickSort(bricks, 0, len(bricks)-1)
}

func intersects(brickA *sandBrick, brickB *sandBrick) bool {
	return math.Max(float64(brickA.from.x), float64(brickB.from.x)) <= math.Min(float64(brickA.to.x), float64(brickB.to.x)) &&
		math.Max(float64(brickA.from.y), float64(brickB.from.y)) <= math.Min(float64(brickA.to.y), float64(brickB.to.y))
}

func determineSupportingBricks(bricks []*sandBrick) {
	for i := 0; i < len(bricks); i++ {
		var brickA *sandBrick = bricks[i]
		for j := 0; j < i; j++ {
			var brickB *sandBrick = bricks[j]
			if intersects(brickA, brickB) && brickA.from.z == brickB.to.z+1 {
				brickA.addToSitsOn(brickB)
				brickB.addToSupports(brickA)
			}
		}
	}
}

func countDisintegratableSandBlocks(bricks []*sandBrick) int {
	var result int = 0
	for _, brick := range bricks {
		var hasMoreThanOneSupporting = true
		for _, supportedBrick := range brick.providesSupportFor {
			if !(len(supportedBrick.sitsOn) > 1) {
				hasMoreThanOneSupporting = false
			}
		}
		if hasMoreThanOneSupporting {
			result++
		}
	}
	return result
}

func calcTotalOfChainReactions(bricks []*sandBrick) int {
	var result int = 0
	for _, brick := range bricks {
		result += len(calcChainReaction(brick, []*sandBrick{brick})) - 1
	}
	return result
}

func calcChainReaction(brick *sandBrick, counted []*sandBrick) []*sandBrick {
	if len(brick.providesSupportFor) == 0 {
		return counted
	}
	var currDecintegratedBricks []*sandBrick = []*sandBrick{}
	for _, supportedBrick := range brick.providesSupportFor {
		if brickHasAlreadyBeenCounted(counted, supportedBrick) {
			continue
		}
		var noOfSatOnBricksThatHaveDecintegrated int = countNoOfSatOnBricksDecintegrated(counted, supportedBrick)
		if len(supportedBrick.sitsOn)-noOfSatOnBricksThatHaveDecintegrated < 1 {
			currDecintegratedBricks = append(currDecintegratedBricks, supportedBrick)
			counted = append(counted, supportedBrick)
		}
	}
	for _, decintegratedBrick := range currDecintegratedBricks {
		counted = calcChainReaction(decintegratedBrick, counted)
	}
	return counted
}

func brickHasAlreadyBeenCounted(seen []*sandBrick, brick *sandBrick) bool {
	var alreadyCounted bool = false
	for _, countedBrick := range seen {
		if countedBrick == brick {
			alreadyCounted = true
			break
		}
	}
	return alreadyCounted
}

func countNoOfSatOnBricksDecintegrated(seen []*sandBrick, brick *sandBrick) int {
	var noOfSatOnBricksThatHaveDecintegrated int = 0
	for _, satOnBrick := range brick.sitsOn {
		for _, decintegratedBrick := range seen {
			if satOnBrick == decintegratedBrick {
				noOfSatOnBricksThatHaveDecintegrated++
			}
		}
	}
	return noOfSatOnBricksThatHaveDecintegrated
}

func printBricks(arr []*sandBrick) {
	for i := 0; i < len(arr); i++ {
		fmt.Print(*arr[i])
		fmt.Print(",")
	}
	fmt.Println()
}

func quickSort(arr []*sandBrick, lo int, hi int) {
	if lo >= hi {
		return
	}
	var pivotIndex int = partition(arr, lo, hi)
	quickSort(arr, lo, pivotIndex-1)
	quickSort(arr, pivotIndex+1, hi)
}

func partition(arr []*sandBrick, lo int, hi int) int {
	var pivot *sandBrick = arr[hi]
	var index int = lo - 1
	for i := lo; i < hi; i++ {
		var lowestZ = math.Min(float64(arr[i].from.z), float64(arr[i].to.z))
		var pivotLowestZ = math.Min(float64(pivot.from.z), float64(pivot.to.z))
		if lowestZ > pivotLowestZ {
			continue
		}

		index++
		var tmp *sandBrick = arr[i]
		arr[i] = arr[index]
		arr[index] = tmp
	}
	index++
	arr[hi] = arr[index]
	arr[index] = pivot
	return index
}

func parseInput(input utils.InputFile) []*sandBrick {
	var bricks []*sandBrick = []*sandBrick{}
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		var positionSplits []string = strings.Split(line, "~")
		var endA position = parseCoordinate(positionSplits[0])
		var endb position = parseCoordinate(positionSplits[1])
		var brick sandBrick = sandBrick{from: endA, to: endb}
		bricks = append(bricks, &brick)
	}
	return bricks
}

func parseCoordinate(unparsed string) position {
	var coords []string = strings.Split(unparsed, ",")
	return position{
		x: utils.StringToInt(coords[0]),
		y: utils.StringToInt(coords[1]),
		z: utils.StringToInt(coords[2]),
	}
}
