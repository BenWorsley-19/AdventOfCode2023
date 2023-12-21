package utils

import (
	"bufio"
	"log"
	"os"
)

// Experimenting with doing this as a struct
type InputFile struct {
	readFile    *os.File
	fileScanner *bufio.Scanner
}

func (t *InputFile) ReadLine() string {
	return t.fileScanner.Text()
}

func (t *InputFile) MoveToNextLine() bool {
	return t.fileScanner.Scan()
}

func (t *InputFile) ToRuneGrid() [][]rune {
	var grid [][]rune = [][]rune{}
	for t.MoveToNextLine() {
		var line string = t.ReadLine()
		var runes []rune = []rune(line)
		grid = append(grid, runes)
	}
	return grid
}

func (t *InputFile) ToLineArray() []string {
	var lines []string
	for t.MoveToNextLine() {
		var line string = t.ReadLine()
		lines = append(lines, line)
	}
	return lines
}

func (t *InputFile) Close() {
	t.readFile.Close()
}

func InitInputFile(fileName string) InputFile {
	var filePath = "../inputs/" + fileName
	readFile, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}

	var fileScanner *bufio.Scanner = bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	return InputFile{readFile: readFile, fileScanner: fileScanner}
}
