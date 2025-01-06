package day02

import (
	"Advent_of_Code_2024/utils"
	"fmt"
	"strconv"
	"strings"
)

func Part1and2() {
	data := utils.ReadFile("./day02/input")
	lines := strings.Split(data, "\n")
	var safeReports int = 0
	var safeReports2 int = 0
	var safe bool
	var index1, index2 int
	for i := 0; i < len(lines); i++ {
		report := strings.Split(lines[i], " ")

		tmpReport := make([]string, len(report))
		tmpReport2 := make([]string, len(report))
		safe, index1, index2 = checkReport(report)
		if safe {
			safeReports++
			safeReports2++
		} else {
			copy(tmpReport, report)
			tmpReport = removeIndex(tmpReport, index1)
			safe, _, _ = checkReport(tmpReport)
			if safe {
				safeReports2++
			} else {
				copy(tmpReport2, report)
				tmpReport2 = removeIndex(tmpReport2, index2)
				safe, _, _ = checkReport(tmpReport2)
				if safe {
					safeReports2++
				}
			}
		}

	}
	fmt.Println(safeReports)
	fmt.Println(safeReports2)
}

func removeIndex(s []string, index int) []string {
	if len(s)-1 == index {
		return append(s[:index])
	} else {
		return append(s[:index], s[index+1:]...)
	}
}

func checkReport(report []string) (bool, int, int) {
	var lastLevel, level, richtung int
	for i := 0; i < len(report); i++ {
		level, _ = strconv.Atoi(report[i])
		if lastLevel == level {
			return false, i - 1, i
		} else if i == 0 {
			lastLevel = level
		} else if utils.Diff(lastLevel, level) < 0 || utils.Diff(lastLevel, level) > 3 {
			return false, i - 1, i
		} else {
			switch true {
			case lastLevel > level && richtung == 0:
				richtung = 1
			case lastLevel < level && richtung == 0:
				richtung = 2
			case lastLevel > level && richtung == 2, lastLevel < level && richtung == 1:
				return false, i - 2, i

			}
		}
		lastLevel = level
	}
	return true, 0, 0
}
