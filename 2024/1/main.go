package main

import (
	"fmt"
	"io"
	"os"
	"slices"
	"strconv"
)

func main() {
	leftList, rightList := readFile()
	var differenz int = 0
	slices.Sort(leftList)
	slices.Sort(rightList)

	for i := 0; i < len(leftList); i++ {
		differenz += diff(leftList[i], rightList[i])
	}
	fmt.Print(differenz)
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func readFile() ([]int, []int) {
	var left, right = make([]int, 0), make([]int, 0)
	var number int
	file, err := os.Open("input") // For read access.
	b1 := make([]byte, 5)

	for i := 0; err == nil; i++ {
		if i == 0 {
			_, err = file.Seek(0, io.SeekStart)
		} else if i%2 == 0 {
			_, err = file.Seek(1, io.SeekCurrent)
		} else {
			_, err = file.Seek(3, io.SeekCurrent)
		}
		if err != nil {
			break
		}

		_, err = file.Read(b1)
		if err != nil {
			break
		}
		number, _ = strconv.Atoi(string(b1))
		if i%2 == 0 {
			left = append(left, number)
		} else {
			right = append(right, number)
		}
	}
	err = file.Close()
	return left, right
}
