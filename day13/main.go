package main

import (
	"adventofcode/utils"
	"fmt"
)

func main() {
	input := utils.InitInputFile("mirrorData.txt")
	defer input.Close()

	var maxNoOfSmudges int = 1

	var pattern [][]rune = getNextPattern(input)
	var rowTotal int = 0
	var colTotal int = 0
	for len(pattern) > 0 {
		reflectionFound, indexFoundAt := findReflection(pattern, maxNoOfSmudges)
		if reflectionFound {
			rowTotal += indexFoundAt
		} else {
			var transposedPattern = utils.Transpose(pattern)
			reflectionFound, indexFoundAt = findReflection(transposedPattern, maxNoOfSmudges)
			if reflectionFound {
				colTotal += indexFoundAt
			} else {
				fmt.Print()
			}
		}
		pattern = getNextPattern(input)
	}

	var partOneResult int = rowTotal*100 + colTotal
	fmt.Println("Part 1 result: ", partOneResult)
}

func findReflection(pattern [][]rune, requiredNoOfSmudges int) (bool, int) {
	var reflectionFound bool = false
	var indexFoundAt int = 0
	var iterationsFromFound int = 0
	var prevRow []rune = []rune{}
	var smudgeCount = 0
	for i := 0; i < len(pattern); i++ {
		var row []rune = pattern[i]
		firstPerfectMatch, firstSmudgeMatch := matchOrSmudgeMatch(row, prevRow, requiredNoOfSmudges)
		if firstSmudgeMatch {
			smudgeCount++
		}
		if firstPerfectMatch || firstSmudgeMatch && smudgeCount <= requiredNoOfSmudges {
			reflectionFound = true
			indexFoundAt = i
			iterationsFromFound = 1
			var smudgeCountBeforeReflectionCheck int = smudgeCount
			// loop again to see if rest reflects. If not carry on where it left off
			for j := indexFoundAt + 1; j < len(pattern); j++ {
				iterationsFromFound++
				var reflectedIndex int = indexFoundAt - iterationsFromFound
				if reflectedIndex < 0 {
					break
				}
				var rowReflectedToRight []rune = pattern[j]
				var rowReflectedToLeft []rune = pattern[reflectedIndex]

				perfectMatch, smudgeMatch := matchOrSmudgeMatch(rowReflectedToRight, rowReflectedToLeft, requiredNoOfSmudges)
				if smudgeMatch {
					smudgeCount++
				}
				if !smudgeMatch && !perfectMatch {
					reflectionFound = false
					break
				}
			}
			if smudgeCount != requiredNoOfSmudges {
				reflectionFound = false
			}
			smudgeCount = smudgeCountBeforeReflectionCheck
			if reflectionFound {
				return true, indexFoundAt
			}
		}
		smudgeCount = 0
		prevRow = row
	}
	return reflectionFound, indexFoundAt
}

func matchOrSmudgeMatch(arrA []rune, arrB []rune, noOfSmudges int) (bool, bool) {
	if len(arrA) != len(arrB) {
		return false, false
	}
	var noOfMatches int = 0
	for i, v := range arrA {
		if v == arrB[i] {
			noOfMatches++
		}
	}
	var perfectMatch = noOfMatches == len(arrA)
	var smudgeMatch = noOfMatches+noOfSmudges == len(arrA)
	return perfectMatch, smudgeMatch
}

func getNextPattern(input utils.InputFile) [][]rune {
	var grid [][]rune = [][]rune{}
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		if line == "" {
			break
		}
		var runes []rune = []rune(line)
		grid = append(grid, runes)
	}
	return grid
}
