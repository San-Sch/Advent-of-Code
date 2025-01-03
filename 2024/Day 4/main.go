package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	puzzle := readFile()
	result := 0
	result2 := 0
	var direction [2]int
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
			if puzzle[i][j] == "X" {
				for x := -1; x <= 1; x++ {
					for y := -1; y <= 1; y++ {
						if x == 0 && y == 0 {
							continue
						}
						direction[0] = x
						direction[1] = y
						if find(puzzle, "M", i+x, j+y, direction) {
							result++
						}
					}
				}
			}
			// Part 2
			if puzzle[i][j] == "M" && i+2 < len(puzzle) && j+2 < len(puzzle[i]) {
				if puzzle[i+1][j+1] == "A" && puzzle[i+2][j+2] == "S" {
					if checkCross(puzzle, i, j) {
						result2++
					}
				}
			}
			if puzzle[i][j] == "S" && i+2 < len(puzzle) && j+2 < len(puzzle[i]) {
				if puzzle[i+1][j+1] == "A" && puzzle[i+2][j+2] == "M" {
					if checkCross(puzzle, i, j) {
						result2++
					}
				}
			}
		}
	}
	fmt.Println(result)
	fmt.Println(result2)

}

func checkCross(puzzle [][]string, i int, j int) bool {
	if (puzzle[i][j+2] == "S" && puzzle[i+2][j] == "M") ||
		(puzzle[i][j+2] == "M" && puzzle[i+2][j] == "S") {
		return true
	}
	return false
}

func find(puzzle [][]string, letter string, x int, y int, direction [2]int) bool {
	if x < 0 || y < 0 || x > len(puzzle)-1 || y > len(puzzle[0])-1 {
		return false
	}
	if puzzle[x][y] == letter {
		switch letter {
		case "M":
			if find(puzzle, "A", x+direction[0], y+direction[1], direction) {
				return true
			}
			break
		case "A":
			if find(puzzle, "S", x+direction[0], y+direction[1], direction) {
				return true
			}
		case "S":
			return true
		}
	}
	return false
}

func readFile() [][]string {
	file, _ := os.Open("2024/Day 4/input")
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
