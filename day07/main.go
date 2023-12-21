package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Hand struct {
	hand     string
	handVal  int
	handType int
}

var cardLabels []rune = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

func main() {
	var filePath = "cardsData.txt"
	readFile, err := os.Open(filePath)
	defer readFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	var handRank []Hand = []Hand{}
	for fileScanner.Scan() {
		var handLine string = fileScanner.Text()
		var handAndVal []string = strings.Split(handLine, " ")
		// get the type of the hand
		var hand string = handAndVal[0]
		var handVal int = stringToNumber(handAndVal[1])
		var handType int = determineHandType(hand)
		handStruct := Hand{
			hand:     hand,
			handVal:  handVal,
			handType: handType,
		}

		handRank = append(handRank, handStruct)
	}
	// sort the type where weakest type is first (rank 1), if same type then get the first different char and order by strongest cardlabel
	// times the hand val by the rank and add to the result
	quickSort(handRank, 0, len(handRank)-1)
	var result int = 0
	for i := 0; i < len(handRank); i++ {
		result += handRank[i].handVal * (i + 1)
	}
	fmt.Println("result: ", result)
}

func quickSort(arr []Hand, lo int, hi int) {
	if lo >= hi {
		return
	}
	var pivotIndex int = partition(arr, lo, hi)
	quickSort(arr, lo, pivotIndex-1)
	quickSort(arr, pivotIndex+1, hi)
}

func partition(arr []Hand, lo int, hi int) int {
	var pivot Hand = arr[hi]
	var index int = lo - 1
	for i := lo; i < hi; i++ {
		if arr[i].handType > pivot.handType {
			continue
		}
		if arr[i].handType == pivot.handType {
			// if the same hand type, then the hand with the highest different card is reater
			var hasStrongerCards bool = false
			runes := []rune(arr[i].hand)
			pivotRunes := []rune(pivot.hand)
			for j := 0; j < len(runes); j++ {
				cardStrength := indexOf(cardLabels, runes[j])
				pivotCardStregth := indexOf(cardLabels, pivotRunes[j])
				if cardStrength < pivotCardStregth {
					hasStrongerCards = true
					break
				}
				// if it's not less than but not equal to, it's greater than so we no hasStrongerCards is false
				if cardStrength != pivotCardStregth {
					break
				}
			}
			if hasStrongerCards {
				continue
			}
		}

		index++
		var tmp Hand = arr[i]
		arr[i] = arr[index]
		arr[index] = tmp
	}
	index++
	arr[hi] = arr[index]
	arr[index] = pivot
	return index
}

// 6 Five of a kind
// 5 Four of a kind
// 4 Full house
// 3 Three of a kind
// 2 Two pair
// 1 One pair
// 0 High card
func determineHandType(hand string) int {
	var cardCount map[rune]int = map[rune]int{}
	for i := 0; i < len(cardLabels); i++ {
		cardCount[cardLabels[i]] = 0
	}

	runes := []rune(hand)
	var noOfJokers int = 0
	for i := 0; i < len(runes); i++ {
		if runes[i] == 'J' {
			noOfJokers++
		} else {
			cardCount[runes[i]] = cardCount[runes[i]] + 1
		}
	}

	// add the counts to a list
	var counts []int = []int{}
	for _, v := range cardCount {
		counts = append(counts, v)
	}

	sort.Ints(counts[:])

	var highest = counts[12]
	var secondHighest = counts[11]

	if highest+noOfJokers >= 5 {
		return 6
	}
	if highest+noOfJokers >= 4 {
		return 5
	}
	if highest+secondHighest+noOfJokers >= 5 {
		return 4
	}
	if highest+noOfJokers >= 3 {
		return 3
	}
	if highest+secondHighest+noOfJokers >= 4 {
		return 2
	}
	if highest+noOfJokers >= 2 {
		return 1
	}
	return 0
}

func indexOf(arr []rune, value rune) int {
	for i, v := range arr {
		if v == value {
			return i
		}
	}
	return -1
}

func stringToNumber(toCovert string) int {
	result, err := strconv.Atoi(toCovert)
	if err != nil {
		log.Fatal("Error during conversion")
	}
	return result
}
