package utils

import (
	"bytes"
	"crypto/sha256"
	"fmt"
)

func ArrayContains[T comparable](arr []T, r T) bool {
	for _, i := range arr {
		if i == r {
			return true
		}
	}
	return false
}

func ArraysAreEqual(arrA []rune, arrB []rune) bool {
	if len(arrA) != len(arrB) {
		return false
	}
	for i, v := range arrA {
		if v != arrB[i] {
			return false
		}
	}
	return true
}

func StringArrayToIntArray(arr []string) []int {
	var result []int = []int{}
	for _, i := range arr {
		result = append(result, StringToInt(i))
	}
	return result
}

func HashRuneArray(arr []rune) [32]byte {
	return (sha256.Sum256([]byte(string(arr))))
}

func HashIntArray(arr []int) [32]byte {
	var buffer bytes.Buffer
	for i, _ := range arr {
		buffer.WriteString(string(arr[i]))
		buffer.WriteString(",")
	}
	return (sha256.Sum256([]byte(buffer.String())))
}

func Transpose(slice [][]rune) [][]rune {
	xl := len(slice[0])
	yl := len(slice)
	result := make([][]rune, xl)
	for i := range result {
		result[i] = make([]rune, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = slice[j][i]
		}
	}
	return result
}

func FlipHorizontal(arr [][]rune) [][]rune {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}

func PrintGrid(arr [][]rune) {
	for row := 0; row < len(arr); row++ {
		for column := 0; column < len(arr[0]); column++ {
			fmt.Print(string(arr[row][column]), " ")
		}
		fmt.Print("\n")
	}
}
