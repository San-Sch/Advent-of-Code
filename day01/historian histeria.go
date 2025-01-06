package day01

import (
	"Advent_of_Code_2024/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func Part1() {
	data := utils.ReadFile("./day01/input")
	leftList, rightList := makeLists(data)

	var differenz int = 0

	for i := 0; i < len(leftList); i++ {
		differenz += utils.Diff(leftList[i], rightList[i])
	}
	fmt.Println(differenz)
}

func Part2() {
	data := utils.ReadFile("./day01/input")
	leftList, rightList := makeLists(data)

	var similarity int = 0

	for i := 0; i < len(leftList); i++ {
		index := slices.Index(rightList, leftList[i])
		count := 0
		if index != -1 {
			count = 1
			index++
			for rightList[index] == leftList[i] {
				count++
				index++
			}
		}
		similarity += leftList[i] * count
	}
	fmt.Println(similarity)
}

func makeLists(data string) (leftList []int, rightList []int) {
	list := strings.Split(data, "\n")

	for _, v := range list {
		tmp := strings.Split(v, "   ")
		left, _ := strconv.Atoi(tmp[0])
		right, _ := strconv.Atoi(tmp[1])
		leftList = append(leftList, left)
		rightList = append(rightList, right)
	}
	slices.Sort(leftList)
	slices.Sort(rightList)
	return
}
