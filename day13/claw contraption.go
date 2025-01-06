package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Pos struct {
	x int
	y int
}

type Contraption struct {
	a     Pos
	b     Pos
	prize Pos
}

const TokenCostA = 3

func main() {
	result := 0
	data := readFile()
	contraptions := makeContraptions(data)
	for _, contraption := range contraptions {
		result += calculate(contraption)
	}
	fmt.Print(result)
}

func calculate(contraption Contraption) int {
	d := contraption.a.x*contraption.b.y - contraption.a.y*contraption.b.x

	dx := contraption.prize.x*contraption.b.y - contraption.prize.y*contraption.b.x
	dy := contraption.a.x*contraption.prize.y - contraption.a.y*contraption.prize.x
	a := dx / d
	b := dy / d
	if contraption.a.x*a+contraption.b.x*b == contraption.prize.x &&
		contraption.a.y*a+contraption.b.y*b == contraption.prize.y {
		return a*TokenCostA + b
	}
	return 0
}

func makeContraptions(data string) []Contraption {
	contraptions := make([]Contraption, 0)
	scanner := bufio.NewScanner(strings.NewReader(data))
	contraptions = append(contraptions, Contraption{})
	for scanner.Scan() {
		str := scanner.Text()

		if strings.Contains(str, "A") {
			fmt.Sscanf(str, "Button A: X+%d, Y+%d", &contraptions[len(contraptions)-1].a.x, &contraptions[len(contraptions)-1].a.y)
		}
		if strings.Contains(str, "B") {
			fmt.Sscanf(str, "Button B: X+%d, Y+%d", &contraptions[len(contraptions)-1].b.x, &contraptions[len(contraptions)-1].b.y)
		}
		if strings.Contains(str, "Prize") {
			fmt.Sscanf(str, "Prize: X=%d, Y=%d", &contraptions[len(contraptions)-1].prize.x, &contraptions[len(contraptions)-1].prize.y)
			contraptions[len(contraptions)-1].prize.x += 10000000000000
			contraptions[len(contraptions)-1].prize.y += 10000000000000
		}

		if str == "" {
			contraptions = append(contraptions, Contraption{})
		}
	}
	return contraptions
}

func readFile() string {
	data, _ := os.ReadFile("2024/Day 13/input")
	return string(data)
}
