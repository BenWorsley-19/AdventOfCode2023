package utils

import (
	"log"
	"strconv"
)

func StringArrToInt64Arr(arr []string) []int64 {
	var result []int64 = []int64{}
	for _, s := range arr {
		result = append(result, StringToInt64(s))
	}
	return result
}

func StringToInt64(toCovert string) int64 {
	result, err := strconv.ParseInt(toCovert, 10, 64)
	if err != nil {
		log.Fatal("Error during conversion")
	}
	return result
}

func StringToInt(toCovert string) int {
	result, err := strconv.Atoi(toCovert)
	if err != nil {
		log.Fatal("Error during conversion")
	}
	return result
}

func HexToInt64(hex string) int64 {
	result, err := strconv.ParseInt(hex, 16, 64)
	if err != nil {
		log.Fatal("Error during conversion")
	}
	return result
}

func RuneToInt(r rune) int {
	return StringToInt(string(r))
}
