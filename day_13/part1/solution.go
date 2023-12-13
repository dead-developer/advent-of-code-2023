package main

import (
	"fmt"
	"os"
	"strings"
)

var maps [][]string

func main() {
	parseData()

	var sum int
	for index, pattern := range maps {

		horizontal, vertical := findMirror(pattern)

		if horizontal < 1 && vertical < 1 {
			panic("No mirror found" + string(rune(index+1)))
		}

		//fmt.Println("Found mirror at", index+1, "horizontal:", horizontal, "vertical:", vertical)

		sum += horizontal * 100
		sum += vertical
	}

	fmt.Println("Sum:", sum)

}

func findMirror(pattern []string) (int, int) {
	var vhit, hhit int

	hhit = getMirrorRow(pattern)
	if hhit > 0 {
		return hhit, 0
	}

	// check vertical
	pattern = rotatePattern(pattern)
	vhit = getMirrorRow(pattern)

	return hhit, vhit
}

func getMirrorRow(pattern []string) int {
	for i := 1; i < len(pattern); i++ {
		size := min(i, len(pattern)-i)

		//fmt.Println("Scanning row", i, "size", size)

		//fmt.Println("i:", i, i-size, i+size)

		if compareArrays(pattern[i-size:i], pattern[i:i+size]) {
			//fmt.Println("Found mirror at", i)
			//fmt.Println(pattern[i-size:i], pattern[i:i+size])
			return i
		}
	}
	return 0
}

func compareArrays(a []string, b []string) bool {

	//fmt.Println(a)
	//fmt.Println(b)
	for i := range a {

		if a[i] != b[len(a)-1-i] {
			//fmt.Println("Not equal", i, len(a)-1-i)
			//fmt.Println(a[i], b[len(a)-1-i])
			return false
		}
	}
	return true
}

func rotatePattern(oldPattern []string) []string {
	// rotate clockwise 90 degrees
	var newPattern []string = make([]string, len(oldPattern[0]))

	for y := len(oldPattern) - 1; y >= 0; y-- {
		for x := 0; x < len(oldPattern[0]); x++ {
			newPattern[x] += string(oldPattern[y][x])
		}
	}

	return newPattern
}

func findDoubleRow(pattern []string, start int) int {
	var past string
	for i := start; i < len(pattern); i++ {
		if pattern[i] == past {
			return i
		}
		past = pattern[i]
	}
	return -1
}

func parseData() {
	lines := readInput("day_13/input.txt")
	var pattern []string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			maps = append(maps, pattern)
			pattern = []string{}
			continue
		}
		pattern = append(pattern, line)
	}
	maps = append(maps, pattern)
}

func readInput(name string) []string {
	content, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []string{}
	}
	lines := strings.Split(string(content), "\n")
	return lines
}

func debugPrint(pattern []string) {
	for _, line := range pattern {
		fmt.Println(line)
	}
	fmt.Println()
}
