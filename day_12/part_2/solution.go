package main

import (
	"AoC2023/framework"
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

// TODO: Finish
// NOT WORKING, WIP

type springs struct {
	springs string
	damaged []int
}

var data []springs

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	readData()

	var sum int
	for _, arrangement := range data {
		//fmt.Println(arrangement)

		var matches = countMatches(arrangement)

		fmt.Println("Matches:", matches)
		sum += matches
	}
	return sum
}

func getRegexp(arrangement springs) string {
	var parts = make([]string, 0)
	for _, item := range arrangement.damaged {
		parts = append(parts, "[#]{"+strconv.Itoa(item)+"}")
	}

	return "(^[.]+|^)" + strings.Join(parts, "[.]+") + "([.]+$|$)"
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

func getAmountOfUnknowns(arrangement springs) int {
	var count int
	for _, item := range arrangement.springs {
		if item == '?' {
			count++
		}
	}
	return count
}

func maxNumber(bits int) uint {
	return uint(math.Pow(2, float64(bits))) - 1
}
func replaceQuestionMarks(src, replacements string) string {
	replacement := "."
	for _, r := range replacements {
		if r == '0' {
			replacement = "."
		} else {
			replacement = "#"
		}
		src = strings.Replace(src, "?", replacement, 1)
	}
	return src
}

func countMatches(arrangement springs) int {

	//build regexp
	rx := regexp.MustCompile(getRegexp(arrangement))

	var length = getAmountOfUnknowns(arrangement)
	var lastNumber = maxNumber(length)

	var matches int

	fmt.Println("Length:", length, "Last number:", lastNumber)
	for i := uint(0); i <= lastNumber; i++ {
		var stringMap = intToStringMap(i, length)
		var compareString = replaceQuestionMarks(arrangement.springs, stringMap)
		//fmt.Println(compareString)

		if rx.MatchString(compareString) {
			//fmt.Println("Matched", compareString)
			matches++
		}
		if (i % 1000000) == 0 {
			fmt.Println("Progress:", i, "of", lastNumber)
		}
	}
	return matches
}

func intToStringMap(value uint, length int) string {
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

func readData() {
	lines := framework.ReadInput("sample.txt")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		var lineData springs = springs{}
		var parts = strings.Split(line, " ")

		// make parts 5 times longer with the same data
		parts[0] = strings.Repeat(parts[0]+"?", 5)
		parts[0] = parts[0][:len(parts[0])-1]

		parts[1] = strings.Repeat(parts[1]+",", 5)
		parts[1] = parts[1][:len(parts[1])-1]

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
