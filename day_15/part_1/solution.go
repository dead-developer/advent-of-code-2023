package main

import (
	"AoC2023/framework"
	"fmt"
	"strings"
)

var steps []string

const dataFile = "input.txt"

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	loadData()

	totalHash := 0
	for _, step := range steps {
		hash := calculateHash(step)

		totalHash += hash
	}

	return totalHash
}

func calculateHash(step string) int {
	hash := 0
	for _, char := range step {
		hash += int(char)
		hash *= 17
		hash %= 256
	}
	return hash
}

func loadData() {
	lines := framework.ReadInput(dataFile)
	line := lines[0]
	steps = strings.Split(line, ",")
}
