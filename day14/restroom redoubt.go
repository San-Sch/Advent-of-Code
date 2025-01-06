package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Robot struct {
	pos Pos
	v   Pos
}

type Map struct {
	height int
	width  int
}

func main() {
	data := readFile()
	robots := initRobots(data)
	var restroom Map
	restroom.height = 103
	restroom.width = 101
	//restroom.height = 7
	//restroom.width = 11
	for i := 0; i < 11000; i++ {
		moveRobots(restroom, robots)
		if safetyFactor(restroom, robots) < 60000000 {

			fmt.Println("After: ", i+1, " Safety Factor: ", safetyFactor(restroom, robots))
			printMap(restroom, robots)

		}
	}
	fmt.Println(safetyFactor(restroom, robots)) //220437504 too low
}

func safetyFactor(room Map, robots []Robot) int {
	slices.SortFunc(robots, func(a, b Robot) int {
		if a.pos.y < b.pos.y {
			return -1
		}
		if a.pos.y > b.pos.y {
			return 1
		}
		if a.pos.x < b.pos.x {
			return -1
		}
		if a.pos.x > b.pos.x {
			return 1
		}
		return 0

	})
	factor := factorQuadrant(0, 0, (room.height-1)/2-1, (room.width-1)/2-1, robots)
	factor *= factorQuadrant((room.height-1)/2+1, 0, room.height-1, (room.width-1)/2-1, robots)
	factor *= factorQuadrant(0, (room.width-1)/2+1, (room.height-1)/2-1, room.width-1, robots)
	factor *= factorQuadrant((room.height-1)/2+1, (room.width-1)/2+1, room.height-1, room.width-1, robots)
	return factor
}

func factorQuadrant(yStart, xStart, yEnd, xEnd int, robots []Robot) int {
	cnt := 0
	for y := yStart; y <= yEnd; y++ {
		for x := xStart; x <= xEnd; x++ {
			idx := slices.IndexFunc(robots, func(robot Robot) bool { return robot.pos.x == x && robot.pos.y == y })
			if idx > -1 {
				for idx < len(robots) && robots[idx].pos.y == y && robots[idx].pos.x == x {
					idx++
					cnt++
				}
			}
		}
	}
	return cnt
}

func printMap(room Map, robots []Robot) {
	slices.SortFunc(robots, func(a, b Robot) int {
		if a.pos.y < b.pos.y {
			return -1
		}
		if a.pos.y > b.pos.y {
			return 1
		}
		if a.pos.x < b.pos.x {
			return -1
		}
		if a.pos.x > b.pos.x {
			return 1
		}
		return 0

	})
	fmt.Println()
	for y := 0; y < room.height; y++ {
		for x := 0; x < room.width; x++ {
			cnt := 0
			idx := slices.IndexFunc(robots, func(robot Robot) bool { return robot.pos.x == x && robot.pos.y == y })
			if idx > -1 {
				for idx < len(robots) && robots[idx].pos.y == y && robots[idx].pos.x == x {
					idx++
					cnt++
				}
			}
			if cnt == 0 {
				fmt.Print(".")
			} else {
				fmt.Print(cnt)
			}
		}
		fmt.Println()
	}
	fmt.Println()
}

func moveRobots(restroom Map, robots []Robot) {
	for idx, robot := range robots {
		newPos := robot.pos
		newPos.x += robot.v.x
		newPos.y += robot.v.y
		if onMap(restroom, newPos) {
			robots[idx].pos = newPos
		} else {
			robots[idx].pos = teleport(restroom, newPos)
		}
	}
}

func teleport(restroom Map, pos Pos) Pos {
	newPos := pos
	if pos.y < 0 {
		newPos.y += restroom.height
	} else if pos.y >= restroom.height {
		newPos.y -= restroom.height
	}
	if pos.x < 0 {
		newPos.x += restroom.width
	} else if pos.x >= restroom.width {
		newPos.x -= restroom.width
	}
	return newPos
}

func onMap(room Map, position Pos) bool {
	if position.y < 0 || position.y >= room.height || position.x < 0 || position.x >= room.width {
		return false
	}
	return true
}

func initRobots(data string) []Robot {
	robots := make([]Robot, 0)
	scanner := bufio.NewScanner(strings.NewReader(data))
	for scanner.Scan() {
		robots = append(robots, Robot{})
		_, err := fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &robots[len(robots)-1].pos.x, &robots[len(robots)-1].pos.y, &robots[len(robots)-1].v.x, &robots[len(robots)-1].v.y)
		if err != nil {
			panic(err)
		}
	}
	return robots
}

func readFile() string {
	data, _ := os.ReadFile("2024/Day 14/input")
	return string(data)
}
