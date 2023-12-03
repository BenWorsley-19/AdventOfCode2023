package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode"
)

func partOne() {
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
			if '.' == r {
				continue
			}
			if unicode.IsNumber(r) {
				indexNumberFoundAt := i
				i++
				var numberString string = string(r)
				var numberDigits int = 1
				// get rest of numbver
				for i != len(runes) && unicode.IsNumber(runes[i]) {
					numberString += string(runes[i])
					numberDigits++
					i++
				}
				// get points to check for symbol
				leftPoint := indexNumberFoundAt - 1
				rightPoint := indexNumberFoundAt + numberDigits
				//check current
				if checkForSymbolAtPoints(leftPoint, rightPoint, runes) {
					result += stringToNumber(numberString)
					continue
				}
				// check above
				previousLineRunes := []rune(previousLine)
				if checkForSymbolAtPoints(leftPoint, rightPoint, previousLineRunes) {
					result += stringToNumber(numberString)
					continue
				}
				// check below
				nextLineRunes := []rune(nextLine)
				if checkForSymbolAtPoints(leftPoint, rightPoint, nextLineRunes) {
					result += stringToNumber(numberString)
					continue
				}
			}
		}
		previousLine = currentLine
		currentLine = nextLine
	}
	fmt.Println("result:")
	fmt.Println(result)
}

func checkForSymbolAtPoints(leftPoint int, rightPoint int, runes []rune) bool {
	for i := leftPoint; i <= rightPoint; i++ {
		if i < 0 || i > len(runes)-1 {
			continue
		}
		if '.' != runes[i] && !unicode.IsLetter(runes[i]) && !unicode.IsNumber(runes[i]) {
			return true
		}
	}
	return false
}
