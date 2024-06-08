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
	color     string
}

var instructions []instruction

var grid = make(map[string]uint8)
var gridMinX, gridMinY int
var gridMaxX, gridMaxY int

var areaWidth = 0
var areaHeight = 0

type coordinate struct {
	x, y int
}

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
		for i := 0; i < instruction.amount; i++ {

			grid[strconv.Itoa(posX)+","+strconv.Itoa(posY)] = 1

			switch instruction.direction {
			case 'L':
				posX--
				if posX < gridMinX {
					gridMinX = posX
				}
			case 'R':
				posX++
				if posX > gridMaxX {
					gridMaxX = posX
				}
			case 'U':
				posY--
				if posY < gridMinY {
					gridMinY = posY
				}
			case 'D':
				posY++
				if posY > gridMaxY {
					gridMaxY = posY
				}
			}
		}
	}
	areaWidth = gridMaxX - gridMinX + 1
	areaHeight = gridMaxY - gridMinY + 1

	//fillInside()
	return calculateArea()
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

func fillInside() {
	x, y := findFillStart()

	if x == -1 {
		panic("no fill start found")
	}
	fillQueue = append(fillQueue, coordinate{x: x, y: y})
	for len(fillQueue) > 0 {

		cell := fillQueue[0]
		fillQueue = fillQueue[1:]
		if grid[strconv.Itoa(cell.x)+","+strconv.Itoa(cell.y)] == 1 {
			continue
		}
		grid[strconv.Itoa(cell.x)+","+strconv.Itoa(cell.y)] = 1

		addCell(cell.x-1, cell.y)
		addCell(cell.x+1, cell.y)
		addCell(cell.x, cell.y-1)
		addCell(cell.x, cell.y+1)

	}

}

func addCell(x, y int) {
	if x < 0 || x >= areaWidth {
		return
	}
	if y < 0 || y > areaHeight {
		return
	}
	if grid[strconv.Itoa(x)+","+strconv.Itoa(y)] == 0 {
		fillQueue = append(fillQueue, coordinate{x: x, y: y})
	}
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

//func loadData() {
//	lines := framework.ReadInput(dataFile)
//	exp := regexp.MustCompile(`^(.)\s(\d+)\s\((.+)\)`)
//	for _, line := range lines {
//		matches := exp.FindStringSubmatch(line)
//		if len(matches) != 4 {
//			log.Panic("invalid instruction")
//		}
//		amount, _ := strconv.Atoi(matches[2])
//		instructions = append(instructions, instruction{direction: rune(matches[1][0]), amount: amount, color: matches[3]})
//	}
//}

func loadData() {
	lines := framework.ReadInput(dataFile)
	exp := regexp.MustCompile(`^(.)\s(\d+)\s\((.+)\)`)
	for _, line := range lines {
		matches := exp.FindStringSubmatch(line)
		if len(matches) != 4 {
			log.Panic("invalid instruction")
		}
		//get direction
		lastChar := string(matches[3][6])
		var direction rune
		if lastChar == "0" {
			direction = 'R'
		} else if lastChar == "2" {
			direction = 'L'
		} else if lastChar == "1" {
			direction = 'D'
		} else if lastChar == "3" {
			direction = 'U'
		}

		//get string from 1 to 5 index
		distanceHex := "0" + matches[3][1:6]

		//convert hex to int
		amount, _ := strconv.ParseInt(distanceHex, 16, 64)
		instructions = append(instructions, instruction{direction: direction, amount: int(amount), color: matches[3]})
	}
}
