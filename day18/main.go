package main

import (
	"adventofcode/utils"
	"fmt"
	"math"
	"strings"
)

type digger struct {
	perimiter    int64
	points       []Point
	currPosition Point
}

func InitDigger() *digger {
	var result *digger = &digger{}
	result.perimiter = 0
	result.currPosition = Point{X: 0, Y: 0}
	result.points = []Point{{X: 0, Y: 0}}
	return result
}

func (t *digger) dig(direction Direction, length int64) {
	t.perimiter += length
	t.currPosition = direction.NextPoint(t.currPosition, length)
	t.points = append(t.points, t.currPosition)
}

func (t *digger) calcSurfaceArea() float64 {
	var shoelaceResult float64 = t.shoelace()
	// Picks theorum
	return shoelaceResult + math.Abs(float64(t.perimiter))/2 + 1
}

func (t *digger) shoelace() float64 {
	var area int64 = 0
	for i := 1; i < len(t.points); i++ {
		firstPoint := t.points[i-1]
		secondPoint := t.points[i]
		area += (firstPoint.X - secondPoint.X) * (firstPoint.Y + secondPoint.Y)
	}
	return math.Abs(float64(area)) / 2
}

func main() {
	input := utils.InitInputFile("digData.txt")
	defer input.Close()

	// result := partOne(input)
	// fmt.Println("Part 1 result: ", int64(result))

	result := partTwo(input)
	fmt.Println("Part 2 result: ", int64(result))
}

func partOne(input utils.InputFile) float64 {
	var digger digger = *InitDigger()
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		var splits []string = strings.Split(line, " ")

		var direction Direction
		switch splits[0] {
		case "U":
			direction = North
		case "R":
			direction = East
		case "D":
			direction = South
		case "L":
			direction = West
		}
		var metersToDig int64 = utils.StringToInt64(splits[1])

		digger.dig(direction, metersToDig)
	}
	return digger.calcSurfaceArea()
}

func partTwo(input utils.InputFile) float64 {
	var digger digger = *InitDigger()
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		var splits []string = strings.Split(line, " ")
		var metersToDigHex string = splits[2][2:7]
		var metersToDig int64 = utils.HexToInt64(metersToDigHex)
		var direction Direction
		switch splits[2][7:8] {
		case "3":
			direction = North
		case "0":
			direction = East
		case "1":
			direction = South
		case "2":
			direction = West
		}
		// var metersToDig int = utils.StringToInt(splits[1])
		// var color string = splits[2]

		digger.dig(direction, metersToDig)
	}
	return digger.calcSurfaceArea()
}
