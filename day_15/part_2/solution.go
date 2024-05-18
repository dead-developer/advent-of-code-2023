package main

import (
	"AoC2023/framework"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const dataFile = "input.txt"

var steps []string

type lens struct {
	label string
	focal int
}

var boxes = [256][]lens{}

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	loadData()

	for _, step := range steps {

		lens, operator := parseLens(step)
		boxNumber := calculateHash(lens.label)
		if operator == "=" {
			addToBox(lens, boxNumber)
		} else if operator == "-" {
			removeFromBox(lens, boxNumber)
		}

	}

	sum := 0
	for boxNumber, box := range boxes {
		if len(box) == 0 {
			continue
		}
		for slotNumber, lens := range box {
			value := (boxNumber + 1) * (slotNumber + 1) * lens.focal
			sum += value
		}
	}

	return sum
}

func addToBox(lens lens, box int) {
	// check if label already exists, if so, replace focal
	for index, l := range boxes[box] {
		if l.label == lens.label {
			boxes[box][index].focal = lens.focal
			return
		}
	}

	boxes[box] = append(boxes[box], lens)
}

func removeFromBox(lens lens, box int) {
	// check if label already exists, if so, remove
	for i, l := range boxes[box] {
		if l.label == lens.label {
			boxes[box] = append(boxes[box][:i], boxes[box][i+1:]...)
			return
		}
	}
}

func parseLens(step string) (lens, string) {

	pattern := regexp.MustCompile("(.*)([=-])(.*)")
	parts := pattern.FindStringSubmatch(step)

	label := parts[1]
	operator := parts[2]
	focal := 0
	if len(parts) > 1 {
		focal, _ = strconv.Atoi(parts[3])
	}
	return lens{label, focal}, operator
}

func calculateHash(step string) int {
	hash := 0
	for _, char := range step {
		hash += int(char)
		hash *= 17
		hash %= 256
	}
	return hash
}

func loadData() {
	lines := framework.ReadInput(dataFile)
	line := lines[0]
	steps = strings.Split(line, ",")
}

func debug() {
	// print boxes with contents and index

	for i, box := range boxes {
		if len(box) == 0 {
			continue
		}
		fmt.Printf("Box %d\n", i)
		for _, l := range box {
			fmt.Printf("  %s: %d\n", l.label, l.focal)
		}
	}
}
