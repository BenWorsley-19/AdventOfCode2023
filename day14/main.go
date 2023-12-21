package main

import (
	"adventofcode/utils"
	"fmt"
	"slices"
	"strings"
)

func main() {
	input := utils.InitInputFile("stoneData.txt")
	defer input.Close()
	var grid [][]rune = input.ToRuneGrid()

	var seen []string = []string{hash(grid)}
	var first int = -1
	for {
		grid = performCycle(grid)
		gridHash := hash(grid)
		first = slices.Index(seen, gridHash)
		if first != -1 {
			break
		}
		seen = append(seen, gridHash)
	}

	var finalGridPosition int = (1000000000-first)%(len(seen)-first) + first
	grid = unhash(seen[finalGridPosition])

	var result int = calcLoadOfStones(grid)
	fmt.Print("Result:", result)
}

func performCycle(grid [][]rune) [][]rune {
	for i := 0; i < 4; i++ {
		tilt(grid)
		grid = utils.FlipHorizontal(grid)
		grid = utils.Transpose(grid)
	}
	return grid
}

func tilt(grid [][]rune) {
	for col := 0; col < len(grid); col++ {
		for row := 0; row < len(grid); row++ {
			if grid[row][col] == '#' {
				continue
			}

			var count int = 0
			var start = row
			for ; row < len(grid); row++ {
				if grid[row][col] == '#' {
					break
				}
				if grid[row][col] == 'O' {
					count++
				}
			}

			for j := start; j < start+count; j++ {
				if grid[j][col] != '#' {
					grid[j][col] = 'O'
				}
			}
			for j := start + count; j < row; j++ {
				if grid[j][col] != '#' {
					grid[j][col] = '.'
				}
			}
		}
	}
}

func calcLoadOfStones(grid [][]rune) int {
	var result int = 0
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			if grid[row][col] == 'O' {
				result += len(grid[0]) - row
			}
		}
	}
	return result

}

func hash(grid [][]rune) string {
	sb := strings.Builder{}
	for row := 0; row < len(grid); row++ {
		for col := 0; col < len(grid[0]); col++ {
			sb.WriteRune(grid[row][col])
		}
		sb.WriteRune(',')
	}

	return sb.String()
}

func unhash(hash string) [][]rune {
	var lines []string = strings.Split(hash, ",")
	var grid [][]rune = [][]rune{}
	for i := 0; i < len(lines)-1; i++ {
		grid = append(grid, []rune(lines[i]))
	}
	return grid
}
