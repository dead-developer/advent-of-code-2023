package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

var matrix [][]string
var validPositions [][]bool

func main() {
	lines := readFile("day_3/input.txt")

	matrix = initializeMatrix(lines)
	validPositions = buildValidPositions()

	//displayValidCells()

	var foundNumbers = []string{}
	var currentNumber = ""
	var validNumber = false

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if isNumber(matrix[x][y]) {
				currentNumber += matrix[x][y]
				if validPositions[x][y] {
					validNumber = true
				}
			} else {
				if len(currentNumber) > 0 {
					if validNumber {
						foundNumbers = append(foundNumbers, currentNumber)
					}
					currentNumber = ""
					validNumber = false

				}
			}
		}
	}

	total := 0
	for _, number := range foundNumbers {
		value, _ := strconv.Atoi(number)
		total += value
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

func isSymbol(str string) bool {
	if str == "." {
		return false
	}
	return !isNumber(str)
}

func isNumber(str string) bool {
	_, err := strconv.Atoi(str)
	if err == nil {
		return true
	}
	return false
}

func markValid(x, y int) {
	if (x < 0) || (y < 0) || (x >= len(validPositions)) || (y >= len(validPositions[0])) {
		return
	}
	validPositions[x][y] = true
}

func initializeMatrix(lines []string) [][]string {
	width := len(strings.TrimSpace(lines[0]))
	height := len(lines)
	var matrix = make([][]string, width)
	for i := 0; i < height; i++ {
		matrix[i] = make([]string, height)
	}
	for y, line := range lines {
		line = strings.TrimSpace(line)
		for x, char := range line {
			matrix[x][y] = string(char)
		}
	}
	return matrix
}

func buildValidPositions() [][]bool {
	validPositions = make([][]bool, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		validPositions[i] = make([]bool, len(matrix))
	}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if isSymbol(matrix[x][y]) {
				markValid(x, y)
				for _, position := range neighbourPositions(x, y) {
					markValid(position[0], position[1])
				}
			}
		}
	}
	return validPositions
}

func neighbourPositions(x, y int) [][]int {
	positions := [][]int{
		{x - 1, y - 1},
		{x, y - 1},
		{x + 1, y - 1},

		{x - 1, y},
		{x + 1, y},

		{x - 1, y + 1},
		{x, y + 1},
		{x + 1, y + 1},
	}
	return positions
}

func displayValidCells() {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if validPositions[x][y] {
				fmt.Print("X")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
	fmt.Println("\n\n")
}
