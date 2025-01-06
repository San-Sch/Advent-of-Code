package day03

import (
	"Advent_of_Code_2024/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func Part1() {
	str := utils.ReadFile("./day03/input")
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	re2 := regexp.MustCompile(`\d{1,3}`)
	mul := re.FindAllString(str, -1)
	var result int
	for i := 0; i < len(mul); i++ {
		numbers := re2.FindAllString(mul[i], -1)
		num1, _ := strconv.Atoi(numbers[0])
		num2, _ := strconv.Atoi(numbers[1])
		result += num1 * num2
	}
	fmt.Println(result)
}

func Part2() {
	str := utils.ReadFile("./day03/input")
	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	re2 := regexp.MustCompile(`\d{1,3}`)
	mul := re.FindAllString(str, -1)
	result := 0
	runLoop := true
	for runLoop {
		indexDont := strings.Index(str, "don't()")
		if indexDont == -1 {
			indexDont = len(str) - 1
			runLoop = false
		}
		tmpStr := str[:indexDont]
		str = str[indexDont:]
		mul = re.FindAllString(tmpStr, -1)
		for i := 0; i < len(mul); i++ {
			numbers := re2.FindAllString(mul[i], -1)
			num1, _ := strconv.Atoi(numbers[0])
			num2, _ := strconv.Atoi(numbers[1])
			result += num1 * num2
		}
		indexDo := strings.Index(str, "do()")
		if indexDo == -1 {
			runLoop = false
		} else {
			str = str[indexDo:]
		}
	}
	fmt.Println(result)
}
