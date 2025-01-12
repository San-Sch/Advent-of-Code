package day16

import (
	"Advent_of_Code_2024/utils"
	"fmt"
	"math"
	"slices"
)

type node struct {
	c   utils.Coordinate
	p   utils.Coordinate
	d   utils.Coordinate
	val int
}

func Part1() {
	data := utils.ReadFile("./day16/test1")
	puzzleMap := utils.CreatePuzzleMap(data)
	nodes, finish := prepareMapData(puzzleMap)
	vNodes := make([]node, 0)
	sortNodes(nodes)
	for nodes[0].c != finish.c {
		calcNextNodes(puzzleMap, nodes[0], nodes)
		vNodes = append(vNodes, nodes[0])
		nodes = nodes[1:]
		sortNodes(nodes)

	}
	fmt.Println(nodes[0].val)
	//fmt.Println(vNodes)
}

func findNode(nodes []node, c utils.Coordinate) int {
	return slices.IndexFunc(nodes, func(node node) bool {
		return node.c == c
	})
}

func calcNextNodes(p [][]string, n node, nodes []node) {
	val := 0
	directions := make([]utils.Coordinate, 0)
	directions = append(directions, n.d, turn(n, "CCW"), turn(n, "CW"))
	for d, v := range directions {
		if d == 0 {
			val = n.val + 1
		} else {
			val = n.val + 1001
		}
		if p[n.c.Y+v.Y][n.c.X+v.X] == "." || p[n.c.Y+v.Y][n.c.X+v.X] == "E" {
			idx := findNode(nodes, utils.Coordinate{X: n.c.X + v.X, Y: n.c.Y + v.Y})
			if idx > -1 && nodes[idx].val > val {
				nodes[idx].val = val
				nodes[idx].p = n.c
				nodes[idx].d = v
			}
		}
	}
}

func turn(n node, dir string) utils.Coordinate {
	directions := make([]utils.Coordinate, 0)
	directions = append(directions, utils.Coordinate{X: 0, Y: 1}, utils.Coordinate{X: 1, Y: 0}, utils.Coordinate{X: 0, Y: -1}, utils.Coordinate{X: -1, Y: 0})
	idx := slices.Index(directions, n.d)
	if dir == "CW" {
		if idx == 0 {
			idx = 3
		} else {
			idx--
		}
	} else if dir == "CCW" {
		if idx == 3 {
			idx = 0
		} else {
			idx++
		}
	}
	return directions[idx]
}

func sortNodes(nodes []node) {
	slices.SortFunc(nodes, func(a, b node) int {
		if a.val < b.val {
			return -1
		}
		if a.val > b.val {
			return 1
		}
		return 0
	})
}

func prepareMapData(puzzleMap [][]string) ([]node, node) {
	nodes := make([]node, 0)
	finish := node{}
	for y := 0; y < len(puzzleMap); y++ {
		for x := 0; x < len(puzzleMap[y]); x++ {
			if puzzleMap[y][x] == "#" {
				continue
			}
			if puzzleMap[y][x] == "." {
				eliminateDeadEnds(puzzleMap, utils.Coordinate{X: x, Y: y})
				nodes = append(nodes, node{c: utils.Coordinate{X: x, Y: y}, val: math.MaxInt})
			}
			if puzzleMap[y][x] == "S" {
				nodes = append(nodes, node{c: utils.Coordinate{X: x, Y: y}, val: 0, p: utils.Coordinate{X: -1, Y: -1}, d: utils.Coordinate{X: 1, Y: 0}})
			}
			if puzzleMap[y][x] == "E" {
				nodes = append(nodes, node{c: utils.Coordinate{X: x, Y: y}, val: math.MaxInt})
				finish = nodes[len(nodes)-1]
			}
		}
	}
	return nodes, finish
}

func eliminateDeadEnds(puzzleMap [][]string, coordinate utils.Coordinate) {
	walls, empty := surroundingTiles(puzzleMap, coordinate, "#")
	if walls >= 3 {
		puzzleMap[coordinate.Y][coordinate.X] = "#"
		if empty.Y != 0 || empty.X != 0 {
			eliminateDeadEnds(puzzleMap, empty)
		}
	}
}

func surroundingTiles(puzzleMap [][]string, coordinate utils.Coordinate, tile string) (int, utils.Coordinate) {
	coordinates := make([]utils.Coordinate, 0)
	coordinates = append(coordinates,
		utils.Coordinate{X: coordinate.X + 1, Y: coordinate.Y},
		utils.Coordinate{X: coordinate.X - 1, Y: coordinate.Y},
		utils.Coordinate{X: coordinate.X, Y: coordinate.Y + 1},
		utils.Coordinate{X: coordinate.X, Y: coordinate.Y - 1})
	cntTiles := 0
	var emptyPos utils.Coordinate
	for _, val := range coordinates {
		if puzzleMap[val.Y][val.X] == tile {
			cntTiles++
		} else if puzzleMap[val.Y][val.X] != "S" && puzzleMap[val.Y][val.X] != "E" {
			emptyPos = val
		}
	}
	return cntTiles, emptyPos
}
