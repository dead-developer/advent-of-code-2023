package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	numberRegex := regexp.MustCompile(`\d+`)

	lines := readInput("day_6/input.txt")

	//remove s

	times := numberRegex.FindAllString(lines[0], -1)
	distances := numberRegex.FindAllString(lines[1], -1)

	var results []int

	for i, time := range times {
		var won = amountOfRecordBeatingTimes(mustConvertToInt(time), mustConvertToInt(distances[i]))
		results = append(results, won)
	}

	//multiply results
	var total = 1
	for _, result := range results {
		total *= result
	}

	fmt.Println(total)

}

func amountOfRecordBeatingTimes(raceTime int, record int) int {
	var counter int
	for holdTime := 0; holdTime < raceTime; holdTime++ {
		distance := distanceTravelled(holdTime, raceTime-holdTime)

		if distance > record {
			counter++
		}
	}

	return counter
}
func distanceTravelled(speed int, travelTime int) int {
	return speed * travelTime
}

func mustConvertToInt(str string) int {
	num, err := strconv.Atoi(str)
	if err != nil {
		panic(err)
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
