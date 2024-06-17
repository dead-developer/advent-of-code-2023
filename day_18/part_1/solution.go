package main

import (
	"AoC2023/framework"
	"fmt"
	"log"
	"regexp"
	"strconv"
)

const dataFile = "input.txt"

type instruction struct {
	direction rune
	amount    int
	color     string
}

var instructions []instruction

var grid = make(map[string]string)
var gridMinX, gridMinY int
var gridMaxX, gridMaxY int

var area = make([][]string, 0)
var areaWidth = 0
var areaHeight = 0

type coordinate struct {
	x, y int
}

var fillQueue []coordinate

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

			areaCoords := fmt.Sprintf("%d,%d", posX, posY)
			grid[areaCoords] = instruction.color

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

	buildArea()

	fillInside()
	//displayArea()
	return calculateArea()
}

// find the first empty cell in the second row with trench above it
func findFillStart() (int, int) {
	for x := 0; x < areaWidth; x++ {
		if area[0][x] == "#" && area[1][x] == "." {
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
		if area[cell.y][cell.x] == "#" {
			continue
		}

		area[cell.y][cell.x] = "#"

		for x := cell.x - 1; x <= cell.x+1; x++ {
			if x < 0 || x >= areaWidth {
				continue
			}
			if area[cell.y][x] == "." {
				fillQueue = append(fillQueue, coordinate{x: x, y: cell.y})
			}

		}

		for y := cell.y - 1; y <= cell.y+1; y++ {
			if y < 0 || y > areaHeight {
				continue
			}
			if area[y][cell.x] == "." {
				fillQueue = append(fillQueue, coordinate{x: cell.x, y: y})
			}
		}
	}

}

func buildArea() {

	areaWidth = gridMaxX - gridMinX + 1
	areaHeight = gridMaxY - gridMinY + 1

	area = make([][]string, areaHeight)
	for y := 0; y < areaHeight; y++ {
		area[y] = make([]string, areaWidth)
	}

	for y := gridMinY; y <= gridMaxY; y++ {
		for x := gridMinX; x <= gridMaxX; x++ {
			if _, ok := grid[fmt.Sprintf("%d,%d", x, y)]; ok {
				area[y-gridMinY][x-gridMinX] = "#"
			} else {
				area[y-gridMinY][x-gridMinX] = "."
			}
		}
	}
}

func calculateArea() int {
	var amount int
	for y := 0; y < areaHeight; y++ {
		for x := 0; x < areaWidth; x++ {
			if area[y][x] == "#" {
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
		instructions = append(instructions, instruction{direction: rune(matches[1][0]), amount: amount, color: matches[3]})
	}
}

//func displayArea() {
//
//	for _, line := range area {
//		for _, element := range line {
//			if element == "." {
//				fmt.Print(".")
//			} else {
//				fmt.Print("#")
//			}
//		}
//		fmt.Println()
//	}
//}
