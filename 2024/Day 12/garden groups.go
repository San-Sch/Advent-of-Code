package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	puzzle := readFile()
	areas := make(map[string]bool)
	areaPlots := make(map[string][]int)
	areaSides := make(map[string]int)

	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle); j++ {
			if len(puzzle[i][j]) == 1 {
				numberStr := "-0"
				for number := 0; areas[puzzle[i][j]+numberStr]; number++ {
					numberStr = "-" + strconv.Itoa(number)
				}
				puzzle[i][j] = puzzle[i][j] + numberStr
				areas[puzzle[i][j]] = true
				spread(puzzle, [2]int{i, j}, areaPlots)
			}
		}
	}
	sumPrice := 0
	for _, val := range areaPlots {
		sumP := 0
		for _, plot := range val {
			sumP += plot
		}
		sumPrice += len(val) * sumP
	}
	fmt.Println(sumPrice)
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle); j++ {
			areaSides[puzzle[i][j]] += corners(puzzle, [2]int{i, j})
		}
	}
	sumPrice = 0
	for idx, val := range areaPlots {
		sumPrice += len(val) * areaSides[idx]
	}
	fmt.Println(sumPrice)
}

func corners(puzzle [][]string, pos [2]int) int {
	count := 0
	if (!onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0] - 1, pos[1]}) ||
		puzzle[pos[0]][pos[1]] != puzzle[pos[0]-1][pos[1]]) &&
		(!onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0], pos[1] - 1}) ||
			puzzle[pos[0]][pos[1]] != puzzle[pos[0]][pos[1]-1]) {
		count++
	}
	if (!onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0] - 1, pos[1]}) ||
		puzzle[pos[0]][pos[1]] != puzzle[pos[0]-1][pos[1]]) &&
		(!onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0], pos[1] + 1}) ||
			puzzle[pos[0]][pos[1]] != puzzle[pos[0]][pos[1]+1]) {
		count++
	}
	if (!onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0] + 1, pos[1]}) ||
		puzzle[pos[0]][pos[1]] != puzzle[pos[0]+1][pos[1]]) &&
		(!onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0], pos[1] - 1}) ||
			puzzle[pos[0]][pos[1]] != puzzle[pos[0]][pos[1]-1]) {
		count++
	}
	if (!onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0] + 1, pos[1]}) ||
		puzzle[pos[0]][pos[1]] != puzzle[pos[0]+1][pos[1]]) &&
		(!onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0], pos[1] + 1}) ||
			puzzle[pos[0]][pos[1]] != puzzle[pos[0]][pos[1]+1]) {
		count++
	}
	if onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0] + 1, pos[1]}) &&
		puzzle[pos[0]][pos[1]] == puzzle[pos[0]+1][pos[1]] &&
		onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0], pos[1] + 1}) &&
		puzzle[pos[0]][pos[1]] == puzzle[pos[0]][pos[1]+1] &&
		puzzle[pos[0]][pos[1]] != puzzle[pos[0]+1][pos[1]+1] {
		count++
	}
	if onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0] - 1, pos[1]}) &&
		puzzle[pos[0]][pos[1]] == puzzle[pos[0]-1][pos[1]] &&
		onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0], pos[1] - 1}) &&
		puzzle[pos[0]][pos[1]] == puzzle[pos[0]][pos[1]-1] &&
		puzzle[pos[0]][pos[1]] != puzzle[pos[0]-1][pos[1]-1] {
		count++
	}
	if onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0] - 1, pos[1]}) &&
		puzzle[pos[0]][pos[1]] == puzzle[pos[0]-1][pos[1]] &&
		onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0], pos[1] + 1}) &&
		puzzle[pos[0]][pos[1]] == puzzle[pos[0]][pos[1]+1] &&
		puzzle[pos[0]][pos[1]] != puzzle[pos[0]-1][pos[1]+1] {
		count++
	}
	if onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0] + 1, pos[1]}) &&
		puzzle[pos[0]][pos[1]] == puzzle[pos[0]+1][pos[1]] &&
		onMap(len(puzzle), len(puzzle[0]), [2]int{pos[0], pos[1] - 1}) &&
		puzzle[pos[0]][pos[1]] == puzzle[pos[0]][pos[1]-1] &&
		puzzle[pos[0]][pos[1]] != puzzle[pos[0]+1][pos[1]-1] {
		count++
	}
	return count
}

func spread(puzzle [][]string, pos [2]int, areaPlots map[string][]int) {
	directions := [4][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}
	areaNeighbours := 0
	for _, val := range directions {
		tmpPos := [2]int{pos[0] + val[0], pos[1] + val[1]}
		if onMap(len(puzzle), len(puzzle[0]), tmpPos) {
			if puzzle[tmpPos[0]][tmpPos[1]] == puzzle[pos[0]][pos[1]][:1] {
				puzzle[tmpPos[0]][tmpPos[1]] = puzzle[pos[0]][pos[1]]
				areaNeighbours++
				spread(puzzle, tmpPos, areaPlots)
			} else if puzzle[tmpPos[0]][tmpPos[1]] == puzzle[pos[0]][pos[1]] {
				areaNeighbours++
			}
		}
	}
	areaPlots[puzzle[pos[0]][pos[1]]] = append(areaPlots[puzzle[pos[0]][pos[1]]], 4-areaNeighbours)
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
	file, _ := os.Open("2024/Day 12/input")
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
