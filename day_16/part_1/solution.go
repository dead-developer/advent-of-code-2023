package main

import (
	"AoC2023/framework"
	"fmt"
	"strings"
)

var area [][]string
var energized [][]string

type direction int

const (
	up    direction = 0
	down  direction = 1
	left  direction = 2
	right direction = 3
)

type head struct {
	x      int
	y      int
	facing direction
}

var areaWidth int
var areaHeight int

var heads []head

const dataFile = "input.txt"

func main() {
	value := solution()
	fmt.Println(value)

}

func solution() int {
	loadData()

	// add initial head
	heads = append(heads, head{x: -1, y: 0, facing: right})
	energizedCount := 0

	// move heads
	for len(heads) > 0 {

		if !moveHead(0) {
			//remove head
			heads = heads[1:]
			continue
		}

		if area[heads[0].y][heads[0].x] == "|" && (heads[0].facing == right || heads[0].facing == left) {
			if energized[heads[0].y][heads[0].x] == "#" { //already split here
				heads = heads[1:]
				continue
			}
			heads[0].facing = up
			heads = append(heads, head{x: heads[0].x, y: heads[0].y, facing: down})
		}

		if area[heads[0].y][heads[0].x] == "-" && (heads[0].facing == down || heads[0].facing == up) {
			if energized[heads[0].y][heads[0].x] == "#" { //already split here
				heads = heads[1:]
				continue
			}
			heads[0].facing = right
			heads = append(heads, head{x: heads[0].x, y: heads[0].y, facing: left})
		}

		if energized[heads[0].y][heads[0].x] != "#" {
			energizedCount++
			energized[heads[0].y][heads[0].x] = "#"
		}

		if area[heads[0].y][heads[0].x] == "/" {
			if heads[0].facing == right {
				heads[0].facing = up
			} else if heads[0].facing == up {
				heads[0].facing = right
			} else if heads[0].facing == left {
				heads[0].facing = down
			} else if heads[0].facing == down {
				heads[0].facing = left
			}
		}

		if area[heads[0].y][heads[0].x] == "\\" {
			if heads[0].facing == left {
				heads[0].facing = up
			} else if heads[0].facing == up {
				heads[0].facing = left
			} else if heads[0].facing == right {
				heads[0].facing = down
			} else if heads[0].facing == down {
				heads[0].facing = right
			}
		}

	}
	return countEnergizedCells()
}

func directionToVector(facing direction) (int, int) {
	switch facing {
	case up:
		return 0, -1
	case down:
		return 0, 1
	case left:
		return -1, 0
	case right:
		return 1, 0
	}
	return 0, 0
}

func moveHead(headIndex int) bool {
	x, y := directionToVector(heads[headIndex].facing)

	heads[headIndex].x += x
	heads[headIndex].y += y

	// check if head is in a valid position
	if heads[headIndex].x < 0 || heads[headIndex].x >= areaWidth || heads[headIndex].y < 0 || heads[headIndex].y >= areaHeight {
		return false
	}

	return true
}

func countEnergizedCells() int {
	var count int
	for y := 0; y < areaHeight; y++ {
		for x := 0; x < areaWidth; x++ {
			if energized[y][x] == "#" {
				count++
			}
		}
	}
	return count
}

func loadData() {
	lines := framework.ReadInput(dataFile)

	areaWidth = len(lines[0])
	for _, line := range lines {
		area = append(area, strings.Split(line, ""))

		energizedLine := make([]string, len(line))
		for i := range energizedLine {
			energizedLine[i] = "."
		}
		energized = append(energized, energizedLine)
	}
	areaHeight = len(area)
}

func displayArea(displayArea [][]string) {
	for y, line := range displayArea {
		for x, element := range line {
			fmt.Printf(element)
			if x == len(line)-1 {
				fmt.Println()
			}
		}
		if y == len(displayArea)-1 {
			fmt.Println()
		}
	}
}
