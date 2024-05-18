package main

import (
	"AoC2023/framework"
	"fmt"
	"strconv"
	"strings"
)

var matchStrings = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	total := solution()
	fmt.Println("result:", total)
}

func solution() int {
	lines := framework.ReadInput("input.txt")

	total := 0
	for _, line := range lines {
		var first = getFirstMatchString(line)
		var last = getLastMatchString(line)

		num, err := strconv.Atoi(first + last)
		if err != nil {
			fmt.Println("Error converting string to int:", err)
		}
		total += num
	}
	return total
}

func getFirstMatchString(str string) string {
	for i := 0; i < len(str); i++ {
		substring := str[i:]
		for _, matchString := range matchStrings {
			if strings.HasPrefix(substring, matchString) {
				return matchString
			}
		}
	}
	return ""
}
func getLastMatchString(str string) string {
	//flip string
	for i := len(str); i >= 0; i-- {
		substring := str[:i]
		for _, matchString := range matchStrings {
			if strings.HasSuffix(substring, matchString) {
				return matchString
			}
		}
	}
	return ""
}
