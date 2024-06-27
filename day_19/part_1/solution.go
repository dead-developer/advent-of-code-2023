package main

import (
	"AoC2023/framework"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

const dataFile = "input.txt"

type part struct {
	x, m, a, s, total int
}

var parts []part

type ruleset struct {
	rules []func(ruleIndex int) string
}

var rulesets = make(map[string]ruleset)

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	loadData()

	var sum int

	for partIndex := range parts {

		var currentRule = "in"

		for {
			result := processRules(partIndex, currentRule)
			if result == "R" {
				break
			}
			if result == "A" {
				sum += parts[partIndex].total
				break
			} else {
				currentRule = result
				continue
			}
		}

	}

	return sum
}

func processRules(partIndex int, ruleName string) string {
	for _, ruleFunc := range rulesets[ruleName].rules {
		ruleResult := ruleFunc(partIndex)
		if ruleResult != "" {
			return ruleResult
		}
	}
	panic("NO RESULT FROM RULE")
	return "FAIL"
}

func loadData() {
	lines := framework.ReadInput(dataFile)

	var rulesDone bool

	for _, line := range lines {
		if line == "" {
			rulesDone = true
		}
		if rulesDone {
			parsePart(line)

		} else {
			parseRule(line)
		}

	}

}

func parsePart(partString string) {
	if partString == "" {
		return
	}

	partExp := regexp.MustCompile(`^{x=(\d+),m=(\d+),a=(\d+),s=(\d+)}$`)
	matches := partExp.FindStringSubmatch(partString)

	x, _ := strconv.Atoi(matches[1])
	m, _ := strconv.Atoi(matches[2])
	a, _ := strconv.Atoi(matches[3])
	s, _ := strconv.Atoi(matches[4])

	newPart := part{
		x:     x,
		m:     m,
		a:     a,
		s:     s,
		total: x + m + a + s,
	}
	parts = append(parts, newPart)

}

func parseRule(line string) {

	exp := regexp.MustCompile("^(.+){(.+)}")
	ruleExp := regexp.MustCompile("^(.+)([<>])(.+):(.+)")

	matches := exp.FindStringSubmatch(line)

	if len(matches) != 3 {
		log.Panic("invalid instruction")
	}

	parseRules := strings.Split(matches[2], ",")
	newRuleset := make([]func(ruleIndex int) string, 0)

	name := matches[1]

	for _, rule := range parseRules {

		comparisonMatches := ruleExp.FindStringSubmatch(rule)
		if len(comparisonMatches) == 5 {
			compareToValue, _ := strconv.Atoi(comparisonMatches[3])
			if compareToValue == 0 {
				panic("invalid rule")
			}
			newRuleset = append(newRuleset, func(ruleIndex int) string {
				var value int

				if comparisonMatches[1] == "s" {
					value = parts[ruleIndex].s
				}
				if comparisonMatches[1] == "x" {
					value = parts[ruleIndex].x
				}
				if comparisonMatches[1] == "m" {
					value = parts[ruleIndex].m
				}
				if comparisonMatches[1] == "a" {
					value = parts[ruleIndex].a
				}

				if comparisonMatches[2] == "<" {
					if value < compareToValue {
						return comparisonMatches[4]
					}
				}
				if comparisonMatches[2] == ">" {
					if value > compareToValue {
						return comparisonMatches[4]
					}
				}
				return ""
			})
			continue

		}

		if rule == "A" {
			newRuleset = append(newRuleset, func(ruleIndex int) string {
				return "A"
			})
			continue
		} else if rule == "R" {
			newRuleset = append(newRuleset, func(ruleIndex int) string {
				return "R"
			})
			continue
		} else {
			// pass through
			newRuleset = append(newRuleset, func(ruleIndex int) string {
				return rule
			})
		}

	}

	rulesets[name] = ruleset{rules: newRuleset}

}
