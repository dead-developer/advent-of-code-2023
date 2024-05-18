package main

import (
	"AoC2023/framework"
	"fmt"
	"reflect"
	"strings"
)

var area [][]byte

var width int
var height int

type stone struct {
	x int
	y int
}

const dataFile = "input.txt"

var stones []stone

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	loadData()

	initStones()

	var log []int

	for i := 0; i < 1000; i++ {
		tiltNorth()
		tiltWest()
		tiltSouth()
		tiltEast()

		weight := calculateWeight()
		log = append(log, weight)

	}
	repeatingSize, pattern := findPattern(log)

	index := (1000000000 - 1001) % repeatingSize

	return pattern[index]

}

func findPattern(log []int) (int, []int) {

	sampleSize := 3

	for {
		for i := 0; i < sampleSize; i++ {
			// get last sampleSize elements
			lastSample := log[len(log)-sampleSize:]
			// get previous sampleSize elements
			prevSample := log[len(log)-sampleSize*2 : len(log)-sampleSize]

			// check if lastSample is equal to prevSample, return sample size and sample
			if reflect.DeepEqual(lastSample, prevSample) {
				return sampleSize, lastSample
			}
		}
		sampleSize++
	}
}

func tiltNorth() {
	reset()
	moveCount := 1
	for moveCount > 0 {
		moveCount = 0
		for index, stone := range stones {
			if area[stone.y][stone.x] == 129 { //if 128 bit is set, skip
				continue
			}
			if stone.y == 0 { //already at top
				area[stone.y][stone.x] = 129
				continue
			}
			moveTarget := area[stone.y-1][stone.x]
			if moveTarget == 2 {
				area[stone.y][stone.x] = 129
				continue
			}
			if moveTarget == 129 { // target already locked in place
				area[stone.y][stone.x] = 129
				continue
			}

			if moveTarget == 0 {
				// move
				moveCount++
				area[stone.y][stone.x] = 0
				area[stone.y-1][stone.x] = 1
				stones[index].y--

			}
		}
	}
}
func tiltSouth() {
	reset()
	moveCount := 1
	for moveCount > 0 {
		moveCount = 0
		for index, stone := range stones {
			if area[stone.y][stone.x] == 129 { //if 128 bit is set, skip
				continue
			}
			if stone.y == height-1 { //already at bottom
				area[stone.y][stone.x] = 129
				continue
			}
			moveTarget := area[stone.y+1][stone.x]
			if moveTarget == 2 { // block
				area[stone.y][stone.x] = 129
				continue
			}
			if moveTarget == 129 { // target already locked in place
				area[stone.y][stone.x] = 129
				continue
			}

			if moveTarget == 0 {
				// move
				moveCount++
				area[stone.y][stone.x] = 0
				area[stone.y+1][stone.x] = 1
				stones[index].y++
			}
		}
	}
}
func tiltWest() {
	reset()
	moveCount := 1
	for moveCount > 0 {
		moveCount = 0
		for index, stone := range stones {
			if area[stone.y][stone.x] == 129 { //if 128 bit is set, skip
				continue
			}
			if stone.x == 0 { //already at west
				area[stone.y][stone.x] = 129
				continue
			}
			moveTarget := area[stone.y][stone.x-1]
			if moveTarget == 2 { // block
				area[stone.y][stone.x] = 129
				continue
			}
			if moveTarget == 129 { // target already locked in place
				area[stone.y][stone.x] = 129
				continue
			}

			if moveTarget == 0 {
				// move
				moveCount++
				area[stone.y][stone.x] = 0
				area[stone.y][stone.x-1] = 1
				stones[index].x--
			}
		}
	}
}

func tiltEast() {
	reset()
	moveCount := 1
	for moveCount > 0 {
		moveCount = 0
		for index, stone := range stones {
			if area[stone.y][stone.x] == 129 { //if 128 bit is set, skip
				continue
			}
			if stone.x == width-1 { //already at east
				area[stone.y][stone.x] = 129
				continue
			}
			moveTarget := area[stone.y][stone.x+1]
			if moveTarget == 2 { // block
				area[stone.y][stone.x] = 129
				continue
			}
			if moveTarget == 129 { // target already locked in place
				area[stone.y][stone.x] = 129
				continue
			}

			if moveTarget == 0 {
				// move
				moveCount++
				area[stone.y][stone.x] = 0
				area[stone.y][stone.x+1] = 1
				stones[index].x++
			}
		}
	}
}

// -------------------------------------------------------
func reset() {
	for _, stone := range stones {
		area[stone.y][stone.x] &= 0x7F
	}

}

func loadData() int {
	lines := framework.ReadInput(dataFile)

	height = len(lines)
	width = len(strings.TrimSpace(lines[0]))

	area = make([][]byte, height)

	stoneCount := 0
	for index, line := range lines {
		trimmed := strings.TrimSpace(line)
		row := make([]byte, len(trimmed))

		for i := 0; i < len(trimmed); i++ {
			if trimmed[i] == '.' {
				row[i] = 0
			}
			if trimmed[i] == 'O' {
				row[i] = 1
				stoneCount++
			}
			if trimmed[i] == '#' {
				row[i] = 2
			}
		}
		area[index] = row
	}
	return stoneCount
}

func initStones() {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if area[y][x] == 1 {
				stones = append(stones, stone{x: x, y: y})
			}
		}
	}
}

func displayArray() {
	for _, line := range area {
		for _, char := range line {
			// unset last bit
			char &= 0x7F
			if char == 0 {
				fmt.Print(".")
			}
			if char == 1 {
				fmt.Print("O")
			}
			if char == 2 {
				fmt.Print("#")
			}

		}
		fmt.Println()
	}
}

func calculateWeight() int {
	weight := 0
	for lineNum, line := range area {
		for _, char := range line {
			char &= 0x7F
			if char == 1 {
				weight += len(area) - lineNum
			}
		}
	}
	return weight
}
