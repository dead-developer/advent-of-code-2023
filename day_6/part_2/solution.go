package main

import (
	"AoC2023/framework"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	numberRegex := regexp.MustCompile(`\d+`)

	lines := framework.ReadInput("input.txt")

	//remove spaces from lines
	for i, line := range lines {
		lines[i] = strings.ReplaceAll(line, " ", "")
	}

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
	return total
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
