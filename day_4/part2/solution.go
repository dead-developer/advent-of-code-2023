package main

import (
	"fmt"
	"github.com/thoas/go-funk"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*

SUPER SLOW RECURSIVE SOLUTION

*/

func main() {
	cards := readInput("day_4/input.txt")

	amountOfCards := 0
	amountOfCards = processCards(cards, cards)

	fmt.Println("Amount of cards:", amountOfCards)

}

func processCards(processList []string, allCards []string) int {
	counter := 0
	for _, line := range processList {
		gameNumber, results, ownNumbers := bruteForceSplitter(line)
		fmt.Println("Processing card", gameNumber)
		amountOfMatches := amountOfMatches(results, ownNumbers)

		if amountOfMatches > 0 {
			lastCard := gameNumber + amountOfMatches
			if lastCard > len(allCards) {
				lastCard = len(allCards)
			}
			counter += processCards(allCards[gameNumber:lastCard], allCards)
		}
		counter++
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

func readInput(name string) []string {
	content, err := os.ReadFile(name)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return []string{}
	}
	lines := strings.Split(string(content), "\n")
	return lines
}
