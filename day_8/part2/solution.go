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
var instructionIndex int = 0

func main() {

	lineRegEx := regexp.MustCompile(`(\w{3}) = \((\w{3}), (\w{3})\)`)

	lines := readInput("day_8/input.txt")

	instructions = strings.TrimSpace(lines[0])

	var startingNodes []string

	var pathResults = make(map[string]int)

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

		if strings.HasSuffix(node.id, "A") {
			startingNodes = append(startingNodes, node.id)
		}
	}

	for _, startingNode := range startingNodes {
		nodeId := []string{startingNode}
		var count int = 0

		for {
			if isFinished(nodeId) {
				break
			} // check if all end with Z

			// update nodes
			direction := nextInstruction()
			for index, currentNode := range nodeId {
				var node = nodes[currentNode]
				if direction == "L" {
					nodeId[index] = node.left
				} else {
					nodeId[index] = node.right
				}
			}
			count++
		}
		pathResults[startingNode] = count

	}

	// calculate lowest common nominator
	// couldn't solve this myself, found hint online and this is made with lot of help from GiPiTi.
	var lowestCommonNominator int = 1
	for _, value := range pathResults {
		lowestCommonNominator = lcm(lowestCommonNominator, value)
	}
	fmt.Println(lowestCommonNominator)

}

func lcm(a, b int) int {
	return a * b / gcd(a, b)
}
func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func isFinished(nodes []string) bool {
	var finish bool = true
	for _, id := range nodes {
		if !strings.HasSuffix(id, "Z") {
			finish = false
			break
		}
	}
	return finish
}

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
