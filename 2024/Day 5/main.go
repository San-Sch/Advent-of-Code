package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	rules, updates := readFile()
	result := 0
	result2 := 0
	for i := 0; i < len(updates); i++ {
		if checkRules(updates[i], rules) {
			page, _ := strconv.Atoi(updates[i][(len(updates[i])-1)/2])
			result += page
		} else {
			page, _ := strconv.Atoi(sortUpdate(updates[i], rules)[(len(updates[i])-1)/2])
			result2 += page
		}
	}
	fmt.Println(result)
	fmt.Println(result2) //4400 too low
}

func sortUpdate(update []string, rules [][]string) []string {
	slices.SortFunc(update, func(a, b string) int {
		for _, rule := range rules {
			if slices.Contains(rule, a) && slices.Contains(rule, b) {
				if slices.Index(rule, a) < slices.Index(rule, b) {
					return -1
				} else {
					return 1
				}
			}
		}
		return 0
	})

	return update
}

func checkRule(update []string, rule []string) bool {
	if slices.Index(update, rule[0]) > slices.Index(update, rule[1]) {
		return false
	}
	return true
}

func checkRules(update []string, rules [][]string) bool {
	for _, rule := range rules {
		if slices.Contains(update, rule[0]) && slices.Contains(update, rule[1]) {
			if !checkRule(update, rule) {
				return false
			}
		}
	}
	return true
}

func readFile() ([][]string, [][]string) {
	rules, updates := make([][]string, 0), make([][]string, 0)
	file, _ := os.Open("2024/Day 5/input")
	scanner := bufio.NewScanner(file)
	for scanner.Scan() && strings.Contains(scanner.Text(), "|") {
		rules = append(rules, strings.Split(scanner.Text(), "|"))
	}

	for scanner.Scan() {
		if strings.Contains(scanner.Text(), ",") {
			updates = append(updates, strings.Split(scanner.Text(), ","))
		}
	}
	return rules, updates
}
