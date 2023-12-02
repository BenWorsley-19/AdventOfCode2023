package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filePath = "gamesData.txt"
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var result int = 0
	for fileScanner.Scan() {
		var gameLine string = fileScanner.Text()
		fmt.Println("line:")
		fmt.Println(gameLine)
		var idAndHands []string = strings.Split(gameLine, ":")
		// var id string = strings.Split(idAndHands[0], " ")[1]
		var hands []string = strings.Split(idAndHands[1], ";")

		var highestRed int = 0
		var highestBlue int = 0
		var highestGreen int = 0
		for i := 0; i < len(hands); i++ {
			var cubes []string = strings.Split(hands[i], ",")
			for j := 0; j < len(cubes); j++ {
				var numberAndColor []string = strings.Split(cubes[j], " ")
				var numberAsString string = numberAndColor[1]
				numberOfCubes, err := strconv.Atoi(numberAsString)
				if err != nil {
					log.Fatal("Error during conversion")
				}
				var color string = numberAndColor[2]
				if "red" == color && numberOfCubes > highestRed {
					highestRed = numberOfCubes
				} else if "blue" == color && numberOfCubes > highestBlue {
					highestBlue = numberOfCubes
				} else if "green" == color && numberOfCubes > highestGreen {
					highestGreen = numberOfCubes
				}
			}
		}
		// result += calcPartOne(highestRed, highestBlue, highestGreen, id)
		result += calcPartTwo(highestRed, highestBlue, highestGreen)
	}
	fmt.Println("result:")
	fmt.Println(result)

	readFile.Close()
}

func calcPartOne(highestRed int, highestBlue int, highestGreen int, id string) int {
	var maxRedCubes int = 12
	var maxGreenCubes int = 13
	var maxBlueCubes int = 14
	if highestRed <= maxRedCubes && highestBlue <= maxBlueCubes && highestGreen <= maxGreenCubes {
		idToAdd, err := strconv.Atoi(id)
		if err != nil {
			log.Fatal("Error during conversion")
		}
		return idToAdd
	}
	return 0
}

func calcPartTwo(highestRed int, highestBlue int, highestGreen int) int {
	var power int = highestRed * highestBlue * highestGreen
	return power
}
