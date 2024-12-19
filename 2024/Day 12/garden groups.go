package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	puzzle := readFile()
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle); j++ {
			fmt.Print(puzzle[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Println()
	printPuzzle(puzzle)
}

func printPuzzle(puzzle [][]string) {
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle); j++ {
			fmt.Print(puzzle[i][j])
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func readFile() [][]string {
	file, _ := os.Open("2024/Day 12/test2")
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	puzzle := make([][]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for i := 0; i < len(lines); i++ {
		puzzle = append(puzzle, make([]string, 0))
		for j := 0; j < len(lines[i]); j++ {
			puzzle[i] = append(puzzle[i], lines[i][j:j+1])
		}
	}
	return puzzle
}
