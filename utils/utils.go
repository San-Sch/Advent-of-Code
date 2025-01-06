package utils

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Coordinate struct {
	X int
	Y int
}

func Diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func CreatePuzzleMap(data string) [][]string {
	scanner := bufio.NewScanner(strings.NewReader(data))
	puzzleMap := make([][]string, 0)
	for scanner.Scan() {
		line := scanner.Text()
		puzzleMap = append(puzzleMap, make([]string, 0))
		for x := 0; x < len(line); x++ {
			puzzleMap[len(puzzleMap)-1] = append(puzzleMap[len(puzzleMap)-1], line[x:x+1])
		}
	}
	return puzzleMap
}

func PrintPuzzleMap(puzzle [][]string) {
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle); j++ {
			fmt.Print(puzzle[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func OnMap(height int, width int, position Coordinate) bool {
	if position.Y < 0 || position.Y >= height || position.X < 0 || position.X >= width {
		return false
	}
	return true
}

func ReadFile(path string) string {
	data, _ := os.ReadFile(path)
	return string(data)
}
