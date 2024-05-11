package main

import (
	"AoC2023/framework"
	"fmt"
	"strings"
)

var matrix [][]string

type galaxy struct {
	x, y int
	id   int
}

var galaxies []galaxy

var emptyX, emptyY []int

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	lines := framework.ReadInput("input.txt")

	for _, line := range lines {
		matrix = append(matrix, strings.Split(strings.TrimSpace(line), ""))
	}

	expandMatrixX()
	expandMatrixY()

	galaxies = findGalaxies()

	var pairs [][]galaxy
	for i := 0; i < len(galaxies); i++ {
		for j := i + 1; j < len(galaxies); j++ {
			pairs = append(pairs, []galaxy{galaxies[i], galaxies[j]})
		}
	}

	var sum int
	for _, pair := range pairs {
		sum += countDistance(pair[0].x, pair[0].y, pair[1].x, pair[1].y)
	}

	return sum
}

func findGalaxies() []galaxy {
	var galaxies []galaxy
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] == "#" {
				galaxies = append(galaxies, galaxy{x, y, len(galaxies) + 1})
			}
		}
	}
	return galaxies
}

func countDistance(x, y, dx, dy int) int {
	var count int

	var multiplier = 1000000

	curX, curY := x, y

	for {
		if curX > dx {
			curX--
			count++
		}
		if curX < dx {
			curX++
			count++
		}
		if curY < dy {
			curY++
			count++
		}
		if curY > dy {
			curY--
			count++
		}
		// check if curX in emptyX
		if inArray(emptyX, curX) {
			count += multiplier - 1
		}
		if inArray(emptyY, curY) {
			count += multiplier - 1
		}
		if (curX == dx) && (curY == dy) {
			break
		}
	}
	return count
}

func inArray(arr []int, val int) bool {
	for _, v := range arr {
		if v == val {
			return true
		}
	}
	return false
}

func expandMatrixX() {
	// if whole row is empty, double it
	x := len(matrix[0]) - 1
	for x >= 0 {
		empty := true
		for y := 0; y < len(matrix); y++ {
			if matrix[y][x] != "." {
				empty = false
				break
			}
		}
		if empty {
			emptyX = append(emptyX, x)
		}
		x--
	}
}
func expandMatrixY() {
	// if whole row is empty, double it
	y := len(matrix) - 1
	for y >= 0 {

		empty := true
		for x := 0; x < len(matrix[y]); x++ {
			if matrix[y][x] != "." {
				empty = false
				break
			}
		}
		if empty {
			emptyY = append(emptyY, y)
		}
		y--
	}
}
