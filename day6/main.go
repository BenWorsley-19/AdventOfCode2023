package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"unicode"
)

func main() {
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

	var raceLength int64 = toRaceNumber(raceLengthsLine)
	var bestDistance int64 = toRaceNumber(bestDistancesLine)

	var possibleWaysToWin int64 = 0
	for j := raceLength - 1; j > 0; j-- {
		pressTime := raceLength - j
		runningTime := raceLength - pressTime
		raceDistance := pressTime * runningTime
		if raceDistance > bestDistance {
			possibleWaysToWin++
		}
	}

	fmt.Println("possible ways to win: ", possibleWaysToWin)
}

func toRaceNumber(line string) int64 {
	var numberString string = ""
	var runes []rune = []rune(line)
	for i := 0; i < len(runes); i++ {
		if unicode.IsNumber(runes[i]) {
			for ; i != len(line); i++ {
				if unicode.IsNumber(runes[i]) {
					numberString += string(line[i])
				}
			}
		}
	}
	return stringToInt64(numberString)
}

func stringToInt64(toCovert string) int64 {
	result, err := strconv.ParseInt(toCovert, 10, 64)
	if err != nil {
		log.Fatal("Error during conversion")
	}
	return result
}
