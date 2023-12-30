package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

type hailstone struct {
	px float64
	py float64
	pz float64
	vx float64
	vy float64
	vz float64
	a  float64
	b  float64
	c  float64
}

func initHailstone(px float64, py float64, pz float64, vx float64, vy float64, vz float64) hailstone {
	var h hailstone = hailstone{}
	h.px = px
	h.py = py
	h.pz = pz

	h.vx = vx
	h.vy = vy
	h.vz = vz

	h.a = vy
	h.b = -vx
	h.c = vy*px - vx*py
	return h
}

func (hA hailstone) determineIntersection(hB hailstone) (float64, float64) {
	var x float64 = (hA.c*hB.b - hB.c*hA.b) / (hA.a*hB.b - hB.a*hA.b)
	var y float64 = (hB.c*hA.a - hA.c*hB.a) / (hA.a*hB.b - hB.a*hA.b)
	return x, y
}

func (hA hailstone) isParralelWith(hB hailstone) bool {
	return hA.a*hB.b == hA.b*hB.a
}

func parse(input utils.InputFile) []hailstone {
	var hailstones []hailstone = []hailstone{}
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		var splits []string = strings.Split(line, " @ ")
		var positions []string = strings.Split(splits[0], ",")
		var velocities []string = strings.Split(splits[1], ",")
		var px float64 = utils.StringToFloat64(strings.ReplaceAll(positions[0], " ", ""))
		var py float64 = utils.StringToFloat64(strings.ReplaceAll(positions[1], " ", ""))
		var pz float64 = utils.StringToFloat64(strings.ReplaceAll(positions[2], " ", ""))
		var vx float64 = utils.StringToFloat64(strings.ReplaceAll(velocities[0], " ", ""))
		var vy float64 = utils.StringToFloat64(strings.ReplaceAll(velocities[1], " ", ""))
		var vz float64 = utils.StringToFloat64(strings.ReplaceAll(velocities[2], " ", ""))
		hailstones = append(hailstones, initHailstone(px, py, pz, vx, vy, vz))
	}
	return hailstones
}

func printSageScriptForPartTwo(hailstones []hailstone) {
	// For part 2 I looked to reddit for clues where I found lots of suggestions for using sage math.
	// I found this script https://github.com/vipul0092/advent-of-code-2023/blob/main/day24/day24.go
	// and add the x y and z of the result from sage math to get the star on this one
	fmt.Println()
	fmt.Println("var('x y z vx vy vz t1 t2 t3')")
	fmt.Println("eq1 = x + (vx * t1) == ", hailstones[0].px, " + (", hailstones[0].vx, " * t1)")
	fmt.Println("eq2 = y + (vy * t1) == ", hailstones[0].py, " + (", hailstones[0].vy, " * t1)")
	fmt.Println("eq3 = z + (vz * t1) == ", hailstones[0].pz, " + (", hailstones[0].vz, " * t1)")
	fmt.Println("eq4 = x + (vx * t2) == ", hailstones[1].px, " + (", hailstones[1].vx, " * t2)")
	fmt.Println("eq5 = y + (vy * t2) == ", hailstones[1].py, " + (", hailstones[1].vy, " * t2)")
	fmt.Println("eq6 = z + (vz * t2) == ", hailstones[1].pz, " + (", hailstones[1].vz, " * t2)")
	fmt.Println("eq7 = x + (vx * t3) == ", hailstones[2].px, " + (", hailstones[2].vx, " * t3)")
	fmt.Println("eq8 = y + (vy * t3) == ", hailstones[2].py, " + (", hailstones[2].vy, " * t3)")
	fmt.Println("eq9 = z + (vz * t3) == ", hailstones[2].pz, " + (", hailstones[2].vz, " * t3)")
	fmt.Println("print(solve([eq1,eq2,eq3,eq4,eq5,eq6,eq7,eq8,eq9],x,y,z,vx,vy,vz,t1,t2,t3))")
	fmt.Println()
}

func sumOfXYZSageResult() int64 {
	return 454587375941126 + 244764814652484 + 249133632375809
}

func countIntersectionsWithinTestArea(hailstones []hailstone, testAreaX float64, testAreaY float64) int {
	var result int = 0
	for i, hailstoneA := range hailstones {
		for j := 0; j < i; j++ {
			var hailstoneB hailstone = hailstones[j]
			if hailstoneA.isParralelWith(hailstoneB) {
				continue
			}
			x, y := hailstoneA.determineIntersection(hailstoneB)
			if withinTestArea(x, y, testAreaX, testAreaY) && hailstonesAreInFuture(x, y, hailstoneA, hailstoneB) {
				result++
			}
		}
	}
	return result
}

func withinTestArea(x, y, testAreaX, testAreaY float64) bool {
	return testAreaX <= x && x <= testAreaY && testAreaX <= y && y <= testAreaY
}

func hailstonesAreInFuture(x, y float64, hA, hB hailstone) bool {
	return (x-hA.px)*hA.vx >= 0 && (y-hA.py)*hA.vy >= 0 &&
		(x-hB.px)*hB.vx >= 0 && (y-hB.py)*hB.vy >= 0
}

func main() {
	input := utils.InitInputFile("hailstoneData.txt")
	defer input.Close()

	var hailstones []hailstone = parse(input)

	var partOneResult int = countIntersectionsWithinTestArea(hailstones, 200000000000000, 400000000000000)
	fmt.Println("Part one result: ", partOneResult)

	printSageScriptForPartTwo(hailstones)
	var partTwoResult int64 = sumOfXYZSageResult()
	fmt.Println("Part two result: ", partTwoResult)
}
