package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

func main() {
	data := readFile()
	puzzleMap, movements, startPos := initData(data)
	position := startPos
	for _, val := range movements {
		position = move(puzzleMap, position, string(val))
	}
	fmt.Println(calcCoordinateSum(puzzleMap))
	puzzleMap, movements, _ = initData(data)
	puzzleMap, startPos = resizePuzzle(puzzleMap)
	position = startPos
	for _, val := range movements {
		position = move(puzzleMap, position, string(val))
	}
	printPuzzle(puzzleMap, position)
	fmt.Println(calcCoordinateSum(puzzleMap))
}

func resizePuzzle(puzzle [][]string) ([][]string, Pos) {
	newMap := make([][]string, 0)
	startPosition := Pos{0, 0}
	for i := 0; i < len(puzzle); i++ {
		newMap = append(newMap, make([]string, 0))
		for j := 0; j < len(puzzle[i]); j++ {
			switch puzzle[i][j] {
			case "#":
				newMap[i] = append(newMap[i], "#")
				newMap[i] = append(newMap[i], "#")
			case "O":
				newMap[i] = append(newMap[i], "[")
				newMap[i] = append(newMap[i], "]")
			case ".":
				newMap[i] = append(newMap[i], ".")
				newMap[i] = append(newMap[i], ".")
			case "@":
				startPosition = Pos{x: len(newMap[i]), y: i}
				newMap[i] = append(newMap[i], "@")
				newMap[i] = append(newMap[i], ".")
			}
		}
	}
	return newMap, startPosition
}

func calcCoordinateSum(puzzle [][]string) int {
	sum := 0
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {
			if puzzle[i][j] == "O" || puzzle[i][j] == "[" {
				sum += (100 * i) + j
			}
		}
	}
	return sum
}

func calcNewPos(pos Pos, direction string) Pos {
	var newPos Pos
	switch direction {
	case "<":
		newPos = Pos{x: pos.x - 1, y: pos.y}
	case ">":
		newPos = Pos{x: pos.x + 1, y: pos.y}
	case "^":
		newPos = Pos{x: pos.x, y: pos.y - 1}
	case "v":
		newPos = Pos{x: pos.x, y: pos.y + 1}
	}
	return newPos
}

func move(puzzleMap [][]string, position Pos, direction string) Pos {
	newPos := calcNewPos(position, direction)

	switch puzzleMap[newPos.y][newPos.x] {

	case ".":
		puzzleMap[position.y][position.x] = "."
		//puzzleMap[newPos.y][newPos.x] = "@"
		return newPos
	case "O":
		if moveBox(puzzleMap, newPos, direction) {
			puzzleMap[position.y][position.x] = "."
			//puzzleMap[newPos.y][newPos.x] = "@"
			return newPos
		} else {
			return position
		}
	case "[", "]":

		if moveBigBox(puzzleMap, newPos, direction) {
			puzzleMap[position.y][position.x] = "."
			//puzzleMap[newPos.y][newPos.x] = "@"
			return newPos
		} else {
			return position
		}

	case "#":
		return position
	}
	return Pos{x: -1, y: -1}
}

func getBoxPositions(puzzleMap [][]string, pos Pos) (Pos, Pos) {
	var posLeft, posRight Pos

	if puzzleMap[pos.y][pos.x] == "[" {
		posLeft = pos
		posRight = Pos{pos.x + 1, pos.y}
	} else if puzzleMap[pos.y][pos.x] == "]" {
		posRight = pos
		posLeft = Pos{pos.x - 1, pos.y}
	}
	return posLeft, posRight
}

func movable(puzzleMap [][]string, pos Pos, direction string) bool {

	posLeft, posRight := getBoxPositions(puzzleMap, pos)

	posLeft, posRight = calcNewPos(posLeft, direction), calcNewPos(posRight, direction)

	if puzzleMap[posRight.y][posRight.x] == "#" || puzzleMap[posLeft.y][posLeft.x] == "#" {
		return false
	}
	if puzzleMap[posLeft.y][posLeft.x] == "[" || puzzleMap[posLeft.y][posLeft.x] == "]" {
		if !movable(puzzleMap, posLeft, direction) {
			return false
		}
	}
	if puzzleMap[posRight.y][posRight.x] == "[" || puzzleMap[posRight.y][posRight.x] == "]" {
		if !movable(puzzleMap, posRight, direction) {
			return false
		}
	}
	return true
}

func moveBigBox(puzzleMap [][]string, pos Pos, direction string) bool {
	newPos := calcNewPos(pos, direction)

	switch direction {
	case "<":
		if puzzleMap[newPos.y][newPos.x-1] == "." {
			puzzleMap[pos.y][pos.x] = "."
			puzzleMap[newPos.y][newPos.x] = "]"
			puzzleMap[newPos.y][newPos.x-1] = "["
			return true
		} else if puzzleMap[newPos.y][newPos.x-1] == "]" {
			if moveBigBox(puzzleMap, calcNewPos(newPos, direction), direction) {
				puzzleMap[pos.y][pos.x] = "."
				puzzleMap[newPos.y][newPos.x] = "]"
				puzzleMap[newPos.y][newPos.x-1] = "["
				return true
			}
		}
	case ">":
		if puzzleMap[newPos.y][newPos.x+1] == "." {

			puzzleMap[pos.y][pos.x] = "."
			puzzleMap[newPos.y][newPos.x] = "["
			puzzleMap[newPos.y][newPos.x+1] = "]"
			return true
		} else if puzzleMap[newPos.y][newPos.x+1] == "[" {
			if moveBigBox(puzzleMap, calcNewPos(newPos, direction), direction) {
				puzzleMap[pos.y][pos.x] = "."
				puzzleMap[newPos.y][newPos.x] = "["
				puzzleMap[newPos.y][newPos.x+1] = "]"
				return true
			}
		}
	case "^", "v":
		if movable(puzzleMap, pos, direction) {
			left, right := getBoxPositions(puzzleMap, pos)
			newLeft, newRight := calcNewPos(left, direction), calcNewPos(right, direction)
			if puzzleMap[newLeft.y][newLeft.x] == "[" || puzzleMap[newLeft.y][newLeft.x] == "]" {
				moveBigBox(puzzleMap, newLeft, direction)
			}
			if puzzleMap[newRight.y][newRight.x] == "[" || puzzleMap[newRight.y][newRight.x] == "]" {
				moveBigBox(puzzleMap, newRight, direction)
			}
			puzzleMap[left.y][left.x], puzzleMap[right.y][right.x] = ".", "."
			puzzleMap[newLeft.y][newLeft.x], puzzleMap[newRight.y][newRight.x] = "[", "]"
			return true
		}
	}
	return false
}

func moveBox(puzzleMap [][]string, position Pos, direction string) bool {
	newPos := calcNewPos(position, direction)

	switch puzzleMap[newPos.y][newPos.x] {

	case ".":
		puzzleMap[position.y][position.x] = "."
		puzzleMap[newPos.y][newPos.x] = "O"
		return true
	case "O":
		if moveBox(puzzleMap, newPos, direction) {
			puzzleMap[position.y][position.x] = "."
			puzzleMap[newPos.y][newPos.x] = "O"
			return true
		} else {
			return false
		}

	case "#":
		return false
	}
	return false
}

func initData(data string) ([][]string, string, Pos) {
	scanner := bufio.NewScanner(strings.NewReader(data))
	readMap := true
	puzzleMap := make([][]string, 0)
	startPos := Pos{0, 0}
	var movements string
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			readMap = false
		}
		if readMap {
			puzzleMap = append(puzzleMap, make([]string, 0))
			for j := 0; j < len(line); j++ {
				if line[j:j+1] == "@" {
					startPos.y = len(puzzleMap) - 1
					startPos.x = j
				}
				puzzleMap[len(puzzleMap)-1] = append(puzzleMap[len(puzzleMap)-1], line[j:j+1])
			}
		} else {
			movements += line
		}
	}
	return puzzleMap, movements, startPos
}

func printPuzzle(puzzle [][]string, pos Pos) {
	for i := 0; i < len(puzzle); i++ {
		for j := 0; j < len(puzzle[i]); j++ {

			if pos.x == j && pos.y == i {
				fmt.Print("@")
			} else {
				fmt.Print(puzzle[i][j])
			}
		}
		fmt.Print("\n")
	}
	fmt.Println()
}

func readFile() string {
	data, _ := os.ReadFile("2024/Day 15/input")
	return string(data)
}
