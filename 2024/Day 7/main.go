package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := readFile()
	results := make([]int, 0)
	values := make([][]int, 0)
	partOne := 0
	partTwo := 0
	for _, val := range lines {
		result, numbers := splitEquations(val)
		results = append(results, result)
		values = append(values, numbers)
		if check(result, numbers) {
			partOne += result
		}
		if checkTwo(result, numbers) {
			partTwo += result
		}

	}
	fmt.Println(partOne)
	fmt.Println(partTwo)
}

func calculate(result int, calcNumber int, numbers []int) bool {
	if len(numbers) > 1 {
		if calculate(result, calcNumber+numbers[0], numbers[1:]) ||
			calculate(result, calcNumber*numbers[0], numbers[1:]) ||
			calculate(result, concatNumbers(calcNumber, numbers[0]), numbers[1:]) {
			return true
		}
	} else {
		if result == calcNumber+numbers[0] ||
			result == calcNumber*numbers[0] ||
			result == concatNumbers(calcNumber, numbers[0]) {
			return true
		}
	}
	return false
}

func check(result int, numbers []int) bool {
	if len(numbers) > 1 {
		if check(result-numbers[len(numbers)-1], numbers[:len(numbers)-1]) {
			return true
		}
		if result%numbers[len(numbers)-1] == 0 {
			if check(result/numbers[len(numbers)-1], numbers[:len(numbers)-1]) {
				return true
			}
		}
	} else {
		if numbers[0] == result {
			return true
		}
	}
	return false
}

func checkTwo(result int, numbers []int) bool {
	if len(numbers) == 2 {
		if result == numbers[0]+numbers[1] ||
			result == numbers[0]*numbers[1] ||
			result == concatNumbers(numbers[0], numbers[1]) {
			return true
		}

	} else {
		if calculate(result, numbers[0]+numbers[1], numbers[2:]) ||
			calculate(result, numbers[0]*numbers[1], numbers[2:]) ||
			calculate(result, concatNumbers(numbers[0], numbers[1]), numbers[2:]) {
			return true
		}
	}
	return false
}

func concatNumbers(i int, j int) int {
	str := strconv.Itoa(i) + strconv.Itoa(j)
	num, _ := strconv.Atoi(str)
	return num
}

func splitEquations(line string) (int, []int) {
	splitLine := strings.Split(line, ":")
	result, _ := strconv.Atoi(splitLine[0])
	splitLine[1] = strings.TrimLeft(splitLine[1], " ")
	splitLine = strings.Split(splitLine[1], " ")
	numbers := make([]int, 0)
	for _, val := range splitLine {
		tmp, _ := strconv.Atoi(val)
		numbers = append(numbers, tmp)
	}
	return result, numbers
}

func readFile() []string {
	file, _ := os.Open("2024/Day 7/input")
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines
}
