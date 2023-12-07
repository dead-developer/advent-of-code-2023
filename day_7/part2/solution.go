package main

import (
	"fmt"
	"os"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var cards = "J23456789TQKA"

func main() {
	lineRegEx := regexp.MustCompile(`(\w{5}) (\d+)`)

	lines := readInput("day_7/input.txt")

	var hands [][]int

	for _, line := range lines {
		parts := lineRegEx.FindStringSubmatch(line)
		if len(parts) == 0 {
			continue
		}
		var hand = []rune(parts[1])
		var bid = mustConvertToInt(parts[2])
		handType := analyzeHand(hand)
		sortIndex := calculateSortIndex(handType, hand)

		hands = append(hands, []int{bid, mustConvertToInt(sortIndex)})
	}

	//sort hands by second index
	sort.Slice(hands, func(i, j int) bool {
		return hands[i][1] < hands[j][1]
	})

	//calculate points
	var total int = 0
	for i, hand := range hands {
		total += hand[0] * int(i+1)
		//fmt.Println("Hand:", hand, "Rank:", i+1, "Points:", hand[0]*int(i+1))
	}
	fmt.Println("Total:", total)
}

func calculateSortIndex(handType int, order []rune) string {
	var index string = fmt.Sprintf("%01d", handType)
	for _, v := range order {
		var cardIndex = getCardIndex(v)
		index += fmt.Sprintf("%02d", cardIndex)
	}
	return index

}

func analyzeHand(hand []rune) int {
	set := map[int]int{}
	jokers := 0
	for _, card := range hand {
		var cardIndex = getCardIndex(card)
		if cardIndex == 0 {
			jokers++
		} else {
			set[cardIndex]++
		}

	}

	var processedSet [][]int
	for k, v := range set {
		processedSet = append(processedSet, []int{int(k), int(v)})
	}
	//sort by second index
	sort.Slice(processedSet, func(i, j int) bool {
		if processedSet[i][1] == processedSet[j][1] {
			return processedSet[i][0] > processedSet[j][0]
		}
		return processedSet[i][1] > processedSet[j][1]
	})
	var sortOrder []int
	for i, _ := range processedSet {
		sortOrder = append(sortOrder, processedSet[i][0])
	}

	if jokers == 5 {
		return 7
	}
	if jokers == 4 {
		return 7
	}
	if processedSet[0][1] == 5 {
		return 7
	}
	if processedSet[0][1]+jokers == 5 {
		return 7
	}

	if jokers == 3 {
		return 6
	}

	if processedSet[0][1] == 4 {
		return 6
	}
	if processedSet[0][1]+jokers == 4 && processedSet[0][0] != 0 {
		return 6
	}

	if processedSet[0][1] == 2 && processedSet[1][1] == 2 && jokers == 1 { // full house
		return 5
	}

	if processedSet[0][1] == 3 && processedSet[1][1] == 2 { // full house
		return 5
	}

	if processedSet[0][1] == 2 && processedSet[1][1] == 2 && jokers == 1 { // full house
		return 5
	}

	if jokers == 2 { // three of a kind
		return 4
	}

	if processedSet[0][1] == 3 || processedSet[0][1]+jokers == 3 { // three of a kind
		return 4
	}

	if processedSet[0][1] == 2 && processedSet[1][1] == 2 { // two pair
		return 3
	}

	if processedSet[0][1]+jokers == 2 { // one pair
		return 2
	}
	if jokers == 1 {
		return 2
	}

	return 1
}

func getCardIndex(card rune) int {
	return strings.Index(cards, string(card))
}

func mustConvertToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
	}
	return int(num)
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
