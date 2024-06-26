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
	parseData()

	var sum int
	for index, pattern := range maps {

		horizontal, vertical := findMirror(pattern)

		if horizontal < 1 && vertical < 1 {
			panic("No mirror found" + string(rune(index+1)))
		}

		sum += horizontal * 100
		sum += vertical
	}
	return sum

}

func findMirror(pattern []string) (int, int) {
	var vHit, hHit int

	hHit = getMirrorRow(pattern)
	if hHit > 0 {
		return hHit, 0
	}

	// check vertical
	pattern = rotatePattern(pattern)
	vHit = getMirrorRow(pattern)

	return hHit, vHit
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
	var newPattern = make([]string, len(oldPattern[0]))

	for y := len(oldPattern) - 1; y >= 0; y-- {
		for x := 0; x < len(oldPattern[0]); x++ {
			newPattern[x] += string(oldPattern[y][x])
		}
	}

	return newPattern
}

func parseData() {
	lines := framework.ReadInput("input.txt")
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
