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
	var filePath = "calibrationData.txt"
	readFile, err := os.Open(filePath)
	if err != nil {
		fmt.Println(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	digitMap := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}

	var digitTree *DigitTrie = InitDigitTrie()

	var result int = 0
	for fileScanner.Scan() {
		var line = fileScanner.Text()
		fmt.Println("line:")
		fmt.Println(line)
		var numberFromLine int = combineFirstAndLastDigitIncludingText(line, digitMap, digitTree)
		fmt.Println("numberfromline:")
		fmt.Println(numberFromLine)
		result += numberFromLine
		fmt.Println("result:")
		fmt.Println(result)
	}

	readFile.Close()
}

// part 1
func combineFirstAndLastDigit(line string) int {
	var firstDigit = "nf"
	var lastDigit = "nf"

	runes := []rune(line)
	for i := 0; i < len(runes); i++ {
		var rune = runes[i]
		if unicode.IsNumber(rune) && firstDigit == "nf" {
			firstDigit = string(rune)
			lastDigit = string(rune)
		} else if unicode.IsNumber(rune) && firstDigit != "nf" {
			lastDigit = string(rune)
		}
	}
	var combinedDigits = firstDigit + lastDigit
	result, err := strconv.Atoi(combinedDigits)
	if err != nil {
		log.Fatal("Error during conversion")
	}
	return result
}

// part 2
func combineFirstAndLastDigitIncludingText(line string, digitMap map[string]string, digitTree *DigitTrie) int {
	var firstDigit = "nf"
	var lastDigit = "nf"

	runes := []rune(line)
	for i := 0; i < len(runes); i++ {
		var rune = runes[i]
		if unicode.IsNumber(rune) && firstDigit == "nf" {
			firstDigit = string(rune)
			lastDigit = string(rune)
			continue
		} else if unicode.IsNumber(rune) && firstDigit != "nf" {
			lastDigit = string(rune)
			continue
		}

		var node *DigitNode = digitTree.IsStartOfDigit(rune)
		if nil == node {
			continue
		}
		var potentialWord string = string(rune)
		for j := i + 1; j < len(runes); j++ {
			potentialWord += string(runes[j])
			if potentialWord == "eigh" {
				fmt.Print("zsdas")
			}
			node = digitTree.IsNextCharInDigit(node, runes[j])
			if nil == node {
				break
			}
			if node.isWord {
				if firstDigit == "nf" {
					firstDigit = digitMap[potentialWord]
					lastDigit = digitMap[potentialWord]
					break
				} else if unicode.IsNumber(rune) && firstDigit != "nf" {
					lastDigit = digitMap[potentialWord]
					break
				}
			}
		}

	}
	var combinedDigits = firstDigit + lastDigit
	result, err := strconv.Atoi(combinedDigits)
	if err != nil {
		log.Fatal("Error during conversion")
	}
	return result
}
