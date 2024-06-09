package main

import (
	"AoC2023/framework"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const dataFile = "clue.txt"

type instruction struct {
	direction rune
	amount    int
}

var instructions []instruction

type vector struct {
	startX, startY int
	endX, endY     int
}

var vectors []vector

type coordinate struct {
	x, y int
}

var points []coordinate
var grid = make(map[string]uint8)
var gridMinX, gridMinY int

var gridMaxX, gridMaxY int
var areaWidth = 0

var areaHeight = 0

var fillQueue []coordinate

func init() {
	loadData()

}

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	loadData()

	var posX, posY int
	posX = 0
	posY = 0

	for _, instruction := range instructions {

		var line = vector{startX: posX, startY: posY, endX: posX, endY: posY}

		points = append(points, coordinate{x: posX, y: posY})
		switch instruction.direction {
		case 'L':
			posX -= instruction.amount
			line.endX = posX
			if posX < gridMinX {
				gridMinX = posX
			}
		case 'R':
			posX += instruction.amount
			line.endX = posX
			if posX > gridMaxX {
				gridMaxX = posX
			}
		case 'U':
			posY -= instruction.amount
			line.endY = posY
			if posY < gridMinY {
				gridMinY = posY
			}
		case 'D':
			posY += instruction.amount
			line.endY = posY
			if posY > gridMaxY {
				gridMaxY = posY
			}
		}
		vectors = append(vectors, line)
	}

	areaWidth = gridMaxX - gridMinX + 1
	areaHeight = gridMaxY - gridMinY + 1

	fmt.Println(vectors)
	fmt.Println(points)

	return 0
}

// find the first empty cell in the second row with trench above it
func findFillStart() (int, int) {
	for x := 0; x < areaWidth; x++ {
		if grid[strconv.Itoa(x)+","+strconv.Itoa(0)] == 1 && grid[strconv.Itoa(x)+","+strconv.Itoa(1)] == 0 {
			return x, 1
		}
	}
	return -1, -1
}

func calculateArea() int {
	var amount int
	for y := 0; y < areaHeight; y++ {
		for x := 0; x < areaWidth; x++ {
			if grid[strconv.Itoa(x)+","+strconv.Itoa(y)] == 1 {
				amount++
			}
		}
	}
	return amount
}

func loadData() {
	lines := framework.ReadInput(dataFile)
	exp := regexp.MustCompile(`^(.)\s(\d+)\s\((.+)\)`)
	for _, line := range lines {
		matches := exp.FindStringSubmatch(line)
		if len(matches) != 4 {
			log.Panic("invalid instruction")
		}
		amount, _ := strconv.Atoi(matches[2])
		instructions = append(instructions, instruction{direction: rune(matches[1][0]), amount: amount})
	}
}

//func loadData() {
//	lines := framework.ReadInput(dataFile)
//	exp := regexp.MustCompile(`^(.)\s(\d+)\s\((.+)\)`)
//	for _, line := range lines {
//		matches := exp.FindStringSubmatch(line)
//		if len(matches) != 4 {
//			log.Panic("invalid instruction")
//		}
//		//get direction
//		lastChar := string(matches[3][6])
//		var direction rune
//		if lastChar == "0" {
//			direction = 'R'
//		} else if lastChar == "2" {
//			direction = 'L'
//		} else if lastChar == "1" {
//			direction = 'D'
//		} else if lastChar == "3" {
//			direction = 'U'
//		}
//
//		//get string from 1 to 5 index
//		distanceHex := "0" + matches[3][1:6]
//
//		//convert hex to int
//		amount, _ := strconv.ParseInt(distanceHex, 16, 64)
//		instructions = append(instructions, instruction{direction: direction, amount: int(amount), color: matches[3]})
//	}
//}
