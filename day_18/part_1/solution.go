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

	fillArea()
	displayArea()
	return 0
}

//
//
//func fillTrenches() {
//
//	startFill := 0
//	for y := 0; y < areaHeight; y++ {
//		for x := 0; x < areaWidth; x++ {
//			if area[y][x] != "" && startFill == 0 {
//				startFill = 1
//				continue
//			}
//
//			if startFill == 1 {
//				if area[y][x] != "" {
//					startFill = 0
//					break
//				}
//				if area[y][x] == "" {
//					area[y][x] = "#"
//				}
//			}
//
//		}
//	}
//
//}

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

func fillArea() {

	areaWidth = gridMaxX - gridMinX + 1
	areaHeight = gridMaxY - gridMinY + 1

	println(areaWidth)
	println(areaHeight)

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

func displayArea() {

	for _, line := range area {
		for _, element := range line {
			if element == "." {
				fmt.Print(".")
			} else {
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}
