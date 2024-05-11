package main

import (
	"AoC2023/framework"
	"fmt"
	"regexp"
	"strconv"
)

var numberRegex = regexp.MustCompile(`-?\d+`)

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	lines := framework.ReadInput("input.txt")

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

		result := predictPreviousNumber(interpolateArray)

		sum += result

	}
	return sum

}

func predictPreviousNumber(interpolateArr [][]int) int {

	var newFirstArray = []int{0}
	var diff int
	for i := len(interpolateArr) - 1; i > 0; i-- {
		diff = interpolateArr[i-1][0] - diff
		newFirstArray = append(newFirstArray, diff)
	}
	return diff
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
