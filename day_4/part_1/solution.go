package main

import (
	"AoC2023/framework"
	"fmt"
	"github.com/thoas/go-funk"
	"math"
	"regexp"
)

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	lines := framework.ReadInput("input.txt")

	total := 0
	for _, line := range lines {

		results, ownNumbers := bruteForceSplitter(line)
		score := calculateScore(results, ownNumbers)

		total += score
	}
	return total
}

func calculateScore(results []string, ownNumbers []string) int {
	intersect := funk.IntersectString(results, ownNumbers)
	score := math.Pow(2.0, float64(len(intersect)-1))
	return int(score)
}

func bruteForceSplitter(str string) ([]string, []string) {
	re := regexp.MustCompile(`\d+`)
	parts := re.FindAllString(str, -1)

	winningNumbers := parts[1:11]
	ownNumbers := parts[11:]

	return winningNumbers, ownNumbers
}
