package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type springs struct {
	springs string
	damaged []int
}

var data []springs

func main() {
	parseData()

	var sum int
	for _, arrangement := range data {
		//fmt.Println(arrangement)

		var length int = len(arrangement.springs)

		var matches int = countMatches(length, arrangement)

		//fmt.Println("Matches:", matches)
		sum += matches
	}
	fmt.Println(sum)

}

func known(arrangement springs) (string, string) {
	var value, value2 string

	for i := 0; i < len(arrangement.springs); i++ {
		if arrangement.springs[i] == '.' {
			value += "1"
		} else {
			value += "0"
		}
		if arrangement.springs[i] == '#' {
			value2 += "1"
		} else {
			value2 += "0"
		}
	}

	return value, value2
}

func countMatches(length int, arrangement springs) int {
	var max int = 1 << uint(length)

	var knownWorking, knownBroken = known(arrangement)

	var matches int
	// loop through all possible arrangements
	for i := 0; i < max; i++ {

		var stringMap = intToStringMap(i, length)

		if skipSet(stringMap, knownBroken) {
			continue
		}
		if skipNotSet(stringMap, knownWorking) {
			continue
		}

		var blocks = getBlocks(stringMap)

		if compareSlices(blocks, arrangement.damaged) {
			//fmt.Println("Found", stringMap, blocks, arrangement.damaged)
			matches++
		}
	}
	return matches
}

func intToStringMap(value int, length int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(length)+"b", value)
}

func skipSet(sourceString string, compareMask string) bool {
	for i := 0; i < len(sourceString); i++ {
		if string(compareMask[i]) == "1" {
			if string(sourceString[i]) == "?" {
				continue
			}
			if string(sourceString[i]) != "1" {
				//fmt.Println("Skipping", sourceString, "because", compareMask, "at", i, "is not 1")
				return true
			}
		}
	}
	return false
}

func skipNotSet(sourceString string, compareMask string) bool {
	for i := 0; i < len(sourceString); i++ {
		if string(compareMask[i]) == "1" {
			if string(sourceString[i]) == "?" {
				continue
			}
			if string(sourceString[i]) != "0" {
				//fmt.Println("Skipping", sourceString, "because", compareMask, "at", i, "is not 0")
				return true
			}
		}
	}
	return false
}

func compareSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func getBlocks(binary string) []int {
	// loop through the arrangement and find continuous blocks of 1
	var blocks []int = make([]int, 0)
	var blockLength int = 0
	for j := 0; j < len(binary); j++ {
		if binary[j] == '1' {
			blockLength++
		} else {
			if blockLength > 0 {
				blocks = append(blocks, blockLength)
				blockLength = 0
			}
		}
	}
	if blockLength > 0 {
		blocks = append(blocks, blockLength)
	}

	return blocks
}

func parseData() {
	lines := readInput("day_12/input.txt")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		var lineData = springs{}
		var parts = strings.Split(line, " ")
		lineData.springs = parts[0]
		lineData.damaged = make([]int, 0)
		for _, item := range strings.Split(parts[1], ",") {
			number := mustConvertToInt(item)
			lineData.damaged = append(lineData.damaged, number)
		}

		data = append(data, lineData)
	}
}

func mustConvertToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return num
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
