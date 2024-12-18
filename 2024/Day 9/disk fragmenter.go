package main

import (
	"fmt"
	"os"
	"slices"
	"strconv"
)

func main() {
	diskMap := readFile()
	blockFile := makeBlockFile(diskMap)
	orderedFile := orderBlocks(slices.Clone(blockFile))
	orderedFile2 := orderBlocks2(slices.Clone(blockFile))
	fmt.Println(checkSum(orderedFile))
	fmt.Println(checkSum(orderedFile2))
}

func checkSum(file []string) int {
	sum := 0
	for i := 0; i < len(file); i++ {
		if file[i] != "." {
			num, _ := strconv.Atoi(file[i])
			sum += i * num
		}
	}
	return sum
}

func orderBlocks2(blockFile []string) []string {
	lastPos := len(blockFile) - 1
	for blockFile[lastPos] == "." {
		lastPos--
	}
	fileIndex, _ := strconv.Atoi(blockFile[lastPos])
	startPos := 0
	index := -1
	for ; fileIndex > 0; fileIndex-- {
		startPos = slices.Index(blockFile, strconv.Itoa(fileIndex))
		lastPos = startPos + 1
		for lastPos < len(blockFile) && blockFile[lastPos] == strconv.Itoa(fileIndex) {
			lastPos++
		}
		index = freeBlockIndex(lastPos-startPos, blockFile)
		if index < startPos && index > -1 {
			for ; startPos < lastPos; startPos++ {
				blockFile[index], blockFile[startPos] = blockFile[startPos], blockFile[index]
				index++
			}
		}
	}
	return blockFile
}

func freeBlockIndex(blocks int, file []string) int {
	index := slices.Index(file, ".")
	freeBlocks := 0
	for ; index < len(file); index++ {
		if file[index] == "." {
			freeBlocks += 1
		} else {
			freeBlocks = 0
		}
		if freeBlocks == blocks {
			break
		}
	}
	if freeBlocks != blocks {
		return -1
	}

	return index - freeBlocks + 1
}

func orderBlocks(blockFile []string) []string {
	lastPos := len(blockFile) - 1
	for blockFile[lastPos] == "." {
		lastPos--
	}

	index := slices.Index(blockFile, ".")

	for index < lastPos && index > -1 {
		blockFile[index], blockFile[lastPos] = blockFile[lastPos], blockFile[index]
		for blockFile[lastPos] == "." {
			lastPos--
		}
		index = slices.Index(blockFile, ".")
	}
	return blockFile
}

func makeBlockFile(diskMap string) []string {
	blockFile := make([]string, 0)
	for idx, val := range diskMap {
		tmpNum, _ := strconv.Atoi(string(val))
		for i := 0; i < tmpNum; i++ {
			if idx%2 == 0 {
				blockFile = append(blockFile, strconv.Itoa(idx/2))
			} else {
				blockFile = append(blockFile, ".")
			}
		}

	}
	return blockFile
}

func readFile() string {
	data, _ := os.ReadFile("2024/Day 9/input")
	return string(data)

}
