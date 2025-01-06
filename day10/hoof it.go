package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	result1 := 0
	result2 := 0
	puzzle := readFile()
	positions := make([][2]int, 0)
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[0]); j++ {
			if puzzle[i][j] == "0" {
				positions = nil
				result2 += countTrailheads(puzzle, [2]int{i, j}, "0", &positions)
				result1 += len(positions)
			}
		}
	}
	fmt.Println(result1)
	fmt.Println(result2)
}

func countTrailheads(puzzle [][]string, position [2]int, height string, positions *[][2]int) int {
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	newHeight := addHeight(height)
	trailheads := 0
	for _, val := range directions {
		pos := [2]int{position[0] + val[0], position[1] + val[1]}
		if onMap(len(puzzle), len(puzzle[0]), pos) {
			if newHeight == puzzle[pos[0]][pos[1]] {
				if newHeight == "9" {
					if !slices.Contains(*positions, [2]int{pos[0], pos[1]}) {
						*positions = append(*positions, [2]int{pos[0], pos[1]})
					}
					trailheads++
				} else {
					trailheads += countTrailheads(puzzle, pos, newHeight, positions)
				}
			}
		}
	}
	return trailheads
}

func addHeight(height string) string {
	heightNum, _ := strconv.Atoi(height)
	return strconv.Itoa(heightNum + 1)
}

func onMap(i int, j int, position [2]int) bool {
	if position[0] < 0 || position[0] >= i || position[1] < 0 || position[1] >= j {
		return false
	}
	return true
}

func printTrailhead(puzzle [][]string, positions [][2]int) {
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle); j++ {
			if slices.Contains(positions, [2]int{i, j}) {
				fmt.Print(puzzle[i][j])
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
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
	file, _ := os.Open("2024/Day 10/input")
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
