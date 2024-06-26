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

type rule struct {
	parameter string
	action    string
	compareTo int
	result    string
}

type ruleset struct {
	rules []rule
}

var rulesets = make(map[string]ruleset)

var results []int

func main() {
	value := solution()
	fmt.Println(value)
}

func solution() int {
	loadData()

	limitData := map[string]int{
		"xMin": 1,
		"xMax": 4000,
		"mMin": 1,
		"mMax": 4000,
		"aMin": 1,
		"aMax": 4000,
		"sMin": 1,
		"sMax": 4000,
	}

	processRules("in", limitData)

	var total int
	for _, result := range results {
		total += result
	}

	return total
}

func processRules(ruleName string, limitData map[string]int) {

	for _, rule := range rulesets[ruleName].rules {

		if rule.action == "GOTO" {
			processRules(rule.result, limitData)
			return
		} else if rule.action == "R" {
			// rejected. Stop
			return
		} else if rule.action == "A" {
			addToResult(limitData)
			return
		} else if rule.action == "<" || rule.action == ">" {
			newLimitData := copyMap(limitData)

			if rule.action == "<" {
				limitData[rule.parameter+"Min"] = rule.compareTo
				newLimitData[rule.parameter+"Max"] = rule.compareTo - 1
			} else {
				limitData[rule.parameter+"Max"] = rule.compareTo
				newLimitData[rule.parameter+"Min"] = rule.compareTo + 1
			}

			if rule.result == "A" {
				addToResult(newLimitData)
			} else if rule.result == "R" {

			} else {
				processRules(rule.result, newLimitData)
			}
		}
	}
}

func copyMap(source map[string]int) map[string]int {
	newMap := make(map[string]int, len(source))
	for key, val := range source {
		newMap[key] = val
	}
	return newMap
}

func addToResult(limitData map[string]int) {
	//calculate combinations

	var combinations int
	combinations = (limitData["xMax"] - limitData["xMin"] + 1) * (limitData["mMax"] - limitData["mMin"] + 1) * (limitData["aMax"] - limitData["aMin"] + 1) * (limitData["sMax"] - limitData["sMin"] + 1)
	results = append(
		results,
		combinations,
	)

}

func loadData() {
	lines := framework.ReadInput(dataFile)

	for _, line := range lines {
		if line == "" {

			break
		}
		parseRule(line)
	}
}

func parseRule(line string) {

	exp := regexp.MustCompile("^(.+){(.+)}")
	ruleExp := regexp.MustCompile("^(.+)([<>])(.+):(.+)")

	matches := exp.FindStringSubmatch(line)
	RuleSetName := matches[1]

	if len(matches) != 3 {
		log.Panic("invalid instruction")
	}

	parseRules := strings.Split(matches[2], ",")

	var newRuleset []rule

	for _, ruleString := range parseRules {

		comparisonMatches := ruleExp.FindStringSubmatch(ruleString)

		if len(comparisonMatches) == 5 {
			compareToValue, err := strconv.Atoi(comparisonMatches[3])
			if err != nil {
				panic("invalid rule")
			}

			newRuleset = append(newRuleset, rule{parameter: comparisonMatches[1], action: comparisonMatches[2], compareTo: compareToValue, result: comparisonMatches[4]})
			continue
		}
		if ruleString == "A" {
			newRuleset = append(newRuleset, rule{parameter: "", action: "A", compareTo: 0, result: ""})
			continue
		} else if ruleString == "R" {
			newRuleset = append(newRuleset, rule{parameter: "", action: "R", compareTo: 0, result: ""})
			continue
		} else {
			newRuleset = append(newRuleset, rule{parameter: "", action: "GOTO", compareTo: 0, result: ruleString})
		}
	}
	rulesets[RuleSetName] = ruleset{rules: newRuleset}

}
