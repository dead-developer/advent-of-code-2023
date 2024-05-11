package main

import (
	"AoC2023/framework"
	"fmt"
	"strconv"
	"strings"
)

var matrix [][]string
var validPositions [][]int

var gears = make(map[string][]int)

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	lines := framework.ReadInput("input.txt")

	matrix = initializeMatrix(lines)
	initValidPositions()
	initGearPositions()

	findValidGears()

	//displayValidCells()

	var foundNumbers []int
	var currentNumber = ""
	var isValidNumber = false
	var gearLocation = ""

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if isNumber(matrix[x][y]) {
				currentNumber += matrix[x][y]
				if validPositions[x][y] > 0 {
					isValidNumber = true
					if validPositions[x][y] == 2 {
						var gearName = findGearPosition(x, y)
						if len(gears[gearName]) == 2 {
							gearLocation = gearName
						}
					}
				}
			} else {
				if len(currentNumber) > 0 {
					var number, _ = strconv.Atoi(currentNumber)
					if isValidNumber {
						if gearLocation != "" {
							number = gears[gearLocation][0] * gears[gearLocation][1]
							gears[gearLocation] = []int{0, 0} //used up.
							foundNumbers = append(foundNumbers, number)
						}

					}
					currentNumber = ""
					isValidNumber = false
					gearLocation = ""
				}
			}
		}
	}

	total := 0
	for _, number := range foundNumbers {
		total += number
	}
	return total
}

func findValidGears() {
	// count numbers next to gears
	var currentNumber = ""
	var currentGear = ""

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if isNumber(matrix[x][y]) {
				currentNumber += matrix[x][y]
				if validPositions[x][y] == 2 {
					currentGear = findGearPosition(x, y)
				}
			} else {
				if len(currentNumber) > 0 {
					number, _ := strconv.Atoi(currentNumber) //ignore error
					if currentGear != "" {
						gears[currentGear] = append(gears[currentGear], number)
					}
					currentNumber = ""
					currentGear = ""
				}
			}
		}
	}
	return
}

func findGearPosition(x, y int) string {
	for _, position := range neighbourPositions(x, y) {
		if (position[0] < 0) || (position[1] < 0) || (position[0] >= len(matrix)) || (position[1] >= len(matrix[0])) {
			continue
		}
		if matrix[position[0]][position[1]] == "*" {

			return strconv.Itoa(position[0]) + "-" + strconv.Itoa(position[1])
		}
	}
	return ""
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

func markValid(x, y, value int) {
	// value: 1=symbol, 2=gear
	if (x < 0) || (y < 0) || (x >= len(validPositions)) || (y >= len(validPositions[0])) {
		return
	}
	validPositions[x][y] = value
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

func initValidPositions() [][]int {
	validPositions = make([][]int, len(matrix[0]))
	for i := 0; i < len(matrix); i++ {
		validPositions[i] = make([]int, len(matrix))
	}

	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if isSymbol(matrix[x][y]) {
				markValid(x, y, 1)
				for _, position := range neighbourPositions(x, y) {
					markValid(position[0], position[1], 1)
				}
			}
		}
	}
	return validPositions
}

func initGearPositions() [][]int {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[0]); x++ {
			if matrix[x][y] == "*" {
				markValid(x, y, 2)
				for _, position := range neighbourPositions(x, y) {
					markValid(position[0], position[1], 2)
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
