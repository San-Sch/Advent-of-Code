package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	puzzle := readFile()
	posMap := make(map[string][][2]int)
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle); j++ {
			if puzzle[i][j] != "." && puzzle[i][j] != "#" {
				posMap[puzzle[i][j]] = append(posMap[puzzle[i][j]], [2]int{i, j})
			}
		}
	}
	for _, val := range posMap {
		for idxPos := range val {
			if idxPos < len(val)-1 {
				for i := idxPos + 1; i < len(val); i++ {
					setAntinodes(val[idxPos], val[i], puzzle)
				}
			}
		}
	}
	fmt.Println(countPositions(puzzle))

	for _, val := range posMap {
		for idxPos := range val {
			if idxPos < len(val)-1 {
				for i := idxPos + 1; i < len(val); i++ {
					setAntinodes2(val[idxPos], val[i], puzzle)
				}
			}
		}
	}
	//printPuzzle(puzzle)
	fmt.Println(countPositions(puzzle))
}

func countPositions(puzzle [][]string) int {
	count := 0
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle); j++ {
			if puzzle[i][j] == "#" {
				count++
			}
		}
	}
	return count
}

func setAntinodes(pos1 [2]int, pos2 [2]int, puzzle [][]string) {
	posAntinode1 := [2]int{pos2[0] + pos2[0] - pos1[0], pos2[1] + pos2[1] - pos1[1]}
	posAntinode2 := [2]int{pos1[0] + pos1[0] - pos2[0], pos1[1] + pos1[1] - pos2[1]}
	if onMap(len(puzzle), len(puzzle[0]), posAntinode1) {
		puzzle[posAntinode1[0]][posAntinode1[1]] = "#"
	}
	if onMap(len(puzzle), len(puzzle[0]), posAntinode2) {
		puzzle[posAntinode2[0]][posAntinode2[1]] = "#"
	}
}

func setAntinodes2(pos1 [2]int, pos2 [2]int, puzzle [][]string) {
	diffY := pos2[0] - pos1[0]
	diffX := pos2[1] - pos1[1]
	pos := pos1
	pos[0] += diffY
	pos[1] += diffX
	puzzle[pos1[0]][pos1[1]] = "#"
	for onMap(len(puzzle), len(puzzle[0]), pos) {
		puzzle[pos[0]][pos[1]] = "#"
		pos[0] += diffY
		pos[1] += diffX
	}
	pos = pos1
	pos[0] -= diffY
	pos[1] -= diffX
	for onMap(len(puzzle), len(puzzle[0]), pos) {
		puzzle[pos[0]][pos[1]] = "#"
		pos[0] -= diffY
		pos[1] -= diffX
	}
}

func onMap(i int, j int, position [2]int) bool {
	if position[0] < 0 || position[0] >= i || position[1] < 0 || position[1] >= j {
		return false
	}
	return true
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
	file, _ := os.Open("2024/Day 8/input")
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
