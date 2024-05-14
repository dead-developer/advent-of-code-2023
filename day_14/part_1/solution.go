package main

import (
	"AoC2023/framework"
	"fmt"
	"strings"
)

var maps [][]string

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	lines := loadData()

	currentLine := 0
	stop := false

	for !stop {

		dropped := dropBoulders(lines, currentLine)
		if dropped > 0 && currentLine > 0 {
			currentLine--
		} else {
			currentLine++
		}

		if currentLine == len(lines)-1 {
			stop = true
		}
	}
	//displayArray(lines)
	return calculateWeight(lines)

}

func dropBoulders(lines [][]string, line int) int {
	dropped := 0
	var width = len(lines[0])
	for x := 0; x < width; x++ {
		if lines[line][x] == "." {
			if lines[line+1][x] == "O" {
				dropped++
				lines[line][x] = "O"
				lines[line+1][x] = "."
			}
		}
	}
	return dropped
}

func loadData() [][]string {
	lines := framework.ReadInput("input.txt")
	var returnValue [][]string
	for _, line := range lines {
		returnValue = append(returnValue, strings.Split(line, ""))
	}

	return returnValue
}

func displayArray(array [][]string) {
	for _, line := range array {
		fmt.Println(line)
	}
}

func calculateWeight(array [][]string) int {
	weight := 0
	for lineNum, line := range array {
		for _, char := range line {
			if char == "O" {
				weight += len(array) - lineNum
			}
		}
	}
	return weight
}
