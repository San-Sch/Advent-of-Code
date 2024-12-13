package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	puzzle, position := readFile()
	direction := [2]int{-1, 0}
	moving := true
	puzzle[position[0]][position[1]] = "X"
	for moving {
		position, direction = move(puzzle, position, direction)
		if !onMap(len(puzzle), len(puzzle[0]), position) {
			moving = false
		} else {
			puzzle[position[0]][position[1]] = "X"
		}
	}
	/*for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle); j++ {
			fmt.Print(puzzle[i][j])
		}
		fmt.Print("\n")
	}*/
	fmt.Println(countPositions(puzzle))
}

func countPositions(puzzle [][]string) int {
	count := 0
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle); j++ {
			if puzzle[i][j] == "X" {
				count++
			}
		}
	}
	return count
}

func onMap(i int, j int, position [2]int) bool {
	if position[0] < 0 || position[0] >= i || position[1] < 0 || position[1] >= j {
		return false
	}
	return true
}

func move(puzzle [][]string, position [2]int, direction [2]int) ([2]int, [2]int) {

	for onMap(len(puzzle), len(puzzle[0]), [2]int{position[0] + direction[0], position[1] + direction[1]}) &&
		puzzle[position[0]+direction[0]][position[1]+direction[1]] == "#" {
		direction = changeDirection(direction)
	}
	return [2]int{position[0] + direction[0], position[1] + direction[1]}, direction
}

func changeDirection(direction [2]int) [2]int {
	newDirection := [2]int{0, 0}
	if direction[0] == 1 {
		newDirection[1] = -1
	} else if direction[0] == -1 {
		newDirection[1] = 1
	} else if direction[1] == -1 {
		newDirection[0] = -1
	} else if direction[1] == 1 {
		newDirection[0] = 1
	}
	return newDirection
}

func readFile() ([][]string, [2]int) {
	file, _ := os.Open("2024/Day 6/input")
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	puzzle := make([][]string, 0)
	startPosition := [2]int{0, 0}
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	for i := 0; i < len(lines); i++ {
		puzzle = append(puzzle, make([]string, 0))
		for j := 0; j < len(lines[i]); j++ {
			if lines[i][j:j+1] == "^" {
				startPosition[0] = i
				startPosition[1] = j
			}
			puzzle[i] = append(puzzle[i], lines[i][j:j+1])
		}
	}
	return puzzle, startPosition
}
