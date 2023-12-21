package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

type lens struct {
	label       string
	focalLength int
}

func main() {
	input := utils.InitInputFile("initSequenceData.txt")
	defer input.Close()

	var boxes map[int64][]*lens = map[int64][]*lens{}
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		words := strings.Split(line, ",")
		for i := 0; i < len(words); i++ {
			var runes []rune = []rune(words[i])
			var opIndex int = findOpIndex(runes)
			var hash int64 = 0
			var label string = ""
			for j := 0; j < opIndex; j++ {
				r := int64(runes[j])
				hash += r
				hash = hash * 17
				hash = hash % 256
				label += string(runes[j])
			}

			var op rune = runes[opIndex]
			var lenses []*lens = boxes[hash]
			if op == '=' {
				var focalLength int = utils.StringToInt(string(runes[len(runes)-1]))
				found, index := getIndexOfLense(lenses, label)
				if found {
					lenses[index] = &lens{label: label, focalLength: focalLength}
				} else {
					lenses = append(lenses, &lens{label: label, focalLength: focalLength})
				}
			} else {
				found, index := getIndexOfLense(lenses, label)
				if found {
					lenses = append(lenses[:index], lenses[index+1:]...)
				}
			}
			if len(lenses) == 0 {
				delete(boxes, hash)
			} else {
				boxes[hash] = lenses
			}
		}
	}

	var result int64 = 0
	for hash, lenses := range boxes {
		var boxNumber int64 = hash + 1
		for i := 0; i < len(lenses); i++ {
			var slotNumber int64 = int64(i + 1)
			var focalLength int64 = int64(lenses[i].focalLength)
			result += boxNumber * slotNumber * focalLength
		}
	}

	fmt.Println("Part 2 result: ", result)
}

func getIndexOfLense(lenses []*lens, label string) (bool, int) {
	for y := 0; y < len(lenses); y++ {
		if lenses[y].label == label {
			return true, y
		}
	}
	return false, -1
}

func findOpIndex(runes []rune) int {
	for i := len(runes) - 1; i >= 0; i-- {
		if runes[i] == '=' {
			return i
		}
		if runes[i] == '-' {
			return i
		}
	}
	panic("= or - required")
}

func partOne() {
	input := utils.InitInputFile("initSequenceData.txt")
	defer input.Close()

	var result int64 = 0
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		words := strings.Split(line, ",")
		var wordTotal int64 = 0
		for i := 0; i < len(words); i++ {
			var runes []rune = []rune(words[i])
			var charTotal int64 = 0
			for j := 0; j < len(runes); j++ {
				r := int64(runes[j])
				charTotal += r
				charTotal = charTotal * 17
				charTotal = charTotal % 256
			}
			wordTotal += charTotal
		}
		result += wordTotal
	}

	fmt.Println("Part 1 result: ", result)
	fmt.Println("Part 2 result: ", result)
}
