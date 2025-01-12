package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
)

func main() {
	puzzle, position := readFile()
	direction := [2]int{-1, 0}
	positions := make([][2][2]int, 0)
	obstaclePositions := make([][2]int, 0)
	moving := true
	result2 := 0
	puzzle[position[0]][position[1]] = "X"
	for moving {
		position, direction = move(puzzle, position, direction)
		if !onMap(len(puzzle), len(puzzle[0]), position) {
			moving = false
		} else {
			puzzle[position[0]][position[1]] = "X"
			positions = append(positions, [2][2]int{position, direction})
		}

	}
	for _, pos := range positions {
		obstaclePosition := [2]int{pos[0][0] + pos[1][0], pos[0][1] + pos[1][1]}
		if !slices.Contains(obstaclePositions, obstaclePosition) {
			if checkForLoop(obstaclePosition, pos) {
				obstaclePositions = append(obstaclePositions, obstaclePosition)
				result2++
				fmt.Println(result2)
			}
		}
	}
	fmt.Println(countPositions(puzzle))
	fmt.Println(result2) //false
}

func checkForLoop(obstaclePos [2]int, pos [2][2]int) bool {
	moving := true
	puzzle, startPosition := readFile()
	position := pos[0]
	direction := pos[1]
	if obstaclePos[0] < 0 || obstaclePos[0] >= len(puzzle) ||
		obstaclePos[1] < 0 || obstaclePos[1] >= len(puzzle[0]) ||
		(obstaclePos[0] == startPosition[0] && obstaclePos[1] == startPosition[1]) {
		return false
	}
	puzzle[obstaclePos[0]][obstaclePos[1]] = "N"
	positions := make([][2][2]int, 0)
	positions = append(positions, [2][2]int{position, direction})
	direction = changeDirection(pos[1])

	positions = append(positions, [2][2]int{position, direction})

	for moving {
		position, direction = move(puzzle, position, direction)
		tmpPos := [2][2]int{position, direction}

		if slices.Contains(positions, tmpPos) {
			return true
		}
		if !onMap(len(puzzle), len(puzzle[0]), position) {
			moving = false
		} else {
			puzzle[position[0]][position[1]] = "X"
			positions = append(positions, [2][2]int{position, direction})

		}

	}
	return false
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
	pos := [2]int{}
	if onMap(len(puzzle), len(puzzle[0]), [2]int{position[0] + direction[0], position[1] + direction[1]}) &&
		(puzzle[position[0]+direction[0]][position[1]+direction[1]] == "#" ||
			puzzle[position[0]+direction[0]][position[1]+direction[1]] == "N") {
		direction = changeDirection(direction)
		pos = position
	} else {
		pos = [2]int{position[0] + direction[0], position[1] + direction[1]}
	}
	return pos, direction
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
