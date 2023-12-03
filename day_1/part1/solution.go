package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var matchStrings = []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}

func main() {
	lines := readFile("day_1/input.txt")

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
	for i := len(str) - 1; i >= 0; i-- {
		substring := str[:i]
		for _, matchString := range matchStrings {
			if strings.HasSuffix(substring, matchString) {
				return matchString
			}
		}
	}
	return ""
}
