package main

import (
	"adventofcode/utils"
	"fmt"
	"strings"
)

var cache map[string]*int = map[string]*int{}

func main() {
	input := utils.InitInputFile("springData.txt")
	defer input.Close()

	var result int = 0
	for input.MoveToNextLine() {
		var line string = input.ReadLine()
		springsAndGroups := strings.Split(line, " ")
		var springs []rune = []rune(springsAndGroups[0])
		var foldedSprings []rune = []rune{}
		for i := 0; i <= 4; i++ {
			if i != 0 {
				foldedSprings = append(foldedSprings, '?')
			}
			foldedSprings = append(foldedSprings, springs...)

		}
		var groups []int = utils.StringArrayToIntArray(strings.Split(springsAndGroups[1], ","))
		var foldedGroups []int = []int{}
		for i := 0; i <= 4; i++ {
			foldedGroups = append(foldedGroups, groups...)
		}
		result += countPossibleArrangements(foldedSprings, foldedGroups)
	}

	fmt.Println("Part 2 result: ", result)
}

func countPossibleArrangements(springs []rune, groups []int) int {
	var total int = 0
	if len(springs) == 0 {
		if len(groups) == 0 {
			return 1
		}
		return 0
	}

	if len(groups) == 0 {
		if utils.ArrayContains(springs, '#') {
			return 0
		}
		return 1
	}

	var springsHash [32]byte = utils.HashRuneArray(springs)
	var groupsHash [32]byte = utils.HashIntArray(groups)
	var cacheId string = string(springsHash[:]) + string(groupsHash[:])
	if cache[cacheId] != nil {
		return *cache[cacheId]
	}

	var sumOfGroups int = 0
	for _, v := range groups {
		sumOfGroups += v
	}
	if len(springs) < sumOfGroups+len(groups)-1 {
		return 0
	}

	if springs[0] == '.' || springs[0] == '?' {
		total += countPossibleArrangements(springs[1:], groups)
	}

	var n int = groups[0]
	if (springs[0] == '#' || springs[0] == '?') &&
		!utils.ArrayContains(springs[:n], '.') &&
		(len(springs) == n || (springs[n] == '.' || springs[n] == '?')) {
		var next = []rune{}
		if n+1 < len(springs) {
			next = springs[n+1:]
		}
		total += countPossibleArrangements(next, groups[1:])
	}

	cache[cacheId] = &total
	return total
}
