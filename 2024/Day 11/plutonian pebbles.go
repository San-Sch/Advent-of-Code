package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	result := 0
	pebbles := getUniquePebbles(strings.Split(readFile(), " "))
	for i := 0; i < 75; i++ {
		pebbles = blink(pebbles)
	}
	for _, val := range pebbles {
		result += val
	}
	fmt.Println(result)
}

func getUniquePebbles(pebbles []string) map[string]int {
	slices.Sort(pebbles)
	pebblesMap := make(map[string]int)
	pebble := pebbles[0]
	for _, s := range pebbles {
		if s == pebble {
			pebblesMap[pebble]++
		} else {
			pebble = s
			pebblesMap[pebble] = 1
		}
	}
	return pebblesMap
}

func blink(pebbles map[string]int) map[string]int {
	pebblesTmp := make([]map[string]int, 0)
	for idx, val := range pebbles {
		if idx == "0" {
			pebblesTmp = append(pebblesTmp, rule0(val))
		} else if len(idx) > 0 && len(idx)%2 == 0 {
			pebblesTmp = append(pebblesTmp, ruleEven(idx, val))
		} else {
			pebblesTmp = append(pebblesTmp, rule2024(idx, val))
		}
	}
	pebblesNew := make(map[string]int)
	for _, val := range pebblesTmp {
		for idx, count := range val {
			pebblesNew[idx] += count
		}
	}
	return pebblesNew
}

func rule0(val int) map[string]int {
	pebbles := make(map[string]int)
	pebbles["1"] = val
	return pebbles
}

func ruleEven(pebble string, val int) map[string]int {
	pebbles := make(map[string]int)
	tmpStr := pebble[len(pebble)/2:]
	tmpStr = strings.TrimLeft(tmpStr, "0")
	if len(tmpStr) == 0 {
		tmpStr = "0"
	}
	pebbles[tmpStr] = val
	pebbles[pebble[:len(pebble)/2]] += val
	return pebbles
}

func rule2024(pebble string, val int) map[string]int {
	pebbles := make(map[string]int)
	pebbleNum, _ := strconv.Atoi(pebble)
	pebbleNum *= 2024
	pebbles[strconv.Itoa(pebbleNum)] = val
	return pebbles
}

func readFile() string {
	data, _ := os.ReadFile("2024/Day 11/input")
	return string(data)
}
