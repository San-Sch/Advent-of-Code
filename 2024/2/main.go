package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	lines := ReadFile()
	var safeReports int = 0
	for i := 0; i < len(lines); i++ {
		report := strings.Split(lines[i], " ")
		safe := checkReport(report)
		if safe {
			safeReports++
		} else {
			fmt.Println(report)
		}

	}
	fmt.Println(safeReports)
}

func checkReport(report []string) bool {
	var lastLevel, level, richtung int
	for i := 0; i < len(report); i++ {
		level, _ = strconv.Atoi(report[i])
		if lastLevel == level {
			return false
		} else if i == 0 {
			lastLevel = level
		} else if diff(lastLevel, level) < 0 || diff(lastLevel, level) > 3 {
			return false
		} else {
			switch true {
			case lastLevel > level && richtung == 0:
				richtung = 1
			case lastLevel < level && richtung == 0:
				richtung = 2
			case lastLevel > level && richtung == 2, lastLevel < level && richtung == 1:
				return false

			}
		}
		lastLevel = level
	}
	return true
}

func diff(a, b int) int {
	if a < b {
		return b - a
	}
	return a - b
}

func ReadFile() []string {
	file, _ := os.Open("2024/2/input")
	scanner := bufio.NewScanner(file)
	lines := make([]string, 0)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}
