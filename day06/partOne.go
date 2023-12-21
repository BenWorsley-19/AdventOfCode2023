package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func partOne() {
	var filePath = "raceData.txt"
	readFile, err := os.Open(filePath)
	defer readFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	raceLengthsLine := fileScanner.Text()
	fileScanner.Scan()
	bestDistancesLine := fileScanner.Text()

	var raceTimeLengths []int = toResultArray(raceLengthsLine)
	var bestDistances []int = toResultArray(bestDistancesLine)

	var results [][]int = [][]int{}
	for i := 0; i < len(raceTimeLengths); i++ {
		var possibleWaysToWin []int = []int{}
		raceLength := raceTimeLengths[i]
		bestDistance := bestDistances[i]
		for j := raceLength - 1; j > 0; j-- {
			pressTime := raceLength - j
			runningTime := raceLength - pressTime
			raceDistance := pressTime * runningTime
			if raceDistance > bestDistance {
				possibleWaysToWin = append(possibleWaysToWin, raceDistance)
			}
		}
		results = append(results, possibleWaysToWin)
	}
	x := 1
	for i := 0; i < len(results); i++ {
		x = x * len(results[i])
	}
	fmt.Println(x)
}

func toResultArray(line string) []int {
	var results []int = []int{}
	var runes []rune = []rune(line)
	for i := 0; i < len(runes); i++ {
		if unicode.IsNumber(runes[i]) {
			var numberString string = ""
			for i != len(line) && unicode.IsNumber(runes[i]) {
				numberString += string(line[i])
				i++
			}
			results = append(results, stringToNumber(numberString))
		}
	}
	return results
}

func stringToNumber(toCovert string) int {
	result, err := strconv.Atoi(toCovert)
	if err != nil {
		log.Fatal("Error during conversion")
	}
	return result
}
