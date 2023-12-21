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
	var filePath = "scratchcardData.txt"
	readFile, err := os.Open(filePath)
	defer readFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var result int = 0
	var copies map[int]int = map[int]int{}
	for fileScanner.Scan() {

		line := fileScanner.Text()
		runes := []rune(line)

		// skip the word Card
		var i int = 4
		var currentRune rune = 'C'
		var cardNumber int = getCardNumber(&i, runes)
		for currentRune != ':' {
			currentRune = runes[i]
			i++
		}

		var winningNumbers map[int]bool = map[int]bool{}
		addWinningNumbersToMap(&i, runes, winningNumbers)
		calculateCopies(&i, runes, winningNumbers, cardNumber, copies)
	}
	result = addUpCopies(copies)
	fmt.Println("result:")
	fmt.Println(result)
}

func getCardNumber(i *int, runes []rune) int {
	for ; *i < len(runes); *i++ {
		var r = runes[*i]
		if !unicode.IsNumber(r) {
			continue
		}
		return getRestOfNumber(i, runes)
	}
	panic("must always be a card number")
}

func calculateCopies(i *int, runes []rune, winningNumbers map[int]bool, cardNumber int, copies map[int]int) {
	// We've always got at least 1 copy
	copies[cardNumber] = copies[cardNumber] + 1
	var noOfFoundNumbers int = 0
	for j := copies[cardNumber]; j > 0; j-- {
		for ; *i < len(runes); *i++ {
			var r = runes[*i]
			if !unicode.IsNumber(r) {
				continue
			}
			var fullNumber int = getRestOfNumber(i, runes)
			if winningNumbers[fullNumber] {
				noOfFoundNumbers++
			}
		}
		for y := cardNumber + noOfFoundNumbers; y > cardNumber; y-- {
			copies[y] = copies[y] + 1
		}
	}
}

func addUpCopies(copies map[int]int) int {
	var total int = 0
	for _, noOfCopies := range copies {
		total += noOfCopies
	}
	return total
}

func partOne() {
	var filePath = "scratchcardData.txt"
	readFile, err := os.Open(filePath)
	defer readFile.Close()
	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var result int = 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		runes := []rune(line)

		// skip the word Card
		var i int = 4
		var currentRune rune = 'C'
		for currentRune != ':' {
			currentRune = runes[i]
			i++
		}

		var winningNumbers map[int]bool = map[int]bool{}
		addWinningNumbersToMap(&i, runes, winningNumbers)
		var lineResult int = calculateWinningNumbers(&i, runes, winningNumbers)
		result += lineResult
	}

	fmt.Println("result:")
	fmt.Println(result)
}

func addWinningNumbersToMap(i *int, runes []rune, winningNumbers map[int]bool) {
	for ; *i < len(runes); *i++ {
		var r = runes[*i]
		if r == '|' {
			break
		}
		if !unicode.IsNumber(r) {
			continue
		}
		var fullNumber int = getRestOfNumber(i, runes)
		winningNumbers[fullNumber] = true
	}
}

func calculateWinningNumbers(i *int, runes []rune, winningNumbers map[int]bool) int {
	var lineResult int = 0
	for ; *i < len(runes); *i++ {
		var r = runes[*i]
		if !unicode.IsNumber(r) {
			continue
		}
		var fullNumber int = getRestOfNumber(i, runes)
		if winningNumbers[fullNumber] {
			if lineResult == 0 {
				lineResult++
			} else {
				lineResult = lineResult * 2
			}
		}
	}
	return lineResult
}

func getRestOfNumber(i *int, runes []rune) int {
	var numberString string = ""
	for *i != len(runes) && unicode.IsNumber(runes[*i]) {
		numberString += string(runes[*i])
		*i++
	}
	return stringToNumber(numberString)
}

func stringToNumber(toCovert string) int {
	result, err := strconv.Atoi(toCovert)
	if err != nil {
		log.Fatal("Error during conversion")
	}
	return result
}
