package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// RGB
var totalCubes = [3]int{12, 13, 14}

func main() {
	lines := readFile("day_2/input.txt")

	total := 0
	for _, line := range lines {
		gameNumber := parseLine(line)
		total += gameNumber

	}
	fmt.Println(total)
}

func readFile(name string) []string {
	content, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []string{}
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func parseLine(line string) int {
	words := strings.Split(line, ":")
	lastPart := words[1]
	handfuls := strings.Split(lastPart, ";")

	for _, handful := range handfuls {
		cubes := strings.Split(handful, ",")
		handfulTotal := [3]int{0, 0, 0}
		for _, cube := range cubes {
			handfulTotal = addArrays(handfulTotal, parseColor(cube))
		}
		if !isPossible(handfulTotal, totalCubes) {
			return 0
		}

	}
	return getGameNumber(line)
}

func getGameNumber(line string) int {
	words := strings.Split(line, ":")
	firstPart := words[0]
	firstPart = strings.ReplaceAll(firstPart, "Game ", "")
	gameNumber, _ := strconv.Atoi(firstPart)
	return gameNumber
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

func isPossible(a, b [3]int) bool {
	if (a[0] > b[0]) || (a[1] > b[1]) || (a[2] > b[2]) {
		return false
	}
	return true
}
