package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Node struct {
	id    string
	left  string
	right string
}

var nodes = make(map[string]Node)

var instructions string

func main() {

	lineRegEx := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

	lines := readInput("day_8/input.txt")

	instructions = strings.TrimSpace(lines[0])

	for _, line := range lines[1:] {
		parts := lineRegEx.FindStringSubmatch(line)
		if len(parts) == 0 {
			continue
		}

		var node = Node{
			id:    parts[1],
			left:  parts[2],
			right: parts[3],
		}
		nodes[parts[1]] = node
	}

	var count int
	var nodeId string = "AAA"
	for {
		direction := nextInstruction()
		node := nodes[nodeId]
		if direction == "L" {
			nodeId = node.left
		} else {
			nodeId = node.right
		}

		//fmt.Println(nodeId)

		count++
		if nodeId == "ZZZ" {
			break
		}

	}
	fmt.Println(count)
}

var instructionIndex int = 0

func nextInstruction() string {
	instruction := string(instructions[instructionIndex])
	instructionIndex++
	if instructionIndex >= len(instructions) {
		instructionIndex = 0
	}
	return instruction
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
