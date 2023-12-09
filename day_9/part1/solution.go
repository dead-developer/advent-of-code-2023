package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var numberRegex = regexp.MustCompile(`-?\d+`)

func main() {
	lines := readInput("day_9/input.txt")

	var sum int
	for _, line := range lines {
		lineNumbers := parseInput(numberRegex.FindAllString(line, -1))

		var finish bool

		var interpolateArray = [][]int{lineNumbers}

		for !finish {
			var diff []int
			diff, finish = getDifference(interpolateArray[len(interpolateArray)-1])
			if len(diff) == 0 {
				break
			}
			interpolateArray = append(interpolateArray, diff)
		}

		//for _, items := range interpolateArray {
		//	fmt.Println(items)
		//}

		result := predictNextNumber(interpolateArray)
		//fmt.Println(result)
		//fmt.Println("--------------------")
		sum += result

	}

	fmt.Println(sum)
}

func predictNextNumber(interpolateArr [][]int) int {
	var number int
	for _, items := range interpolateArr {
		number += items[len(items)-1]
	}
	return number
}

func getDifference(numbers []int) ([]int, bool) {
	// build new array of numbers with the difference between each number
	// and the next number
	if len(numbers) == 1 {
		return []int{}, true
	}

	result := make([]int, len(numbers)-1)
	finish := true
	for i := 0; i < len(numbers)-1; i++ {
		result[i] = numbers[i+1] - numbers[i]
		if result[i] != 0 {
			finish = false
		}
	}
	return result, finish
}

func parseInput(lines []string) []int {
	numbers := make([]int, len(lines))
	for i, line := range lines {
		numbers[i] = mustConvertToInt(line)
	}
	return numbers
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
