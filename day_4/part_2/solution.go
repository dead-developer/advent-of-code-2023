package main

import (
	"AoC2023/framework"
	"fmt"
	"github.com/thoas/go-funk"
	"regexp"
	"strconv"
)

// WARNING: this is a slow brute forcesolution
// TODO: make better

var hitCache map[int]int
var cardNumbers []int

func main() {
	value := solution()
	fmt.Println("Amount of cards:", value)
}

func solution() int {
	cards := framework.ReadInput("input.txt")
	cardNumbers = make([]int, len(cards))

	// make cache for match amounts
	hitCache = make(map[int]int)
	for _, line := range cards {
		gameNumber, results, ownNumbers := bruteForceSplitter(line)
		hitCache[gameNumber] = amountOfMatches(results, ownNumbers)
		cardNumbers[gameNumber-1] = gameNumber
	}

	amountOfCards := processCards(cardNumbers)
	return amountOfCards
}

func processCards(cardList []int) int {
	var counter int
	for _, gameNumber := range cardList {
		counter++
		if hitCache[gameNumber] == 0 {
			continue
		}
		lastCard := gameNumber + hitCache[gameNumber]
		if lastCard > len(cardNumbers) {
			lastCard = len(cardNumbers)
		}
		counter += processCards(cardNumbers[gameNumber:lastCard])
		//fmt.Println("Processing card", gameNumber)
	}
	return counter
}

func amountOfMatches(results []string, ownNumbers []string) int {
	intersect := funk.IntersectString(results, ownNumbers)
	return len(intersect)
}

func bruteForceSplitter(str string) (int, []string, []string) {
	re := regexp.MustCompile(`\d+`)
	parts := re.FindAllString(str, -1)

	gameNumber := convertToNumber(parts[0])
	winningNumbers := parts[1:11]
	ownNumbers := parts[11:]

	// for example
	//winningNumbers := parts[1:6]
	//ownNumbers := parts[6:]

	return gameNumber, winningNumbers, ownNumbers
}

func convertToNumber(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		fmt.Println("Error converting string to int:", err)
	}
	return num
}
