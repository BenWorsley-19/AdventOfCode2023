package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

func main() {
	input := utils.InitInputFile("oasisData.txt")
	defer input.Close()

	var result int64 = 0
	for input.MoveToNextLine() {
		var readingsLine string = input.ReadLine()
		var readingsAsString []string = strings.Split(readingsLine, " ")
		var readings []int64 = utils.StringArrToInt64Arr(readingsAsString)
		// part 1
		// var nextVal int64 = determineNextValue(readings)
		var nextVal int64 = determinePreviousValue(readings)
		result += nextVal
	}

	fmt.Println("result: ", result)
}

func determinePreviousValue(readings []int64) int64 {
	var noDifference bool = true
	for _, v := range readings {
		if v != 0 {
			noDifference = false
			break
		}
	}
	if noDifference {
		return 0
	}

	differences := make([]int64, len(readings)-1)
	var lastReadingIndex int = len(readings) - 1
	for i := lastReadingIndex; i >= 1; i-- {
		var reading int64 = readings[i]
		var previousReading int64 = readings[i-1]
		difference := reading - previousReading
		differences[i-1] = difference
	}
	var prevVal int64 = determinePreviousValue(differences)
	return readings[0] - prevVal
}

func determineNextValue(readings []int64) int64 {
	var noDifference bool = true
	for _, v := range readings {
		if v != 0 {
			noDifference = false
			break
		}
	}
	if noDifference {
		return 0
	}

	differences := make([]int64, len(readings)-1)
	var lastReadingIndex int = len(readings) - 1
	for i := lastReadingIndex; i >= 1; i-- {
		var reading int64 = readings[i]
		var previousReading int64 = readings[i-1]
		difference := reading - previousReading
		differences[i-1] = difference
	}
	var nextVal int64 = determineNextValue(differences)
	return nextVal + readings[lastReadingIndex]
}
