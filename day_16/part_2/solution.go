package main

import (
	"AoC2023/framework"
	"fmt"
	"sort"
	"strings"
)

var area [][]string
var energized [][]int

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

	results := make([]int, 0)

	for x := 0; x < areaWidth; x++ {
		results = append(results, energize(head{x: x, y: -1, facing: down}))
		results = append(results, energize(head{x: x, y: areaHeight, facing: up}))
	}
	for y := 0; y < areaWidth; y++ {
		results = append(results, energize(head{x: -1, y: y, facing: right}))
		results = append(results, energize(head{x: areaWidth, y: y, facing: left}))
	}

	sort.Sort(sort.Reverse(sort.IntSlice(results)))

	return results[0]
}

func energize(startPoint head) int {
	heads = make([]head, 0)
	heads = append(heads, startPoint)
	resetEnergized()

	// move heads
	for len(heads) > 0 {

		if !moveHead(0) {
			//remove head
			heads = heads[1:]
			continue
		}

		if area[heads[0].y][heads[0].x] == "|" && (heads[0].facing == right || heads[0].facing == left) {
			if energized[heads[0].y][heads[0].x] == 1 { //already split here
				heads = heads[1:]
				continue
			}
			heads[0].facing = up
			heads = append(heads, head{x: heads[0].x, y: heads[0].y, facing: down})
		}

		if area[heads[0].y][heads[0].x] == "-" && (heads[0].facing == down || heads[0].facing == up) {
			if energized[heads[0].y][heads[0].x] == 1 { //already split here
				heads = heads[1:]
				continue
			}
			heads[0].facing = right
			heads = append(heads, head{x: heads[0].x, y: heads[0].y, facing: left})
		}

		energized[heads[0].y][heads[0].x] = 1

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
			if energized[y][x] == 1 {
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
	}
	resetEnergized()
	areaHeight = len(area)
}

func resetEnergized() {
	energized = make([][]int, 0)

	for y := 0; y < areaWidth; y++ {
		energized = append(energized, make([]int, areaWidth))
	}
}
