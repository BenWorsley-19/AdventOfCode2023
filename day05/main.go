package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var filePath = "soilData.txt"
	readFile, err := os.Open(filePath)
	defer readFile.Close()
	if err != nil {
		log.Fatal(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	seedsLine := fileScanner.Text()
	seeds := strings.Split(seedsLine, " ")

	var lowestLocalDest int64 = math.MaxInt64
	for i := 1; i < len(seeds); i += 2 {
		fmt.Println("seed", i)
		startOfRange := stringToInt64(seeds[i])
		rangeLength := stringToInt64(seeds[i+1])
		endOfRange := startOfRange + rangeLength - 1
		for j := startOfRange; j <= endOfRange; j++ {
			readFile, err := os.Open(filePath)
			if err != nil {
				log.Fatal(err)
			}

			fileScanner := bufio.NewScanner(readFile)
			fileScanner.Split(bufio.ScanLines)

			fileScanner.Scan()
			// skip the empty line
			fileScanner.Scan()

			soilRanges, seedRanges := buildRangeMappings(fileScanner)
			soilDest := getCorrespondingDestination(j, soilRanges, seedRanges)
			fertalizerRanges, soilRanges := buildRangeMappings(fileScanner)
			fertalizerDest := getCorrespondingDestination(soilDest, fertalizerRanges, soilRanges)
			waterRanges, fertalizerRanges := buildRangeMappings(fileScanner)
			waterDest := getCorrespondingDestination(fertalizerDest, waterRanges, fertalizerRanges)
			lightRanges, waterRanges := buildRangeMappings(fileScanner)
			lightDest := getCorrespondingDestination(waterDest, lightRanges, waterRanges)
			tempRanges, lightRanges := buildRangeMappings(fileScanner)
			tempDest := getCorrespondingDestination(lightDest, tempRanges, lightRanges)
			humRanges, tempRanges := buildRangeMappings(fileScanner)
			humDest := getCorrespondingDestination(tempDest, humRanges, tempRanges)
			locRanges, humRanges := buildRangeMappings(fileScanner)
			locDest := getCorrespondingDestination(humDest, locRanges, humRanges)
			if locDest < lowestLocalDest {
				lowestLocalDest = locDest
			}

			readFile.Close()
		}
	}

	fmt.Println(lowestLocalDest)

	// get the soil to fertalizer mappings
	// destRanges2, sourceRanges2 := buildRangeMappings(fileScanner)
}

func getCorrespondingDestination(itemToSearch int64, destRanges map[int]RangeData, sourceRanges map[int]RangeData) int64 {
	for fk, sourceRangeData := range sourceRanges {
		if sourceRangeData.isInRange(itemToSearch) {
			numberToAdd := itemToSearch - sourceRangeData.rangeStart
			return destRanges[fk].rangeStart + numberToAdd
		}
	}
	return itemToSearch
}

// 	fmt.Println("result:")
// 	fmt.Println(result)
// }

func buildRangeMappings(fileScanner *bufio.Scanner) (map[int]RangeData, map[int]RangeData) {
	var destRanges map[int]RangeData = map[int]RangeData{}
	var sourceRanges map[int]RangeData = map[int]RangeData{}
	i := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if line == "" {
			break
		}
		rangePoints := strings.Split(line, " ")

		// check if it's the range data
		destRangeStartString := rangePoints[0]
		_, err := strconv.ParseInt(destRangeStartString, 10, 64)
		if err != nil {
			continue
		}

		var rangeAmount int64 = stringToInt64(rangePoints[2])
		var destRangeStart int64 = stringToInt64(destRangeStartString)
		destRange := RangeData{
			rangeAmount: rangeAmount,
			rangeStart:  destRangeStart,
			rangeEnd:    destRangeStart + rangeAmount - 1,
		}
		destRanges[i] = destRange
		var sourceRangeStart int64 = stringToInt64(rangePoints[1])
		sourceRange := RangeData{
			rangeAmount: rangeAmount,
			rangeStart:  sourceRangeStart,
			rangeEnd:    sourceRangeStart + rangeAmount - 1,
		}
		sourceRanges[i] = sourceRange
		i++
	}
	return destRanges, sourceRanges
}

type RangeData struct {
	rangeAmount int64
	rangeStart  int64
	rangeEnd    int64
}

func (n *RangeData) isInRange(sourceItem int64) bool {
	return sourceItem >= n.rangeStart && sourceItem <= n.rangeEnd
}

func stringToInt64(toCovert string) int64 {
	result, err := strconv.ParseInt(toCovert, 10, 64)
	if err != nil {
		log.Fatal("Error during conversion")
	}
	return result
}
