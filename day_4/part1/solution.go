package main

import (
	"fmt"
	"github.com/thoas/go-funk"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	lines := readInput("day_4/input.txt")

	total := 0
	for _, line := range lines {

		gameNumber, results, ownNumbers := bruteForceSplitter(line)
		score := calculateScore(gameNumber, results, ownNumbers)

		total += score
	}
	fmt.Println(total)
}

func calculateScore(gameNumber int, results []string, ownNumbers []string) int {
	intersect := funk.IntersectString(results, ownNumbers)
	score := math.Pow(2.0, float64(len(intersect)-1))
	return int(score)
}

func bruteForceSplitter(str string) (int, []string, []string) {
	re := regexp.MustCompile(`\d+`)
	parts := re.FindAllString(str, -1)

	gameNumber := convertToNumber(parts[0])
	winningNumbers := parts[1:11]
	ownNumbers := parts[11:]

	return gameNumber, winningNumbers, ownNumbers
}

func convertToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}
	return num
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
