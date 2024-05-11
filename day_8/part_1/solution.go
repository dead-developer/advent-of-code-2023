package main

import (
	"AoC2023/framework"
	"fmt"
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
	value := solution()
	fmt.Println(value)
}

func solution() int {

	lineRegEx := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

	lines := framework.ReadInput("input.txt")

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
	var nodeId = "AAA"
	for {
		direction := nextInstruction()
		node := nodes[nodeId]
		if direction == "L" {
			nodeId = node.left
		} else {
			nodeId = node.right
		}

		count++
		if nodeId == "ZZZ" {
			break
		}

	}
	return count
}

var instructionIndex = 0

func nextInstruction() string {
	instruction := string(instructions[instructionIndex])
	instructionIndex++
	if instructionIndex >= len(instructions) {
		instructionIndex = 0
	}
	return instruction
}
