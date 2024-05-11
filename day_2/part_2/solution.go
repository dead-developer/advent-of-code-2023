package main

import (
	"AoC2023/framework"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	total := solution()
	fmt.Println(total)
}

func solution() int {
	lines := framework.ReadInput("input.txt")

	total := 0
	for _, line := range lines {
		gameNumber := parseLine(line)
		total += gameNumber
	}
	return total
}

func parseLine(line string) int {
	words := strings.Split(line, ":")
	lastPart := words[1]
	handfuls := strings.Split(lastPart, ";")

	highestAmount := [3]int{0, 0, 0}

	for _, handful := range handfuls {
		cubes := strings.Split(handful, ",")
		handfulTotal := [3]int{0, 0, 0}
		for _, cube := range cubes {
			handfulTotal = addArrays(handfulTotal, parseColor(cube))
		}
		highestAmount = highestArray(handfulTotal, highestAmount)
	}
	return calculatePower(highestAmount)
}

func parseColor(str string) [3]int {
	str = strings.TrimSpace(str)
	words := strings.Split(str, " ")

	amount, _ := strconv.Atoi(words[0])

	if words[1] == "red" {
		return [3]int{amount, 0, 0}
	}
	if words[1] == "green" {
		return [3]int{0, amount, 0}
	}
	if words[1] == "blue" {
		return [3]int{0, 0, amount}
	}
	return [3]int{0, 0, 0}
}

func addArrays(a, b [3]int) [3]int {
	return [3]int{a[0] + b[0], a[1] + b[1], a[2] + b[2]}
}

func highestArray(highest, newVals [3]int) [3]int {
	if newVals[0] > highest[0] {
		highest[0] = newVals[0]
	}
	if newVals[1] > highest[1] {
		highest[1] = newVals[1]
	}
	if newVals[2] > highest[2] {
		highest[2] = newVals[2]
	}

	return highest
}

func calculatePower(vals [3]int) int {
	return vals[0] * vals[1] * vals[2]
}
