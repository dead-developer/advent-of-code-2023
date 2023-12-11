package main

import (
	"fmt"
	"os"
	"strings"
)

var matrix [][]string

type galaxy struct {
	x, y int
	id   int
}

var galaxies []galaxy

func main() {
	lines := readInput("day_11/input.txt")

	for _, line := range lines {
		matrix = append(matrix, strings.Split(strings.TrimSpace(line), ""))
	}

	expandMatrixX()
	expandMatrixY()

	galaxies = findGalaxies()
	//fmt.Println("Galaxies:", galaxies)

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

	fmt.Println("Sum:", sum)

	//debugPrint()

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
		if (curX == dx) && (curY == dy) {
			break
		}
	}
	return count
}

func debugPrint() {
	for y := 0; y < len(matrix); y++ {
		for x := 0; x < len(matrix[y]); x++ {
			fmt.Print(matrix[y][x])
		}
		fmt.Println()
	}
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
			for y := 0; y < len(matrix); y++ {
				matrix[y] = append(matrix[y][:x], append([]string{"."}, matrix[y][x:]...)...)
			}
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
			// add matrix[y] to matrix[y:]
			matrix = append(matrix[:y], append([][]string{matrix[y]}, matrix[y:]...)...)
		}
		y--
	}
}

func readInput(name string) []string {
	content, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []string{}
	}
	lines := strings.Split(string(content), "\n")
	return lines
}
