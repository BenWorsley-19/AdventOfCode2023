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
	partTwo()
}

func partTwo() {
	var filePath = "engineData.txt"
	readFile, err := os.Open(filePath)
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var result int = 0
	previousLine := ""
	currentLine := ""
	nextLine := ""

	if fileScanner.Scan() {
		currentLine = fileScanner.Text()
	}

	for "" != currentLine {
		if fileScanner.Scan() {
			nextLine = fileScanner.Text()
		} else {
			nextLine = ""
		}
		runes := []rune(currentLine)
		for i := 0; i < len(runes); i++ {
			var r = runes[i]
			if '*' != r {
				continue
			}
			var rightPoint = i + 1
			var leftPoint = i - 1
			var currentLineFoundNumbers []int = checkLineForNumbers(runes, rightPoint, leftPoint)
			// get numbers from previous line
			var previousLineRunes []rune = []rune(previousLine)
			var previousLineFoundNumbers []int = checkLineForNumbers(previousLineRunes, rightPoint, leftPoint)
			// get numbers from next line
			var nextLineRunes []rune = []rune(nextLine)
			var nextLineFoundNumbers []int = checkLineForNumbers(nextLineRunes, rightPoint, leftPoint)
			var noOfFindNumbers int = len(currentLineFoundNumbers) + len(previousLineFoundNumbers) + len(nextLineFoundNumbers)
			if noOfFindNumbers != 2 {
				continue
			}
			result += multiplyNumbers(currentLineFoundNumbers, previousLineFoundNumbers, nextLineFoundNumbers)
		}
		previousLine = currentLine
		currentLine = nextLine
	}
	fmt.Println("result:")
	fmt.Println(result)
}

func checkLineForNumbers(line []rune, rightPoint int, leftPoint int) []int {
	var results []int = []int{}
	for i := leftPoint; i <= rightPoint; i++ {
		if i < 0 || i > len(line)-1 {
			continue
		}
		if unicode.IsNumber(line[i]) {
			var numberString string = ""
			var indexNumberFoundAt int = i
			for i != len(line) && unicode.IsNumber(line[i]) {
				numberString += string(line[i])
				// increment i so that when we loop i the outer loop, it starts from where we looped through here
				i++
			}
			var j int = indexNumberFoundAt - 1
			var precedingDigits []rune = []rune{}
			for j >= 0 && unicode.IsNumber(line[j]) {
				precedingDigits = append(precedingDigits, line[j])
				j--
			}
			var precedingNumberString string = ""
			for y := len(precedingDigits) - 1; y >= 0; y-- {
				precedingNumberString += string(precedingDigits[y])
			}
			var foundNumber int = stringToNumber(precedingNumberString + numberString)
			results = append(results, foundNumber)
		}
	}
	return results
}

func multiplyNumbers(currentLineNumbers []int, previousLineNumbers []int, nextLineNumbers []int) int {
	numbers := append(currentLineNumbers, previousLineNumbers...)
	numbers = append(numbers, nextLineNumbers...)
	return numbers[0] * numbers[1]
}

func stringToNumber(toCovert string) int {
	result, err := strconv.Atoi(toCovert)
	if err != nil {
		log.Fatal("Error during conversion")
	}
	return result
}
