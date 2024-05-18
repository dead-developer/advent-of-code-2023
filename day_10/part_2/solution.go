package main

import (
	"AoC2023/framework"
	"fmt"
)

type vector struct {
	x int
	y int
}

type pipe struct {
	distance   int
	location   vector
	directions []vector
}

var maze = make(map[vector]pipe)
var start vector

var fillMatrix = make(map[vector]int)

var size vector

var queue []vector
var fillMatrixQueue []vector

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	lines := framework.ReadInput("input.txt")
	readMaze(lines)
	updateDistances(maze[start])

	for len(queue) > 0 {
		var location = queue[0]
		updateDistances(maze[location])
		queue = queue[1:]
	}

	size = vector{len(lines[0]), len(lines)}
	buildFillMatrix()

	for y := 0; y < size.y*2; y++ {
		fill(vector{0, y})
		fill(vector{size.x*2 - 1, y})
	}
	for x := 0; x < size.y*2; x++ {
		fill(vector{x, 0})
		fill(vector{x, size.y*2 - 1})
	}

	//count empties in coordinates divided by 2
	enclosed := 0
	for y := 0; y < size.y*2; y++ {
		for x := 0; x < size.x*2; x++ {
			if x%2 == 0 && y%2 == 0 {
				if fillMatrix[vector{x, y}] == 0 {
					enclosed++
				}
			}

		}
	}
	return enclosed

}

func fill(startPos vector) {
	fillMatrixQueue = append(fillMatrixQueue, startPos)
	for len(fillMatrixQueue) > 0 {
		currentPos := fillMatrixQueue[0]
		fillMatrixQueue = fillMatrixQueue[1:]
		if fillMatrix[currentPos] > 0 {
			continue
		}
		fillMatrix[currentPos] = 2
		if currentPos.y > 0 {
			fillMatrixQueue = append(fillMatrixQueue, vector{currentPos.x, currentPos.y - 1})
		}
		if currentPos.y < size.y*2 {
			fillMatrixQueue = append(fillMatrixQueue, vector{currentPos.x, currentPos.y + 1})
		}
		if currentPos.x > 0 {
			fillMatrixQueue = append(fillMatrixQueue, vector{currentPos.x - 1, currentPos.y})
		}
		if currentPos.x < size.x*2 {
			fillMatrixQueue = append(fillMatrixQueue, vector{currentPos.x + 1, currentPos.y})
		}
	}
}

func buildFillMatrix() {
	for y := 0; y < size.y; y++ {
		for x := 0; x < size.x; x++ {
			if currentPipe, ok := maze[vector{x, y}]; ok {
				if currentPipe.distance == -1 {
					continue
				}
				fillMatrix[vector{x * 2, y * 2}] = 1
				if isConnectedTo(currentPipe, vector{x, y - 1}) {
					fillMatrix[vector{x * 2, y*2 - 1}] = 1
				}
				if isConnectedTo(currentPipe, vector{x, y + 1}) {
					fillMatrix[vector{x * 2, y*2 + 1}] = 1
				}
				if isConnectedTo(currentPipe, vector{x + 1, y}) {
					fillMatrix[vector{x*2 + 1, y * 2}] = 1
				}
				if isConnectedTo(currentPipe, vector{x - 1, y}) {
					fillMatrix[vector{x*2 - 1, y * 2}] = 1
				}
			}
		}
	}

}

func updateDistances(currentPipe pipe) {
	for _, direction := range currentPipe.directions {
		if pipe, ok := maze[direction]; ok {
			if pipe.distance == -1 {
				pipe.distance = currentPipe.distance + 1
				maze[pipe.location] = pipe
				queue = append(queue, pipe.location)
			}
		}
	}
}

func updateStartPipe() {
	startPipe, _ := maze[start]
	startPipe.directions = getReverseNeighbours(startPipe)
	startPipe.distance = 0
	maze[start] = startPipe
}

func getPipe(location vector) (pipe, bool) {
	if pipe, ok := maze[location]; ok {
		return pipe, true
	}
	return pipe{}, false
}
func getReverseNeighbours(target pipe) []vector {
	var neighbours []vector

	// check up
	if sourcePipe, ok := getPipe(vector{target.location.x, target.location.y - 1}); ok {
		if isConnectedTo(sourcePipe, start) {
			neighbours = append(neighbours, sourcePipe.location)
		}
	}
	// check down
	if sourcePipe, ok := getPipe(vector{target.location.x, target.location.y + 1}); ok {
		if isConnectedTo(sourcePipe, start) {
			neighbours = append(neighbours, sourcePipe.location)
		}
	}
	if sourcePipe, ok := getPipe(vector{target.location.x + 1, target.location.y}); ok {

		if isConnectedTo(sourcePipe, start) {
			neighbours = append(neighbours, sourcePipe.location)
		}
	}
	if sourcePipe, ok := getPipe(vector{target.location.x - 1, target.location.y}); ok {

		if isConnectedTo(sourcePipe, start) {
			neighbours = append(neighbours, sourcePipe.location)
		}
	}

	return neighbours
}

func isConnectedTo(sourcePipe pipe, target vector) bool {
	for _, direction := range sourcePipe.directions {
		if direction == target {
			return true
		}
	}
	return false
}

func readMaze(lines []string) {
	for y, line := range lines {
		for x, element := range line {
			char := string(element)
			var directions []vector
			if char == "." {
				continue
			}
			if char == "S" {
				start = vector{x, y}
			}
			if char == "|" {
				directions = []vector{
					{x, y - 1},
					{x, y + 1},
				}
			}
			if char == "-" {
				directions = []vector{
					{x - 1, y},
					{x + 1, y},
				}
			}
			if char == "L" {
				directions = []vector{
					{x, y - 1},
					{x + 1, y},
				}
			}
			if char == "J" {
				directions = []vector{
					{x, y - 1},
					{x - 1, y},
				}
			}
			if char == "7" {
				directions = []vector{
					{x, y + 1},
					{x - 1, y},
				}
			}
			if char == "F" {
				directions = []vector{
					{x, y + 1},
					{x + 1, y},
				}
			}
			//fmt.Println(string(element))

			maze[vector{x, y}] = pipe{
				location:   vector{x, y},
				directions: directions,
				distance:   -1,
			}
			//fmt.Println(x, y, element)
		}
	}
	updateStartPipe()
}
